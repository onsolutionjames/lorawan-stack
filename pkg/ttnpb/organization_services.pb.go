// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/organization_services.proto

package ttnpb

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() {
	proto.RegisterFile("lorawan-stack/api/organization_services.proto", fileDescriptor_1a990e3af7846fd3)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/organization_services.proto", fileDescriptor_1a990e3af7846fd3)
}

var fileDescriptor_1a990e3af7846fd3 = []byte{
	// 774 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0x4f, 0x4f, 0x13, 0x4f,
	0x18, 0xc7, 0x7f, 0xfb, 0xfb, 0x53, 0xf2, 0x1b, 0x09, 0xc4, 0x89, 0xd1, 0x64, 0xc5, 0xc6, 0xac,
	0x06, 0x4a, 0x63, 0x67, 0x94, 0x12, 0x12, 0x08, 0x86, 0xf0, 0x2f, 0x15, 0x21, 0x4a, 0x4a, 0xbc,
	0x70, 0x21, 0xdb, 0xf2, 0xb0, 0x9d, 0xb4, 0xec, 0xac, 0x33, 0x53, 0xb0, 0x22, 0x1e, 0x3c, 0x78,
	0x35, 0xd1, 0x9b, 0x07, 0x3d, 0x1b, 0x4f, 0xbe, 0x03, 0xdf, 0x83, 0xf1, 0xe8, 0x41, 0xf4, 0xe0,
	0xd1, 0x97, 0x60, 0x76, 0x76, 0xab, 0xbb, 0x6d, 0xe9, 0x52, 0xca, 0xad, 0x3b, 0xfb, 0xcc, 0xf3,
	0xfd, 0x3c, 0xcf, 0xcc, 0xf3, 0xdd, 0xa2, 0x5c, 0x8d, 0x0b, 0x7b, 0xdf, 0x76, 0x73, 0x52, 0xd9,
	0xe5, 0x2a, 0xb5, 0x3d, 0x46, 0xb9, 0x70, 0x6c, 0x97, 0x3d, 0xb6, 0x15, 0xe3, 0xee, 0x96, 0x04,
	0xb1, 0xc7, 0xca, 0x20, 0x89, 0x27, 0xb8, 0xe2, 0x78, 0x48, 0x29, 0x97, 0x84, 0x5b, 0xc8, 0x5e,
	0xde, 0xcc, 0x39, 0x4c, 0x55, 0xea, 0x25, 0x52, 0xe6, 0xbb, 0xd4, 0xe1, 0x0e, 0xa7, 0x3a, 0xac,
	0x54, 0xdf, 0xd1, 0x4f, 0xfa, 0x41, 0xff, 0x0a, 0xb6, 0x9b, 0x23, 0x0e, 0xe7, 0x4e, 0x0d, 0xb4,
	0x8c, 0xed, 0xba, 0x5c, 0x69, 0x91, 0x30, 0xb9, 0x79, 0x39, 0x7c, 0xfb, 0x3b, 0x07, 0xec, 0x7a,
	0xaa, 0x11, 0xbe, 0xbc, 0xd6, 0x0e, 0xca, 0xb6, 0xc1, 0x55, 0x6c, 0x87, 0x81, 0x68, 0x66, 0xb8,
	0xde, 0xbd, 0x9a, 0x30, 0x2a, 0xdd, 0x1e, 0x25, 0x98, 0x53, 0x51, 0x61, 0x96, 0x89, 0xcf, 0x03,
	0xe8, 0xc2, 0xfd, 0xc8, 0xb6, 0x22, 0x38, 0x4c, 0x2a, 0xd1, 0xc0, 0x2f, 0x0d, 0x94, 0x5a, 0x14,
	0x60, 0x2b, 0xc0, 0xe3, 0x24, 0xde, 0x09, 0x12, 0xac, 0xc7, 0xb7, 0x3d, 0xac, 0x83, 0x54, 0xe6,
	0x48, 0x6b, 0x68, 0x34, 0xc8, 0x9a, 0x7b, 0xf6, 0xe9, 0xfb, 0xab, 0xbf, 0xa7, 0xad, 0x49, 0x5a,
	0x97, 0x20, 0x24, 0x3d, 0x28, 0xf3, 0x5a, 0xcd, 0x2e, 0x71, 0x61, 0x2b, 0x2e, 0x88, 0xbf, 0xb6,
	0xc5, 0xb6, 0x65, 0xf3, 0xc7, 0x61, 0xac, 0x1e, 0x39, 0x63, 0x64, 0xf1, 0x73, 0x03, 0xfd, 0x53,
	0x00, 0x85, 0x47, 0x5b, 0x65, 0x0a, 0xa0, 0x7a, 0xc7, 0x99, 0xd6, 0x38, 0x79, 0x7c, 0x2b, 0x2e,
	0x44, 0x0f, 0x62, 0xb7, 0xc2, 0x27, 0x6a, 0x59, 0x38, 0xc4, 0x6f, 0x0d, 0xf4, 0xef, 0x1a, 0x93,
	0x0a, 0x67, 0x5a, 0x15, 0xfc, 0xd5, 0xa8, 0x8a, 0x6c, 0xb2, 0x5c, 0xe9, 0xc6, 0x22, 0xad, 0x7b,
	0x1a, 0xe6, 0x0e, 0x1e, 0x8a, 0xc3, 0x6c, 0x4e, 0xe1, 0x53, 0x75, 0x0b, 0xbf, 0x30, 0x50, 0xea,
	0x81, 0xb7, 0xdd, 0xf1, 0xfc, 0x82, 0xf5, 0xde, 0x1b, 0x36, 0xab, 0x19, 0xa7, 0xcc, 0xae, 0x0d,
	0x23, 0x9d, 0x1a, 0xe6, 0x1f, 0x9e, 0x44, 0xa9, 0x25, 0xa8, 0x81, 0x02, 0x3c, 0xd6, 0x4d, 0x65,
	0xe5, 0xcf, 0x4d, 0x37, 0x2f, 0x92, 0x60, 0x4c, 0x48, 0x73, 0x4c, 0xc8, 0xb2, 0x3f, 0x26, 0x56,
	0x46, 0x83, 0x58, 0xd9, 0xab, 0x09, 0x27, 0x77, 0x88, 0x9f, 0xa0, 0x81, 0x22, 0x48, 0xc5, 0xc5,
	0x19, 0xa8, 0xde, 0xd4, 0xaa, 0x59, 0x2b, 0x93, 0xa4, 0x4a, 0x45, 0x28, 0xf9, 0x08, 0xfd, 0xb7,
	0x5e, 0x17, 0xce, 0x19, 0x68, 0x13, 0xad, 0x9d, 0xc9, 0x8e, 0x26, 0x6a, 0x7b, 0xbe, 0xe0, 0xc4,
	0x57, 0x84, 0x70, 0x54, 0x63, 0xbe, 0x5c, 0x06, 0x29, 0xf1, 0x53, 0x84, 0xfc, 0x0b, 0x5a, 0xd4,
	0x16, 0xd0, 0x0b, 0x55, 0x4b, 0x60, 0x90, 0xc0, 0xa2, 0x9a, 0x6a, 0x1c, 0x8f, 0x25, 0x77, 0x24,
	0x50, 0x7c, 0x63, 0xa0, 0xc1, 0xc0, 0x3d, 0xe6, 0xd7, 0x57, 0x56, 0xa1, 0x81, 0x69, 0xb2, 0xb7,
	0x04, 0x91, 0xcd, 0x1b, 0xda, 0x86, 0x12, 0xbc, 0xb6, 0x96, 0x35, 0xca, 0x9c, 0x35, 0xd3, 0xf3,
	0x30, 0xfb, 0xa6, 0x98, 0xab, 0x42, 0x43, 0x3b, 0xcc, 0x6b, 0x03, 0x9d, 0xf3, 0x3b, 0x14, 0x64,
	0x95, 0x98, 0x24, 0xcd, 0x77, 0x18, 0xd8, 0xc4, 0xbb, 0xd4, 0x19, 0x4f, 0x5a, 0x0b, 0x9a, 0x6f,
	0x16, 0xf7, 0xc1, 0xe7, 0x77, 0xef, 0xff, 0x02, 0x84, 0x6c, 0xf8, 0x46, 0x82, 0x09, 0x9e, 0xac,
	0x6f, 0xab, 0x9a, 0x6b, 0x19, 0x2f, 0x9e, 0x9e, 0x8b, 0x1e, 0x54, 0xa1, 0xa1, 0xa7, 0xed, 0xbd,
	0x81, 0x06, 0x03, 0x73, 0x39, 0xee, 0x78, 0xdb, 0xad, 0xe7, 0x64, 0x98, 0x45, 0x8d, 0xb9, 0x66,
	0x16, 0xfa, 0xc1, 0xb4, 0x3d, 0xb6, 0x55, 0x85, 0x06, 0x09, 0x0d, 0xe9, 0x8b, 0x81, 0x86, 0x0b,
	0xa0, 0x16, 0x23, 0xb6, 0x8a, 0x27, 0x12, 0x9a, 0x1a, 0x0d, 0x6e, 0x32, 0x8f, 0x75, 0xd8, 0x13,
	0x8f, 0x93, 0x1e, 0x77, 0x25, 0x58, 0xbb, 0xba, 0x08, 0x67, 0x13, 0x70, 0xb9, 0xf7, 0x32, 0xa2,
	0xee, 0xaf, 0xbf, 0x08, 0x49, 0x1f, 0x04, 0xfc, 0xce, 0x40, 0xc3, 0x1b, 0x49, 0xf5, 0x6d, 0x24,
	0xd7, 0x77, 0x9c, 0x27, 0xdd, 0xd5, 0xe5, 0x2c, 0x99, 0x73, 0xfd, 0x15, 0xa3, 0xe7, 0xee, 0x83,
	0x81, 0xce, 0xfb, 0xa3, 0x15, 0xd5, 0x97, 0x78, 0x32, 0x69, 0xfa, 0x62, 0xe1, 0xc7, 0x7e, 0x69,
	0x63, 0x51, 0x56, 0x41, 0x63, 0xcf, 0xe3, 0x7e, 0xb1, 0x17, 0x6e, 0xff, 0x38, 0x4a, 0xff, 0xf5,
	0xf3, 0x28, 0x6d, 0x7c, 0xfc, 0x96, 0x36, 0x36, 0xa9, 0xc3, 0x89, 0xaa, 0x80, 0xaa, 0x30, 0xd7,
	0x91, 0xc4, 0x05, 0xb5, 0xcf, 0x45, 0x95, 0xc6, 0xff, 0x7e, 0xed, 0xe5, 0xa9, 0x57, 0x75, 0xa8,
	0x52, 0xae, 0x57, 0x2a, 0xa5, 0x74, 0x3b, 0xf3, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x33, 0x73,
	0x68, 0x02, 0x97, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrganizationRegistryClient is the client API for OrganizationRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrganizationRegistryClient interface {
	// Create a new organization. This also sets the given user as
	// first collaborator with all possible rights.
	Create(ctx context.Context, in *CreateOrganizationRequest, opts ...grpc.CallOption) (*Organization, error)
	// Get the organization with the given identifiers, selecting the fields specified
	// in the field mask.
	// More or less fields may be returned, depending on the rights of the caller.
	Get(ctx context.Context, in *GetOrganizationRequest, opts ...grpc.CallOption) (*Organization, error)
	// List organizations where the given user or organization is a direct collaborator.
	// If no user or organization is given, this returns the organizations the caller
	// has access to.
	// Similar to Get, this selects the fields given by the field mask.
	// More or less fields may be returned, depending on the rights of the caller.
	List(ctx context.Context, in *ListOrganizationsRequest, opts ...grpc.CallOption) (*Organizations, error)
	// Update the organization, changing the fields specified by the field mask to the provided values.
	Update(ctx context.Context, in *UpdateOrganizationRequest, opts ...grpc.CallOption) (*Organization, error)
	// Delete the organization. This may not release the organization ID for reuse.
	Delete(ctx context.Context, in *OrganizationIdentifiers, opts ...grpc.CallOption) (*types.Empty, error)
	// Restore a recently deleted organization.
	//
	// Deployment configuration may specify if, and for how long after deletion,
	// entities can be restored.
	Restore(ctx context.Context, in *OrganizationIdentifiers, opts ...grpc.CallOption) (*types.Empty, error)
	// Purge the organization. This will release the organization ID for reuse.
	// The user is responsible for clearing data from any (external) integrations
	// that may store and expose data by user or organization ID.
	Purge(ctx context.Context, in *OrganizationIdentifiers, opts ...grpc.CallOption) (*types.Empty, error)
}

type organizationRegistryClient struct {
	cc *grpc.ClientConn
}

func NewOrganizationRegistryClient(cc *grpc.ClientConn) OrganizationRegistryClient {
	return &organizationRegistryClient{cc}
}

func (c *organizationRegistryClient) Create(ctx context.Context, in *CreateOrganizationRequest, opts ...grpc.CallOption) (*Organization, error) {
	out := new(Organization)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationRegistry/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationRegistryClient) Get(ctx context.Context, in *GetOrganizationRequest, opts ...grpc.CallOption) (*Organization, error) {
	out := new(Organization)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationRegistry/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationRegistryClient) List(ctx context.Context, in *ListOrganizationsRequest, opts ...grpc.CallOption) (*Organizations, error) {
	out := new(Organizations)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationRegistry/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationRegistryClient) Update(ctx context.Context, in *UpdateOrganizationRequest, opts ...grpc.CallOption) (*Organization, error) {
	out := new(Organization)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationRegistry/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationRegistryClient) Delete(ctx context.Context, in *OrganizationIdentifiers, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationRegistry/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationRegistryClient) Restore(ctx context.Context, in *OrganizationIdentifiers, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationRegistry/Restore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationRegistryClient) Purge(ctx context.Context, in *OrganizationIdentifiers, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationRegistry/Purge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrganizationRegistryServer is the server API for OrganizationRegistry service.
type OrganizationRegistryServer interface {
	// Create a new organization. This also sets the given user as
	// first collaborator with all possible rights.
	Create(context.Context, *CreateOrganizationRequest) (*Organization, error)
	// Get the organization with the given identifiers, selecting the fields specified
	// in the field mask.
	// More or less fields may be returned, depending on the rights of the caller.
	Get(context.Context, *GetOrganizationRequest) (*Organization, error)
	// List organizations where the given user or organization is a direct collaborator.
	// If no user or organization is given, this returns the organizations the caller
	// has access to.
	// Similar to Get, this selects the fields given by the field mask.
	// More or less fields may be returned, depending on the rights of the caller.
	List(context.Context, *ListOrganizationsRequest) (*Organizations, error)
	// Update the organization, changing the fields specified by the field mask to the provided values.
	Update(context.Context, *UpdateOrganizationRequest) (*Organization, error)
	// Delete the organization. This may not release the organization ID for reuse.
	Delete(context.Context, *OrganizationIdentifiers) (*types.Empty, error)
	// Restore a recently deleted organization.
	//
	// Deployment configuration may specify if, and for how long after deletion,
	// entities can be restored.
	Restore(context.Context, *OrganizationIdentifiers) (*types.Empty, error)
	// Purge the organization. This will release the organization ID for reuse.
	// The user is responsible for clearing data from any (external) integrations
	// that may store and expose data by user or organization ID.
	Purge(context.Context, *OrganizationIdentifiers) (*types.Empty, error)
}

// UnimplementedOrganizationRegistryServer can be embedded to have forward compatible implementations.
type UnimplementedOrganizationRegistryServer struct {
}

func (*UnimplementedOrganizationRegistryServer) Create(ctx context.Context, req *CreateOrganizationRequest) (*Organization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedOrganizationRegistryServer) Get(ctx context.Context, req *GetOrganizationRequest) (*Organization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedOrganizationRegistryServer) List(ctx context.Context, req *ListOrganizationsRequest) (*Organizations, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedOrganizationRegistryServer) Update(ctx context.Context, req *UpdateOrganizationRequest) (*Organization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedOrganizationRegistryServer) Delete(ctx context.Context, req *OrganizationIdentifiers) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedOrganizationRegistryServer) Restore(ctx context.Context, req *OrganizationIdentifiers) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (*UnimplementedOrganizationRegistryServer) Purge(ctx context.Context, req *OrganizationIdentifiers) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Purge not implemented")
}

func RegisterOrganizationRegistryServer(s *grpc.Server, srv OrganizationRegistryServer) {
	s.RegisterService(&_OrganizationRegistry_serviceDesc, srv)
}

func _OrganizationRegistry_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationRegistryServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationRegistry/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationRegistryServer).Create(ctx, req.(*CreateOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationRegistry_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationRegistryServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationRegistry/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationRegistryServer).Get(ctx, req.(*GetOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationRegistry_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrganizationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationRegistryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationRegistry/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationRegistryServer).List(ctx, req.(*ListOrganizationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationRegistry_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationRegistryServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationRegistry/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationRegistryServer).Update(ctx, req.(*UpdateOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationRegistry_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationRegistryServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationRegistry/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationRegistryServer).Delete(ctx, req.(*OrganizationIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationRegistry_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationRegistryServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationRegistry/Restore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationRegistryServer).Restore(ctx, req.(*OrganizationIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationRegistry_Purge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationRegistryServer).Purge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationRegistry/Purge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationRegistryServer).Purge(ctx, req.(*OrganizationIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrganizationRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.OrganizationRegistry",
	HandlerType: (*OrganizationRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _OrganizationRegistry_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _OrganizationRegistry_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _OrganizationRegistry_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _OrganizationRegistry_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _OrganizationRegistry_Delete_Handler,
		},
		{
			MethodName: "Restore",
			Handler:    _OrganizationRegistry_Restore_Handler,
		},
		{
			MethodName: "Purge",
			Handler:    _OrganizationRegistry_Purge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/organization_services.proto",
}

// OrganizationAccessClient is the client API for OrganizationAccess service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrganizationAccessClient interface {
	// List the rights the caller has on this organization.
	ListRights(ctx context.Context, in *OrganizationIdentifiers, opts ...grpc.CallOption) (*Rights, error)
	// Create an API key scoped to this organization.
	// Organization API keys can give access to the organization itself, as well as
	// any application, gateway and OAuth client this organization is a collaborator of.
	CreateAPIKey(ctx context.Context, in *CreateOrganizationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error)
	// List the API keys for this organization.
	ListAPIKeys(ctx context.Context, in *ListOrganizationAPIKeysRequest, opts ...grpc.CallOption) (*APIKeys, error)
	// Get a single API key of this organization.
	GetAPIKey(ctx context.Context, in *GetOrganizationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error)
	// Update the rights of an API key of the organization.
	// This method can also be used to delete the API key, by giving it no rights.
	// The caller is required to have all assigned or/and removed rights.
	UpdateAPIKey(ctx context.Context, in *UpdateOrganizationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error)
	// Get the rights of a collaborator (member) of the organization.
	// Pseudo-rights in the response (such as the "_ALL" right) are not expanded.
	GetCollaborator(ctx context.Context, in *GetOrganizationCollaboratorRequest, opts ...grpc.CallOption) (*GetCollaboratorResponse, error)
	// Set the rights of a collaborator (member) on the organization.
	// Organization collaborators can get access to the organization itself, as well as
	// any application, gateway and OAuth client this organization is a collaborator of.
	// This method can also be used to delete the collaborator, by giving them no rights.
	// The caller is required to have all assigned or/and removed rights.
	SetCollaborator(ctx context.Context, in *SetOrganizationCollaboratorRequest, opts ...grpc.CallOption) (*types.Empty, error)
	// List the collaborators on this organization.
	ListCollaborators(ctx context.Context, in *ListOrganizationCollaboratorsRequest, opts ...grpc.CallOption) (*Collaborators, error)
}

type organizationAccessClient struct {
	cc *grpc.ClientConn
}

func NewOrganizationAccessClient(cc *grpc.ClientConn) OrganizationAccessClient {
	return &organizationAccessClient{cc}
}

func (c *organizationAccessClient) ListRights(ctx context.Context, in *OrganizationIdentifiers, opts ...grpc.CallOption) (*Rights, error) {
	out := new(Rights)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationAccess/ListRights", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationAccessClient) CreateAPIKey(ctx context.Context, in *CreateOrganizationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error) {
	out := new(APIKey)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationAccess/CreateAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationAccessClient) ListAPIKeys(ctx context.Context, in *ListOrganizationAPIKeysRequest, opts ...grpc.CallOption) (*APIKeys, error) {
	out := new(APIKeys)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationAccess/ListAPIKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationAccessClient) GetAPIKey(ctx context.Context, in *GetOrganizationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error) {
	out := new(APIKey)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationAccess/GetAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationAccessClient) UpdateAPIKey(ctx context.Context, in *UpdateOrganizationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error) {
	out := new(APIKey)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationAccess/UpdateAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationAccessClient) GetCollaborator(ctx context.Context, in *GetOrganizationCollaboratorRequest, opts ...grpc.CallOption) (*GetCollaboratorResponse, error) {
	out := new(GetCollaboratorResponse)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationAccess/GetCollaborator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationAccessClient) SetCollaborator(ctx context.Context, in *SetOrganizationCollaboratorRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationAccess/SetCollaborator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationAccessClient) ListCollaborators(ctx context.Context, in *ListOrganizationCollaboratorsRequest, opts ...grpc.CallOption) (*Collaborators, error) {
	out := new(Collaborators)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationAccess/ListCollaborators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrganizationAccessServer is the server API for OrganizationAccess service.
type OrganizationAccessServer interface {
	// List the rights the caller has on this organization.
	ListRights(context.Context, *OrganizationIdentifiers) (*Rights, error)
	// Create an API key scoped to this organization.
	// Organization API keys can give access to the organization itself, as well as
	// any application, gateway and OAuth client this organization is a collaborator of.
	CreateAPIKey(context.Context, *CreateOrganizationAPIKeyRequest) (*APIKey, error)
	// List the API keys for this organization.
	ListAPIKeys(context.Context, *ListOrganizationAPIKeysRequest) (*APIKeys, error)
	// Get a single API key of this organization.
	GetAPIKey(context.Context, *GetOrganizationAPIKeyRequest) (*APIKey, error)
	// Update the rights of an API key of the organization.
	// This method can also be used to delete the API key, by giving it no rights.
	// The caller is required to have all assigned or/and removed rights.
	UpdateAPIKey(context.Context, *UpdateOrganizationAPIKeyRequest) (*APIKey, error)
	// Get the rights of a collaborator (member) of the organization.
	// Pseudo-rights in the response (such as the "_ALL" right) are not expanded.
	GetCollaborator(context.Context, *GetOrganizationCollaboratorRequest) (*GetCollaboratorResponse, error)
	// Set the rights of a collaborator (member) on the organization.
	// Organization collaborators can get access to the organization itself, as well as
	// any application, gateway and OAuth client this organization is a collaborator of.
	// This method can also be used to delete the collaborator, by giving them no rights.
	// The caller is required to have all assigned or/and removed rights.
	SetCollaborator(context.Context, *SetOrganizationCollaboratorRequest) (*types.Empty, error)
	// List the collaborators on this organization.
	ListCollaborators(context.Context, *ListOrganizationCollaboratorsRequest) (*Collaborators, error)
}

// UnimplementedOrganizationAccessServer can be embedded to have forward compatible implementations.
type UnimplementedOrganizationAccessServer struct {
}

func (*UnimplementedOrganizationAccessServer) ListRights(ctx context.Context, req *OrganizationIdentifiers) (*Rights, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRights not implemented")
}
func (*UnimplementedOrganizationAccessServer) CreateAPIKey(ctx context.Context, req *CreateOrganizationAPIKeyRequest) (*APIKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAPIKey not implemented")
}
func (*UnimplementedOrganizationAccessServer) ListAPIKeys(ctx context.Context, req *ListOrganizationAPIKeysRequest) (*APIKeys, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAPIKeys not implemented")
}
func (*UnimplementedOrganizationAccessServer) GetAPIKey(ctx context.Context, req *GetOrganizationAPIKeyRequest) (*APIKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAPIKey not implemented")
}
func (*UnimplementedOrganizationAccessServer) UpdateAPIKey(ctx context.Context, req *UpdateOrganizationAPIKeyRequest) (*APIKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAPIKey not implemented")
}
func (*UnimplementedOrganizationAccessServer) GetCollaborator(ctx context.Context, req *GetOrganizationCollaboratorRequest) (*GetCollaboratorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollaborator not implemented")
}
func (*UnimplementedOrganizationAccessServer) SetCollaborator(ctx context.Context, req *SetOrganizationCollaboratorRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCollaborator not implemented")
}
func (*UnimplementedOrganizationAccessServer) ListCollaborators(ctx context.Context, req *ListOrganizationCollaboratorsRequest) (*Collaborators, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCollaborators not implemented")
}

func RegisterOrganizationAccessServer(s *grpc.Server, srv OrganizationAccessServer) {
	s.RegisterService(&_OrganizationAccess_serviceDesc, srv)
}

func _OrganizationAccess_ListRights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationAccessServer).ListRights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationAccess/ListRights",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationAccessServer).ListRights(ctx, req.(*OrganizationIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationAccess_CreateAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrganizationAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationAccessServer).CreateAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationAccess/CreateAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationAccessServer).CreateAPIKey(ctx, req.(*CreateOrganizationAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationAccess_ListAPIKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrganizationAPIKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationAccessServer).ListAPIKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationAccess/ListAPIKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationAccessServer).ListAPIKeys(ctx, req.(*ListOrganizationAPIKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationAccess_GetAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrganizationAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationAccessServer).GetAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationAccess/GetAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationAccessServer).GetAPIKey(ctx, req.(*GetOrganizationAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationAccess_UpdateAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrganizationAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationAccessServer).UpdateAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationAccess/UpdateAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationAccessServer).UpdateAPIKey(ctx, req.(*UpdateOrganizationAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationAccess_GetCollaborator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrganizationCollaboratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationAccessServer).GetCollaborator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationAccess/GetCollaborator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationAccessServer).GetCollaborator(ctx, req.(*GetOrganizationCollaboratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationAccess_SetCollaborator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetOrganizationCollaboratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationAccessServer).SetCollaborator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationAccess/SetCollaborator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationAccessServer).SetCollaborator(ctx, req.(*SetOrganizationCollaboratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationAccess_ListCollaborators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrganizationCollaboratorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationAccessServer).ListCollaborators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationAccess/ListCollaborators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationAccessServer).ListCollaborators(ctx, req.(*ListOrganizationCollaboratorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrganizationAccess_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.OrganizationAccess",
	HandlerType: (*OrganizationAccessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRights",
			Handler:    _OrganizationAccess_ListRights_Handler,
		},
		{
			MethodName: "CreateAPIKey",
			Handler:    _OrganizationAccess_CreateAPIKey_Handler,
		},
		{
			MethodName: "ListAPIKeys",
			Handler:    _OrganizationAccess_ListAPIKeys_Handler,
		},
		{
			MethodName: "GetAPIKey",
			Handler:    _OrganizationAccess_GetAPIKey_Handler,
		},
		{
			MethodName: "UpdateAPIKey",
			Handler:    _OrganizationAccess_UpdateAPIKey_Handler,
		},
		{
			MethodName: "GetCollaborator",
			Handler:    _OrganizationAccess_GetCollaborator_Handler,
		},
		{
			MethodName: "SetCollaborator",
			Handler:    _OrganizationAccess_SetCollaborator_Handler,
		},
		{
			MethodName: "ListCollaborators",
			Handler:    _OrganizationAccess_ListCollaborators_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/organization_services.proto",
}
