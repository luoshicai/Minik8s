// grpc server is API Server and client is Kubelet

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: proto/apiserver_for_kubelet.proto

package proto

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
	ApiServerKubeletService_SayHello_FullMethodName        = "/apiserver_for_kubelet.ApiServerKubeletService/SayHello"
	ApiServerKubeletService_RegisterNode_FullMethodName    = "/apiserver_for_kubelet.ApiServerKubeletService/RegisterNode"
	ApiServerKubeletService_UpdatePodStatus_FullMethodName = "/apiserver_for_kubelet.ApiServerKubeletService/UpdatePodStatus"
)

// ApiServerKubeletServiceClient is the client API for ApiServerKubeletService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiServerKubeletServiceClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	RegisterNode(ctx context.Context, in *RegisterNodeRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	UpdatePodStatus(ctx context.Context, in *UpdatePodStatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type apiServerKubeletServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiServerKubeletServiceClient(cc grpc.ClientConnInterface) ApiServerKubeletServiceClient {
	return &apiServerKubeletServiceClient{cc}
}

func (c *apiServerKubeletServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, ApiServerKubeletService_SayHello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServerKubeletServiceClient) RegisterNode(ctx context.Context, in *RegisterNodeRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, ApiServerKubeletService_RegisterNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServerKubeletServiceClient) UpdatePodStatus(ctx context.Context, in *UpdatePodStatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, ApiServerKubeletService_UpdatePodStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServerKubeletServiceServer is the server API for ApiServerKubeletService service.
// All implementations must embed UnimplementedApiServerKubeletServiceServer
// for forward compatibility
type ApiServerKubeletServiceServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
	RegisterNode(context.Context, *RegisterNodeRequest) (*StatusResponse, error)
	UpdatePodStatus(context.Context, *UpdatePodStatusRequest) (*StatusResponse, error)
	mustEmbedUnimplementedApiServerKubeletServiceServer()
}

// UnimplementedApiServerKubeletServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApiServerKubeletServiceServer struct {
}

func (UnimplementedApiServerKubeletServiceServer) SayHello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedApiServerKubeletServiceServer) RegisterNode(context.Context, *RegisterNodeRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterNode not implemented")
}
func (UnimplementedApiServerKubeletServiceServer) UpdatePodStatus(context.Context, *UpdatePodStatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePodStatus not implemented")
}
func (UnimplementedApiServerKubeletServiceServer) mustEmbedUnimplementedApiServerKubeletServiceServer() {
}

// UnsafeApiServerKubeletServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServerKubeletServiceServer will
// result in compilation errors.
type UnsafeApiServerKubeletServiceServer interface {
	mustEmbedUnimplementedApiServerKubeletServiceServer()
}

func RegisterApiServerKubeletServiceServer(s grpc.ServiceRegistrar, srv ApiServerKubeletServiceServer) {
	s.RegisterService(&ApiServerKubeletService_ServiceDesc, srv)
}

func _ApiServerKubeletService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServerKubeletServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiServerKubeletService_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServerKubeletServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiServerKubeletService_RegisterNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServerKubeletServiceServer).RegisterNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiServerKubeletService_RegisterNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServerKubeletServiceServer).RegisterNode(ctx, req.(*RegisterNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiServerKubeletService_UpdatePodStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePodStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServerKubeletServiceServer).UpdatePodStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiServerKubeletService_UpdatePodStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServerKubeletServiceServer).UpdatePodStatus(ctx, req.(*UpdatePodStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiServerKubeletService_ServiceDesc is the grpc.ServiceDesc for ApiServerKubeletService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiServerKubeletService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apiserver_for_kubelet.ApiServerKubeletService",
	HandlerType: (*ApiServerKubeletServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _ApiServerKubeletService_SayHello_Handler,
		},
		{
			MethodName: "RegisterNode",
			Handler:    _ApiServerKubeletService_RegisterNode_Handler,
		},
		{
			MethodName: "UpdatePodStatus",
			Handler:    _ApiServerKubeletService_UpdatePodStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/apiserver_for_kubelet.proto",
}
