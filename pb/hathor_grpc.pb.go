// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: pb/hathor.proto

package hathor

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Hathor_Ping_FullMethodName        = "/hathor.Hathor/Ping"
	Hathor_CreateEvent_FullMethodName = "/hathor.Hathor/CreateEvent"
	Hathor_GetEvent_FullMethodName    = "/hathor.Hathor/GetEvent"
	Hathor_UpdateEvent_FullMethodName = "/hathor.Hathor/UpdateEvent"
	Hathor_DeleteEvent_FullMethodName = "/hathor.Hathor/DeleteEvent"
)

// HathorClient is the client API for Hathor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HathorClient interface {
	Ping(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	// CreateEvent creates and event given an Event message
	CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*Event, error)
	// GetEvent returns the relevant info for an event given an ID
	GetEvent(ctx context.Context, in *EventIDRequest, opts ...grpc.CallOption) (*Event, error)
	// UpdateEvent finds and updates a given event.
	UpdateEvent(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*Event, error)
	// DeleteEvent removes an event from the repository given an ID
	DeleteEvent(ctx context.Context, in *EventIDRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type hathorClient struct {
	cc grpc.ClientConnInterface
}

func NewHathorClient(cc grpc.ClientConnInterface) HathorClient {
	return &hathorClient{cc}
}

func (c *hathorClient) Ping(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, Hathor_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hathorClient) CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, Hathor_CreateEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hathorClient) GetEvent(ctx context.Context, in *EventIDRequest, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, Hathor_GetEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hathorClient) UpdateEvent(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, Hathor_UpdateEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hathorClient) DeleteEvent(ctx context.Context, in *EventIDRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, Hathor_DeleteEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HathorServer is the server API for Hathor service.
// All implementations must embed UnimplementedHathorServer
// for forward compatibility
type HathorServer interface {
	Ping(context.Context, *EmptyRequest) (*EmptyResponse, error)
	// CreateEvent creates and event given an Event message
	CreateEvent(context.Context, *CreateEventRequest) (*Event, error)
	// GetEvent returns the relevant info for an event given an ID
	GetEvent(context.Context, *EventIDRequest) (*Event, error)
	// UpdateEvent finds and updates a given event.
	UpdateEvent(context.Context, *UpdateEventRequest) (*Event, error)
	// DeleteEvent removes an event from the repository given an ID
	DeleteEvent(context.Context, *EventIDRequest) (*EmptyResponse, error)
	mustEmbedUnimplementedHathorServer()
}

// UnimplementedHathorServer must be embedded to have forward compatible implementations.
type UnimplementedHathorServer struct {
}

func (UnimplementedHathorServer) Ping(context.Context, *EmptyRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedHathorServer) CreateEvent(context.Context, *CreateEventRequest) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (UnimplementedHathorServer) GetEvent(context.Context, *EventIDRequest) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (UnimplementedHathorServer) UpdateEvent(context.Context, *UpdateEventRequest) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}
func (UnimplementedHathorServer) DeleteEvent(context.Context, *EventIDRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}
func (UnimplementedHathorServer) mustEmbedUnimplementedHathorServer() {}

// UnsafeHathorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HathorServer will
// result in compilation errors.
type UnsafeHathorServer interface {
	mustEmbedUnimplementedHathorServer()
}

func RegisterHathorServer(s grpc.ServiceRegistrar, srv HathorServer) {
	s.RegisterService(&Hathor_ServiceDesc, srv)
}

func _Hathor_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HathorServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hathor_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HathorServer).Ping(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hathor_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HathorServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hathor_CreateEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HathorServer).CreateEvent(ctx, req.(*CreateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hathor_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HathorServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hathor_GetEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HathorServer).GetEvent(ctx, req.(*EventIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hathor_UpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HathorServer).UpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hathor_UpdateEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HathorServer).UpdateEvent(ctx, req.(*UpdateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hathor_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HathorServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hathor_DeleteEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HathorServer).DeleteEvent(ctx, req.(*EventIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Hathor_ServiceDesc is the grpc.ServiceDesc for Hathor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hathor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hathor.Hathor",
	HandlerType: (*HathorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Hathor_Ping_Handler,
		},
		{
			MethodName: "CreateEvent",
			Handler:    _Hathor_CreateEvent_Handler,
		},
		{
			MethodName: "GetEvent",
			Handler:    _Hathor_GetEvent_Handler,
		},
		{
			MethodName: "UpdateEvent",
			Handler:    _Hathor_UpdateEvent_Handler,
		},
		{
			MethodName: "DeleteEvent",
			Handler:    _Hathor_DeleteEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/hathor.proto",
}
