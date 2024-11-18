// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: goadesign_goagen_helios_oracle_rpc.proto

package helios_oracle_rpcpb

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

// InjectiveOracleRPCClient is the client API for InjectiveOracleRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InjectiveOracleRPCClient interface {
	// List all oracles
	OracleList(ctx context.Context, in *OracleListRequest, opts ...grpc.CallOption) (*OracleListResponse, error)
	// Gets the price of the oracle
	Price(ctx context.Context, in *PriceRequest, opts ...grpc.CallOption) (*PriceResponse, error)
	// StreamPrices streams new price changes for a specified oracle. If no oracles
	// are provided, all price changes are streamed.
	StreamPrices(ctx context.Context, in *StreamPricesRequest, opts ...grpc.CallOption) (InjectiveOracleRPC_StreamPricesClient, error)
	// StreamPrices streams new price changes markets
	StreamPricesByMarkets(ctx context.Context, in *StreamPricesByMarketsRequest, opts ...grpc.CallOption) (InjectiveOracleRPC_StreamPricesByMarketsClient, error)
}

type heliosOracleRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewInjectiveOracleRPCClient(cc grpc.ClientConnInterface) InjectiveOracleRPCClient {
	return &heliosOracleRPCClient{cc}
}

func (c *heliosOracleRPCClient) OracleList(ctx context.Context, in *OracleListRequest, opts ...grpc.CallOption) (*OracleListResponse, error) {
	out := new(OracleListResponse)
	err := c.cc.Invoke(ctx, "/helios_oracle_rpc.InjectiveOracleRPC/OracleList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heliosOracleRPCClient) Price(ctx context.Context, in *PriceRequest, opts ...grpc.CallOption) (*PriceResponse, error) {
	out := new(PriceResponse)
	err := c.cc.Invoke(ctx, "/helios_oracle_rpc.InjectiveOracleRPC/Price", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heliosOracleRPCClient) StreamPrices(ctx context.Context, in *StreamPricesRequest, opts ...grpc.CallOption) (InjectiveOracleRPC_StreamPricesClient, error) {
	stream, err := c.cc.NewStream(ctx, &InjectiveOracleRPC_ServiceDesc.Streams[0], "/helios_oracle_rpc.InjectiveOracleRPC/StreamPrices", opts...)
	if err != nil {
		return nil, err
	}
	x := &heliosOracleRPCStreamPricesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InjectiveOracleRPC_StreamPricesClient interface {
	Recv() (*StreamPricesResponse, error)
	grpc.ClientStream
}

type heliosOracleRPCStreamPricesClient struct {
	grpc.ClientStream
}

func (x *heliosOracleRPCStreamPricesClient) Recv() (*StreamPricesResponse, error) {
	m := new(StreamPricesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *heliosOracleRPCClient) StreamPricesByMarkets(ctx context.Context, in *StreamPricesByMarketsRequest, opts ...grpc.CallOption) (InjectiveOracleRPC_StreamPricesByMarketsClient, error) {
	stream, err := c.cc.NewStream(ctx, &InjectiveOracleRPC_ServiceDesc.Streams[1], "/helios_oracle_rpc.InjectiveOracleRPC/StreamPricesByMarkets", opts...)
	if err != nil {
		return nil, err
	}
	x := &heliosOracleRPCStreamPricesByMarketsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InjectiveOracleRPC_StreamPricesByMarketsClient interface {
	Recv() (*StreamPricesByMarketsResponse, error)
	grpc.ClientStream
}

type heliosOracleRPCStreamPricesByMarketsClient struct {
	grpc.ClientStream
}

func (x *heliosOracleRPCStreamPricesByMarketsClient) Recv() (*StreamPricesByMarketsResponse, error) {
	m := new(StreamPricesByMarketsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// InjectiveOracleRPCServer is the server API for InjectiveOracleRPC service.
// All implementations must embed UnimplementedInjectiveOracleRPCServer
// for forward compatibility
type InjectiveOracleRPCServer interface {
	// List all oracles
	OracleList(context.Context, *OracleListRequest) (*OracleListResponse, error)
	// Gets the price of the oracle
	Price(context.Context, *PriceRequest) (*PriceResponse, error)
	// StreamPrices streams new price changes for a specified oracle. If no oracles
	// are provided, all price changes are streamed.
	StreamPrices(*StreamPricesRequest, InjectiveOracleRPC_StreamPricesServer) error
	// StreamPrices streams new price changes markets
	StreamPricesByMarkets(*StreamPricesByMarketsRequest, InjectiveOracleRPC_StreamPricesByMarketsServer) error
	mustEmbedUnimplementedInjectiveOracleRPCServer()
}

// UnimplementedInjectiveOracleRPCServer must be embedded to have forward compatible implementations.
type UnimplementedInjectiveOracleRPCServer struct {
}

func (UnimplementedInjectiveOracleRPCServer) OracleList(context.Context, *OracleListRequest) (*OracleListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OracleList not implemented")
}
func (UnimplementedInjectiveOracleRPCServer) Price(context.Context, *PriceRequest) (*PriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Price not implemented")
}
func (UnimplementedInjectiveOracleRPCServer) StreamPrices(*StreamPricesRequest, InjectiveOracleRPC_StreamPricesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamPrices not implemented")
}
func (UnimplementedInjectiveOracleRPCServer) StreamPricesByMarkets(*StreamPricesByMarketsRequest, InjectiveOracleRPC_StreamPricesByMarketsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamPricesByMarkets not implemented")
}
func (UnimplementedInjectiveOracleRPCServer) mustEmbedUnimplementedInjectiveOracleRPCServer() {}

// UnsafeInjectiveOracleRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InjectiveOracleRPCServer will
// result in compilation errors.
type UnsafeInjectiveOracleRPCServer interface {
	mustEmbedUnimplementedInjectiveOracleRPCServer()
}

func RegisterInjectiveOracleRPCServer(s grpc.ServiceRegistrar, srv InjectiveOracleRPCServer) {
	s.RegisterService(&InjectiveOracleRPC_ServiceDesc, srv)
}

func _InjectiveOracleRPC_OracleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OracleListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveOracleRPCServer).OracleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helios_oracle_rpc.InjectiveOracleRPC/OracleList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveOracleRPCServer).OracleList(ctx, req.(*OracleListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveOracleRPC_Price_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveOracleRPCServer).Price(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helios_oracle_rpc.InjectiveOracleRPC/Price",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveOracleRPCServer).Price(ctx, req.(*PriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveOracleRPC_StreamPrices_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamPricesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InjectiveOracleRPCServer).StreamPrices(m, &heliosOracleRPCStreamPricesServer{stream})
}

type InjectiveOracleRPC_StreamPricesServer interface {
	Send(*StreamPricesResponse) error
	grpc.ServerStream
}

type heliosOracleRPCStreamPricesServer struct {
	grpc.ServerStream
}

func (x *heliosOracleRPCStreamPricesServer) Send(m *StreamPricesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _InjectiveOracleRPC_StreamPricesByMarkets_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamPricesByMarketsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InjectiveOracleRPCServer).StreamPricesByMarkets(m, &heliosOracleRPCStreamPricesByMarketsServer{stream})
}

type InjectiveOracleRPC_StreamPricesByMarketsServer interface {
	Send(*StreamPricesByMarketsResponse) error
	grpc.ServerStream
}

type heliosOracleRPCStreamPricesByMarketsServer struct {
	grpc.ServerStream
}

func (x *heliosOracleRPCStreamPricesByMarketsServer) Send(m *StreamPricesByMarketsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// InjectiveOracleRPC_ServiceDesc is the grpc.ServiceDesc for InjectiveOracleRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InjectiveOracleRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helios_oracle_rpc.InjectiveOracleRPC",
	HandlerType: (*InjectiveOracleRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OracleList",
			Handler:    _InjectiveOracleRPC_OracleList_Handler,
		},
		{
			MethodName: "Price",
			Handler:    _InjectiveOracleRPC_Price_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamPrices",
			Handler:       _InjectiveOracleRPC_StreamPrices_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamPricesByMarkets",
			Handler:       _InjectiveOracleRPC_StreamPricesByMarkets_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "goadesign_goagen_helios_oracle_rpc.proto",
}
