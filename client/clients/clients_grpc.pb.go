// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: clients.proto

package clients

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ClientServiceClient is the client API for ClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientServiceClient interface {
	Create(ctx context.Context, in *ClientRequest, opts ...grpc.CallOption) (*Client, error)
	List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (ClientService_ListClient, error)
	Delete(ctx context.Context, opts ...grpc.CallOption) (ClientService_DeleteClient, error)
	Update(ctx context.Context, in *Client, opts ...grpc.CallOption) (*Client, error)
	IsOwner(ctx context.Context, in *IsOwnerRequest, opts ...grpc.CallOption) (*IsOwnerResponse, error)
	CreateRole(ctx context.Context, opts ...grpc.CallOption) (ClientService_CreateRoleClient, error)
	UpdateRoles(ctx context.Context, opts ...grpc.CallOption) (ClientService_UpdateRolesClient, error)
	Retrieve(ctx context.Context, in *DeleteClientRequest, opts ...grpc.CallOption) (ClientService_RetrieveClient, error)
}

type clientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientServiceClient(cc grpc.ClientConnInterface) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) Create(ctx context.Context, in *ClientRequest, opts ...grpc.CallOption) (*Client, error) {
	out := new(Client)
	err := c.cc.Invoke(ctx, "/ClientService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (ClientService_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientService_ServiceDesc.Streams[0], "/ClientService/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ClientService_ListClient interface {
	Recv() (*Client, error)
	grpc.ClientStream
}

type clientServiceListClient struct {
	grpc.ClientStream
}

func (x *clientServiceListClient) Recv() (*Client, error) {
	m := new(Client)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientServiceClient) Delete(ctx context.Context, opts ...grpc.CallOption) (ClientService_DeleteClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientService_ServiceDesc.Streams[1], "/ClientService/Delete", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientServiceDeleteClient{stream}
	return x, nil
}

type ClientService_DeleteClient interface {
	Send(*DeleteClientRequest) error
	Recv() (*Client, error)
	grpc.ClientStream
}

type clientServiceDeleteClient struct {
	grpc.ClientStream
}

func (x *clientServiceDeleteClient) Send(m *DeleteClientRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clientServiceDeleteClient) Recv() (*Client, error) {
	m := new(Client)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientServiceClient) Update(ctx context.Context, in *Client, opts ...grpc.CallOption) (*Client, error) {
	out := new(Client)
	err := c.cc.Invoke(ctx, "/ClientService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) IsOwner(ctx context.Context, in *IsOwnerRequest, opts ...grpc.CallOption) (*IsOwnerResponse, error) {
	out := new(IsOwnerResponse)
	err := c.cc.Invoke(ctx, "/ClientService/IsOwner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) CreateRole(ctx context.Context, opts ...grpc.CallOption) (ClientService_CreateRoleClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientService_ServiceDesc.Streams[2], "/ClientService/CreateRole", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientServiceCreateRoleClient{stream}
	return x, nil
}

type ClientService_CreateRoleClient interface {
	Send(*CreateRoleRequest) error
	Recv() (*Role, error)
	grpc.ClientStream
}

type clientServiceCreateRoleClient struct {
	grpc.ClientStream
}

func (x *clientServiceCreateRoleClient) Send(m *CreateRoleRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clientServiceCreateRoleClient) Recv() (*Role, error) {
	m := new(Role)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientServiceClient) UpdateRoles(ctx context.Context, opts ...grpc.CallOption) (ClientService_UpdateRolesClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientService_ServiceDesc.Streams[3], "/ClientService/UpdateRoles", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientServiceUpdateRolesClient{stream}
	return x, nil
}

type ClientService_UpdateRolesClient interface {
	Send(*Role) error
	Recv() (*Role, error)
	grpc.ClientStream
}

type clientServiceUpdateRolesClient struct {
	grpc.ClientStream
}

func (x *clientServiceUpdateRolesClient) Send(m *Role) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clientServiceUpdateRolesClient) Recv() (*Role, error) {
	m := new(Role)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientServiceClient) Retrieve(ctx context.Context, in *DeleteClientRequest, opts ...grpc.CallOption) (ClientService_RetrieveClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientService_ServiceDesc.Streams[4], "/ClientService/Retrieve", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientServiceRetrieveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ClientService_RetrieveClient interface {
	Recv() (*RetrieveResourceScope, error)
	grpc.ClientStream
}

type clientServiceRetrieveClient struct {
	grpc.ClientStream
}

func (x *clientServiceRetrieveClient) Recv() (*RetrieveResourceScope, error) {
	m := new(RetrieveResourceScope)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClientServiceServer is the server API for ClientService service.
// All implementations must embed UnimplementedClientServiceServer
// for forward compatibility
type ClientServiceServer interface {
	Create(context.Context, *ClientRequest) (*Client, error)
	List(*empty.Empty, ClientService_ListServer) error
	Delete(ClientService_DeleteServer) error
	Update(context.Context, *Client) (*Client, error)
	IsOwner(context.Context, *IsOwnerRequest) (*IsOwnerResponse, error)
	CreateRole(ClientService_CreateRoleServer) error
	UpdateRoles(ClientService_UpdateRolesServer) error
	Retrieve(*DeleteClientRequest, ClientService_RetrieveServer) error
	mustEmbedUnimplementedClientServiceServer()
}

// UnimplementedClientServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientServiceServer struct {
}

func (UnimplementedClientServiceServer) Create(context.Context, *ClientRequest) (*Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedClientServiceServer) List(*empty.Empty, ClientService_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedClientServiceServer) Delete(ClientService_DeleteServer) error {
	return status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedClientServiceServer) Update(context.Context, *Client) (*Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedClientServiceServer) IsOwner(context.Context, *IsOwnerRequest) (*IsOwnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsOwner not implemented")
}
func (UnimplementedClientServiceServer) CreateRole(ClientService_CreateRoleServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedClientServiceServer) UpdateRoles(ClientService_UpdateRolesServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateRoles not implemented")
}
func (UnimplementedClientServiceServer) Retrieve(*DeleteClientRequest, ClientService_RetrieveServer) error {
	return status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}
func (UnimplementedClientServiceServer) mustEmbedUnimplementedClientServiceServer() {}

// UnsafeClientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientServiceServer will
// result in compilation errors.
type UnsafeClientServiceServer interface {
	mustEmbedUnimplementedClientServiceServer()
}

func RegisterClientServiceServer(s grpc.ServiceRegistrar, srv ClientServiceServer) {
	s.RegisterService(&ClientService_ServiceDesc, srv)
}

func _ClientService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ClientService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Create(ctx, req.(*ClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClientServiceServer).List(m, &clientServiceListServer{stream})
}

type ClientService_ListServer interface {
	Send(*Client) error
	grpc.ServerStream
}

type clientServiceListServer struct {
	grpc.ServerStream
}

func (x *clientServiceListServer) Send(m *Client) error {
	return x.ServerStream.SendMsg(m)
}

func _ClientService_Delete_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClientServiceServer).Delete(&clientServiceDeleteServer{stream})
}

type ClientService_DeleteServer interface {
	Send(*Client) error
	Recv() (*DeleteClientRequest, error)
	grpc.ServerStream
}

type clientServiceDeleteServer struct {
	grpc.ServerStream
}

func (x *clientServiceDeleteServer) Send(m *Client) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clientServiceDeleteServer) Recv() (*DeleteClientRequest, error) {
	m := new(DeleteClientRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ClientService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Client)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ClientService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Update(ctx, req.(*Client))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_IsOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).IsOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ClientService/IsOwner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).IsOwner(ctx, req.(*IsOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_CreateRole_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClientServiceServer).CreateRole(&clientServiceCreateRoleServer{stream})
}

type ClientService_CreateRoleServer interface {
	Send(*Role) error
	Recv() (*CreateRoleRequest, error)
	grpc.ServerStream
}

type clientServiceCreateRoleServer struct {
	grpc.ServerStream
}

func (x *clientServiceCreateRoleServer) Send(m *Role) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clientServiceCreateRoleServer) Recv() (*CreateRoleRequest, error) {
	m := new(CreateRoleRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ClientService_UpdateRoles_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClientServiceServer).UpdateRoles(&clientServiceUpdateRolesServer{stream})
}

type ClientService_UpdateRolesServer interface {
	Send(*Role) error
	Recv() (*Role, error)
	grpc.ServerStream
}

type clientServiceUpdateRolesServer struct {
	grpc.ServerStream
}

func (x *clientServiceUpdateRolesServer) Send(m *Role) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clientServiceUpdateRolesServer) Recv() (*Role, error) {
	m := new(Role)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ClientService_Retrieve_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DeleteClientRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClientServiceServer).Retrieve(m, &clientServiceRetrieveServer{stream})
}

type ClientService_RetrieveServer interface {
	Send(*RetrieveResourceScope) error
	grpc.ServerStream
}

type clientServiceRetrieveServer struct {
	grpc.ServerStream
}

func (x *clientServiceRetrieveServer) Send(m *RetrieveResourceScope) error {
	return x.ServerStream.SendMsg(m)
}

// ClientService_ServiceDesc is the grpc.ServiceDesc for ClientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ClientService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ClientService_Update_Handler,
		},
		{
			MethodName: "IsOwner",
			Handler:    _ClientService_IsOwner_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _ClientService_List_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Delete",
			Handler:       _ClientService_Delete_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "CreateRole",
			Handler:       _ClientService_CreateRole_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "UpdateRoles",
			Handler:       _ClientService_UpdateRoles_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Retrieve",
			Handler:       _ClientService_Retrieve_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "clients.proto",
}
