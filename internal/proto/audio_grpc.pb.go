// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: internal/proto/audio.proto

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
	AudioService_GetAudioFile_FullMethodName       = "/audioservice.AudioService/GetAudioFile"
	AudioService_ReceiveNewsContent_FullMethodName = "/audioservice.AudioService/ReceiveNewsContent"
)

// AudioServiceClient is the client API for AudioService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AudioServiceClient interface {
	GetAudioFile(ctx context.Context, in *AudioRequest, opts ...grpc.CallOption) (*AudioResponse, error)
	ReceiveNewsContent(ctx context.Context, in *NewsContentRequest, opts ...grpc.CallOption) (*NewsContentResponse, error)
}

type audioServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAudioServiceClient(cc grpc.ClientConnInterface) AudioServiceClient {
	return &audioServiceClient{cc}
}

func (c *audioServiceClient) GetAudioFile(ctx context.Context, in *AudioRequest, opts ...grpc.CallOption) (*AudioResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AudioResponse)
	err := c.cc.Invoke(ctx, AudioService_GetAudioFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *audioServiceClient) ReceiveNewsContent(ctx context.Context, in *NewsContentRequest, opts ...grpc.CallOption) (*NewsContentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NewsContentResponse)
	err := c.cc.Invoke(ctx, AudioService_ReceiveNewsContent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AudioServiceServer is the server API for AudioService service.
// All implementations must embed UnimplementedAudioServiceServer
// for forward compatibility.
type AudioServiceServer interface {
	GetAudioFile(context.Context, *AudioRequest) (*AudioResponse, error)
	ReceiveNewsContent(context.Context, *NewsContentRequest) (*NewsContentResponse, error)
	mustEmbedUnimplementedAudioServiceServer()
}

// UnimplementedAudioServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAudioServiceServer struct{}

func (UnimplementedAudioServiceServer) GetAudioFile(context.Context, *AudioRequest) (*AudioResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAudioFile not implemented")
}
func (UnimplementedAudioServiceServer) ReceiveNewsContent(context.Context, *NewsContentRequest) (*NewsContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveNewsContent not implemented")
}
func (UnimplementedAudioServiceServer) mustEmbedUnimplementedAudioServiceServer() {}
func (UnimplementedAudioServiceServer) testEmbeddedByValue()                      {}

// UnsafeAudioServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AudioServiceServer will
// result in compilation errors.
type UnsafeAudioServiceServer interface {
	mustEmbedUnimplementedAudioServiceServer()
}

func RegisterAudioServiceServer(s grpc.ServiceRegistrar, srv AudioServiceServer) {
	// If the following call pancis, it indicates UnimplementedAudioServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AudioService_ServiceDesc, srv)
}

func _AudioService_GetAudioFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AudioRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudioServiceServer).GetAudioFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AudioService_GetAudioFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudioServiceServer).GetAudioFile(ctx, req.(*AudioRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AudioService_ReceiveNewsContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewsContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudioServiceServer).ReceiveNewsContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AudioService_ReceiveNewsContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudioServiceServer).ReceiveNewsContent(ctx, req.(*NewsContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AudioService_ServiceDesc is the grpc.ServiceDesc for AudioService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AudioService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "audioservice.AudioService",
	HandlerType: (*AudioServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAudioFile",
			Handler:    _AudioService_GetAudioFile_Handler,
		},
		{
			MethodName: "ReceiveNewsContent",
			Handler:    _AudioService_ReceiveNewsContent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/proto/audio.proto",
}
