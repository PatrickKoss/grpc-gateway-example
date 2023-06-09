// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: examples/internal/proto/examplepb/non_standard_names.proto

package examplepb

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

// NonStandardServiceClient is the client API for NonStandardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NonStandardServiceClient interface {
	// Apply field mask to empty NonStandardMessage and return result.
	Update(ctx context.Context, in *NonStandardUpdateRequest, opts ...grpc.CallOption) (*NonStandardMessage, error)
	// Apply field mask to empty NonStandardMessageWithJSONNames and return result.
	UpdateWithJSONNames(ctx context.Context, in *NonStandardWithJSONNamesUpdateRequest, opts ...grpc.CallOption) (*NonStandardMessageWithJSONNames, error)
}

type nonStandardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNonStandardServiceClient(cc grpc.ClientConnInterface) NonStandardServiceClient {
	return &nonStandardServiceClient{cc}
}

func (c *nonStandardServiceClient) Update(ctx context.Context, in *NonStandardUpdateRequest, opts ...grpc.CallOption) (*NonStandardMessage, error) {
	out := new(NonStandardMessage)
	err := c.cc.Invoke(ctx, "/grpc.gateway.examples.internal.proto.examplepb.NonStandardService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nonStandardServiceClient) UpdateWithJSONNames(ctx context.Context, in *NonStandardWithJSONNamesUpdateRequest, opts ...grpc.CallOption) (*NonStandardMessageWithJSONNames, error) {
	out := new(NonStandardMessageWithJSONNames)
	err := c.cc.Invoke(ctx, "/grpc.gateway.examples.internal.proto.examplepb.NonStandardService/UpdateWithJSONNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NonStandardServiceServer is the server API for NonStandardService service.
// All implementations should embed UnimplementedNonStandardServiceServer
// for forward compatibility
type NonStandardServiceServer interface {
	// Apply field mask to empty NonStandardMessage and return result.
	Update(context.Context, *NonStandardUpdateRequest) (*NonStandardMessage, error)
	// Apply field mask to empty NonStandardMessageWithJSONNames and return result.
	UpdateWithJSONNames(context.Context, *NonStandardWithJSONNamesUpdateRequest) (*NonStandardMessageWithJSONNames, error)
}

// UnimplementedNonStandardServiceServer should be embedded to have forward compatible implementations.
type UnimplementedNonStandardServiceServer struct {
}

func (UnimplementedNonStandardServiceServer) Update(context.Context, *NonStandardUpdateRequest) (*NonStandardMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedNonStandardServiceServer) UpdateWithJSONNames(context.Context, *NonStandardWithJSONNamesUpdateRequest) (*NonStandardMessageWithJSONNames, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWithJSONNames not implemented")
}

// UnsafeNonStandardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NonStandardServiceServer will
// result in compilation errors.
type UnsafeNonStandardServiceServer interface {
	mustEmbedUnimplementedNonStandardServiceServer()
}

func RegisterNonStandardServiceServer(s grpc.ServiceRegistrar, srv NonStandardServiceServer) {
	s.RegisterService(&NonStandardService_ServiceDesc, srv)
}

func _NonStandardService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NonStandardUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NonStandardServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.internal.proto.examplepb.NonStandardService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NonStandardServiceServer).Update(ctx, req.(*NonStandardUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NonStandardService_UpdateWithJSONNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NonStandardWithJSONNamesUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NonStandardServiceServer).UpdateWithJSONNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.internal.proto.examplepb.NonStandardService/UpdateWithJSONNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NonStandardServiceServer).UpdateWithJSONNames(ctx, req.(*NonStandardWithJSONNamesUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NonStandardService_ServiceDesc is the grpc.ServiceDesc for NonStandardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NonStandardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gateway.examples.internal.proto.examplepb.NonStandardService",
	HandlerType: (*NonStandardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Update",
			Handler:    _NonStandardService_Update_Handler,
		},
		{
			MethodName: "UpdateWithJSONNames",
			Handler:    _NonStandardService_UpdateWithJSONNames_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "examples/internal/proto/examplepb/non_standard_names.proto",
}
