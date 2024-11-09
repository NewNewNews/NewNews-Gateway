// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: internal/proto/compare.proto

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
	ComparisonService_GetComparison_FullMethodName = "/comparisonservice.ComparisonService/GetComparison"
)

// ComparisonServiceClient is the client API for ComparisonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComparisonServiceClient interface {
	GetComparison(ctx context.Context, in *GetComparisonRequest, opts ...grpc.CallOption) (*Comparison, error)
}

type comparisonServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewComparisonServiceClient(cc grpc.ClientConnInterface) ComparisonServiceClient {
	return &comparisonServiceClient{cc}
}

func (c *comparisonServiceClient) GetComparison(ctx context.Context, in *GetComparisonRequest, opts ...grpc.CallOption) (*Comparison, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Comparison)
	err := c.cc.Invoke(ctx, ComparisonService_GetComparison_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComparisonServiceServer is the server API for ComparisonService service.
// All implementations must embed UnimplementedComparisonServiceServer
// for forward compatibility.
type ComparisonServiceServer interface {
	GetComparison(context.Context, *GetComparisonRequest) (*Comparison, error)
	mustEmbedUnimplementedComparisonServiceServer()
}

// UnimplementedComparisonServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedComparisonServiceServer struct{}

func (UnimplementedComparisonServiceServer) GetComparison(context.Context, *GetComparisonRequest) (*Comparison, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComparison not implemented")
}
func (UnimplementedComparisonServiceServer) mustEmbedUnimplementedComparisonServiceServer() {}
func (UnimplementedComparisonServiceServer) testEmbeddedByValue()                           {}

// UnsafeComparisonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComparisonServiceServer will
// result in compilation errors.
type UnsafeComparisonServiceServer interface {
	mustEmbedUnimplementedComparisonServiceServer()
}

func RegisterComparisonServiceServer(s grpc.ServiceRegistrar, srv ComparisonServiceServer) {
	// If the following call pancis, it indicates UnimplementedComparisonServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ComparisonService_ServiceDesc, srv)
}

func _ComparisonService_GetComparison_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetComparisonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComparisonServiceServer).GetComparison(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComparisonService_GetComparison_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComparisonServiceServer).GetComparison(ctx, req.(*GetComparisonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ComparisonService_ServiceDesc is the grpc.ServiceDesc for ComparisonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ComparisonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comparisonservice.ComparisonService",
	HandlerType: (*ComparisonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetComparison",
			Handler:    _ComparisonService_GetComparison_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/proto/compare.proto",
}