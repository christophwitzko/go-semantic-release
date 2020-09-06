// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package hooks

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HooksPluginClient is the client API for HooksPlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HooksPluginClient interface {
	Init(ctx context.Context, in *HooksInit_Request, opts ...grpc.CallOption) (*HooksInit_Response, error)
	Name(ctx context.Context, in *HooksName_Request, opts ...grpc.CallOption) (*HooksName_Response, error)
	Version(ctx context.Context, in *HooksVersion_Request, opts ...grpc.CallOption) (*HooksVersion_Response, error)
	Success(ctx context.Context, in *SuccessHook_Request, opts ...grpc.CallOption) (*SuccessHook_Response, error)
	NoRelease(ctx context.Context, in *NoReleaseHook_Request, opts ...grpc.CallOption) (*NoReleaseHook_Response, error)
}

type hooksPluginClient struct {
	cc grpc.ClientConnInterface
}

func NewHooksPluginClient(cc grpc.ClientConnInterface) HooksPluginClient {
	return &hooksPluginClient{cc}
}

func (c *hooksPluginClient) Init(ctx context.Context, in *HooksInit_Request, opts ...grpc.CallOption) (*HooksInit_Response, error) {
	out := new(HooksInit_Response)
	err := c.cc.Invoke(ctx, "/HooksPlugin/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hooksPluginClient) Name(ctx context.Context, in *HooksName_Request, opts ...grpc.CallOption) (*HooksName_Response, error) {
	out := new(HooksName_Response)
	err := c.cc.Invoke(ctx, "/HooksPlugin/Name", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hooksPluginClient) Version(ctx context.Context, in *HooksVersion_Request, opts ...grpc.CallOption) (*HooksVersion_Response, error) {
	out := new(HooksVersion_Response)
	err := c.cc.Invoke(ctx, "/HooksPlugin/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hooksPluginClient) Success(ctx context.Context, in *SuccessHook_Request, opts ...grpc.CallOption) (*SuccessHook_Response, error) {
	out := new(SuccessHook_Response)
	err := c.cc.Invoke(ctx, "/HooksPlugin/Success", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hooksPluginClient) NoRelease(ctx context.Context, in *NoReleaseHook_Request, opts ...grpc.CallOption) (*NoReleaseHook_Response, error) {
	out := new(NoReleaseHook_Response)
	err := c.cc.Invoke(ctx, "/HooksPlugin/NoRelease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HooksPluginServer is the server API for HooksPlugin service.
// All implementations must embed UnimplementedHooksPluginServer
// for forward compatibility
type HooksPluginServer interface {
	Init(context.Context, *HooksInit_Request) (*HooksInit_Response, error)
	Name(context.Context, *HooksName_Request) (*HooksName_Response, error)
	Version(context.Context, *HooksVersion_Request) (*HooksVersion_Response, error)
	Success(context.Context, *SuccessHook_Request) (*SuccessHook_Response, error)
	NoRelease(context.Context, *NoReleaseHook_Request) (*NoReleaseHook_Response, error)
	mustEmbedUnimplementedHooksPluginServer()
}

// UnimplementedHooksPluginServer must be embedded to have forward compatible implementations.
type UnimplementedHooksPluginServer struct {
}

func (*UnimplementedHooksPluginServer) Init(context.Context, *HooksInit_Request) (*HooksInit_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (*UnimplementedHooksPluginServer) Name(context.Context, *HooksName_Request) (*HooksName_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Name not implemented")
}
func (*UnimplementedHooksPluginServer) Version(context.Context, *HooksVersion_Request) (*HooksVersion_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (*UnimplementedHooksPluginServer) Success(context.Context, *SuccessHook_Request) (*SuccessHook_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Success not implemented")
}
func (*UnimplementedHooksPluginServer) NoRelease(context.Context, *NoReleaseHook_Request) (*NoReleaseHook_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NoRelease not implemented")
}
func (*UnimplementedHooksPluginServer) mustEmbedUnimplementedHooksPluginServer() {}

func RegisterHooksPluginServer(s *grpc.Server, srv HooksPluginServer) {
	s.RegisterService(&_HooksPlugin_serviceDesc, srv)
}

func _HooksPlugin_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HooksInit_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HooksPluginServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HooksPlugin/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HooksPluginServer).Init(ctx, req.(*HooksInit_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _HooksPlugin_Name_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HooksName_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HooksPluginServer).Name(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HooksPlugin/Name",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HooksPluginServer).Name(ctx, req.(*HooksName_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _HooksPlugin_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HooksVersion_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HooksPluginServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HooksPlugin/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HooksPluginServer).Version(ctx, req.(*HooksVersion_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _HooksPlugin_Success_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuccessHook_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HooksPluginServer).Success(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HooksPlugin/Success",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HooksPluginServer).Success(ctx, req.(*SuccessHook_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _HooksPlugin_NoRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoReleaseHook_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HooksPluginServer).NoRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HooksPlugin/NoRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HooksPluginServer).NoRelease(ctx, req.(*NoReleaseHook_Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _HooksPlugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "HooksPlugin",
	HandlerType: (*HooksPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _HooksPlugin_Init_Handler,
		},
		{
			MethodName: "Name",
			Handler:    _HooksPlugin_Name_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _HooksPlugin_Version_Handler,
		},
		{
			MethodName: "Success",
			Handler:    _HooksPlugin_Success_Handler,
		},
		{
			MethodName: "NoRelease",
			Handler:    _HooksPlugin_NoRelease_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/hooks/hooks.proto",
}
