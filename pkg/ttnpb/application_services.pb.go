// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/application_services.proto

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
	proto.RegisterFile("lorawan-stack/api/application_services.proto", fileDescriptor_f6c42f4fe8e3c902)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/application_services.proto", fileDescriptor_f6c42f4fe8e3c902)
}

var fileDescriptor_f6c42f4fe8e3c902 = []byte{
	// 854 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x4f, 0x4f, 0x1b, 0x47,
	0x1c, 0xed, 0xd0, 0xd6, 0x15, 0x03, 0x2d, 0x62, 0x2a, 0xb5, 0xd2, 0x42, 0xad, 0x6a, 0x29, 0x06,
	0x51, 0x3c, 0xdb, 0x62, 0xb5, 0x15, 0x15, 0xfd, 0x03, 0x18, 0xb9, 0x16, 0xad, 0x8a, 0x8c, 0xb8,
	0xf8, 0x42, 0xd7, 0xf6, 0xb0, 0x1e, 0xd9, 0xec, 0x6c, 0x77, 0x66, 0x4d, 0x8d, 0xc5, 0xa1, 0x39,
	0x72, 0x8b, 0xa2, 0xe4, 0x10, 0xe5, 0x1a, 0xe5, 0x94, 0xcf, 0x90, 0x7c, 0x87, 0x1c, 0x72, 0x8d,
	0x20, 0x52, 0x72, 0xe0, 0x90, 0x8f, 0x10, 0xed, 0xec, 0x6e, 0xd8, 0xb5, 0xcd, 0xae, 0xff, 0xe4,
	0xe6, 0x9d, 0x79, 0xf3, 0x7b, 0x6f, 0xde, 0xee, 0xef, 0xfd, 0x0c, 0x57, 0x9b, 0xcc, 0xd6, 0x4f,
	0x74, 0x33, 0xcb, 0x85, 0x5e, 0x6d, 0x68, 0xba, 0x45, 0x35, 0xdd, 0xb2, 0x9a, 0xb4, 0xaa, 0x0b,
	0xca, 0xcc, 0x43, 0x4e, 0xec, 0x16, 0xad, 0x12, 0x8e, 0x2d, 0x9b, 0x09, 0x86, 0x3e, 0x13, 0xc2,
	0xc4, 0xfe, 0x09, 0xdc, 0xca, 0x29, 0x59, 0x83, 0x8a, 0xba, 0x53, 0xc1, 0x55, 0x76, 0xac, 0x19,
	0xcc, 0x60, 0x9a, 0x84, 0x55, 0x9c, 0x23, 0xf9, 0x24, 0x1f, 0xe4, 0x2f, 0xef, 0xb8, 0x32, 0x6f,
	0x30, 0x66, 0x34, 0x89, 0xc7, 0x62, 0x9a, 0x4c, 0x48, 0x12, 0xbf, 0xb8, 0x32, 0xe7, 0xef, 0xbe,
	0xab, 0x41, 0x8e, 0x2d, 0xd1, 0xf6, 0x37, 0x17, 0x62, 0x75, 0xde, 0x0c, 0xa2, 0x35, 0x62, 0x0a,
	0x7a, 0x44, 0x89, 0x1d, 0xd0, 0xa4, 0x7b, 0x41, 0x36, 0x35, 0xea, 0xc2, 0xdf, 0x5f, 0xbb, 0x9a,
	0x84, 0x9f, 0x6f, 0x5e, 0x97, 0x2e, 0x11, 0x83, 0x72, 0x61, 0xb7, 0xd1, 0x25, 0x80, 0xa9, 0x6d,
	0x9b, 0xe8, 0x82, 0xa0, 0x65, 0x1c, 0xf5, 0x01, 0x7b, 0xeb, 0x91, 0x53, 0xff, 0x3a, 0x84, 0x0b,
	0x65, 0xae, 0x1b, 0x19, 0xc2, 0xa8, 0xb7, 0xc1, 0xad, 0x67, 0x2f, 0xef, 0x4c, 0x9c, 0x03, 0x35,
	0xa7, 0x39, 0x9c, 0xd8, 0x5c, 0xeb, 0x54, 0x59, 0xb3, 0xa9, 0x57, 0x98, 0xad, 0x0b, 0x66, 0x63,
	0x77, 0xed, 0x90, 0xd6, 0x78, 0xf0, 0xe3, 0x2c, 0x7c, 0x65, 0xfe, 0x33, 0x58, 0x29, 0xef, 0xa9,
	0xbb, 0x1a, 0xb3, 0x0d, 0xdd, 0xa4, 0xa7, 0xde, 0x62, 0x57, 0x85, 0xf0, 0x9e, 0xac, 0xd4, 0xb5,
	0xd0, 0x53, 0x11, 0xfd, 0x0f, 0xe0, 0x87, 0x05, 0x22, 0xd0, 0x62, 0xb7, 0xf0, 0x02, 0x11, 0xc3,
	0xde, 0xef, 0x47, 0x79, 0xbd, 0xef, 0x10, 0x8e, 0xb0, 0x68, 0x9d, 0xf0, 0x07, 0xe6, 0x8a, 0x8a,
	0x3e, 0x9f, 0xa1, 0x2b, 0x00, 0x3f, 0xfa, 0x93, 0x72, 0x81, 0x96, 0xba, 0xab, 0xbb, 0xab, 0x21,
	0x06, 0x1e, 0xc8, 0x98, 0x8f, 0x91, 0xc1, 0xd5, 0x07, 0x9e, 0xcf, 0x77, 0x01, 0xfa, 0x34, 0xa2,
	0xa4, 0xfc, 0x03, 0x1a, 0xc5, 0xf8, 0xf2, 0x5f, 0xe8, 0x7d, 0xba, 0x8e, 0xce, 0x01, 0x4c, 0x1d,
	0x58, 0xb5, 0xbe, 0x1f, 0x96, 0xb7, 0x3e, 0xac, 0xf1, 0xeb, 0xf2, 0xbe, 0x39, 0x25, 0xc6, 0x78,
	0xdc, 0xc7, 0x78, 0xf7, 0xfd, 0x5b, 0x30, 0x95, 0x27, 0x4d, 0x22, 0x08, 0xca, 0xc4, 0x30, 0x14,
	0xaf, 0xbb, 0x4a, 0xf9, 0x02, 0x7b, 0x7d, 0x8b, 0x83, 0xbe, 0xc5, 0x3b, 0x6e, 0xdf, 0xaa, 0x19,
	0x29, 0xe2, 0xeb, 0x95, 0x74, 0xec, 0xdb, 0x3f, 0x43, 0x6d, 0xf8, 0x49, 0x89, 0x70, 0xc1, 0xec,
	0xf1, 0x29, 0xb1, 0xa4, 0x5c, 0x56, 0x33, 0xf1, 0x94, 0x9a, 0xed, 0xf3, 0x39, 0xf0, 0xe3, 0x3d,
	0xc7, 0x36, 0xc6, 0x27, 0x5e, 0x95, 0xc4, 0x99, 0x95, 0x6f, 0x12, 0x88, 0x2d, 0xc9, 0x76, 0x0e,
	0xe0, 0x54, 0x91, 0x73, 0x87, 0xe4, 0x49, 0x6b, 0xe7, 0xa0, 0x38, 0x30, 0xfb, 0x42, 0x37, 0x2e,
	0x54, 0xa4, 0x44, 0xb8, 0xc5, 0x4c, 0x4e, 0x06, 0xf6, 0xa0, 0x46, 0x5a, 0x59, 0xe2, 0xd0, 0xb5,
	0x27, 0x53, 0x70, 0x36, 0xc4, 0xb7, 0x59, 0xad, 0x12, 0xce, 0x51, 0x07, 0x42, 0xb7, 0xd7, 0x4a,
	0x32, 0x18, 0x87, 0xb0, 0xa7, 0x0b, 0xe7, 0x9d, 0x57, 0xb3, 0x52, 0xd3, 0x12, 0x5a, 0x4c, 0x7a,
	0x2f, 0x1e, 0xdd, 0x7d, 0x00, 0xa7, 0xfd, 0x44, 0xdd, 0x2b, 0xee, 0x92, 0x36, 0xc2, 0x89, 0x79,
	0xeb, 0x01, 0x83, 0xe6, 0xe8, 0xd1, 0xe1, 0x6d, 0xab, 0x5b, 0x52, 0xc7, 0x86, 0xfa, 0xd3, 0x70,
	0x81, 0xe4, 0xce, 0x88, 0x6c, 0x83, 0xb4, 0x65, 0x40, 0xde, 0x03, 0x70, 0x4a, 0xc6, 0x90, 0x2c,
	0xc9, 0x51, 0x36, 0x21, 0xa3, 0x7c, 0x5c, 0x20, 0xed, 0xcb, 0xfe, 0xd2, 0xb8, 0xfa, 0x9b, 0xd4,
	0xb6, 0x8e, 0x46, 0xd5, 0xe6, 0xba, 0x36, 0xe9, 0x86, 0xb4, 0x67, 0xd9, 0xb7, 0xf1, 0xf9, 0x3d,
	0x98, 0x5f, 0x7f, 0x48, 0x4d, 0x5b, 0xe8, 0xf7, 0x11, 0x35, 0x69, 0x9d, 0x06, 0x69, 0xcb, 0x26,
	0x7f, 0x04, 0xe0, 0xb4, 0x9f, 0x65, 0x37, 0xbc, 0xd2, 0x9e, 0xa4, 0x1b, 0x4c, 0xe2, 0xdf, 0x52,
	0x62, 0x51, 0xc9, 0x8f, 0x2c, 0x51, 0xb7, 0xe8, 0x61, 0x83, 0xb4, 0xb1, 0x1f, 0x80, 0xcf, 0x27,
	0xe0, 0x4c, 0x81, 0x88, 0xed, 0x50, 0xa0, 0xa3, 0xef, 0xe3, 0xcd, 0x0c, 0x63, 0x03, 0xbd, 0x4b,
	0x7d, 0x8e, 0x44, 0x71, 0x7e, 0xbf, 0xbe, 0xf2, 0x86, 0xd3, 0x0b, 0x50, 0xae, 0xa0, 0x7f, 0x86,
	0xbc, 0x44, 0x78, 0xea, 0xc8, 0x41, 0x96, 0x34, 0xc7, 0xca, 0xa7, 0xe8, 0xbf, 0x71, 0x38, 0xc2,
	0x83, 0x6c, 0xd8, 0xa1, 0x87, 0x1e, 0x02, 0x38, 0xb3, 0x9f, 0xe4, 0xec, 0x7e, 0xa2, 0xb3, 0x37,
	0x65, 0x70, 0x41, 0xfa, 0xb8, 0xa9, 0x6c, 0x8c, 0x71, 0x41, 0xd9, 0xe1, 0x8f, 0x01, 0x9c, 0x75,
	0x9b, 0x38, 0x4c, 0xce, 0x51, 0x2e, 0xa1, 0xcf, 0x23, 0xe8, 0x40, 0xeb, 0x57, 0x3d, 0xc1, 0x15,
	0x46, 0xa9, 0x79, 0x29, 0xf9, 0x57, 0x34, 0x96, 0xe4, 0xad, 0x5f, 0x5e, 0x5f, 0xa4, 0x3f, 0x78,
	0x73, 0x91, 0x06, 0x4f, 0x2f, 0xd3, 0xa0, 0xac, 0x19, 0x0c, 0x8b, 0x3a, 0x11, 0x75, 0x6a, 0x1a,
	0x1c, 0x9b, 0x44, 0x9c, 0x30, 0xbb, 0xa1, 0x45, 0xff, 0xf3, 0xb6, 0x72, 0x9a, 0xd5, 0x30, 0x34,
	0x21, 0x4c, 0xab, 0x52, 0x49, 0x49, 0x1f, 0x73, 0x6f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x74, 0xb1,
	0xdb, 0xc5, 0x0a, 0x0c, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ApplicationRegistryClient is the client API for ApplicationRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApplicationRegistryClient interface {
	// Create a new application. This also sets the given organization or user as
	// first collaborator with all possible rights.
	Create(ctx context.Context, in *CreateApplicationRequest, opts ...grpc.CallOption) (*Application, error)
	// Get the application with the given identifiers, selecting the fields specified
	// in the field mask.
	// More or less fields may be returned, depending on the rights of the caller.
	Get(ctx context.Context, in *GetApplicationRequest, opts ...grpc.CallOption) (*Application, error)
	// List applications where the given user or organization is a direct collaborator.
	// If no user or organization is given, this returns the applications the caller
	// has access to.
	// Similar to Get, this selects the fields given by the field mask.
	// More or less fields may be returned, depending on the rights of the caller.
	List(ctx context.Context, in *ListApplicationsRequest, opts ...grpc.CallOption) (*Applications, error)
	// Update the application, changing the fields specified by the field mask to the provided values.
	Update(ctx context.Context, in *UpdateApplicationRequest, opts ...grpc.CallOption) (*Application, error)
	// Delete the application. This may not release the application ID for reuse.
	// All end devices must be deleted from the application before it can be deleted.
	Delete(ctx context.Context, in *ApplicationIdentifiers, opts ...grpc.CallOption) (*types.Empty, error)
	// Restore a recently deleted application.
	//
	// Deployment configuration may specify if, and for how long after deletion,
	// entities can be restored.
	Restore(ctx context.Context, in *ApplicationIdentifiers, opts ...grpc.CallOption) (*types.Empty, error)
	// Purge the application. This will release the application ID for reuse.
	// All end devices must be deleted from the application before it can be deleted.
	// The application owner is responsible for clearing data from any (external) integrations
	// that may store and expose data by application ID
	Purge(ctx context.Context, in *ApplicationIdentifiers, opts ...grpc.CallOption) (*types.Empty, error)
	// Request DevEUI from the configured address block for a device inside the application.
	// The maximum number of DevEUI's issued per application can be configured.
	IssueDevEUI(ctx context.Context, in *ApplicationIdentifiers, opts ...grpc.CallOption) (*IssueDevEUIResponse, error)
}

type applicationRegistryClient struct {
	cc *grpc.ClientConn
}

func NewApplicationRegistryClient(cc *grpc.ClientConn) ApplicationRegistryClient {
	return &applicationRegistryClient{cc}
}

func (c *applicationRegistryClient) Create(ctx context.Context, in *CreateApplicationRequest, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationRegistry/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationRegistryClient) Get(ctx context.Context, in *GetApplicationRequest, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationRegistry/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationRegistryClient) List(ctx context.Context, in *ListApplicationsRequest, opts ...grpc.CallOption) (*Applications, error) {
	out := new(Applications)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationRegistry/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationRegistryClient) Update(ctx context.Context, in *UpdateApplicationRequest, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationRegistry/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationRegistryClient) Delete(ctx context.Context, in *ApplicationIdentifiers, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationRegistry/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationRegistryClient) Restore(ctx context.Context, in *ApplicationIdentifiers, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationRegistry/Restore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationRegistryClient) Purge(ctx context.Context, in *ApplicationIdentifiers, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationRegistry/Purge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationRegistryClient) IssueDevEUI(ctx context.Context, in *ApplicationIdentifiers, opts ...grpc.CallOption) (*IssueDevEUIResponse, error) {
	out := new(IssueDevEUIResponse)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationRegistry/IssueDevEUI", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplicationRegistryServer is the server API for ApplicationRegistry service.
type ApplicationRegistryServer interface {
	// Create a new application. This also sets the given organization or user as
	// first collaborator with all possible rights.
	Create(context.Context, *CreateApplicationRequest) (*Application, error)
	// Get the application with the given identifiers, selecting the fields specified
	// in the field mask.
	// More or less fields may be returned, depending on the rights of the caller.
	Get(context.Context, *GetApplicationRequest) (*Application, error)
	// List applications where the given user or organization is a direct collaborator.
	// If no user or organization is given, this returns the applications the caller
	// has access to.
	// Similar to Get, this selects the fields given by the field mask.
	// More or less fields may be returned, depending on the rights of the caller.
	List(context.Context, *ListApplicationsRequest) (*Applications, error)
	// Update the application, changing the fields specified by the field mask to the provided values.
	Update(context.Context, *UpdateApplicationRequest) (*Application, error)
	// Delete the application. This may not release the application ID for reuse.
	// All end devices must be deleted from the application before it can be deleted.
	Delete(context.Context, *ApplicationIdentifiers) (*types.Empty, error)
	// Restore a recently deleted application.
	//
	// Deployment configuration may specify if, and for how long after deletion,
	// entities can be restored.
	Restore(context.Context, *ApplicationIdentifiers) (*types.Empty, error)
	// Purge the application. This will release the application ID for reuse.
	// All end devices must be deleted from the application before it can be deleted.
	// The application owner is responsible for clearing data from any (external) integrations
	// that may store and expose data by application ID
	Purge(context.Context, *ApplicationIdentifiers) (*types.Empty, error)
	// Request DevEUI from the configured address block for a device inside the application.
	// The maximum number of DevEUI's issued per application can be configured.
	IssueDevEUI(context.Context, *ApplicationIdentifiers) (*IssueDevEUIResponse, error)
}

// UnimplementedApplicationRegistryServer can be embedded to have forward compatible implementations.
type UnimplementedApplicationRegistryServer struct {
}

func (*UnimplementedApplicationRegistryServer) Create(ctx context.Context, req *CreateApplicationRequest) (*Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedApplicationRegistryServer) Get(ctx context.Context, req *GetApplicationRequest) (*Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedApplicationRegistryServer) List(ctx context.Context, req *ListApplicationsRequest) (*Applications, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedApplicationRegistryServer) Update(ctx context.Context, req *UpdateApplicationRequest) (*Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedApplicationRegistryServer) Delete(ctx context.Context, req *ApplicationIdentifiers) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedApplicationRegistryServer) Restore(ctx context.Context, req *ApplicationIdentifiers) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (*UnimplementedApplicationRegistryServer) Purge(ctx context.Context, req *ApplicationIdentifiers) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Purge not implemented")
}
func (*UnimplementedApplicationRegistryServer) IssueDevEUI(ctx context.Context, req *ApplicationIdentifiers) (*IssueDevEUIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssueDevEUI not implemented")
}

func RegisterApplicationRegistryServer(s *grpc.Server, srv ApplicationRegistryServer) {
	s.RegisterService(&_ApplicationRegistry_serviceDesc, srv)
}

func _ApplicationRegistry_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationRegistryServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationRegistry/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationRegistryServer).Create(ctx, req.(*CreateApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationRegistry_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationRegistryServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationRegistry/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationRegistryServer).Get(ctx, req.(*GetApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationRegistry_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApplicationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationRegistryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationRegistry/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationRegistryServer).List(ctx, req.(*ListApplicationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationRegistry_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationRegistryServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationRegistry/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationRegistryServer).Update(ctx, req.(*UpdateApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationRegistry_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationRegistryServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationRegistry/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationRegistryServer).Delete(ctx, req.(*ApplicationIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationRegistry_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationRegistryServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationRegistry/Restore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationRegistryServer).Restore(ctx, req.(*ApplicationIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationRegistry_Purge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationRegistryServer).Purge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationRegistry/Purge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationRegistryServer).Purge(ctx, req.(*ApplicationIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationRegistry_IssueDevEUI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationRegistryServer).IssueDevEUI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationRegistry/IssueDevEUI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationRegistryServer).IssueDevEUI(ctx, req.(*ApplicationIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApplicationRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.ApplicationRegistry",
	HandlerType: (*ApplicationRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ApplicationRegistry_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ApplicationRegistry_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ApplicationRegistry_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ApplicationRegistry_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ApplicationRegistry_Delete_Handler,
		},
		{
			MethodName: "Restore",
			Handler:    _ApplicationRegistry_Restore_Handler,
		},
		{
			MethodName: "Purge",
			Handler:    _ApplicationRegistry_Purge_Handler,
		},
		{
			MethodName: "IssueDevEUI",
			Handler:    _ApplicationRegistry_IssueDevEUI_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/application_services.proto",
}

// ApplicationAccessClient is the client API for ApplicationAccess service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApplicationAccessClient interface {
	// List the rights the caller has on this application.
	ListRights(ctx context.Context, in *ApplicationIdentifiers, opts ...grpc.CallOption) (*Rights, error)
	// Create an API key scoped to this application.
	CreateAPIKey(ctx context.Context, in *CreateApplicationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error)
	// List the API keys for this application.
	ListAPIKeys(ctx context.Context, in *ListApplicationAPIKeysRequest, opts ...grpc.CallOption) (*APIKeys, error)
	// Get a single API key of this application.
	GetAPIKey(ctx context.Context, in *GetApplicationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error)
	// Update the rights of an API key of the application.
	// This method can also be used to delete the API key, by giving it no rights.
	// The caller is required to have all assigned or/and removed rights.
	UpdateAPIKey(ctx context.Context, in *UpdateApplicationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error)
	// Get the rights of a collaborator (member) of the application.
	// Pseudo-rights in the response (such as the "_ALL" right) are not expanded.
	GetCollaborator(ctx context.Context, in *GetApplicationCollaboratorRequest, opts ...grpc.CallOption) (*GetCollaboratorResponse, error)
	// Set the rights of a collaborator (member) on the application.
	// This method can also be used to delete the collaborator, by giving them no rights.
	// The caller is required to have all assigned or/and removed rights.
	SetCollaborator(ctx context.Context, in *SetApplicationCollaboratorRequest, opts ...grpc.CallOption) (*types.Empty, error)
	// List the collaborators on this application.
	ListCollaborators(ctx context.Context, in *ListApplicationCollaboratorsRequest, opts ...grpc.CallOption) (*Collaborators, error)
}

type applicationAccessClient struct {
	cc *grpc.ClientConn
}

func NewApplicationAccessClient(cc *grpc.ClientConn) ApplicationAccessClient {
	return &applicationAccessClient{cc}
}

func (c *applicationAccessClient) ListRights(ctx context.Context, in *ApplicationIdentifiers, opts ...grpc.CallOption) (*Rights, error) {
	out := new(Rights)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationAccess/ListRights", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationAccessClient) CreateAPIKey(ctx context.Context, in *CreateApplicationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error) {
	out := new(APIKey)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationAccess/CreateAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationAccessClient) ListAPIKeys(ctx context.Context, in *ListApplicationAPIKeysRequest, opts ...grpc.CallOption) (*APIKeys, error) {
	out := new(APIKeys)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationAccess/ListAPIKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationAccessClient) GetAPIKey(ctx context.Context, in *GetApplicationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error) {
	out := new(APIKey)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationAccess/GetAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationAccessClient) UpdateAPIKey(ctx context.Context, in *UpdateApplicationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error) {
	out := new(APIKey)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationAccess/UpdateAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationAccessClient) GetCollaborator(ctx context.Context, in *GetApplicationCollaboratorRequest, opts ...grpc.CallOption) (*GetCollaboratorResponse, error) {
	out := new(GetCollaboratorResponse)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationAccess/GetCollaborator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationAccessClient) SetCollaborator(ctx context.Context, in *SetApplicationCollaboratorRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationAccess/SetCollaborator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationAccessClient) ListCollaborators(ctx context.Context, in *ListApplicationCollaboratorsRequest, opts ...grpc.CallOption) (*Collaborators, error) {
	out := new(Collaborators)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationAccess/ListCollaborators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplicationAccessServer is the server API for ApplicationAccess service.
type ApplicationAccessServer interface {
	// List the rights the caller has on this application.
	ListRights(context.Context, *ApplicationIdentifiers) (*Rights, error)
	// Create an API key scoped to this application.
	CreateAPIKey(context.Context, *CreateApplicationAPIKeyRequest) (*APIKey, error)
	// List the API keys for this application.
	ListAPIKeys(context.Context, *ListApplicationAPIKeysRequest) (*APIKeys, error)
	// Get a single API key of this application.
	GetAPIKey(context.Context, *GetApplicationAPIKeyRequest) (*APIKey, error)
	// Update the rights of an API key of the application.
	// This method can also be used to delete the API key, by giving it no rights.
	// The caller is required to have all assigned or/and removed rights.
	UpdateAPIKey(context.Context, *UpdateApplicationAPIKeyRequest) (*APIKey, error)
	// Get the rights of a collaborator (member) of the application.
	// Pseudo-rights in the response (such as the "_ALL" right) are not expanded.
	GetCollaborator(context.Context, *GetApplicationCollaboratorRequest) (*GetCollaboratorResponse, error)
	// Set the rights of a collaborator (member) on the application.
	// This method can also be used to delete the collaborator, by giving them no rights.
	// The caller is required to have all assigned or/and removed rights.
	SetCollaborator(context.Context, *SetApplicationCollaboratorRequest) (*types.Empty, error)
	// List the collaborators on this application.
	ListCollaborators(context.Context, *ListApplicationCollaboratorsRequest) (*Collaborators, error)
}

// UnimplementedApplicationAccessServer can be embedded to have forward compatible implementations.
type UnimplementedApplicationAccessServer struct {
}

func (*UnimplementedApplicationAccessServer) ListRights(ctx context.Context, req *ApplicationIdentifiers) (*Rights, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRights not implemented")
}
func (*UnimplementedApplicationAccessServer) CreateAPIKey(ctx context.Context, req *CreateApplicationAPIKeyRequest) (*APIKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAPIKey not implemented")
}
func (*UnimplementedApplicationAccessServer) ListAPIKeys(ctx context.Context, req *ListApplicationAPIKeysRequest) (*APIKeys, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAPIKeys not implemented")
}
func (*UnimplementedApplicationAccessServer) GetAPIKey(ctx context.Context, req *GetApplicationAPIKeyRequest) (*APIKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAPIKey not implemented")
}
func (*UnimplementedApplicationAccessServer) UpdateAPIKey(ctx context.Context, req *UpdateApplicationAPIKeyRequest) (*APIKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAPIKey not implemented")
}
func (*UnimplementedApplicationAccessServer) GetCollaborator(ctx context.Context, req *GetApplicationCollaboratorRequest) (*GetCollaboratorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollaborator not implemented")
}
func (*UnimplementedApplicationAccessServer) SetCollaborator(ctx context.Context, req *SetApplicationCollaboratorRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCollaborator not implemented")
}
func (*UnimplementedApplicationAccessServer) ListCollaborators(ctx context.Context, req *ListApplicationCollaboratorsRequest) (*Collaborators, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCollaborators not implemented")
}

func RegisterApplicationAccessServer(s *grpc.Server, srv ApplicationAccessServer) {
	s.RegisterService(&_ApplicationAccess_serviceDesc, srv)
}

func _ApplicationAccess_ListRights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationAccessServer).ListRights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationAccess/ListRights",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationAccessServer).ListRights(ctx, req.(*ApplicationIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationAccess_CreateAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApplicationAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationAccessServer).CreateAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationAccess/CreateAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationAccessServer).CreateAPIKey(ctx, req.(*CreateApplicationAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationAccess_ListAPIKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApplicationAPIKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationAccessServer).ListAPIKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationAccess/ListAPIKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationAccessServer).ListAPIKeys(ctx, req.(*ListApplicationAPIKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationAccess_GetAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApplicationAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationAccessServer).GetAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationAccess/GetAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationAccessServer).GetAPIKey(ctx, req.(*GetApplicationAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationAccess_UpdateAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateApplicationAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationAccessServer).UpdateAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationAccess/UpdateAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationAccessServer).UpdateAPIKey(ctx, req.(*UpdateApplicationAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationAccess_GetCollaborator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApplicationCollaboratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationAccessServer).GetCollaborator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationAccess/GetCollaborator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationAccessServer).GetCollaborator(ctx, req.(*GetApplicationCollaboratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationAccess_SetCollaborator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetApplicationCollaboratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationAccessServer).SetCollaborator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationAccess/SetCollaborator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationAccessServer).SetCollaborator(ctx, req.(*SetApplicationCollaboratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationAccess_ListCollaborators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApplicationCollaboratorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationAccessServer).ListCollaborators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationAccess/ListCollaborators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationAccessServer).ListCollaborators(ctx, req.(*ListApplicationCollaboratorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApplicationAccess_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.ApplicationAccess",
	HandlerType: (*ApplicationAccessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRights",
			Handler:    _ApplicationAccess_ListRights_Handler,
		},
		{
			MethodName: "CreateAPIKey",
			Handler:    _ApplicationAccess_CreateAPIKey_Handler,
		},
		{
			MethodName: "ListAPIKeys",
			Handler:    _ApplicationAccess_ListAPIKeys_Handler,
		},
		{
			MethodName: "GetAPIKey",
			Handler:    _ApplicationAccess_GetAPIKey_Handler,
		},
		{
			MethodName: "UpdateAPIKey",
			Handler:    _ApplicationAccess_UpdateAPIKey_Handler,
		},
		{
			MethodName: "GetCollaborator",
			Handler:    _ApplicationAccess_GetCollaborator_Handler,
		},
		{
			MethodName: "SetCollaborator",
			Handler:    _ApplicationAccess_SetCollaborator_Handler,
		},
		{
			MethodName: "ListCollaborators",
			Handler:    _ApplicationAccess_ListCollaborators_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/application_services.proto",
}
