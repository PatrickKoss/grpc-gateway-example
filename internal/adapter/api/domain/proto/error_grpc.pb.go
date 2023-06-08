// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: internal/adapter/api/domain/proto/error.proto

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
	ErrorService_CreateProduct_FullMethodName = "/grpc.gateway.example.internal.adapter.api.domain.ErrorService/CreateProduct"
)

// ErrorServiceClient is the client API for ErrorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ErrorServiceClient interface {
	CreateProduct(ctx context.Context, in *ErrorMessage, opts ...grpc.CallOption) (*ErrorMessage, error)
}

type errorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewErrorServiceClient(cc grpc.ClientConnInterface) ErrorServiceClient {
	return &errorServiceClient{cc}
}

func (c *errorServiceClient) CreateProduct(ctx context.Context, in *ErrorMessage, opts ...grpc.CallOption) (*ErrorMessage, error) {
	out := new(ErrorMessage)
	err := c.cc.Invoke(ctx, ErrorService_CreateProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ErrorServiceServer is the server API for ErrorService service.
// All implementations must embed UnimplementedErrorServiceServer
// for forward compatibility
type ErrorServiceServer interface {
	CreateProduct(context.Context, *ErrorMessage) (*ErrorMessage, error)
	mustEmbedUnimplementedErrorServiceServer()
}

// UnimplementedErrorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedErrorServiceServer struct {
}

func (UnimplementedErrorServiceServer) CreateProduct(context.Context, *ErrorMessage) (*ErrorMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedErrorServiceServer) mustEmbedUnimplementedErrorServiceServer() {}

// UnsafeErrorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ErrorServiceServer will
// result in compilation errors.
type UnsafeErrorServiceServer interface {
	mustEmbedUnimplementedErrorServiceServer()
}

func RegisterErrorServiceServer(s grpc.ServiceRegistrar, srv ErrorServiceServer) {
	s.RegisterService(&ErrorService_ServiceDesc, srv)
}

func _ErrorService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ErrorMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErrorServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ErrorService_CreateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErrorServiceServer).CreateProduct(ctx, req.(*ErrorMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// ErrorService_ServiceDesc is the grpc.ServiceDesc for ErrorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ErrorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gateway.example.internal.adapter.api.domain.ErrorService",
	HandlerType: (*ErrorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _ErrorService_CreateProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/adapter/api/domain/proto/error.proto",
}
