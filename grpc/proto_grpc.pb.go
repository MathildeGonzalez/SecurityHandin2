// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: grpc/proto.proto

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
	ShareSendingService_SendShare_FullMethodName = "/HandinTwo.ShareSendingService/SendShare"
)

// ShareSendingServiceClient is the client API for ShareSendingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShareSendingServiceClient interface {
	SendShare(ctx context.Context, in *Share, opts ...grpc.CallOption) (*Acknowledge, error)
}

type shareSendingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShareSendingServiceClient(cc grpc.ClientConnInterface) ShareSendingServiceClient {
	return &shareSendingServiceClient{cc}
}

func (c *shareSendingServiceClient) SendShare(ctx context.Context, in *Share, opts ...grpc.CallOption) (*Acknowledge, error) {
	out := new(Acknowledge)
	err := c.cc.Invoke(ctx, ShareSendingService_SendShare_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShareSendingServiceServer is the server API for ShareSendingService service.
// All implementations must embed UnimplementedShareSendingServiceServer
// for forward compatibility
type ShareSendingServiceServer interface {
	SendShare(context.Context, *Share) (*Acknowledge, error)
	mustEmbedUnimplementedShareSendingServiceServer()
}

// UnimplementedShareSendingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedShareSendingServiceServer struct {
}

func (UnimplementedShareSendingServiceServer) SendShare(context.Context, *Share) (*Acknowledge, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendShare not implemented")
}
func (UnimplementedShareSendingServiceServer) mustEmbedUnimplementedShareSendingServiceServer() {}

// UnsafeShareSendingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShareSendingServiceServer will
// result in compilation errors.
type UnsafeShareSendingServiceServer interface {
	mustEmbedUnimplementedShareSendingServiceServer()
}

func RegisterShareSendingServiceServer(s grpc.ServiceRegistrar, srv ShareSendingServiceServer) {
	s.RegisterService(&ShareSendingService_ServiceDesc, srv)
}

func _ShareSendingService_SendShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Share)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShareSendingServiceServer).SendShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShareSendingService_SendShare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShareSendingServiceServer).SendShare(ctx, req.(*Share))
	}
	return interceptor(ctx, in, info, handler)
}

// ShareSendingService_ServiceDesc is the grpc.ServiceDesc for ShareSendingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShareSendingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "HandinTwo.ShareSendingService",
	HandlerType: (*ShareSendingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendShare",
			Handler:    _ShareSendingService_SendShare_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/proto.proto",
}

const (
	AggregatedShareService_SendAggregatedShare_FullMethodName = "/HandinTwo.AggregatedShareService/SendAggregatedShare"
)

// AggregatedShareServiceClient is the client API for AggregatedShareService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AggregatedShareServiceClient interface {
	SendAggregatedShare(ctx context.Context, in *AggregatedShare, opts ...grpc.CallOption) (*Acknowledge, error)
}

type aggregatedShareServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAggregatedShareServiceClient(cc grpc.ClientConnInterface) AggregatedShareServiceClient {
	return &aggregatedShareServiceClient{cc}
}

func (c *aggregatedShareServiceClient) SendAggregatedShare(ctx context.Context, in *AggregatedShare, opts ...grpc.CallOption) (*Acknowledge, error) {
	out := new(Acknowledge)
	err := c.cc.Invoke(ctx, AggregatedShareService_SendAggregatedShare_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AggregatedShareServiceServer is the server API for AggregatedShareService service.
// All implementations must embed UnimplementedAggregatedShareServiceServer
// for forward compatibility
type AggregatedShareServiceServer interface {
	SendAggregatedShare(context.Context, *AggregatedShare) (*Acknowledge, error)
	mustEmbedUnimplementedAggregatedShareServiceServer()
}

// UnimplementedAggregatedShareServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAggregatedShareServiceServer struct {
}

func (UnimplementedAggregatedShareServiceServer) SendAggregatedShare(context.Context, *AggregatedShare) (*Acknowledge, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendAggregatedShare not implemented")
}
func (UnimplementedAggregatedShareServiceServer) mustEmbedUnimplementedAggregatedShareServiceServer() {
}

// UnsafeAggregatedShareServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AggregatedShareServiceServer will
// result in compilation errors.
type UnsafeAggregatedShareServiceServer interface {
	mustEmbedUnimplementedAggregatedShareServiceServer()
}

func RegisterAggregatedShareServiceServer(s grpc.ServiceRegistrar, srv AggregatedShareServiceServer) {
	s.RegisterService(&AggregatedShareService_ServiceDesc, srv)
}

func _AggregatedShareService_SendAggregatedShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AggregatedShare)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatedShareServiceServer).SendAggregatedShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AggregatedShareService_SendAggregatedShare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatedShareServiceServer).SendAggregatedShare(ctx, req.(*AggregatedShare))
	}
	return interceptor(ctx, in, info, handler)
}

// AggregatedShareService_ServiceDesc is the grpc.ServiceDesc for AggregatedShareService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AggregatedShareService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "HandinTwo.AggregatedShareService",
	HandlerType: (*AggregatedShareServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendAggregatedShare",
			Handler:    _AggregatedShareService_SendAggregatedShare_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/proto.proto",
}
