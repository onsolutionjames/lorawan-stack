// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/rights.proto

package ttnpb

import (
	fmt "fmt"
	_ "github.com/TheThingsIndustries/protoc-gen-go-json/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Right is the enum that defines all the different rights to do something in the network.
type Right int32

const (
	Right_right_invalid Right = 0
	// The right to view user information.
	Right_RIGHT_USER_INFO Right = 1
	// The right to edit basic user settings.
	Right_RIGHT_USER_SETTINGS_BASIC Right = 2
	// The right to view and edit user API keys.
	Right_RIGHT_USER_SETTINGS_API_KEYS Right = 3
	// The right to delete user account.
	Right_RIGHT_USER_DELETE Right = 4
	// The right to view and edit authorized OAuth clients of the user.
	Right_RIGHT_USER_AUTHORIZED_CLIENTS Right = 5
	// The right to list applications the user is a collaborator of.
	Right_RIGHT_USER_APPLICATIONS_LIST Right = 6
	// The right to create an application under the user account.
	Right_RIGHT_USER_APPLICATIONS_CREATE Right = 7
	// The right to list gateways the user is a collaborator of.
	Right_RIGHT_USER_GATEWAYS_LIST Right = 8
	// The right to create a gateway under the account of the user.
	Right_RIGHT_USER_GATEWAYS_CREATE Right = 9
	// The right to list OAuth clients the user is a collaborator of.
	Right_RIGHT_USER_CLIENTS_LIST Right = 10
	// The right to create an OAuth client under the account of the user.
	Right_RIGHT_USER_CLIENTS_CREATE Right = 11
	// The right to list organizations the user is a member of.
	Right_RIGHT_USER_ORGANIZATIONS_LIST Right = 12
	// The right to create an organization under the user account.
	Right_RIGHT_USER_ORGANIZATIONS_CREATE Right = 13
	// The pseudo-right for all (current and future) user rights.
	Right_RIGHT_USER_ALL Right = 14
	// The right to view application information.
	Right_RIGHT_APPLICATION_INFO Right = 15
	// The right to edit basic application settings.
	Right_RIGHT_APPLICATION_SETTINGS_BASIC Right = 16
	// The right to view and edit application API keys.
	Right_RIGHT_APPLICATION_SETTINGS_API_KEYS Right = 17
	// The right to view and edit application collaborators.
	Right_RIGHT_APPLICATION_SETTINGS_COLLABORATORS Right = 18
	// The right to view and edit application packages and associations.
	Right_RIGHT_APPLICATION_SETTINGS_PACKAGES Right = 56
	// The right to delete application.
	Right_RIGHT_APPLICATION_DELETE Right = 19
	// The right to view devices in application.
	Right_RIGHT_APPLICATION_DEVICES_READ Right = 20
	// The right to create devices in application.
	Right_RIGHT_APPLICATION_DEVICES_WRITE Right = 21
	// The right to view device keys in application.
	// Note that keys may not be stored in a way that supports viewing them.
	Right_RIGHT_APPLICATION_DEVICES_READ_KEYS Right = 22
	// The right to edit device keys in application.
	Right_RIGHT_APPLICATION_DEVICES_WRITE_KEYS Right = 23
	// The right to read application traffic (uplink and downlink).
	Right_RIGHT_APPLICATION_TRAFFIC_READ Right = 24
	// The right to write uplink application traffic.
	Right_RIGHT_APPLICATION_TRAFFIC_UP_WRITE Right = 25
	// The right to write downlink application traffic.
	Right_RIGHT_APPLICATION_TRAFFIC_DOWN_WRITE Right = 26
	// The right to link as Application to a Network Server for traffic exchange,
	// i.e. read uplink and write downlink (API keys only).
	// This right is typically only given to an Application Server.
	// This right implies RIGHT_APPLICATION_INFO, RIGHT_APPLICATION_TRAFFIC_READ,
	// and RIGHT_APPLICATION_TRAFFIC_DOWN_WRITE.
	Right_RIGHT_APPLICATION_LINK Right = 27
	// The pseudo-right for all (current and future) application rights.
	Right_RIGHT_APPLICATION_ALL Right = 28
	// The pseudo-right for all (current and future) OAuth client rights.
	Right_RIGHT_CLIENT_ALL Right = 29
	// The right to view gateway information.
	Right_RIGHT_GATEWAY_INFO Right = 30
	// The right to edit basic gateway settings.
	Right_RIGHT_GATEWAY_SETTINGS_BASIC Right = 31
	// The right to view and edit gateway API keys.
	Right_RIGHT_GATEWAY_SETTINGS_API_KEYS Right = 32
	// The right to view and edit gateway collaborators.
	Right_RIGHT_GATEWAY_SETTINGS_COLLABORATORS Right = 33
	// The right to delete gateway.
	Right_RIGHT_GATEWAY_DELETE Right = 34
	// The right to read gateway traffic.
	Right_RIGHT_GATEWAY_TRAFFIC_READ Right = 35
	// The right to write downlink gateway traffic.
	Right_RIGHT_GATEWAY_TRAFFIC_DOWN_WRITE Right = 36
	// The right to link as Gateway to a Gateway Server for traffic exchange,
	// i.e. write uplink and read downlink (API keys only)
	// This right is typically only given to a gateway.
	// This right implies RIGHT_GATEWAY_INFO.
	Right_RIGHT_GATEWAY_LINK Right = 37
	// The right to view gateway status.
	Right_RIGHT_GATEWAY_STATUS_READ Right = 38
	// The right to view view gateway location.
	Right_RIGHT_GATEWAY_LOCATION_READ Right = 39
	// The right to store secrets associated with this gateway.
	Right_RIGHT_GATEWAY_WRITE_SECRETS Right = 57
	// The right to retrieve secrets associated with this gateway.
	Right_RIGHT_GATEWAY_READ_SECRETS Right = 58
	// The pseudo-right for all (current and future) gateway rights.
	Right_RIGHT_GATEWAY_ALL Right = 40
	// The right to view organization information.
	Right_RIGHT_ORGANIZATION_INFO Right = 41
	// The right to edit basic organization settings.
	Right_RIGHT_ORGANIZATION_SETTINGS_BASIC Right = 42
	// The right to view and edit organization API keys.
	Right_RIGHT_ORGANIZATION_SETTINGS_API_KEYS Right = 43
	// The right to view and edit organization members.
	Right_RIGHT_ORGANIZATION_SETTINGS_MEMBERS Right = 44
	// The right to delete organization.
	Right_RIGHT_ORGANIZATION_DELETE Right = 45
	// The right to list the applications the organization is a collaborator of.
	Right_RIGHT_ORGANIZATION_APPLICATIONS_LIST Right = 46
	// The right to create an application under the organization.
	Right_RIGHT_ORGANIZATION_APPLICATIONS_CREATE Right = 47
	// The right to list the gateways the organization is a collaborator of.
	Right_RIGHT_ORGANIZATION_GATEWAYS_LIST Right = 48
	// The right to create a gateway under the organization.
	Right_RIGHT_ORGANIZATION_GATEWAYS_CREATE Right = 49
	// The right to list the OAuth clients the organization is a collaborator of.
	Right_RIGHT_ORGANIZATION_CLIENTS_LIST Right = 50
	// The right to create an OAuth client under the organization.
	Right_RIGHT_ORGANIZATION_CLIENTS_CREATE Right = 51
	// The right to add the organization as a collaborator on an existing entity.
	Right_RIGHT_ORGANIZATION_ADD_AS_COLLABORATOR Right = 52
	// The pseudo-right for all (current and future) organization rights.
	Right_RIGHT_ORGANIZATION_ALL Right = 53
	// The right to send invites to new users.
	// Note that this is not prefixed with "USER_"; it is not a right on the user entity.
	Right_RIGHT_SEND_INVITES Right = 54
	// The pseudo-right for all (current and future) possible rights.
	Right_RIGHT_ALL Right = 55
)

var Right_name = map[int32]string{
	0:  "right_invalid",
	1:  "RIGHT_USER_INFO",
	2:  "RIGHT_USER_SETTINGS_BASIC",
	3:  "RIGHT_USER_SETTINGS_API_KEYS",
	4:  "RIGHT_USER_DELETE",
	5:  "RIGHT_USER_AUTHORIZED_CLIENTS",
	6:  "RIGHT_USER_APPLICATIONS_LIST",
	7:  "RIGHT_USER_APPLICATIONS_CREATE",
	8:  "RIGHT_USER_GATEWAYS_LIST",
	9:  "RIGHT_USER_GATEWAYS_CREATE",
	10: "RIGHT_USER_CLIENTS_LIST",
	11: "RIGHT_USER_CLIENTS_CREATE",
	12: "RIGHT_USER_ORGANIZATIONS_LIST",
	13: "RIGHT_USER_ORGANIZATIONS_CREATE",
	14: "RIGHT_USER_ALL",
	15: "RIGHT_APPLICATION_INFO",
	16: "RIGHT_APPLICATION_SETTINGS_BASIC",
	17: "RIGHT_APPLICATION_SETTINGS_API_KEYS",
	18: "RIGHT_APPLICATION_SETTINGS_COLLABORATORS",
	56: "RIGHT_APPLICATION_SETTINGS_PACKAGES",
	19: "RIGHT_APPLICATION_DELETE",
	20: "RIGHT_APPLICATION_DEVICES_READ",
	21: "RIGHT_APPLICATION_DEVICES_WRITE",
	22: "RIGHT_APPLICATION_DEVICES_READ_KEYS",
	23: "RIGHT_APPLICATION_DEVICES_WRITE_KEYS",
	24: "RIGHT_APPLICATION_TRAFFIC_READ",
	25: "RIGHT_APPLICATION_TRAFFIC_UP_WRITE",
	26: "RIGHT_APPLICATION_TRAFFIC_DOWN_WRITE",
	27: "RIGHT_APPLICATION_LINK",
	28: "RIGHT_APPLICATION_ALL",
	29: "RIGHT_CLIENT_ALL",
	30: "RIGHT_GATEWAY_INFO",
	31: "RIGHT_GATEWAY_SETTINGS_BASIC",
	32: "RIGHT_GATEWAY_SETTINGS_API_KEYS",
	33: "RIGHT_GATEWAY_SETTINGS_COLLABORATORS",
	34: "RIGHT_GATEWAY_DELETE",
	35: "RIGHT_GATEWAY_TRAFFIC_READ",
	36: "RIGHT_GATEWAY_TRAFFIC_DOWN_WRITE",
	37: "RIGHT_GATEWAY_LINK",
	38: "RIGHT_GATEWAY_STATUS_READ",
	39: "RIGHT_GATEWAY_LOCATION_READ",
	57: "RIGHT_GATEWAY_WRITE_SECRETS",
	58: "RIGHT_GATEWAY_READ_SECRETS",
	40: "RIGHT_GATEWAY_ALL",
	41: "RIGHT_ORGANIZATION_INFO",
	42: "RIGHT_ORGANIZATION_SETTINGS_BASIC",
	43: "RIGHT_ORGANIZATION_SETTINGS_API_KEYS",
	44: "RIGHT_ORGANIZATION_SETTINGS_MEMBERS",
	45: "RIGHT_ORGANIZATION_DELETE",
	46: "RIGHT_ORGANIZATION_APPLICATIONS_LIST",
	47: "RIGHT_ORGANIZATION_APPLICATIONS_CREATE",
	48: "RIGHT_ORGANIZATION_GATEWAYS_LIST",
	49: "RIGHT_ORGANIZATION_GATEWAYS_CREATE",
	50: "RIGHT_ORGANIZATION_CLIENTS_LIST",
	51: "RIGHT_ORGANIZATION_CLIENTS_CREATE",
	52: "RIGHT_ORGANIZATION_ADD_AS_COLLABORATOR",
	53: "RIGHT_ORGANIZATION_ALL",
	54: "RIGHT_SEND_INVITES",
	55: "RIGHT_ALL",
}

var Right_value = map[string]int32{
	"right_invalid":                            0,
	"RIGHT_USER_INFO":                          1,
	"RIGHT_USER_SETTINGS_BASIC":                2,
	"RIGHT_USER_SETTINGS_API_KEYS":             3,
	"RIGHT_USER_DELETE":                        4,
	"RIGHT_USER_AUTHORIZED_CLIENTS":            5,
	"RIGHT_USER_APPLICATIONS_LIST":             6,
	"RIGHT_USER_APPLICATIONS_CREATE":           7,
	"RIGHT_USER_GATEWAYS_LIST":                 8,
	"RIGHT_USER_GATEWAYS_CREATE":               9,
	"RIGHT_USER_CLIENTS_LIST":                  10,
	"RIGHT_USER_CLIENTS_CREATE":                11,
	"RIGHT_USER_ORGANIZATIONS_LIST":            12,
	"RIGHT_USER_ORGANIZATIONS_CREATE":          13,
	"RIGHT_USER_ALL":                           14,
	"RIGHT_APPLICATION_INFO":                   15,
	"RIGHT_APPLICATION_SETTINGS_BASIC":         16,
	"RIGHT_APPLICATION_SETTINGS_API_KEYS":      17,
	"RIGHT_APPLICATION_SETTINGS_COLLABORATORS": 18,
	"RIGHT_APPLICATION_SETTINGS_PACKAGES":      56,
	"RIGHT_APPLICATION_DELETE":                 19,
	"RIGHT_APPLICATION_DEVICES_READ":           20,
	"RIGHT_APPLICATION_DEVICES_WRITE":          21,
	"RIGHT_APPLICATION_DEVICES_READ_KEYS":      22,
	"RIGHT_APPLICATION_DEVICES_WRITE_KEYS":     23,
	"RIGHT_APPLICATION_TRAFFIC_READ":           24,
	"RIGHT_APPLICATION_TRAFFIC_UP_WRITE":       25,
	"RIGHT_APPLICATION_TRAFFIC_DOWN_WRITE":     26,
	"RIGHT_APPLICATION_LINK":                   27,
	"RIGHT_APPLICATION_ALL":                    28,
	"RIGHT_CLIENT_ALL":                         29,
	"RIGHT_GATEWAY_INFO":                       30,
	"RIGHT_GATEWAY_SETTINGS_BASIC":             31,
	"RIGHT_GATEWAY_SETTINGS_API_KEYS":          32,
	"RIGHT_GATEWAY_SETTINGS_COLLABORATORS":     33,
	"RIGHT_GATEWAY_DELETE":                     34,
	"RIGHT_GATEWAY_TRAFFIC_READ":               35,
	"RIGHT_GATEWAY_TRAFFIC_DOWN_WRITE":         36,
	"RIGHT_GATEWAY_LINK":                       37,
	"RIGHT_GATEWAY_STATUS_READ":                38,
	"RIGHT_GATEWAY_LOCATION_READ":              39,
	"RIGHT_GATEWAY_WRITE_SECRETS":              57,
	"RIGHT_GATEWAY_READ_SECRETS":               58,
	"RIGHT_GATEWAY_ALL":                        40,
	"RIGHT_ORGANIZATION_INFO":                  41,
	"RIGHT_ORGANIZATION_SETTINGS_BASIC":        42,
	"RIGHT_ORGANIZATION_SETTINGS_API_KEYS":     43,
	"RIGHT_ORGANIZATION_SETTINGS_MEMBERS":      44,
	"RIGHT_ORGANIZATION_DELETE":                45,
	"RIGHT_ORGANIZATION_APPLICATIONS_LIST":     46,
	"RIGHT_ORGANIZATION_APPLICATIONS_CREATE":   47,
	"RIGHT_ORGANIZATION_GATEWAYS_LIST":         48,
	"RIGHT_ORGANIZATION_GATEWAYS_CREATE":       49,
	"RIGHT_ORGANIZATION_CLIENTS_LIST":          50,
	"RIGHT_ORGANIZATION_CLIENTS_CREATE":        51,
	"RIGHT_ORGANIZATION_ADD_AS_COLLABORATOR":   52,
	"RIGHT_ORGANIZATION_ALL":                   53,
	"RIGHT_SEND_INVITES":                       54,
	"RIGHT_ALL":                                55,
}

func (x Right) String() string {
	return proto.EnumName(Right_name, int32(x))
}

func (Right) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9bb69af2cf8904c5, []int{0}
}

type Rights struct {
	Rights               []Right  `protobuf:"varint,1,rep,packed,name=rights,proto3,enum=ttn.lorawan.v3.Right" json:"rights,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rights) Reset()         { *m = Rights{} }
func (m *Rights) String() string { return proto.CompactTextString(m) }
func (*Rights) ProtoMessage()    {}
func (*Rights) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bb69af2cf8904c5, []int{0}
}
func (m *Rights) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rights.Unmarshal(m, b)
}
func (m *Rights) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rights.Marshal(b, m, deterministic)
}
func (m *Rights) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rights.Merge(m, src)
}
func (m *Rights) XXX_Size() int {
	return xxx_messageInfo_Rights.Size(m)
}
func (m *Rights) XXX_DiscardUnknown() {
	xxx_messageInfo_Rights.DiscardUnknown(m)
}

var xxx_messageInfo_Rights proto.InternalMessageInfo

func (m *Rights) GetRights() []Right {
	if m != nil {
		return m.Rights
	}
	return nil
}

type APIKey struct {
	// Immutable and unique public identifier for the API key.
	// Generated by the Access Server.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Immutable and unique secret value of the API key.
	// Generated by the Access Server.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// User-defined (friendly) name for the API key.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Rights that are granted to this API key.
	Rights               []Right          `protobuf:"varint,4,rep,packed,name=rights,proto3,enum=ttn.lorawan.v3.Right" json:"rights,omitempty"`
	CreatedAt            *types.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *types.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	ExpiresAt            *types.Timestamp `protobuf:"bytes,7,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *APIKey) Reset()         { *m = APIKey{} }
func (m *APIKey) String() string { return proto.CompactTextString(m) }
func (*APIKey) ProtoMessage()    {}
func (*APIKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bb69af2cf8904c5, []int{1}
}
func (m *APIKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_APIKey.Unmarshal(m, b)
}
func (m *APIKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_APIKey.Marshal(b, m, deterministic)
}
func (m *APIKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_APIKey.Merge(m, src)
}
func (m *APIKey) XXX_Size() int {
	return xxx_messageInfo_APIKey.Size(m)
}
func (m *APIKey) XXX_DiscardUnknown() {
	xxx_messageInfo_APIKey.DiscardUnknown(m)
}

var xxx_messageInfo_APIKey proto.InternalMessageInfo

func (m *APIKey) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *APIKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *APIKey) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *APIKey) GetRights() []Right {
	if m != nil {
		return m.Rights
	}
	return nil
}

func (m *APIKey) GetCreatedAt() *types.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *APIKey) GetUpdatedAt() *types.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *APIKey) GetExpiresAt() *types.Timestamp {
	if m != nil {
		return m.ExpiresAt
	}
	return nil
}

type APIKeys struct {
	ApiKeys              []*APIKey `protobuf:"bytes,1,rep,name=api_keys,json=apiKeys,proto3" json:"api_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *APIKeys) Reset()         { *m = APIKeys{} }
func (m *APIKeys) String() string { return proto.CompactTextString(m) }
func (*APIKeys) ProtoMessage()    {}
func (*APIKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bb69af2cf8904c5, []int{2}
}
func (m *APIKeys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_APIKeys.Unmarshal(m, b)
}
func (m *APIKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_APIKeys.Marshal(b, m, deterministic)
}
func (m *APIKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_APIKeys.Merge(m, src)
}
func (m *APIKeys) XXX_Size() int {
	return xxx_messageInfo_APIKeys.Size(m)
}
func (m *APIKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_APIKeys.DiscardUnknown(m)
}

var xxx_messageInfo_APIKeys proto.InternalMessageInfo

func (m *APIKeys) GetApiKeys() []*APIKey {
	if m != nil {
		return m.ApiKeys
	}
	return nil
}

type Collaborator struct {
	Ids                  *OrganizationOrUserIdentifiers `protobuf:"bytes,1,opt,name=ids,proto3" json:"ids,omitempty"`
	Rights               []Right                        `protobuf:"varint,2,rep,packed,name=rights,proto3,enum=ttn.lorawan.v3.Right" json:"rights,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *Collaborator) Reset()         { *m = Collaborator{} }
func (m *Collaborator) String() string { return proto.CompactTextString(m) }
func (*Collaborator) ProtoMessage()    {}
func (*Collaborator) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bb69af2cf8904c5, []int{3}
}
func (m *Collaborator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Collaborator.Unmarshal(m, b)
}
func (m *Collaborator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Collaborator.Marshal(b, m, deterministic)
}
func (m *Collaborator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Collaborator.Merge(m, src)
}
func (m *Collaborator) XXX_Size() int {
	return xxx_messageInfo_Collaborator.Size(m)
}
func (m *Collaborator) XXX_DiscardUnknown() {
	xxx_messageInfo_Collaborator.DiscardUnknown(m)
}

var xxx_messageInfo_Collaborator proto.InternalMessageInfo

func (m *Collaborator) GetIds() *OrganizationOrUserIdentifiers {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *Collaborator) GetRights() []Right {
	if m != nil {
		return m.Rights
	}
	return nil
}

type GetCollaboratorResponse struct {
	Ids                  *OrganizationOrUserIdentifiers `protobuf:"bytes,1,opt,name=ids,proto3" json:"ids,omitempty"`
	Rights               []Right                        `protobuf:"varint,2,rep,packed,name=rights,proto3,enum=ttn.lorawan.v3.Right" json:"rights,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *GetCollaboratorResponse) Reset()         { *m = GetCollaboratorResponse{} }
func (m *GetCollaboratorResponse) String() string { return proto.CompactTextString(m) }
func (*GetCollaboratorResponse) ProtoMessage()    {}
func (*GetCollaboratorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bb69af2cf8904c5, []int{4}
}
func (m *GetCollaboratorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCollaboratorResponse.Unmarshal(m, b)
}
func (m *GetCollaboratorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCollaboratorResponse.Marshal(b, m, deterministic)
}
func (m *GetCollaboratorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCollaboratorResponse.Merge(m, src)
}
func (m *GetCollaboratorResponse) XXX_Size() int {
	return xxx_messageInfo_GetCollaboratorResponse.Size(m)
}
func (m *GetCollaboratorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCollaboratorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCollaboratorResponse proto.InternalMessageInfo

func (m *GetCollaboratorResponse) GetIds() *OrganizationOrUserIdentifiers {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *GetCollaboratorResponse) GetRights() []Right {
	if m != nil {
		return m.Rights
	}
	return nil
}

type Collaborators struct {
	Collaborators        []*Collaborator `protobuf:"bytes,1,rep,name=collaborators,proto3" json:"collaborators,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Collaborators) Reset()         { *m = Collaborators{} }
func (m *Collaborators) String() string { return proto.CompactTextString(m) }
func (*Collaborators) ProtoMessage()    {}
func (*Collaborators) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bb69af2cf8904c5, []int{5}
}
func (m *Collaborators) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Collaborators.Unmarshal(m, b)
}
func (m *Collaborators) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Collaborators.Marshal(b, m, deterministic)
}
func (m *Collaborators) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Collaborators.Merge(m, src)
}
func (m *Collaborators) XXX_Size() int {
	return xxx_messageInfo_Collaborators.Size(m)
}
func (m *Collaborators) XXX_DiscardUnknown() {
	xxx_messageInfo_Collaborators.DiscardUnknown(m)
}

var xxx_messageInfo_Collaborators proto.InternalMessageInfo

func (m *Collaborators) GetCollaborators() []*Collaborator {
	if m != nil {
		return m.Collaborators
	}
	return nil
}

func init() {
	proto.RegisterEnum("ttn.lorawan.v3.Right", Right_name, Right_value)
	golang_proto.RegisterEnum("ttn.lorawan.v3.Right", Right_name, Right_value)
	proto.RegisterType((*Rights)(nil), "ttn.lorawan.v3.Rights")
	golang_proto.RegisterType((*Rights)(nil), "ttn.lorawan.v3.Rights")
	proto.RegisterType((*APIKey)(nil), "ttn.lorawan.v3.APIKey")
	golang_proto.RegisterType((*APIKey)(nil), "ttn.lorawan.v3.APIKey")
	proto.RegisterType((*APIKeys)(nil), "ttn.lorawan.v3.APIKeys")
	golang_proto.RegisterType((*APIKeys)(nil), "ttn.lorawan.v3.APIKeys")
	proto.RegisterType((*Collaborator)(nil), "ttn.lorawan.v3.Collaborator")
	golang_proto.RegisterType((*Collaborator)(nil), "ttn.lorawan.v3.Collaborator")
	proto.RegisterType((*GetCollaboratorResponse)(nil), "ttn.lorawan.v3.GetCollaboratorResponse")
	golang_proto.RegisterType((*GetCollaboratorResponse)(nil), "ttn.lorawan.v3.GetCollaboratorResponse")
	proto.RegisterType((*Collaborators)(nil), "ttn.lorawan.v3.Collaborators")
	golang_proto.RegisterType((*Collaborators)(nil), "ttn.lorawan.v3.Collaborators")
}

func init() { proto.RegisterFile("lorawan-stack/api/rights.proto", fileDescriptor_9bb69af2cf8904c5) }
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/rights.proto", fileDescriptor_9bb69af2cf8904c5)
}

var fileDescriptor_9bb69af2cf8904c5 = []byte{
	// 1221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xcd, 0x52, 0xdb, 0x56,
	0x14, 0xae, 0x8c, 0xff, 0x38, 0xc4, 0xe4, 0xe6, 0x26, 0x10, 0x63, 0x08, 0x38, 0x26, 0x3f, 0x2e,
	0x8d, 0xe5, 0xc6, 0x34, 0x6d, 0xd3, 0xc9, 0x4c, 0x2b, 0xcb, 0xc2, 0x11, 0x28, 0xb6, 0x47, 0x12,
	0xc9, 0x24, 0x1b, 0x8d, 0xc0, 0x8a, 0x50, 0x01, 0xc9, 0x23, 0x5d, 0x48, 0xe8, 0xb2, 0xcb, 0x2e,
	0xfb, 0x06, 0xdd, 0x76, 0xd3, 0x99, 0x3e, 0x41, 0xa6, 0xaf, 0xd2, 0x5d, 0x1f, 0x81, 0x55, 0xc7,
	0xd2, 0x35, 0xfa, 0xb1, 0x1d, 0x37, 0xdd, 0xd9, 0xe7, 0x7c, 0xe7, 0xdc, 0x73, 0xbe, 0xef, 0xbb,
	0x77, 0x04, 0xeb, 0x27, 0x8e, 0xab, 0xbf, 0xd3, 0xed, 0x9a, 0x47, 0xf4, 0xc3, 0xe3, 0xba, 0x3e,
	0xb0, 0xea, 0xae, 0x65, 0x1e, 0x11, 0x8f, 0x1d, 0xb8, 0x0e, 0x71, 0xf0, 0x22, 0x21, 0x36, 0x4b,
	0x31, 0xec, 0xf9, 0x76, 0x89, 0x33, 0x2d, 0x72, 0x74, 0x76, 0xc0, 0x1e, 0x3a, 0xa7, 0x75, 0xc3,
	0x3e, 0x77, 0x2e, 0x06, 0xae, 0xf3, 0xfe, 0xa2, 0xee, 0x83, 0x0f, 0x6b, 0xa6, 0x61, 0xd7, 0xce,
	0xf5, 0x13, 0xab, 0xaf, 0x13, 0xa3, 0x3e, 0xf6, 0x23, 0x68, 0x59, 0xaa, 0x45, 0x5a, 0x98, 0x8e,
	0xe9, 0x04, 0xc5, 0x07, 0x67, 0x6f, 0xfd, 0x7f, 0xfe, 0x1f, 0xff, 0x17, 0x85, 0xf3, 0x11, 0xb8,
	0x7a, 0x64, 0xa8, 0x47, 0x96, 0x6d, 0x7a, 0xa2, 0xdd, 0x3f, 0xf3, 0x88, 0x6b, 0x19, 0x5e, 0xf4,
	0x68, 0xd3, 0xa9, 0xfd, 0xe8, 0x39, 0x76, 0x5d, 0xb7, 0x6d, 0x87, 0xe8, 0xc4, 0x72, 0x6c, 0xba,
	0x46, 0x69, 0xc3, 0x74, 0x1c, 0xf3, 0xc4, 0x08, 0x8f, 0x22, 0xd6, 0xa9, 0xe1, 0x11, 0xfd, 0x74,
	0x40, 0x01, 0x9b, 0xe3, 0x3c, 0x58, 0x7d, 0xc3, 0x26, 0xd6, 0x5b, 0xcb, 0x70, 0x69, 0x97, 0xca,
	0x0e, 0x64, 0x65, 0x9f, 0x1c, 0xfc, 0x0c, 0xb2, 0x01, 0x4d, 0x45, 0xa6, 0x3c, 0x57, 0x5d, 0x6c,
	0x2c, 0xb1, 0x71, 0x9e, 0x58, 0x1f, 0xd7, 0x2c, 0x5c, 0x36, 0xe1, 0x57, 0x26, 0x57, 0xc9, 0xfc,
	0xcc, 0xa4, 0x10, 0x23, 0xd3, 0x9a, 0xca, 0x87, 0x14, 0x64, 0xb9, 0x9e, 0xb8, 0x67, 0x5c, 0xe0,
	0x45, 0x48, 0x59, 0xfd, 0x22, 0x53, 0x66, 0xaa, 0xf3, 0x72, 0xca, 0xea, 0x63, 0x04, 0x73, 0xc7,
	0xc6, 0x45, 0x31, 0xe5, 0x07, 0x86, 0x3f, 0xf1, 0x2a, 0xa4, 0x6d, 0xfd, 0xd4, 0x28, 0xce, 0x0d,
	0x43, 0xcd, 0xdc, 0x65, 0x33, 0xed, 0xa6, 0x8a, 0x0d, 0xd9, 0x0f, 0x46, 0xe6, 0x48, 0x7f, 0xfa,
	0x1c, 0xf8, 0x29, 0xc0, 0xa1, 0x6b, 0xe8, 0xc4, 0xe8, 0x6b, 0x3a, 0x29, 0x66, 0xca, 0x4c, 0x75,
	0xa1, 0x51, 0x62, 0x03, 0xaa, 0xd8, 0x11, 0x55, 0xac, 0x3a, 0xa2, 0x4a, 0x9e, 0xa7, 0x68, 0x8e,
	0x0c, 0x4b, 0xcf, 0x06, 0xfd, 0x51, 0x69, 0x76, 0x76, 0x29, 0x45, 0x73, 0x04, 0xf3, 0x00, 0xc6,
	0xfb, 0x81, 0xe5, 0x1a, 0xde, 0xb0, 0x34, 0x37, 0xab, 0xb4, 0x99, 0xbf, 0x6c, 0x66, 0xfe, 0x64,
	0x52, 0x3f, 0x30, 0xf2, 0x3c, 0xad, 0xe3, 0x48, 0xe5, 0x19, 0xe4, 0x02, 0x06, 0x3d, 0xfc, 0x18,
	0xf2, 0xfa, 0xc0, 0xd2, 0x8e, 0x8d, 0x8b, 0x40, 0x8d, 0x85, 0xc6, 0x72, 0x92, 0x85, 0x00, 0x2a,
	0xe7, 0xf4, 0x81, 0x35, 0x2c, 0xa9, 0xfc, 0xc1, 0xc0, 0x35, 0xde, 0x39, 0x39, 0xd1, 0x0f, 0x1c,
	0x57, 0x27, 0x8e, 0x8b, 0x45, 0x98, 0xb3, 0xfa, 0x9e, 0xaf, 0xc3, 0x42, 0xa3, 0x96, 0x2c, 0xef,
	0xba, 0xa6, 0x6e, 0x5b, 0x3f, 0xf9, 0x8e, 0xea, 0xba, 0xfb, 0x9e, 0xe1, 0x8a, 0xa1, 0x37, 0xfc,
	0xf9, 0x7e, 0xf1, 0x79, 0x1d, 0xf6, 0x88, 0x48, 0x92, 0xfa, 0x74, 0x49, 0x76, 0xd3, 0xf9, 0x39,
	0x94, 0xde, 0x4d, 0xe7, 0xd3, 0x28, 0xb3, 0x9b, 0xce, 0x67, 0x50, 0x76, 0x37, 0x9d, 0xcf, 0xa2,
	0x5c, 0xe5, 0x37, 0x06, 0x6e, 0xb7, 0x0d, 0x12, 0x1d, 0x5a, 0x36, 0xbc, 0x81, 0x63, 0x7b, 0x06,
	0xfe, 0xfe, 0xff, 0x0f, 0x1f, 0x8c, 0x5c, 0xfb, 0x4f, 0x23, 0xcf, 0x9c, 0x51, 0x81, 0x42, 0x74,
	0x3e, 0x0f, 0x37, 0xa1, 0x70, 0x18, 0x0d, 0x50, 0x79, 0xd6, 0x92, 0xed, 0x63, 0x5b, 0xc5, 0x4b,
	0xb6, 0xfe, 0xba, 0x0e, 0x19, 0xff, 0x78, 0x7c, 0x03, 0x0a, 0xfe, 0x00, 0x9a, 0x65, 0xfb, 0x2f,
	0x0a, 0xfa, 0x0c, 0xdf, 0x84, 0xeb, 0xb2, 0xd8, 0x7e, 0xae, 0x6a, 0xfb, 0x8a, 0x20, 0x6b, 0x62,
	0x67, 0xa7, 0x8b, 0x18, 0x7c, 0x07, 0x56, 0x22, 0x41, 0x45, 0x50, 0x55, 0xb1, 0xd3, 0x56, 0xb4,
	0x26, 0xa7, 0x88, 0x3c, 0x4a, 0xe1, 0x32, 0xac, 0x4d, 0x4a, 0x73, 0x3d, 0x51, 0xdb, 0x13, 0x5e,
	0x2b, 0x68, 0x0e, 0x2f, 0xc1, 0x8d, 0x08, 0xa2, 0x25, 0x48, 0x82, 0x2a, 0xa0, 0x34, 0xbe, 0x0b,
	0x77, 0x22, 0x61, 0x6e, 0x5f, 0x7d, 0xde, 0x95, 0xc5, 0x37, 0x42, 0x4b, 0xe3, 0x25, 0x51, 0xe8,
	0xa8, 0x0a, 0xca, 0x24, 0x7a, 0x73, 0xbd, 0x9e, 0x24, 0xf2, 0x9c, 0x2a, 0x76, 0x3b, 0x8a, 0x26,
	0x89, 0x8a, 0x8a, 0xb2, 0xb8, 0x02, 0xeb, 0xd3, 0x10, 0xbc, 0x2c, 0x70, 0xaa, 0x80, 0x72, 0x78,
	0x0d, 0x8a, 0x11, 0x4c, 0x9b, 0x53, 0x85, 0x57, 0xdc, 0x6b, 0xda, 0x21, 0x8f, 0xd7, 0xa1, 0x34,
	0x29, 0x4b, 0xab, 0xe7, 0xf1, 0x2a, 0xdc, 0x8e, 0xe4, 0xe9, 0x6c, 0x41, 0x31, 0x24, 0xb8, 0x19,
	0x25, 0x69, 0xed, 0x42, 0x62, 0xc5, 0xae, 0xdc, 0xe6, 0x3a, 0xe2, 0x9b, 0xe8, 0x02, 0xd7, 0xf0,
	0x26, 0x6c, 0x4c, 0x85, 0xd0, 0x3e, 0x05, 0x8c, 0x61, 0x31, 0xba, 0xa5, 0x24, 0xa1, 0x45, 0x5c,
	0x82, 0xe5, 0x20, 0x16, 0x59, 0x3a, 0x90, 0xec, 0x3a, 0xbe, 0x07, 0xe5, 0xf1, 0x5c, 0x42, 0x39,
	0x84, 0x1f, 0xc2, 0xe6, 0x47, 0x50, 0x57, 0x02, 0xde, 0xc0, 0x8f, 0xa0, 0xfa, 0x11, 0x20, 0xdf,
	0x95, 0x24, 0xae, 0xd9, 0x95, 0x39, 0xb5, 0x2b, 0x2b, 0x08, 0xcf, 0x68, 0xdb, 0xe3, 0xf8, 0x3d,
	0xae, 0x2d, 0x28, 0xe8, 0xdb, 0x50, 0x97, 0x28, 0x90, 0xda, 0xe3, 0x66, 0xa8, 0x6c, 0x3c, 0xfb,
	0x52, 0xe4, 0x05, 0x45, 0x93, 0x05, 0xae, 0x85, 0x6e, 0x85, 0xe4, 0x4d, 0xc2, 0xbc, 0x92, 0x45,
	0x55, 0x40, 0x4b, 0x93, 0xe7, 0x89, 0x36, 0x0a, 0xd6, 0x5c, 0xc6, 0x55, 0xb8, 0x37, 0xa3, 0x5b,
	0x80, 0xbc, 0x3d, 0x79, 0x36, 0x55, 0xe6, 0x76, 0x76, 0x44, 0x3e, 0x98, 0xad, 0x88, 0x1f, 0x40,
	0x65, 0x3a, 0x66, 0xbf, 0x47, 0xc7, 0x5b, 0x99, 0x7c, 0xea, 0x08, 0xd7, 0xea, 0xbe, 0xea, 0x50,
	0x64, 0x69, 0xb2, 0xe2, 0x92, 0xd8, 0xd9, 0x43, 0xab, 0x78, 0x05, 0x96, 0xc6, 0x73, 0x43, 0xa3,
	0xac, 0xe1, 0x5b, 0x80, 0x82, 0x54, 0x60, 0x4f, 0x3f, 0x7a, 0x07, 0x2f, 0x03, 0x0e, 0xa2, 0xd4,
	0xf1, 0x81, 0x75, 0xd6, 0xc3, 0x2b, 0x37, 0x8a, 0x27, 0x6c, 0xb3, 0x11, 0x92, 0x3e, 0x86, 0xb8,
	0xb2, 0x4c, 0x39, 0xdc, 0x6a, 0x0c, 0x14, 0xb7, 0xcb, 0x5d, 0x5c, 0x84, 0x5b, 0x71, 0x24, 0x75,
	0x40, 0x25, 0xbc, 0x99, 0xa3, 0x4c, 0x8c, 0xe1, 0xcd, 0xd0, 0xe5, 0xc9, 0x7c, 0x84, 0xb5, 0x7b,
	0xe3, 0x8b, 0xfa, 0x8c, 0xdd, 0x0f, 0xaf, 0xee, 0xd5, 0x84, 0x2a, 0xa7, 0xee, 0x53, 0x6b, 0x3d,
	0xc0, 0x1b, 0xb0, 0x9a, 0x28, 0xeb, 0x52, 0x56, 0x7d, 0xc0, 0xc3, 0x71, 0x40, 0xe0, 0x10, 0x45,
	0xe0, 0x65, 0x41, 0x55, 0xd0, 0xd3, 0xf1, 0xf1, 0x7d, 0xaf, 0x8d, 0xf2, 0xdf, 0x85, 0xcf, 0xe2,
	0x28, 0x3f, 0x14, 0xa6, 0x1a, 0xbe, 0x37, 0xd1, 0xb7, 0x20, 0x50, 0xe7, 0x73, 0x7c, 0x1f, 0xee,
	0x4e, 0x48, 0x26, 0x24, 0xda, 0x0a, 0xd9, 0x9f, 0x0c, 0xbb, 0xd2, 0xe9, 0x8b, 0xf0, 0x72, 0x4c,
	0x46, 0xbe, 0x10, 0x5e, 0x34, 0x05, 0x59, 0x41, 0x8f, 0x42, 0xba, 0x62, 0x40, 0xaa, 0x55, 0x6d,
	0xca, 0x89, 0xe3, 0x2f, 0x36, 0x8b, 0xb7, 0xe0, 0xc1, 0x2c, 0x24, 0x7d, 0xf7, 0xea, 0xa1, 0xc2,
	0x31, 0x6c, 0xfc, 0x05, 0xff, 0x32, 0xbc, 0x69, 0x93, 0x51, 0xb4, 0xdb, 0xe3, 0xd0, 0xb8, 0x31,
	0x5c, 0xec, 0x45, 0x6f, 0x4c, 0x61, 0x38, 0xf1, 0xb2, 0x6f, 0x4f, 0xdb, 0xa2, 0xd5, 0xd2, 0xb8,
	0xb8, 0xc5, 0xd1, 0x57, 0xe1, 0xbd, 0x8d, 0x63, 0x25, 0x09, 0x3d, 0x09, 0xdd, 0xa9, 0x08, 0x9d,
	0x96, 0x26, 0x76, 0x5e, 0x8a, 0xaa, 0xa0, 0xa0, 0xaf, 0x71, 0x01, 0xe6, 0xe9, 0x7d, 0x96, 0x24,
	0xf4, 0x4d, 0xa9, 0xf0, 0xcf, 0xef, 0x2b, 0xf3, 0x45, 0x66, 0x2b, 0xe3, 0x07, 0x9b, 0x4f, 0x3e,
	0xfc, 0xbd, 0xce, 0xbc, 0xa9, 0x9b, 0x0e, 0x4b, 0x8e, 0x0c, 0xe2, 0x7f, 0xc1, 0xb3, 0xb6, 0x41,
	0xde, 0x39, 0xee, 0x71, 0x3d, 0xfe, 0xe1, 0x7d, 0xbe, 0x5d, 0x1f, 0x1c, 0x9b, 0x75, 0x42, 0xec,
	0xc1, 0xc1, 0x41, 0xd6, 0xff, 0x1a, 0xdc, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x30, 0x40,
	0xb7, 0xa5, 0x0c, 0x00, 0x00,
}
