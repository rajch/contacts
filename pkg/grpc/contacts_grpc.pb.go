// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ContactServiceClient is the client API for ContactService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContactServiceClient interface {
	NewContact(ctx context.Context, in *Contact, opts ...grpc.CallOption) (*Contact, error)
	GetAllContacts(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (ContactService_GetAllContactsClient, error)
	GetContactById(ctx context.Context, in *GetContactInput, opts ...grpc.CallOption) (*Contact, error)
}

type contactServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContactServiceClient(cc grpc.ClientConnInterface) ContactServiceClient {
	return &contactServiceClient{cc}
}

func (c *contactServiceClient) NewContact(ctx context.Context, in *Contact, opts ...grpc.CallOption) (*Contact, error) {
	out := new(Contact)
	err := c.cc.Invoke(ctx, "/contacts.ContactService/NewContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) GetAllContacts(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (ContactService_GetAllContactsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ContactService_ServiceDesc.Streams[0], "/contacts.ContactService/GetAllContacts", opts...)
	if err != nil {
		return nil, err
	}
	x := &contactServiceGetAllContactsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ContactService_GetAllContactsClient interface {
	Recv() (*Contact, error)
	grpc.ClientStream
}

type contactServiceGetAllContactsClient struct {
	grpc.ClientStream
}

func (x *contactServiceGetAllContactsClient) Recv() (*Contact, error) {
	m := new(Contact)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *contactServiceClient) GetContactById(ctx context.Context, in *GetContactInput, opts ...grpc.CallOption) (*Contact, error) {
	out := new(Contact)
	err := c.cc.Invoke(ctx, "/contacts.ContactService/GetContactById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContactServiceServer is the server API for ContactService service.
// All implementations must embed UnimplementedContactServiceServer
// for forward compatibility
type ContactServiceServer interface {
	NewContact(context.Context, *Contact) (*Contact, error)
	GetAllContacts(*emptypb.Empty, ContactService_GetAllContactsServer) error
	GetContactById(context.Context, *GetContactInput) (*Contact, error)
	mustEmbedUnimplementedContactServiceServer()
}

// UnimplementedContactServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContactServiceServer struct {
}

func (UnimplementedContactServiceServer) NewContact(context.Context, *Contact) (*Contact, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewContact not implemented")
}
func (UnimplementedContactServiceServer) GetAllContacts(*emptypb.Empty, ContactService_GetAllContactsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllContacts not implemented")
}
func (UnimplementedContactServiceServer) GetContactById(context.Context, *GetContactInput) (*Contact, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContactById not implemented")
}
func (UnimplementedContactServiceServer) mustEmbedUnimplementedContactServiceServer() {}

// UnsafeContactServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContactServiceServer will
// result in compilation errors.
type UnsafeContactServiceServer interface {
	mustEmbedUnimplementedContactServiceServer()
}

func RegisterContactServiceServer(s grpc.ServiceRegistrar, srv ContactServiceServer) {
	s.RegisterService(&ContactService_ServiceDesc, srv)
}

func _ContactService_NewContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Contact)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).NewContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contacts.ContactService/NewContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).NewContact(ctx, req.(*Contact))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_GetAllContacts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ContactServiceServer).GetAllContacts(m, &contactServiceGetAllContactsServer{stream})
}

type ContactService_GetAllContactsServer interface {
	Send(*Contact) error
	grpc.ServerStream
}

type contactServiceGetAllContactsServer struct {
	grpc.ServerStream
}

func (x *contactServiceGetAllContactsServer) Send(m *Contact) error {
	return x.ServerStream.SendMsg(m)
}

func _ContactService_GetContactById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContactInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).GetContactById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contacts.ContactService/GetContactById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).GetContactById(ctx, req.(*GetContactInput))
	}
	return interceptor(ctx, in, info, handler)
}

// ContactService_ServiceDesc is the grpc.ServiceDesc for ContactService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContactService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "contacts.ContactService",
	HandlerType: (*ContactServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewContact",
			Handler:    _ContactService_NewContact_Handler,
		},
		{
			MethodName: "GetContactById",
			Handler:    _ContactService_GetContactById_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllContacts",
			Handler:       _ContactService_GetAllContacts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "contacts.proto",
}
