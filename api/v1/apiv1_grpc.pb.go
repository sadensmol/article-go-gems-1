// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: apiv1.proto

package v1

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
	APIV1Service_Test_FullMethodName = "/api.v1.APIV1Service/Test"
)

// APIV1ServiceClient is the client API for APIV1Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type APIV1ServiceClient interface {
	Test(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error)
}

type aPIV1ServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIV1ServiceClient(cc grpc.ClientConnInterface) APIV1ServiceClient {
	return &aPIV1ServiceClient{cc}
}

func (c *aPIV1ServiceClient) Test(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error) {
	out := new(TestResponse)
	err := c.cc.Invoke(ctx, APIV1Service_Test_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIV1ServiceServer is the server API for APIV1Service service.
// All implementations must embed UnimplementedAPIV1ServiceServer
// for forward compatibility
type APIV1ServiceServer interface {
	Test(context.Context, *TestRequest) (*TestResponse, error)
	mustEmbedUnimplementedAPIV1ServiceServer()
}

// UnimplementedAPIV1ServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAPIV1ServiceServer struct {
}

func (UnimplementedAPIV1ServiceServer) Test(context.Context, *TestRequest) (*TestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}
func (UnimplementedAPIV1ServiceServer) mustEmbedUnimplementedAPIV1ServiceServer() {}

// UnsafeAPIV1ServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to APIV1ServiceServer will
// result in compilation errors.
type UnsafeAPIV1ServiceServer interface {
	mustEmbedUnimplementedAPIV1ServiceServer()
}

func RegisterAPIV1ServiceServer(s grpc.ServiceRegistrar, srv APIV1ServiceServer) {
	s.RegisterService(&APIV1Service_ServiceDesc, srv)
}

func _APIV1Service_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIV1ServiceServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: APIV1Service_Test_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIV1ServiceServer).Test(ctx, req.(*TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// APIV1Service_ServiceDesc is the grpc.ServiceDesc for APIV1Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var APIV1Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.APIV1Service",
	HandlerType: (*APIV1ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Test",
			Handler:    _APIV1Service_Test_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apiv1.proto",
}
