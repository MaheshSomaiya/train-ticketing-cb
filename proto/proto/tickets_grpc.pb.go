// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0
// source: proto/tickets.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TrainService_PurchaseTicket_FullMethodName = "/tickets.TrainService/PurchaseTicket"
	TrainService_ViewTicket_FullMethodName     = "/tickets.TrainService/ViewTicket"
	TrainService_ViewUsers_FullMethodName      = "/tickets.TrainService/ViewUsers"
	TrainService_RemoveUser_FullMethodName     = "/tickets.TrainService/RemoveUser"
	TrainService_ModifySeat_FullMethodName     = "/tickets.TrainService/ModifySeat"
)

// TrainServiceClient is the client API for TrainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrainServiceClient interface {
	PurchaseTicket(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*TicketResponse, error)
	ViewTicket(ctx context.Context, in *ViewRequest, opts ...grpc.CallOption) (*TicketResponse, error)
	ViewUsers(ctx context.Context, in *ViewRequest, opts ...grpc.CallOption) (*UsersResponse, error)
	RemoveUser(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error)
	ModifySeat(ctx context.Context, in *ModifySeatRequest, opts ...grpc.CallOption) (*TicketResponse, error)
}

type trainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTrainServiceClient(cc grpc.ClientConnInterface) TrainServiceClient {
	return &trainServiceClient{cc}
}

func (c *trainServiceClient) PurchaseTicket(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*TicketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TicketResponse)
	err := c.cc.Invoke(ctx, TrainService_PurchaseTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainServiceClient) ViewTicket(ctx context.Context, in *ViewRequest, opts ...grpc.CallOption) (*TicketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TicketResponse)
	err := c.cc.Invoke(ctx, TrainService_ViewTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainServiceClient) ViewUsers(ctx context.Context, in *ViewRequest, opts ...grpc.CallOption) (*UsersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UsersResponse)
	err := c.cc.Invoke(ctx, TrainService_ViewUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainServiceClient) RemoveUser(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveResponse)
	err := c.cc.Invoke(ctx, TrainService_RemoveUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainServiceClient) ModifySeat(ctx context.Context, in *ModifySeatRequest, opts ...grpc.CallOption) (*TicketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TicketResponse)
	err := c.cc.Invoke(ctx, TrainService_ModifySeat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrainServiceServer is the server API for TrainService service.
// All implementations must embed UnimplementedTrainServiceServer
// for forward compatibility.
type TrainServiceServer interface {
	PurchaseTicket(context.Context, *PurchaseRequest) (*TicketResponse, error)
	ViewTicket(context.Context, *ViewRequest) (*TicketResponse, error)
	ViewUsers(context.Context, *ViewRequest) (*UsersResponse, error)
	RemoveUser(context.Context, *RemoveRequest) (*RemoveResponse, error)
	ModifySeat(context.Context, *ModifySeatRequest) (*TicketResponse, error)
	mustEmbedUnimplementedTrainServiceServer()
}

// UnimplementedTrainServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTrainServiceServer struct{}

func (UnimplementedTrainServiceServer) PurchaseTicket(context.Context, *PurchaseRequest) (*TicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurchaseTicket not implemented")
}
func (UnimplementedTrainServiceServer) ViewTicket(context.Context, *ViewRequest) (*TicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewTicket not implemented")
}
func (UnimplementedTrainServiceServer) ViewUsers(context.Context, *ViewRequest) (*UsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewUsers not implemented")
}
func (UnimplementedTrainServiceServer) RemoveUser(context.Context, *RemoveRequest) (*RemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUser not implemented")
}
func (UnimplementedTrainServiceServer) ModifySeat(context.Context, *ModifySeatRequest) (*TicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifySeat not implemented")
}
func (UnimplementedTrainServiceServer) mustEmbedUnimplementedTrainServiceServer() {}
func (UnimplementedTrainServiceServer) testEmbeddedByValue()                      {}

// UnsafeTrainServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrainServiceServer will
// result in compilation errors.
type UnsafeTrainServiceServer interface {
	mustEmbedUnimplementedTrainServiceServer()
}

func RegisterTrainServiceServer(s grpc.ServiceRegistrar, srv TrainServiceServer) {
	// If the following call pancis, it indicates UnimplementedTrainServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TrainService_ServiceDesc, srv)
}

func _TrainService_PurchaseTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainServiceServer).PurchaseTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainService_PurchaseTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainServiceServer).PurchaseTicket(ctx, req.(*PurchaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainService_ViewTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainServiceServer).ViewTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainService_ViewTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainServiceServer).ViewTicket(ctx, req.(*ViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainService_ViewUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainServiceServer).ViewUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainService_ViewUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainServiceServer).ViewUsers(ctx, req.(*ViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainService_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainServiceServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainService_RemoveUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainServiceServer).RemoveUser(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainService_ModifySeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifySeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainServiceServer).ModifySeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainService_ModifySeat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainServiceServer).ModifySeat(ctx, req.(*ModifySeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TrainService_ServiceDesc is the grpc.ServiceDesc for TrainService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrainService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tickets.TrainService",
	HandlerType: (*TrainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PurchaseTicket",
			Handler:    _TrainService_PurchaseTicket_Handler,
		},
		{
			MethodName: "ViewTicket",
			Handler:    _TrainService_ViewTicket_Handler,
		},
		{
			MethodName: "ViewUsers",
			Handler:    _TrainService_ViewUsers_Handler,
		},
		{
			MethodName: "RemoveUser",
			Handler:    _TrainService_RemoveUser_Handler,
		},
		{
			MethodName: "ModifySeat",
			Handler:    _TrainService_ModifySeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/tickets.proto",
}
