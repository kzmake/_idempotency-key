// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package time

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

// TimeClient is the client API for Time service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TimeClient interface {
	Now(ctx context.Context, in *NowRequest, opts ...grpc.CallOption) (*NowResponse, error)
}

type timeClient struct {
	cc grpc.ClientConnInterface
}

func NewTimeClient(cc grpc.ClientConnInterface) TimeClient {
	return &timeClient{cc}
}

func (c *timeClient) Now(ctx context.Context, in *NowRequest, opts ...grpc.CallOption) (*NowResponse, error) {
	out := new(NowResponse)
	err := c.cc.Invoke(ctx, "/kzmake.time.v1.Time/Now", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimeServer is the server API for Time service.
// All implementations should embed UnimplementedTimeServer
// for forward compatibility
type TimeServer interface {
	Now(context.Context, *NowRequest) (*NowResponse, error)
}

// UnimplementedTimeServer should be embedded to have forward compatible implementations.
type UnimplementedTimeServer struct {
}

func (UnimplementedTimeServer) Now(context.Context, *NowRequest) (*NowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Now not implemented")
}

// UnsafeTimeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TimeServer will
// result in compilation errors.
type UnsafeTimeServer interface {
	mustEmbedUnimplementedTimeServer()
}

func RegisterTimeServer(s grpc.ServiceRegistrar, srv TimeServer) {
	s.RegisterService(&Time_ServiceDesc, srv)
}

func _Time_Now_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeServer).Now(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kzmake.time.v1.Time/Now",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeServer).Now(ctx, req.(*NowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Time_ServiceDesc is the grpc.ServiceDesc for Time service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Time_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kzmake.time.v1.Time",
	HandlerType: (*TimeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Now",
			Handler:    _Time_Now_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "time/v1/time.proto",
}
