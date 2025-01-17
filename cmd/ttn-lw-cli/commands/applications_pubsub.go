// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"encoding/hex"
	"os"
	"strings"

	pbtypes "github.com/gogo/protobuf/types"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"go.thethings.network/lorawan-stack/v3/cmd/internal/io"
	"go.thethings.network/lorawan-stack/v3/cmd/ttn-lw-cli/internal/api"
	"go.thethings.network/lorawan-stack/v3/cmd/ttn-lw-cli/internal/util"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

var (
	selectApplicationPubSubFlags         = util.FieldMaskFlags(&ttnpb.ApplicationPubSub{})
	setApplicationPubSubFlags            = util.FieldFlags(&ttnpb.ApplicationPubSub{})
	natsProviderApplicationPubSubFlags   = util.HideFlagSet(util.FieldFlags(&ttnpb.ApplicationPubSub_NATSProvider{}, "nats"))
	mqttProviderApplicationPubSubFlags   = util.HideFlagSet(util.FieldFlags(&ttnpb.ApplicationPubSub_MQTTProvider{}, "mqtt"))
	awsiotProviderApplicationPubSubFlags = util.HideFlagSet(util.FieldFlags(&ttnpb.ApplicationPubSub_AWSIoTProvider{}, "aws_iot"))
	awsiotDefaultIntegrationPubSubFlags  = util.HideFlagSet(util.FieldFlags(&ttnpb.ApplicationPubSub_AWSIoTProvider_DefaultIntegration{}, "aws_iot", "deployment", "default"))

	selectAllApplicationPubSubFlags = util.SelectAllFlagSet("application pub/sub")
)

func applicationPubSubIDFlags() *pflag.FlagSet {
	flagSet := &pflag.FlagSet{}
	flagSet.String("application-id", "", "")
	flagSet.String("pubsub-id", "", "")
	return flagSet
}

func applicationPubSubProviderFlags() *pflag.FlagSet {
	flagSet := &pflag.FlagSet{}
	flagSet.Bool("nats", false, "use the NATS provider")
	util.HideFlag(flagSet, "nats")
	flagSet.AddFlagSet(natsProviderApplicationPubSubFlags)
	flagSet.Bool("mqtt", false, "use the MQTT provider")
	util.HideFlag(flagSet, "mqtt")
	flagSet.AddFlagSet(mqttProviderApplicationPubSubFlags)
	flagSet.AddFlagSet(util.HideFlagSet(dataFlags("mqtt.tls-ca", "")))
	flagSet.AddFlagSet(util.HideFlagSet(dataFlags("mqtt.tls-client-cert", "")))
	flagSet.AddFlagSet(util.HideFlagSet(dataFlags("mqtt.tls-client-key", "")))
	flagSet.Bool("aws-iot", false, "use the AWS IoT provider")
	util.HideFlag(flagSet, "aws-iot")
	flagSet.AddFlagSet(awsiotProviderApplicationPubSubFlags)
	flagSet.AddFlagSet(awsiotDefaultIntegrationPubSubFlags)
	addDeprecatedProviderFlags(flagSet)
	return flagSet
}

func addDeprecatedProviderFlags(flagSet *pflag.FlagSet) {
	util.DeprecateFlag(flagSet, "nats_server_url", "nats.server_url")
}

func forwardDeprecatedProviderFlags(flagSet *pflag.FlagSet) {
	util.ForwardFlag(flagSet, "nats_server_url", "nats.server_url")
}

var errNoPubSubID = errors.DefineInvalidArgument("no_pub_sub_id", "no pub/sub ID set")

func getApplicationPubSubID(flagSet *pflag.FlagSet, args []string) (*ttnpb.ApplicationPubSubIdentifiers, error) {
	forwardDeprecatedProviderFlags(flagSet)
	applicationID, _ := flagSet.GetString("application-id")
	pubsubID, _ := flagSet.GetString("pubsub-id")
	switch len(args) {
	case 0:
	case 1:
		logger.Warn("Only single ID found in arguments, not considering arguments")
	case 2:
		applicationID = args[0]
		pubsubID = args[1]
	default:
		logger.Warn("Multiple IDs found in arguments, considering the first")
		applicationID = args[0]
		pubsubID = args[1]
	}
	if applicationID == "" {
		return nil, errNoApplicationID.New()
	}
	if pubsubID == "" {
		return nil, errNoPubSubID.New()
	}
	return &ttnpb.ApplicationPubSubIdentifiers{
		ApplicationIds: &ttnpb.ApplicationIdentifiers{ApplicationId: applicationID},
		PubSubId:       pubsubID,
	}, nil
}

var (
	applicationsPubSubsCommand = &cobra.Command{
		Use:     "pubsubs",
		Aliases: []string{"pubsub", "ps"},
		Short:   "Application pub/sub commands",
	}
	applicationsPubSubsGetFormatsCommand = &cobra.Command{
		Use:     "get-formats",
		Aliases: []string{"formats", "list-formats"},
		Short:   "Get the available formats for application pubsubs",
		RunE: func(cmd *cobra.Command, args []string) error {
			as, err := api.Dial(ctx, config.ApplicationServerGRPCAddress)
			if err != nil {
				return err
			}
			res, err := ttnpb.NewApplicationPubSubRegistryClient(as).GetFormats(ctx, ttnpb.Empty)
			if err != nil {
				return err
			}

			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	applicationsPubSubsGetCommand = &cobra.Command{
		Use:     "get [application-id] [pubsub-id]",
		Aliases: []string{"info"},
		Short:   "Get the properties of an application pub/sub",
		RunE: func(cmd *cobra.Command, args []string) error {
			forwardDeprecatedProviderFlags(cmd.Flags())
			pubsubID, err := getApplicationPubSubID(cmd.Flags(), args)
			if err != nil {
				return err
			}
			paths := util.SelectFieldMask(cmd.Flags(), selectApplicationPubSubFlags)
			if len(paths) == 0 {
				logger.Warn("No fields selected, will select everything")
				selectApplicationPubSubFlags.VisitAll(func(flag *pflag.Flag) {
					paths = append(paths, strings.Replace(flag.Name, "-", "_", -1))
				})
			}
			paths = ttnpb.AllowedFields(paths, ttnpb.RPCFieldMaskPaths["/ttn.lorawan.v3.ApplicationPubSubRegistry/Get"].Allowed)

			as, err := api.Dial(ctx, config.ApplicationServerGRPCAddress)
			if err != nil {
				return err
			}
			res, err := ttnpb.NewApplicationPubSubRegistryClient(as).Get(ctx, &ttnpb.GetApplicationPubSubRequest{
				Ids:       pubsubID,
				FieldMask: &pbtypes.FieldMask{Paths: paths},
			})
			if err != nil {
				return err
			}

			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	applicationsPubSubsListCommand = &cobra.Command{
		Use:     "list [application-id]",
		Aliases: []string{"ls"},
		Short:   "List application pubsubs",
		RunE: func(cmd *cobra.Command, args []string) error {
			forwardDeprecatedProviderFlags(cmd.Flags())
			appID := getApplicationID(cmd.Flags(), args)
			if appID == nil {
				return errNoApplicationID.New()
			}
			paths := util.SelectFieldMask(cmd.Flags(), selectApplicationPubSubFlags)
			if len(paths) == 0 {
				logger.Warn("No fields selected, will select everything")
				selectApplicationPubSubFlags.VisitAll(func(flag *pflag.Flag) {
					paths = append(paths, strings.Replace(flag.Name, "-", "_", -1))
				})
			}
			paths = ttnpb.AllowedFields(paths, ttnpb.RPCFieldMaskPaths["/ttn.lorawan.v3.ApplicationPubSubRegistry/List"].Allowed)

			as, err := api.Dial(ctx, config.ApplicationServerGRPCAddress)
			if err != nil {
				return err
			}
			res, err := ttnpb.NewApplicationPubSubRegistryClient(as).List(ctx, &ttnpb.ListApplicationPubSubsRequest{
				ApplicationIds: appID,
				FieldMask:      &pbtypes.FieldMask{Paths: paths},
			})
			if err != nil {
				return err
			}

			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	applicationsPubSubsSetCommand = &cobra.Command{
		Use:     "set [application-id] [pubsub-id]",
		Aliases: []string{"update"},
		Short:   "Set the properties of an application pub/sub",
		RunE: func(cmd *cobra.Command, args []string) error {
			forwardDeprecatedProviderFlags(cmd.Flags())
			pubsubID, err := getApplicationPubSubID(cmd.Flags(), args)
			if err != nil {
				return err
			}
			paths := util.UpdateFieldMask(cmd.Flags(), setApplicationPubSubFlags)

			as, err := api.Dial(ctx, config.ApplicationServerGRPCAddress)
			if err != nil {
				return err
			}
			pubsub, err := ttnpb.NewApplicationPubSubRegistryClient(as).Get(ctx, &ttnpb.GetApplicationPubSubRequest{
				Ids:       pubsubID,
				FieldMask: &pbtypes.FieldMask{Paths: paths},
			})
			if err != nil && !errors.IsNotFound(err) {
				return err
			}
			if pubsub == nil {
				pubsub = &ttnpb.ApplicationPubSub{Ids: pubsubID}
			}

			if err = util.SetFields(pubsub, setApplicationPubSubFlags); err != nil {
				return err
			}

			if nats, _ := cmd.Flags().GetBool("nats"); nats {
				if pubsub.GetNats() == nil {
					paths = append(paths, "provider")
					pubsub.Provider = &ttnpb.ApplicationPubSub_Nats{
						Nats: &ttnpb.ApplicationPubSub_NATSProvider{},
					}
				} else {
					providerPaths := util.UpdateFieldMask(cmd.Flags(), natsProviderApplicationPubSubFlags)
					providerPaths = ttnpb.FieldsWithPrefix("provider", providerPaths...)
					paths = append(paths, providerPaths...)
				}
				if err = util.SetFields(pubsub.GetNats(), natsProviderApplicationPubSubFlags, "nats"); err != nil {
					return err
				}
			}

			if mqtt, _ := cmd.Flags().GetBool("mqtt"); mqtt {
				if pubsub.GetMqtt() == nil {
					paths = append(paths, "provider")
					pubsub.Provider = &ttnpb.ApplicationPubSub_Mqtt{
						Mqtt: &ttnpb.ApplicationPubSub_MQTTProvider{},
					}
				} else {
					providerPaths := util.UpdateFieldMask(cmd.Flags(), mqttProviderApplicationPubSubFlags)
					providerPaths = ttnpb.FieldsWithPrefix("provider", providerPaths...)
					paths = append(paths, providerPaths...)
				}
				if useTLS, _ := cmd.Flags().GetBool("mqtt.use-tls"); useTLS {
					for _, name := range []string{
						"mqtt.tls-ca",
						"mqtt.tls-client-cert",
						"mqtt.tls-client-key",
					} {
						data, err := getDataBytes(name, cmd.Flags())
						if err != nil {
							return err
						}
						err = cmd.Flags().Set(name, hex.EncodeToString(data))
						if err != nil {
							return err
						}
					}
				}
				if err = util.SetFields(pubsub.GetMqtt(), mqttProviderApplicationPubSubFlags, "mqtt"); err != nil {
					return err
				}
			}

			if awsiot, _ := cmd.Flags().GetBool("aws-iot"); awsiot {
				if pubsub.GetAwsIot() == nil {
					paths = append(paths, "provider")
					pubsub.Provider = &ttnpb.ApplicationPubSub_AwsIot{
						AwsIot: &ttnpb.ApplicationPubSub_AWSIoTProvider{},
					}
				} else {
					providerPaths := util.UpdateFieldMask(cmd.Flags(), awsiotProviderApplicationPubSubFlags)
					providerPaths = ttnpb.FieldsWithPrefix("provider", providerPaths...)
					paths = append(paths, providerPaths...)
				}
				if err = util.SetFields(pubsub.GetAwsIot(), awsiotProviderApplicationPubSubFlags, "aws_iot"); err != nil {
					return err
				}
				if defaultStackName, _ := cmd.Flags().GetString("aws-iot.deployment.default.stack-name"); defaultStackName != "" {
					defaultPaths := util.UpdateFieldMask(cmd.Flags(), awsiotDefaultIntegrationPubSubFlags)
					defaultPaths = ttnpb.FieldsWithPrefix("provider", defaultPaths...)
					paths = append(paths, defaultPaths...)
					defaultIntegration := &ttnpb.ApplicationPubSub_AWSIoTProvider_DefaultIntegration{}
					if err = util.SetFields(defaultIntegration, awsiotDefaultIntegrationPubSubFlags, "aws_iot", "deployment", "default"); err != nil {
						return err
					}
					pubsub.GetAwsIot().Deployment = &ttnpb.ApplicationPubSub_AWSIoTProvider_Default{
						Default: defaultIntegration,
					}
				}
			}

			res, err := ttnpb.NewApplicationPubSubRegistryClient(as).Set(ctx, &ttnpb.SetApplicationPubSubRequest{
				Pubsub:    pubsub,
				FieldMask: &pbtypes.FieldMask{Paths: paths},
			})
			if err != nil {
				return err
			}

			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	applicationsPubSubsDeleteCommand = &cobra.Command{
		Use:     "delete [application-id] [pubsub-id]",
		Aliases: []string{"del", "remove", "rm"},
		Short:   "Delete an application pub/sub",
		RunE: func(cmd *cobra.Command, args []string) error {
			pubsubID, err := getApplicationPubSubID(cmd.Flags(), args)
			if err != nil {
				return err
			}

			as, err := api.Dial(ctx, config.ApplicationServerGRPCAddress)
			if err != nil {
				return err
			}
			_, err = ttnpb.NewApplicationPubSubRegistryClient(as).Delete(ctx, pubsubID)
			if err != nil {
				return err
			}

			return nil
		},
	}
)

func init() {
	applicationsPubSubsCommand.AddCommand(applicationsPubSubsGetFormatsCommand)
	applicationsPubSubsGetCommand.Flags().AddFlagSet(applicationPubSubIDFlags())
	applicationsPubSubsGetCommand.Flags().AddFlagSet(selectApplicationPubSubFlags)
	applicationsPubSubsGetCommand.Flags().AddFlagSet(selectAllApplicationPubSubFlags)
	applicationsPubSubsCommand.AddCommand(applicationsPubSubsGetCommand)
	applicationsPubSubsListCommand.Flags().AddFlagSet(applicationIDFlags())
	applicationsPubSubsListCommand.Flags().AddFlagSet(selectApplicationPubSubFlags)
	applicationsPubSubsListCommand.Flags().AddFlagSet(selectAllApplicationPubSubFlags)
	applicationsPubSubsCommand.AddCommand(applicationsPubSubsListCommand)
	applicationsPubSubsSetCommand.Flags().AddFlagSet(applicationPubSubIDFlags())
	applicationsPubSubsSetCommand.Flags().AddFlagSet(setApplicationPubSubFlags)
	applicationsPubSubsSetCommand.Flags().AddFlagSet(applicationPubSubProviderFlags())
	applicationsPubSubsCommand.AddCommand(applicationsPubSubsSetCommand)
	applicationsPubSubsDeleteCommand.Flags().AddFlagSet(applicationPubSubIDFlags())
	applicationsPubSubsCommand.AddCommand(applicationsPubSubsDeleteCommand)
	applicationsCommand.AddCommand(applicationsPubSubsCommand)
}
