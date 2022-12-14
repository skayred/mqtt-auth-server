// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package auth

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

// MqttAuthClient is the client API for MqttAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MqttAuthClient interface {
	CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*Response, error)
	UpdateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*Response, error)
	ActivateDevice(ctx context.Context, in *DeviceRequest, opts ...grpc.CallOption) (*Response, error)
	DeactivateDevice(ctx context.Context, in *DeviceRequest, opts ...grpc.CallOption) (*Response, error)
}

type mqttAuthClient struct {
	cc grpc.ClientConnInterface
}

func NewMqttAuthClient(cc grpc.ClientConnInterface) MqttAuthClient {
	return &mqttAuthClient{cc}
}

func (c *mqttAuthClient) CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/MqttAuth/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqttAuthClient) UpdateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/MqttAuth/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqttAuthClient) ActivateDevice(ctx context.Context, in *DeviceRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/MqttAuth/ActivateDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqttAuthClient) DeactivateDevice(ctx context.Context, in *DeviceRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/MqttAuth/DeactivateDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MqttAuthServer is the server API for MqttAuth service.
// All implementations must embed UnimplementedMqttAuthServer
// for forward compatibility
type MqttAuthServer interface {
	CreateUser(context.Context, *UserRequest) (*Response, error)
	UpdateUser(context.Context, *UserRequest) (*Response, error)
	ActivateDevice(context.Context, *DeviceRequest) (*Response, error)
	DeactivateDevice(context.Context, *DeviceRequest) (*Response, error)
	mustEmbedUnimplementedMqttAuthServer()
}

// UnimplementedMqttAuthServer must be embedded to have forward compatible implementations.
type UnimplementedMqttAuthServer struct {
}

func (UnimplementedMqttAuthServer) CreateUser(context.Context, *UserRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedMqttAuthServer) UpdateUser(context.Context, *UserRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedMqttAuthServer) ActivateDevice(context.Context, *DeviceRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateDevice not implemented")
}
func (UnimplementedMqttAuthServer) DeactivateDevice(context.Context, *DeviceRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactivateDevice not implemented")
}
func (UnimplementedMqttAuthServer) mustEmbedUnimplementedMqttAuthServer() {}

// UnsafeMqttAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MqttAuthServer will
// result in compilation errors.
type UnsafeMqttAuthServer interface {
	mustEmbedUnimplementedMqttAuthServer()
}

func RegisterMqttAuthServer(s grpc.ServiceRegistrar, srv MqttAuthServer) {
	s.RegisterService(&MqttAuth_ServiceDesc, srv)
}

func _MqttAuth_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqttAuthServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MqttAuth/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqttAuthServer).CreateUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqttAuth_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqttAuthServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MqttAuth/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqttAuthServer).UpdateUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqttAuth_ActivateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqttAuthServer).ActivateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MqttAuth/ActivateDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqttAuthServer).ActivateDevice(ctx, req.(*DeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqttAuth_DeactivateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqttAuthServer).DeactivateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MqttAuth/DeactivateDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqttAuthServer).DeactivateDevice(ctx, req.(*DeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MqttAuth_ServiceDesc is the grpc.ServiceDesc for MqttAuth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MqttAuth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MqttAuth",
	HandlerType: (*MqttAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _MqttAuth_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _MqttAuth_UpdateUser_Handler,
		},
		{
			MethodName: "ActivateDevice",
			Handler:    _MqttAuth_ActivateDevice_Handler,
		},
		{
			MethodName: "DeactivateDevice",
			Handler:    _MqttAuth_DeactivateDevice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
