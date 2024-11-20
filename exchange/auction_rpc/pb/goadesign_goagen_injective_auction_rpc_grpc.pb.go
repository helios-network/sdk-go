// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: goadesign_goagen_helios_auction_rpc.proto

package helios_auction_rpcpb

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

// HeliosAuctionRPCClient is the client API for HeliosAuctionRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HeliosAuctionRPCClient interface {
	// Provide historical auction info for a given auction
	AuctionEndpoint(ctx context.Context, in *AuctionEndpointRequest, opts ...grpc.CallOption) (*AuctionEndpointResponse, error)
	// Provide the historical auctions info
	Auctions(ctx context.Context, in *AuctionsRequest, opts ...grpc.CallOption) (*AuctionsResponse, error)
	// StreamBids streams new bids of an auction.
	StreamBids(ctx context.Context, in *StreamBidsRequest, opts ...grpc.CallOption) (HeliosAuctionRPC_StreamBidsClient, error)
}

type heliosAuctionRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewHeliosAuctionRPCClient(cc grpc.ClientConnInterface) HeliosAuctionRPCClient {
	return &heliosAuctionRPCClient{cc}
}

func (c *heliosAuctionRPCClient) AuctionEndpoint(ctx context.Context, in *AuctionEndpointRequest, opts ...grpc.CallOption) (*AuctionEndpointResponse, error) {
	out := new(AuctionEndpointResponse)
	err := c.cc.Invoke(ctx, "/helios_auction_rpc.HeliosAuctionRPC/AuctionEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heliosAuctionRPCClient) Auctions(ctx context.Context, in *AuctionsRequest, opts ...grpc.CallOption) (*AuctionsResponse, error) {
	out := new(AuctionsResponse)
	err := c.cc.Invoke(ctx, "/helios_auction_rpc.HeliosAuctionRPC/Auctions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heliosAuctionRPCClient) StreamBids(ctx context.Context, in *StreamBidsRequest, opts ...grpc.CallOption) (HeliosAuctionRPC_StreamBidsClient, error) {
	stream, err := c.cc.NewStream(ctx, &HeliosAuctionRPC_ServiceDesc.Streams[0], "/helios_auction_rpc.HeliosAuctionRPC/StreamBids", opts...)
	if err != nil {
		return nil, err
	}
	x := &heliosAuctionRPCStreamBidsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HeliosAuctionRPC_StreamBidsClient interface {
	Recv() (*StreamBidsResponse, error)
	grpc.ClientStream
}

type heliosAuctionRPCStreamBidsClient struct {
	grpc.ClientStream
}

func (x *heliosAuctionRPCStreamBidsClient) Recv() (*StreamBidsResponse, error) {
	m := new(StreamBidsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HeliosAuctionRPCServer is the server API for HeliosAuctionRPC service.
// All implementations must embed UnimplementedHeliosAuctionRPCServer
// for forward compatibility
type HeliosAuctionRPCServer interface {
	// Provide historical auction info for a given auction
	AuctionEndpoint(context.Context, *AuctionEndpointRequest) (*AuctionEndpointResponse, error)
	// Provide the historical auctions info
	Auctions(context.Context, *AuctionsRequest) (*AuctionsResponse, error)
	// StreamBids streams new bids of an auction.
	StreamBids(*StreamBidsRequest, HeliosAuctionRPC_StreamBidsServer) error
	mustEmbedUnimplementedHeliosAuctionRPCServer()
}

// UnimplementedHeliosAuctionRPCServer must be embedded to have forward compatible implementations.
type UnimplementedHeliosAuctionRPCServer struct {
}

func (UnimplementedHeliosAuctionRPCServer) AuctionEndpoint(context.Context, *AuctionEndpointRequest) (*AuctionEndpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuctionEndpoint not implemented")
}
func (UnimplementedHeliosAuctionRPCServer) Auctions(context.Context, *AuctionsRequest) (*AuctionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auctions not implemented")
}
func (UnimplementedHeliosAuctionRPCServer) StreamBids(*StreamBidsRequest, HeliosAuctionRPC_StreamBidsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamBids not implemented")
}
func (UnimplementedHeliosAuctionRPCServer) mustEmbedUnimplementedHeliosAuctionRPCServer() {}

// UnsafeHeliosAuctionRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HeliosAuctionRPCServer will
// result in compilation errors.
type UnsafeHeliosAuctionRPCServer interface {
	mustEmbedUnimplementedHeliosAuctionRPCServer()
}

func RegisterHeliosAuctionRPCServer(s grpc.ServiceRegistrar, srv HeliosAuctionRPCServer) {
	s.RegisterService(&HeliosAuctionRPC_ServiceDesc, srv)
}

func _HeliosAuctionRPC_AuctionEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuctionEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeliosAuctionRPCServer).AuctionEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helios_auction_rpc.HeliosAuctionRPC/AuctionEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeliosAuctionRPCServer).AuctionEndpoint(ctx, req.(*AuctionEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeliosAuctionRPC_Auctions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuctionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeliosAuctionRPCServer).Auctions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helios_auction_rpc.HeliosAuctionRPC/Auctions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeliosAuctionRPCServer).Auctions(ctx, req.(*AuctionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeliosAuctionRPC_StreamBids_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamBidsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HeliosAuctionRPCServer).StreamBids(m, &heliosAuctionRPCStreamBidsServer{stream})
}

type HeliosAuctionRPC_StreamBidsServer interface {
	Send(*StreamBidsResponse) error
	grpc.ServerStream
}

type heliosAuctionRPCStreamBidsServer struct {
	grpc.ServerStream
}

func (x *heliosAuctionRPCStreamBidsServer) Send(m *StreamBidsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// HeliosAuctionRPC_ServiceDesc is the grpc.ServiceDesc for HeliosAuctionRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HeliosAuctionRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helios_auction_rpc.HeliosAuctionRPC",
	HandlerType: (*HeliosAuctionRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuctionEndpoint",
			Handler:    _HeliosAuctionRPC_AuctionEndpoint_Handler,
		},
		{
			MethodName: "Auctions",
			Handler:    _HeliosAuctionRPC_Auctions_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamBids",
			Handler:       _HeliosAuctionRPC_StreamBids_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "goadesign_goagen_helios_auction_rpc.proto",
}
