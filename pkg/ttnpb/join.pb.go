// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/join.proto

package ttnpb

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	go_thethings_network_lorawan_stack_v3_pkg_types "go.thethings.network/lorawan-stack/v3/pkg/types"
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

type JoinRequest struct {
	RawPayload         []byte                                                  `protobuf:"bytes,1,opt,name=raw_payload,json=rawPayload,proto3" json:"raw_payload,omitempty"`
	Payload            *Message                                                `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	DevAddr            go_thethings_network_lorawan_stack_v3_pkg_types.DevAddr `protobuf:"bytes,3,opt,name=dev_addr,json=devAddr,proto3,customtype=go.thethings.network/lorawan-stack/v3/pkg/types.DevAddr" json:"dev_addr"`
	SelectedMacVersion MACVersion                                              `protobuf:"varint,4,opt,name=selected_mac_version,json=selectedMacVersion,proto3,enum=ttn.lorawan.v3.MACVersion" json:"selected_mac_version,omitempty"`
	NetId              go_thethings_network_lorawan_stack_v3_pkg_types.NetID   `protobuf:"bytes,5,opt,name=net_id,json=netId,proto3,customtype=go.thethings.network/lorawan-stack/v3/pkg/types.NetID" json:"net_id"`
	DownlinkSettings   *DLSettings                                             `protobuf:"bytes,6,opt,name=downlink_settings,json=downlinkSettings,proto3" json:"downlink_settings,omitempty"`
	RxDelay            RxDelay                                                 `protobuf:"varint,7,opt,name=rx_delay,json=rxDelay,proto3,enum=ttn.lorawan.v3.RxDelay" json:"rx_delay,omitempty"`
	// Optional CFList.
	CfList         *CFList  `protobuf:"bytes,8,opt,name=cf_list,json=cfList,proto3" json:"cf_list,omitempty"`
	CorrelationIds []string `protobuf:"bytes,10,rep,name=correlation_ids,json=correlationIds,proto3" json:"correlation_ids,omitempty"`
	// Consumed airtime for the transmission of the join request. Calculated by Network Server using the RawPayload size and the transmission settings.
	ConsumedAirtime      *types.Duration `protobuf:"bytes,11,opt,name=consumed_airtime,json=consumedAirtime,proto3" json:"consumed_airtime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *JoinRequest) Reset()         { *m = JoinRequest{} }
func (m *JoinRequest) String() string { return proto.CompactTextString(m) }
func (*JoinRequest) ProtoMessage()    {}
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd69b88666e72e14, []int{0}
}
func (m *JoinRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRequest.Unmarshal(m, b)
}
func (m *JoinRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRequest.Marshal(b, m, deterministic)
}
func (m *JoinRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRequest.Merge(m, src)
}
func (m *JoinRequest) XXX_Size() int {
	return xxx_messageInfo_JoinRequest.Size(m)
}
func (m *JoinRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRequest proto.InternalMessageInfo

func (m *JoinRequest) GetRawPayload() []byte {
	if m != nil {
		return m.RawPayload
	}
	return nil
}

func (m *JoinRequest) GetPayload() *Message {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *JoinRequest) GetSelectedMacVersion() MACVersion {
	if m != nil {
		return m.SelectedMacVersion
	}
	return MAC_UNKNOWN
}

func (m *JoinRequest) GetDownlinkSettings() *DLSettings {
	if m != nil {
		return m.DownlinkSettings
	}
	return nil
}

func (m *JoinRequest) GetRxDelay() RxDelay {
	if m != nil {
		return m.RxDelay
	}
	return RX_DELAY_0
}

func (m *JoinRequest) GetCfList() *CFList {
	if m != nil {
		return m.CfList
	}
	return nil
}

func (m *JoinRequest) GetCorrelationIds() []string {
	if m != nil {
		return m.CorrelationIds
	}
	return nil
}

func (m *JoinRequest) GetConsumedAirtime() *types.Duration {
	if m != nil {
		return m.ConsumedAirtime
	}
	return nil
}

type JoinResponse struct {
	RawPayload           []byte          `protobuf:"bytes,1,opt,name=raw_payload,json=rawPayload,proto3" json:"raw_payload,omitempty"`
	SessionKeys          *SessionKeys    `protobuf:"bytes,2,opt,name=session_keys,json=sessionKeys,proto3" json:"session_keys,omitempty"`
	Lifetime             *types.Duration `protobuf:"bytes,3,opt,name=lifetime,proto3" json:"lifetime,omitempty"`
	CorrelationIds       []string        `protobuf:"bytes,4,rep,name=correlation_ids,json=correlationIds,proto3" json:"correlation_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *JoinResponse) Reset()         { *m = JoinResponse{} }
func (m *JoinResponse) String() string { return proto.CompactTextString(m) }
func (*JoinResponse) ProtoMessage()    {}
func (*JoinResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd69b88666e72e14, []int{1}
}
func (m *JoinResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinResponse.Unmarshal(m, b)
}
func (m *JoinResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinResponse.Marshal(b, m, deterministic)
}
func (m *JoinResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinResponse.Merge(m, src)
}
func (m *JoinResponse) XXX_Size() int {
	return xxx_messageInfo_JoinResponse.Size(m)
}
func (m *JoinResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JoinResponse proto.InternalMessageInfo

func (m *JoinResponse) GetRawPayload() []byte {
	if m != nil {
		return m.RawPayload
	}
	return nil
}

func (m *JoinResponse) GetSessionKeys() *SessionKeys {
	if m != nil {
		return m.SessionKeys
	}
	return nil
}

func (m *JoinResponse) GetLifetime() *types.Duration {
	if m != nil {
		return m.Lifetime
	}
	return nil
}

func (m *JoinResponse) GetCorrelationIds() []string {
	if m != nil {
		return m.CorrelationIds
	}
	return nil
}

func init() {
	proto.RegisterType((*JoinRequest)(nil), "ttn.lorawan.v3.JoinRequest")
	golang_proto.RegisterType((*JoinRequest)(nil), "ttn.lorawan.v3.JoinRequest")
	proto.RegisterType((*JoinResponse)(nil), "ttn.lorawan.v3.JoinResponse")
	golang_proto.RegisterType((*JoinResponse)(nil), "ttn.lorawan.v3.JoinResponse")
}

func init() { proto.RegisterFile("lorawan-stack/api/join.proto", fileDescriptor_dd69b88666e72e14) }
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/join.proto", fileDescriptor_dd69b88666e72e14)
}

var fileDescriptor_dd69b88666e72e14 = []byte{
	// 667 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xbd, 0x6e, 0xe3, 0x46,
	0x10, 0xc7, 0x4d, 0x59, 0x9f, 0x2b, 0xc1, 0x96, 0x89, 0x20, 0x66, 0x94, 0x40, 0x56, 0x5c, 0x09,
	0x01, 0x44, 0x22, 0x36, 0x8c, 0x34, 0x31, 0x02, 0xd1, 0x42, 0x10, 0x3b, 0x76, 0x90, 0xd0, 0x41,
	0x0a, 0x37, 0xc4, 0x8a, 0x3b, 0xa2, 0x36, 0xa2, 0x76, 0x79, 0xbb, 0x2b, 0xc9, 0x72, 0x79, 0xe5,
	0x95, 0xf7, 0x22, 0xf7, 0x0a, 0x57, 0xde, 0x33, 0x5c, 0xe1, 0xc2, 0xd7, 0x5c, 0x79, 0xb8, 0x52,
	0xd5, 0x41, 0x4b, 0xf2, 0xfc, 0x21, 0x03, 0x86, 0x2b, 0xed, 0xce, 0xfc, 0x66, 0xf4, 0x9f, 0xd9,
	0x3f, 0xd1, 0x0f, 0x11, 0x17, 0x78, 0x86, 0x59, 0x47, 0x2a, 0x1c, 0x8c, 0x1c, 0x1c, 0x53, 0xe7,
	0x7f, 0x4e, 0x99, 0x1d, 0x0b, 0xae, 0xb8, 0xb9, 0xa1, 0x14, 0xb3, 0x53, 0xc2, 0x9e, 0xee, 0x37,
	0xba, 0x21, 0x55, 0xc3, 0x49, 0xdf, 0x0e, 0xf8, 0xd8, 0x01, 0x36, 0xe5, 0xf3, 0x58, 0xf0, 0xcb,
	0xb9, 0xa3, 0xe1, 0xa0, 0x13, 0x02, 0xeb, 0x4c, 0x71, 0x44, 0x09, 0x56, 0xe0, 0xac, 0x1c, 0x92,
	0x96, 0x8d, 0xce, 0x9d, 0x16, 0x21, 0x0f, 0x79, 0x52, 0xdc, 0x9f, 0x0c, 0xf4, 0x4d, 0x5f, 0xf4,
	0x29, 0xc5, 0x9b, 0x21, 0xe7, 0x61, 0x04, 0xb7, 0x14, 0x99, 0x08, 0xac, 0x28, 0x4f, 0x15, 0x36,
	0x1e, 0xd1, 0x3f, 0x82, 0xb9, 0x4c, 0xb3, 0x3b, 0xab, 0xd9, 0x6c, 0x1a, 0x0d, 0xec, 0xbe, 0x29,
	0xa0, 0xea, 0x09, 0xa7, 0xcc, 0x83, 0x17, 0x13, 0x90, 0xca, 0x6c, 0xa3, 0xaa, 0xc0, 0x33, 0x3f,
	0xc6, 0xf3, 0x88, 0x63, 0x62, 0x19, 0x2d, 0xa3, 0x5d, 0x73, 0x4b, 0x0b, 0x37, 0x7f, 0x95, 0x1b,
	0x6e, 0x7b, 0x48, 0xe0, 0xd9, 0xdf, 0x49, 0xca, 0xfc, 0x19, 0x95, 0x32, 0x2a, 0xd7, 0x32, 0xda,
	0xd5, 0xbd, 0x6d, 0xfb, 0xfe, 0xb2, 0xec, 0x33, 0x90, 0x12, 0x87, 0xe0, 0x65, 0x9c, 0x79, 0x81,
	0xca, 0x04, 0xa6, 0x3e, 0x26, 0x44, 0x58, 0xeb, 0xba, 0xf3, 0x6f, 0xef, 0xae, 0x77, 0xd6, 0xde,
	0x5f, 0xef, 0xfc, 0x12, 0x72, 0x5b, 0x0d, 0x41, 0x0d, 0x29, 0x0b, 0xa5, 0xcd, 0x40, 0xcd, 0xb8,
	0x18, 0x39, 0xf7, 0xc5, 0x4f, 0xf7, 0x9d, 0x78, 0x14, 0x3a, 0x6a, 0x1e, 0x83, 0xb4, 0x7b, 0x30,
	0xed, 0x12, 0x22, 0xbc, 0x12, 0x49, 0x0e, 0xe6, 0x29, 0xfa, 0x46, 0x42, 0x04, 0x81, 0x02, 0xe2,
	0x8f, 0x71, 0xe0, 0x4f, 0x41, 0x48, 0xca, 0x99, 0x95, 0x6f, 0x19, 0xed, 0x8d, 0xbd, 0xc6, 0x8a,
	0xb6, 0xee, 0xd1, 0x7f, 0x09, 0xe1, 0x99, 0x59, 0xdd, 0x19, 0x0e, 0xd2, 0x98, 0xf9, 0x2f, 0x2a,
	0x32, 0x50, 0x3e, 0x25, 0x56, 0x41, 0xeb, 0x3c, 0x4c, 0x75, 0x1e, 0x3c, 0x57, 0xe7, 0x5f, 0xa0,
	0x8e, 0x7b, 0x5e, 0x81, 0x81, 0x3a, 0x26, 0xe6, 0x3f, 0x68, 0x8b, 0xf0, 0x19, 0x8b, 0x28, 0x1b,
	0xf9, 0x12, 0x94, 0x5a, 0x36, 0xb1, 0x8a, 0x7a, 0x79, 0x2b, 0x02, 0x7b, 0xa7, 0xe7, 0x29, 0xe1,
	0x96, 0x17, 0x6e, 0xe1, 0x95, 0x91, 0xab, 0x1b, 0x5e, 0x3d, 0x2b, 0xcf, 0x72, 0xe6, 0xaf, 0xa8,
	0x2c, 0x2e, 0x7d, 0x02, 0x11, 0x9e, 0x5b, 0x25, 0x3d, 0xea, 0xca, 0x33, 0x78, 0x97, 0xbd, 0x65,
	0x5a, 0xb7, 0x79, 0xa9, 0xdb, 0x94, 0x44, 0x12, 0x32, 0x1d, 0x54, 0x0a, 0x06, 0x7e, 0x44, 0xa5,
	0xb2, 0xca, 0x5a, 0xc6, 0xb7, 0x0f, 0x8b, 0x8f, 0x7e, 0x3f, 0xa5, 0x52, 0x79, 0xc5, 0x60, 0xb0,
	0xfc, 0x35, 0x0f, 0xd0, 0x66, 0xc0, 0x85, 0x80, 0x48, 0x5b, 0xd0, 0xa7, 0x44, 0x5a, 0xa8, 0xb5,
	0xde, 0xae, 0xb8, 0xb5, 0x85, 0x5b, 0x79, 0x6d, 0x14, 0x77, 0xf3, 0x22, 0x67, 0x11, 0x6f, 0xe3,
	0x0e, 0x74, 0x4c, 0xa4, 0xd9, 0x43, 0xf5, 0x80, 0x33, 0x39, 0x19, 0x03, 0xf1, 0x31, 0x15, 0x8a,
	0x8e, 0xc1, 0xaa, 0xea, 0x3f, 0xfc, 0xce, 0x4e, 0xfc, 0x6d, 0x67, 0xfe, 0xb6, 0x7b, 0xa9, 0xbf,
	0xbd, 0xcd, 0xac, 0xa4, 0x9b, 0x54, 0x9c, 0xe4, 0xcb, 0x95, 0x3a, 0xda, 0xfd, 0x6c, 0xa0, 0x5a,
	0xe2, 0x58, 0x19, 0x73, 0x26, 0xc1, 0xfc, 0xe9, 0x31, 0xcb, 0x56, 0x16, 0x6e, 0xf1, 0x2a, 0x5f,
	0xdf, 0xb2, 0x7e, 0xbc, 0x67, 0xda, 0x3f, 0x50, 0x4d, 0x82, 0x5c, 0x3e, 0xb1, 0xbf, 0xfc, 0x4a,
	0x52, 0xe7, 0x7e, 0xff, 0x70, 0xea, 0xf3, 0x84, 0xf9, 0x13, 0xe6, 0x77, 0xb7, 0x5f, 0x95, 0xb7,
	0x61, 0xf3, 0x00, 0x95, 0x23, 0x3a, 0x00, 0x3d, 0xca, 0xfa, 0x53, 0xa3, 0x7c, 0x45, 0x1f, 0x5b,
	0x60, 0xfe, 0xe9, 0x05, 0xba, 0x87, 0x1f, 0x6f, 0x9a, 0x6b, 0x9f, 0x6e, 0x9a, 0xc6, 0xdb, 0x0f,
	0x4d, 0xe3, 0xc2, 0x79, 0x86, 0x0b, 0x15, 0x8b, 0xfb, 0xfd, 0xa2, 0x96, 0xb4, 0xff, 0x25, 0x00,
	0x00, 0xff, 0xff, 0x2c, 0x6d, 0xe6, 0xf3, 0xed, 0x04, 0x00, 0x00,
}
