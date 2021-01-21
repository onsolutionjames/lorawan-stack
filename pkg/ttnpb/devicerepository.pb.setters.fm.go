// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import (
	fmt "fmt"

	types "github.com/gogo/protobuf/types"
)

func (dst *EndDeviceBrand) SetFields(src *EndDeviceBrand, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "brand_id":
			if len(subs) > 0 {
				return fmt.Errorf("'brand_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.BrandID = src.BrandID
			} else {
				var zero string
				dst.BrandID = zero
			}
		case "name":
			if len(subs) > 0 {
				return fmt.Errorf("'name' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Name = src.Name
			} else {
				var zero string
				dst.Name = zero
			}
		case "private_enterprise_number":
			if len(subs) > 0 {
				return fmt.Errorf("'private_enterprise_number' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.PrivateEnterpriseNumber = src.PrivateEnterpriseNumber
			} else {
				var zero uint32
				dst.PrivateEnterpriseNumber = zero
			}
		case "organization_unique_identifiers":
			if len(subs) > 0 {
				return fmt.Errorf("'organization_unique_identifiers' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.OrganizationUniqueIdentifiers = src.OrganizationUniqueIdentifiers
			} else {
				dst.OrganizationUniqueIdentifiers = nil
			}
		case "lora_alliance_vendor_id":
			if len(subs) > 0 {
				return fmt.Errorf("'lora_alliance_vendor_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.LoRaAllianceVendorID = src.LoRaAllianceVendorID
			} else {
				var zero uint32
				dst.LoRaAllianceVendorID = zero
			}
		case "website":
			if len(subs) > 0 {
				return fmt.Errorf("'website' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Website = src.Website
			} else {
				var zero string
				dst.Website = zero
			}
		case "email":
			if len(subs) > 0 {
				return fmt.Errorf("'email' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Email = src.Email
			} else {
				var zero string
				dst.Email = zero
			}
		case "logo":
			if len(subs) > 0 {
				return fmt.Errorf("'logo' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Logo = src.Logo
			} else {
				var zero string
				dst.Logo = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *EndDeviceModel) SetFields(src *EndDeviceModel, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "brand_id":
			if len(subs) > 0 {
				return fmt.Errorf("'brand_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.BrandID = src.BrandID
			} else {
				var zero string
				dst.BrandID = zero
			}
		case "model_id":
			if len(subs) > 0 {
				return fmt.Errorf("'model_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ModelID = src.ModelID
			} else {
				var zero string
				dst.ModelID = zero
			}
		case "name":
			if len(subs) > 0 {
				return fmt.Errorf("'name' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Name = src.Name
			} else {
				var zero string
				dst.Name = zero
			}
		case "description":
			if len(subs) > 0 {
				return fmt.Errorf("'description' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Description = src.Description
			} else {
				var zero string
				dst.Description = zero
			}
		case "hardware_versions":
			if len(subs) > 0 {
				return fmt.Errorf("'hardware_versions' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.HardwareVersions = src.HardwareVersions
			} else {
				dst.HardwareVersions = nil
			}
		case "firmware_versions":
			if len(subs) > 0 {
				return fmt.Errorf("'firmware_versions' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FirmwareVersions = src.FirmwareVersions
			} else {
				dst.FirmwareVersions = nil
			}
		case "sensors":
			if len(subs) > 0 {
				return fmt.Errorf("'sensors' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Sensors = src.Sensors
			} else {
				dst.Sensors = nil
			}
		case "dimensions":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceModel_Dimensions
				if (src == nil || src.Dimensions == nil) && dst.Dimensions == nil {
					continue
				}
				if src != nil {
					newSrc = src.Dimensions
				}
				if dst.Dimensions != nil {
					newDst = dst.Dimensions
				} else {
					newDst = &EndDeviceModel_Dimensions{}
					dst.Dimensions = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Dimensions = src.Dimensions
				} else {
					dst.Dimensions = nil
				}
			}
		case "weight":
			if len(subs) > 0 {
				return fmt.Errorf("'weight' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Weight = src.Weight
			} else {
				dst.Weight = nil
			}
		case "battery":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceModel_Battery
				if (src == nil || src.Battery == nil) && dst.Battery == nil {
					continue
				}
				if src != nil {
					newSrc = src.Battery
				}
				if dst.Battery != nil {
					newDst = dst.Battery
				} else {
					newDst = &EndDeviceModel_Battery{}
					dst.Battery = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Battery = src.Battery
				} else {
					dst.Battery = nil
				}
			}
		case "operating_conditions":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceModel_OperatingConditions
				if (src == nil || src.OperatingConditions == nil) && dst.OperatingConditions == nil {
					continue
				}
				if src != nil {
					newSrc = src.OperatingConditions
				}
				if dst.OperatingConditions != nil {
					newDst = dst.OperatingConditions
				} else {
					newDst = &EndDeviceModel_OperatingConditions{}
					dst.OperatingConditions = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.OperatingConditions = src.OperatingConditions
				} else {
					dst.OperatingConditions = nil
				}
			}
		case "ip_code":
			if len(subs) > 0 {
				return fmt.Errorf("'ip_code' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.IPCode = src.IPCode
			} else {
				var zero string
				dst.IPCode = zero
			}
		case "key_provisioning":
			if len(subs) > 0 {
				return fmt.Errorf("'key_provisioning' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.KeyProvisioning = src.KeyProvisioning
			} else {
				dst.KeyProvisioning = nil
			}
		case "key_security":
			if len(subs) > 0 {
				return fmt.Errorf("'key_security' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.KeySecurity = src.KeySecurity
			} else {
				var zero KeySecurity
				dst.KeySecurity = zero
			}
		case "photos":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceModel_Photos
				if (src == nil || src.Photos == nil) && dst.Photos == nil {
					continue
				}
				if src != nil {
					newSrc = src.Photos
				}
				if dst.Photos != nil {
					newDst = dst.Photos
				} else {
					newDst = &EndDeviceModel_Photos{}
					dst.Photos = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Photos = src.Photos
				} else {
					dst.Photos = nil
				}
			}
		case "videos":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceModel_Videos
				if (src == nil || src.Videos == nil) && dst.Videos == nil {
					continue
				}
				if src != nil {
					newSrc = src.Videos
				}
				if dst.Videos != nil {
					newDst = dst.Videos
				} else {
					newDst = &EndDeviceModel_Videos{}
					dst.Videos = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Videos = src.Videos
				} else {
					dst.Videos = nil
				}
			}
		case "product_url":
			if len(subs) > 0 {
				return fmt.Errorf("'product_url' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ProductURL = src.ProductURL
			} else {
				var zero string
				dst.ProductURL = zero
			}
		case "datasheet_url":
			if len(subs) > 0 {
				return fmt.Errorf("'datasheet_url' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DatasheetURL = src.DatasheetURL
			} else {
				var zero string
				dst.DatasheetURL = zero
			}
		case "resellers":
			if len(subs) > 0 {
				return fmt.Errorf("'resellers' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Resellers = src.Resellers
			} else {
				dst.Resellers = nil
			}
		case "compliances":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceModel_Compliances
				if (src == nil || src.Compliances == nil) && dst.Compliances == nil {
					continue
				}
				if src != nil {
					newSrc = src.Compliances
				}
				if dst.Compliances != nil {
					newDst = dst.Compliances
				} else {
					newDst = &EndDeviceModel_Compliances{}
					dst.Compliances = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Compliances = src.Compliances
				} else {
					dst.Compliances = nil
				}
			}
		case "additional_radios":
			if len(subs) > 0 {
				return fmt.Errorf("'additional_radios' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.AdditionalRadios = src.AdditionalRadios
			} else {
				dst.AdditionalRadios = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetEndDeviceBrandRequest) SetFields(src *GetEndDeviceBrandRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "application_ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationIdentifiers
				if src != nil {
					newSrc = &src.ApplicationIDs
				}
				newDst = &dst.ApplicationIDs
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationIDs = src.ApplicationIDs
				} else {
					var zero ApplicationIdentifiers
					dst.ApplicationIDs = zero
				}
			}
		case "brand_id":
			if len(subs) > 0 {
				return fmt.Errorf("'brand_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.BrandID = src.BrandID
			} else {
				var zero string
				dst.BrandID = zero
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				var zero types.FieldMask
				dst.FieldMask = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListEndDeviceBrandsRequest) SetFields(src *ListEndDeviceBrandsRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "application_ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationIdentifiers
				if src != nil {
					newSrc = &src.ApplicationIDs
				}
				newDst = &dst.ApplicationIDs
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationIDs = src.ApplicationIDs
				} else {
					var zero ApplicationIdentifiers
					dst.ApplicationIDs = zero
				}
			}
		case "limit":
			if len(subs) > 0 {
				return fmt.Errorf("'limit' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Limit = src.Limit
			} else {
				var zero uint32
				dst.Limit = zero
			}
		case "page":
			if len(subs) > 0 {
				return fmt.Errorf("'page' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Page = src.Page
			} else {
				var zero uint32
				dst.Page = zero
			}
		case "order_by":
			if len(subs) > 0 {
				return fmt.Errorf("'order_by' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.OrderBy = src.OrderBy
			} else {
				var zero string
				dst.OrderBy = zero
			}
		case "search":
			if len(subs) > 0 {
				return fmt.Errorf("'search' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Search = src.Search
			} else {
				var zero string
				dst.Search = zero
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				var zero types.FieldMask
				dst.FieldMask = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetEndDeviceModelRequest) SetFields(src *GetEndDeviceModelRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "application_ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationIdentifiers
				if src != nil {
					newSrc = &src.ApplicationIDs
				}
				newDst = &dst.ApplicationIDs
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationIDs = src.ApplicationIDs
				} else {
					var zero ApplicationIdentifiers
					dst.ApplicationIDs = zero
				}
			}
		case "brand_id":
			if len(subs) > 0 {
				return fmt.Errorf("'brand_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.BrandID = src.BrandID
			} else {
				var zero string
				dst.BrandID = zero
			}
		case "model_id":
			if len(subs) > 0 {
				return fmt.Errorf("'model_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ModelID = src.ModelID
			} else {
				var zero string
				dst.ModelID = zero
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				var zero types.FieldMask
				dst.FieldMask = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListEndDeviceModelsRequest) SetFields(src *ListEndDeviceModelsRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "application_ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationIdentifiers
				if src != nil {
					newSrc = &src.ApplicationIDs
				}
				newDst = &dst.ApplicationIDs
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationIDs = src.ApplicationIDs
				} else {
					var zero ApplicationIdentifiers
					dst.ApplicationIDs = zero
				}
			}
		case "brand_id":
			if len(subs) > 0 {
				return fmt.Errorf("'brand_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.BrandID = src.BrandID
			} else {
				var zero string
				dst.BrandID = zero
			}
		case "limit":
			if len(subs) > 0 {
				return fmt.Errorf("'limit' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Limit = src.Limit
			} else {
				var zero uint32
				dst.Limit = zero
			}
		case "page":
			if len(subs) > 0 {
				return fmt.Errorf("'page' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Page = src.Page
			} else {
				var zero uint32
				dst.Page = zero
			}
		case "order_by":
			if len(subs) > 0 {
				return fmt.Errorf("'order_by' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.OrderBy = src.OrderBy
			} else {
				var zero string
				dst.OrderBy = zero
			}
		case "search":
			if len(subs) > 0 {
				return fmt.Errorf("'search' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Search = src.Search
			} else {
				var zero string
				dst.Search = zero
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				var zero types.FieldMask
				dst.FieldMask = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetTemplateRequest) SetFields(src *GetTemplateRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "application_ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationIdentifiers
				if src != nil {
					newSrc = &src.ApplicationIDs
				}
				newDst = &dst.ApplicationIDs
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationIDs = src.ApplicationIDs
				} else {
					var zero ApplicationIdentifiers
					dst.ApplicationIDs = zero
				}
			}
		case "version_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceVersionIdentifiers
				if (src == nil || src.VersionIDs == nil) && dst.VersionIDs == nil {
					continue
				}
				if src != nil {
					newSrc = src.VersionIDs
				}
				if dst.VersionIDs != nil {
					newDst = dst.VersionIDs
				} else {
					newDst = &EndDeviceVersionIdentifiers{}
					dst.VersionIDs = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.VersionIDs = src.VersionIDs
				} else {
					dst.VersionIDs = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetPayloadFormatterRequest) SetFields(src *GetPayloadFormatterRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "application_ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationIdentifiers
				if src != nil {
					newSrc = &src.ApplicationIDs
				}
				newDst = &dst.ApplicationIDs
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationIDs = src.ApplicationIDs
				} else {
					var zero ApplicationIdentifiers
					dst.ApplicationIDs = zero
				}
			}
		case "version_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceVersionIdentifiers
				if (src == nil || src.VersionIDs == nil) && dst.VersionIDs == nil {
					continue
				}
				if src != nil {
					newSrc = src.VersionIDs
				}
				if dst.VersionIDs != nil {
					newDst = dst.VersionIDs
				} else {
					newDst = &EndDeviceVersionIdentifiers{}
					dst.VersionIDs = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.VersionIDs = src.VersionIDs
				} else {
					dst.VersionIDs = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListEndDeviceBrandsResponse) SetFields(src *ListEndDeviceBrandsResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "brands":
			if len(subs) > 0 {
				return fmt.Errorf("'brands' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Brands = src.Brands
			} else {
				dst.Brands = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListEndDeviceModelsResponse) SetFields(src *ListEndDeviceModelsResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "models":
			if len(subs) > 0 {
				return fmt.Errorf("'models' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Models = src.Models
			} else {
				dst.Models = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *MessagePayloadFormatter) SetFields(src *MessagePayloadFormatter, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "formatter":
			if len(subs) > 0 {
				return fmt.Errorf("'formatter' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Formatter = src.Formatter
			} else {
				var zero PayloadFormatter
				dst.Formatter = zero
			}
		case "formatter_parameter":
			if len(subs) > 0 {
				return fmt.Errorf("'formatter_parameter' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FormatterParameter = src.FormatterParameter
			} else {
				var zero string
				dst.FormatterParameter = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *EndDeviceModel_HardwareVersion) SetFields(src *EndDeviceModel_HardwareVersion, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "version":
			if len(subs) > 0 {
				return fmt.Errorf("'version' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Version = src.Version
			} else {
				var zero string
				dst.Version = zero
			}
		case "numeric":
			if len(subs) > 0 {
				return fmt.Errorf("'numeric' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Numeric = src.Numeric
			} else {
				var zero uint32
				dst.Numeric = zero
			}
		case "part_number":
			if len(subs) > 0 {
				return fmt.Errorf("'part_number' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.PartNumber = src.PartNumber
			} else {
				var zero string
				dst.PartNumber = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *EndDeviceModel_FirmwareVersion) SetFields(src *EndDeviceModel_FirmwareVersion, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "version":
			if len(subs) > 0 {
				return fmt.Errorf("'version' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Version = src.Version
			} else {
				var zero string
				dst.Version = zero
			}
		case "numeric":
			if len(subs) > 0 {
				return fmt.Errorf("'numeric' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Numeric = src.Numeric
			} else {
				var zero uint32
				dst.Numeric = zero
			}
		case "supported_hardware_versions":
			if len(subs) > 0 {
				return fmt.Errorf("'supported_hardware_versions' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.SupportedHardwareVersions = src.SupportedHardwareVersions
			} else {
				dst.SupportedHardwareVersions = nil
			}
		case "profiles":
			if len(subs) > 0 {
				return fmt.Errorf("'profiles' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Profiles = src.Profiles
			} else {
				dst.Profiles = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *EndDeviceModel_Dimensions) SetFields(src *EndDeviceModel_Dimensions, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "width":
			if len(subs) > 0 {
				return fmt.Errorf("'width' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Width = src.Width
			} else {
				dst.Width = nil
			}
		case "height":
			if len(subs) > 0 {
				return fmt.Errorf("'height' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Height = src.Height
			} else {
				dst.Height = nil
			}
		case "diameter":
			if len(subs) > 0 {
				return fmt.Errorf("'diameter' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Diameter = src.Diameter
			} else {
				dst.Diameter = nil
			}
		case "length":
			if len(subs) > 0 {
				return fmt.Errorf("'length' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Length = src.Length
			} else {
				dst.Length = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *EndDeviceModel_Battery) SetFields(src *EndDeviceModel_Battery, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "replaceable":
			if len(subs) > 0 {
				return fmt.Errorf("'replaceable' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Replaceable = src.Replaceable
			} else {
				dst.Replaceable = nil
			}
		case "type":
			if len(subs) > 0 {
				return fmt.Errorf("'type' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Type = src.Type
			} else {
				var zero string
				dst.Type = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *EndDeviceModel_OperatingConditions) SetFields(src *EndDeviceModel_OperatingConditions, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "temperature":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceModel_OperatingConditions_Limits
				if (src == nil || src.Temperature == nil) && dst.Temperature == nil {
					continue
				}
				if src != nil {
					newSrc = src.Temperature
				}
				if dst.Temperature != nil {
					newDst = dst.Temperature
				} else {
					newDst = &EndDeviceModel_OperatingConditions_Limits{}
					dst.Temperature = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Temperature = src.Temperature
				} else {
					dst.Temperature = nil
				}
			}
		case "relative_humidity":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceModel_OperatingConditions_Limits
				if (src == nil || src.RelativeHumidity == nil) && dst.RelativeHumidity == nil {
					continue
				}
				if src != nil {
					newSrc = src.RelativeHumidity
				}
				if dst.RelativeHumidity != nil {
					newDst = dst.RelativeHumidity
				} else {
					newDst = &EndDeviceModel_OperatingConditions_Limits{}
					dst.RelativeHumidity = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.RelativeHumidity = src.RelativeHumidity
				} else {
					dst.RelativeHumidity = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *EndDeviceModel_Photos) SetFields(src *EndDeviceModel_Photos, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "main":
			if len(subs) > 0 {
				return fmt.Errorf("'main' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Main = src.Main
			} else {
				var zero string
				dst.Main = zero
			}
		case "other":
			if len(subs) > 0 {
				return fmt.Errorf("'other' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Other = src.Other
			} else {
				dst.Other = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *EndDeviceModel_Videos) SetFields(src *EndDeviceModel_Videos, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "main":
			if len(subs) > 0 {
				return fmt.Errorf("'main' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Main = src.Main
			} else {
				var zero string
				dst.Main = zero
			}
		case "other":
			if len(subs) > 0 {
				return fmt.Errorf("'other' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Other = src.Other
			} else {
				dst.Other = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *EndDeviceModel_Reseller) SetFields(src *EndDeviceModel_Reseller, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "name":
			if len(subs) > 0 {
				return fmt.Errorf("'name' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Name = src.Name
			} else {
				var zero string
				dst.Name = zero
			}
		case "region":
			if len(subs) > 0 {
				return fmt.Errorf("'region' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Region = src.Region
			} else {
				dst.Region = nil
			}
		case "url":
			if len(subs) > 0 {
				return fmt.Errorf("'url' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.URL = src.URL
			} else {
				var zero string
				dst.URL = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *EndDeviceModel_Compliances) SetFields(src *EndDeviceModel_Compliances, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "safety":
			if len(subs) > 0 {
				return fmt.Errorf("'safety' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Safety = src.Safety
			} else {
				dst.Safety = nil
			}
		case "radio_equipment":
			if len(subs) > 0 {
				return fmt.Errorf("'radio_equipment' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.RadioEquipment = src.RadioEquipment
			} else {
				dst.RadioEquipment = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *EndDeviceModel_FirmwareVersion_Profile) SetFields(src *EndDeviceModel_FirmwareVersion_Profile, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "profile_id":
			if len(subs) > 0 {
				return fmt.Errorf("'profile_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ProfileID = src.ProfileID
			} else {
				var zero string
				dst.ProfileID = zero
			}
		case "lorawan_certified":
			if len(subs) > 0 {
				return fmt.Errorf("'lorawan_certified' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.LoRaWANCertified = src.LoRaWANCertified
			} else {
				var zero bool
				dst.LoRaWANCertified = zero
			}
		case "codec_id":
			if len(subs) > 0 {
				return fmt.Errorf("'codec_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CodecID = src.CodecID
			} else {
				var zero string
				dst.CodecID = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *EndDeviceModel_OperatingConditions_Limits) SetFields(src *EndDeviceModel_OperatingConditions_Limits, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "min":
			if len(subs) > 0 {
				return fmt.Errorf("'min' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Min = src.Min
			} else {
				dst.Min = nil
			}
		case "max":
			if len(subs) > 0 {
				return fmt.Errorf("'max' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Max = src.Max
			} else {
				dst.Max = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *EndDeviceModel_Compliances_Compliance) SetFields(src *EndDeviceModel_Compliances_Compliance, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "body":
			if len(subs) > 0 {
				return fmt.Errorf("'body' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Body = src.Body
			} else {
				var zero string
				dst.Body = zero
			}
		case "norm":
			if len(subs) > 0 {
				return fmt.Errorf("'norm' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Norm = src.Norm
			} else {
				var zero string
				dst.Norm = zero
			}
		case "standard":
			if len(subs) > 0 {
				return fmt.Errorf("'standard' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Standard = src.Standard
			} else {
				var zero string
				dst.Standard = zero
			}
		case "version":
			if len(subs) > 0 {
				return fmt.Errorf("'version' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Version = src.Version
			} else {
				var zero string
				dst.Version = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
