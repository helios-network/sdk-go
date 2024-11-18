// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: goadesign_goagen_helios_spot_exchange_rpc.proto

package helios_spot_exchange_rpcpb

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

// InjectiveSpotExchangeRPCClient is the client API for InjectiveSpotExchangeRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InjectiveSpotExchangeRPCClient interface {
	// Get a list of Spot Markets
	Markets(ctx context.Context, in *MarketsRequest, opts ...grpc.CallOption) (*MarketsResponse, error)
	// Get details of a single spot market
	Market(ctx context.Context, in *MarketRequest, opts ...grpc.CallOption) (*MarketResponse, error)
	// Stream live updates of selected spot markets
	StreamMarkets(ctx context.Context, in *StreamMarketsRequest, opts ...grpc.CallOption) (InjectiveSpotExchangeRPC_StreamMarketsClient, error)
	// Orderbook of a Spot Market
	OrderbookV2(ctx context.Context, in *OrderbookV2Request, opts ...grpc.CallOption) (*OrderbookV2Response, error)
	// Orderbook of Spot Markets
	OrderbooksV2(ctx context.Context, in *OrderbooksV2Request, opts ...grpc.CallOption) (*OrderbooksV2Response, error)
	// Stream live snapshot updates of selected spot market orderbook
	StreamOrderbookV2(ctx context.Context, in *StreamOrderbookV2Request, opts ...grpc.CallOption) (InjectiveSpotExchangeRPC_StreamOrderbookV2Client, error)
	// Stream live level updates of selected spot market orderbook
	StreamOrderbookUpdate(ctx context.Context, in *StreamOrderbookUpdateRequest, opts ...grpc.CallOption) (InjectiveSpotExchangeRPC_StreamOrderbookUpdateClient, error)
	// Orders of a Spot Market
	Orders(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (*OrdersResponse, error)
	// Stream updates to individual orders of a Spot Market
	StreamOrders(ctx context.Context, in *StreamOrdersRequest, opts ...grpc.CallOption) (InjectiveSpotExchangeRPC_StreamOrdersClient, error)
	// Trades of a Spot Market
	Trades(ctx context.Context, in *TradesRequest, opts ...grpc.CallOption) (*TradesResponse, error)
	// Stream newly executed trades from Spot Market
	StreamTrades(ctx context.Context, in *StreamTradesRequest, opts ...grpc.CallOption) (InjectiveSpotExchangeRPC_StreamTradesClient, error)
	// Trades of a Spot Market
	TradesV2(ctx context.Context, in *TradesV2Request, opts ...grpc.CallOption) (*TradesV2Response, error)
	// Stream newly executed trades from Spot Market
	StreamTradesV2(ctx context.Context, in *StreamTradesV2Request, opts ...grpc.CallOption) (InjectiveSpotExchangeRPC_StreamTradesV2Client, error)
	// List orders posted from this subaccount
	SubaccountOrdersList(ctx context.Context, in *SubaccountOrdersListRequest, opts ...grpc.CallOption) (*SubaccountOrdersListResponse, error)
	// List trades executed by this subaccount
	SubaccountTradesList(ctx context.Context, in *SubaccountTradesListRequest, opts ...grpc.CallOption) (*SubaccountTradesListResponse, error)
	// Lists history orders posted from this subaccount
	OrdersHistory(ctx context.Context, in *OrdersHistoryRequest, opts ...grpc.CallOption) (*OrdersHistoryResponse, error)
	// Stream updates to historical orders of a spot Market
	StreamOrdersHistory(ctx context.Context, in *StreamOrdersHistoryRequest, opts ...grpc.CallOption) (InjectiveSpotExchangeRPC_StreamOrdersHistoryClient, error)
	// Get historical atomic swaps
	AtomicSwapHistory(ctx context.Context, in *AtomicSwapHistoryRequest, opts ...grpc.CallOption) (*AtomicSwapHistoryResponse, error)
}

type heliosSpotExchangeRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewInjectiveSpotExchangeRPCClient(cc grpc.ClientConnInterface) InjectiveSpotExchangeRPCClient {
	return &heliosSpotExchangeRPCClient{cc}
}

func (c *heliosSpotExchangeRPCClient) Markets(ctx context.Context, in *MarketsRequest, opts ...grpc.CallOption) (*MarketsResponse, error) {
	out := new(MarketsResponse)
	err := c.cc.Invoke(ctx, "/helios_spot_exchange_rpc.InjectiveSpotExchangeRPC/Markets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heliosSpotExchangeRPCClient) Market(ctx context.Context, in *MarketRequest, opts ...grpc.CallOption) (*MarketResponse, error) {
	out := new(MarketResponse)
	err := c.cc.Invoke(ctx, "/helios_spot_exchange_rpc.InjectiveSpotExchangeRPC/Market", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heliosSpotExchangeRPCClient) StreamMarkets(ctx context.Context, in *StreamMarketsRequest, opts ...grpc.CallOption) (InjectiveSpotExchangeRPC_StreamMarketsClient, error) {
	stream, err := c.cc.NewStream(ctx, &InjectiveSpotExchangeRPC_ServiceDesc.Streams[0], "/helios_spot_exchange_rpc.InjectiveSpotExchangeRPC/StreamMarkets", opts...)
	if err != nil {
		return nil, err
	}
	x := &heliosSpotExchangeRPCStreamMarketsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InjectiveSpotExchangeRPC_StreamMarketsClient interface {
	Recv() (*StreamMarketsResponse, error)
	grpc.ClientStream
}

type heliosSpotExchangeRPCStreamMarketsClient struct {
	grpc.ClientStream
}

func (x *heliosSpotExchangeRPCStreamMarketsClient) Recv() (*StreamMarketsResponse, error) {
	m := new(StreamMarketsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *heliosSpotExchangeRPCClient) OrderbookV2(ctx context.Context, in *OrderbookV2Request, opts ...grpc.CallOption) (*OrderbookV2Response, error) {
	out := new(OrderbookV2Response)
	err := c.cc.Invoke(ctx, "/helios_spot_exchange_rpc.InjectiveSpotExchangeRPC/OrderbookV2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heliosSpotExchangeRPCClient) OrderbooksV2(ctx context.Context, in *OrderbooksV2Request, opts ...grpc.CallOption) (*OrderbooksV2Response, error) {
	out := new(OrderbooksV2Response)
	err := c.cc.Invoke(ctx, "/helios_spot_exchange_rpc.InjectiveSpotExchangeRPC/OrderbooksV2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heliosSpotExchangeRPCClient) StreamOrderbookV2(ctx context.Context, in *StreamOrderbookV2Request, opts ...grpc.CallOption) (InjectiveSpotExchangeRPC_StreamOrderbookV2Client, error) {
	stream, err := c.cc.NewStream(ctx, &InjectiveSpotExchangeRPC_ServiceDesc.Streams[1], "/helios_spot_exchange_rpc.InjectiveSpotExchangeRPC/StreamOrderbookV2", opts...)
	if err != nil {
		return nil, err
	}
	x := &heliosSpotExchangeRPCStreamOrderbookV2Client{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InjectiveSpotExchangeRPC_StreamOrderbookV2Client interface {
	Recv() (*StreamOrderbookV2Response, error)
	grpc.ClientStream
}

type heliosSpotExchangeRPCStreamOrderbookV2Client struct {
	grpc.ClientStream
}

func (x *heliosSpotExchangeRPCStreamOrderbookV2Client) Recv() (*StreamOrderbookV2Response, error) {
	m := new(StreamOrderbookV2Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *heliosSpotExchangeRPCClient) StreamOrderbookUpdate(ctx context.Context, in *StreamOrderbookUpdateRequest, opts ...grpc.CallOption) (InjectiveSpotExchangeRPC_StreamOrderbookUpdateClient, error) {
	stream, err := c.cc.NewStream(ctx, &InjectiveSpotExchangeRPC_ServiceDesc.Streams[2], "/helios_spot_exchange_rpc.InjectiveSpotExchangeRPC/StreamOrderbookUpdate", opts...)
	if err != nil {
		return nil, err
	}
	x := &heliosSpotExchangeRPCStreamOrderbookUpdateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InjectiveSpotExchangeRPC_StreamOrderbookUpdateClient interface {
	Recv() (*StreamOrderbookUpdateResponse, error)
	grpc.ClientStream
}

type heliosSpotExchangeRPCStreamOrderbookUpdateClient struct {
	grpc.ClientStream
}

func (x *heliosSpotExchangeRPCStreamOrderbookUpdateClient) Recv() (*StreamOrderbookUpdateResponse, error) {
	m := new(StreamOrderbookUpdateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *heliosSpotExchangeRPCClient) Orders(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (*OrdersResponse, error) {
	out := new(OrdersResponse)
	err := c.cc.Invoke(ctx, "/helios_spot_exchange_rpc.InjectiveSpotExchangeRPC/Orders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heliosSpotExchangeRPCClient) StreamOrders(ctx context.Context, in *StreamOrdersRequest, opts ...grpc.CallOption) (InjectiveSpotExchangeRPC_StreamOrdersClient, error) {
	stream, err := c.cc.NewStream(ctx, &InjectiveSpotExchangeRPC_ServiceDesc.Streams[3], "/helios_spot_exchange_rpc.InjectiveSpotExchangeRPC/StreamOrders", opts...)
	if err != nil {
		return nil, err
	}
	x := &heliosSpotExchangeRPCStreamOrdersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InjectiveSpotExchangeRPC_StreamOrdersClient interface {
	Recv() (*StreamOrdersResponse, error)
	grpc.ClientStream
}

type heliosSpotExchangeRPCStreamOrdersClient struct {
	grpc.ClientStream
}

func (x *heliosSpotExchangeRPCStreamOrdersClient) Recv() (*StreamOrdersResponse, error) {
	m := new(StreamOrdersResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *heliosSpotExchangeRPCClient) Trades(ctx context.Context, in *TradesRequest, opts ...grpc.CallOption) (*TradesResponse, error) {
	out := new(TradesResponse)
	err := c.cc.Invoke(ctx, "/helios_spot_exchange_rpc.InjectiveSpotExchangeRPC/Trades", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heliosSpotExchangeRPCClient) StreamTrades(ctx context.Context, in *StreamTradesRequest, opts ...grpc.CallOption) (InjectiveSpotExchangeRPC_StreamTradesClient, error) {
	stream, err := c.cc.NewStream(ctx, &InjectiveSpotExchangeRPC_ServiceDesc.Streams[4], "/helios_spot_exchange_rpc.InjectiveSpotExchangeRPC/StreamTrades", opts...)
	if err != nil {
		return nil, err
	}
	x := &heliosSpotExchangeRPCStreamTradesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InjectiveSpotExchangeRPC_StreamTradesClient interface {
	Recv() (*StreamTradesResponse, error)
	grpc.ClientStream
}

type heliosSpotExchangeRPCStreamTradesClient struct {
	grpc.ClientStream
}

func (x *heliosSpotExchangeRPCStreamTradesClient) Recv() (*StreamTradesResponse, error) {
	m := new(StreamTradesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *heliosSpotExchangeRPCClient) TradesV2(ctx context.Context, in *TradesV2Request, opts ...grpc.CallOption) (*TradesV2Response, error) {
	out := new(TradesV2Response)
	err := c.cc.Invoke(ctx, "/helios_spot_exchange_rpc.InjectiveSpotExchangeRPC/TradesV2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heliosSpotExchangeRPCClient) StreamTradesV2(ctx context.Context, in *StreamTradesV2Request, opts ...grpc.CallOption) (InjectiveSpotExchangeRPC_StreamTradesV2Client, error) {
	stream, err := c.cc.NewStream(ctx, &InjectiveSpotExchangeRPC_ServiceDesc.Streams[5], "/helios_spot_exchange_rpc.InjectiveSpotExchangeRPC/StreamTradesV2", opts...)
	if err != nil {
		return nil, err
	}
	x := &heliosSpotExchangeRPCStreamTradesV2Client{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InjectiveSpotExchangeRPC_StreamTradesV2Client interface {
	Recv() (*StreamTradesV2Response, error)
	grpc.ClientStream
}

type heliosSpotExchangeRPCStreamTradesV2Client struct {
	grpc.ClientStream
}

func (x *heliosSpotExchangeRPCStreamTradesV2Client) Recv() (*StreamTradesV2Response, error) {
	m := new(StreamTradesV2Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *heliosSpotExchangeRPCClient) SubaccountOrdersList(ctx context.Context, in *SubaccountOrdersListRequest, opts ...grpc.CallOption) (*SubaccountOrdersListResponse, error) {
	out := new(SubaccountOrdersListResponse)
	err := c.cc.Invoke(ctx, "/helios_spot_exchange_rpc.InjectiveSpotExchangeRPC/SubaccountOrdersList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heliosSpotExchangeRPCClient) SubaccountTradesList(ctx context.Context, in *SubaccountTradesListRequest, opts ...grpc.CallOption) (*SubaccountTradesListResponse, error) {
	out := new(SubaccountTradesListResponse)
	err := c.cc.Invoke(ctx, "/helios_spot_exchange_rpc.InjectiveSpotExchangeRPC/SubaccountTradesList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heliosSpotExchangeRPCClient) OrdersHistory(ctx context.Context, in *OrdersHistoryRequest, opts ...grpc.CallOption) (*OrdersHistoryResponse, error) {
	out := new(OrdersHistoryResponse)
	err := c.cc.Invoke(ctx, "/helios_spot_exchange_rpc.InjectiveSpotExchangeRPC/OrdersHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heliosSpotExchangeRPCClient) StreamOrdersHistory(ctx context.Context, in *StreamOrdersHistoryRequest, opts ...grpc.CallOption) (InjectiveSpotExchangeRPC_StreamOrdersHistoryClient, error) {
	stream, err := c.cc.NewStream(ctx, &InjectiveSpotExchangeRPC_ServiceDesc.Streams[6], "/helios_spot_exchange_rpc.InjectiveSpotExchangeRPC/StreamOrdersHistory", opts...)
	if err != nil {
		return nil, err
	}
	x := &heliosSpotExchangeRPCStreamOrdersHistoryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InjectiveSpotExchangeRPC_StreamOrdersHistoryClient interface {
	Recv() (*StreamOrdersHistoryResponse, error)
	grpc.ClientStream
}

type heliosSpotExchangeRPCStreamOrdersHistoryClient struct {
	grpc.ClientStream
}

func (x *heliosSpotExchangeRPCStreamOrdersHistoryClient) Recv() (*StreamOrdersHistoryResponse, error) {
	m := new(StreamOrdersHistoryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *heliosSpotExchangeRPCClient) AtomicSwapHistory(ctx context.Context, in *AtomicSwapHistoryRequest, opts ...grpc.CallOption) (*AtomicSwapHistoryResponse, error) {
	out := new(AtomicSwapHistoryResponse)
	err := c.cc.Invoke(ctx, "/helios_spot_exchange_rpc.InjectiveSpotExchangeRPC/AtomicSwapHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InjectiveSpotExchangeRPCServer is the server API for InjectiveSpotExchangeRPC service.
// All implementations must embed UnimplementedInjectiveSpotExchangeRPCServer
// for forward compatibility
type InjectiveSpotExchangeRPCServer interface {
	// Get a list of Spot Markets
	Markets(context.Context, *MarketsRequest) (*MarketsResponse, error)
	// Get details of a single spot market
	Market(context.Context, *MarketRequest) (*MarketResponse, error)
	// Stream live updates of selected spot markets
	StreamMarkets(*StreamMarketsRequest, InjectiveSpotExchangeRPC_StreamMarketsServer) error
	// Orderbook of a Spot Market
	OrderbookV2(context.Context, *OrderbookV2Request) (*OrderbookV2Response, error)
	// Orderbook of Spot Markets
	OrderbooksV2(context.Context, *OrderbooksV2Request) (*OrderbooksV2Response, error)
	// Stream live snapshot updates of selected spot market orderbook
	StreamOrderbookV2(*StreamOrderbookV2Request, InjectiveSpotExchangeRPC_StreamOrderbookV2Server) error
	// Stream live level updates of selected spot market orderbook
	StreamOrderbookUpdate(*StreamOrderbookUpdateRequest, InjectiveSpotExchangeRPC_StreamOrderbookUpdateServer) error
	// Orders of a Spot Market
	Orders(context.Context, *OrdersRequest) (*OrdersResponse, error)
	// Stream updates to individual orders of a Spot Market
	StreamOrders(*StreamOrdersRequest, InjectiveSpotExchangeRPC_StreamOrdersServer) error
	// Trades of a Spot Market
	Trades(context.Context, *TradesRequest) (*TradesResponse, error)
	// Stream newly executed trades from Spot Market
	StreamTrades(*StreamTradesRequest, InjectiveSpotExchangeRPC_StreamTradesServer) error
	// Trades of a Spot Market
	TradesV2(context.Context, *TradesV2Request) (*TradesV2Response, error)
	// Stream newly executed trades from Spot Market
	StreamTradesV2(*StreamTradesV2Request, InjectiveSpotExchangeRPC_StreamTradesV2Server) error
	// List orders posted from this subaccount
	SubaccountOrdersList(context.Context, *SubaccountOrdersListRequest) (*SubaccountOrdersListResponse, error)
	// List trades executed by this subaccount
	SubaccountTradesList(context.Context, *SubaccountTradesListRequest) (*SubaccountTradesListResponse, error)
	// Lists history orders posted from this subaccount
	OrdersHistory(context.Context, *OrdersHistoryRequest) (*OrdersHistoryResponse, error)
	// Stream updates to historical orders of a spot Market
	StreamOrdersHistory(*StreamOrdersHistoryRequest, InjectiveSpotExchangeRPC_StreamOrdersHistoryServer) error
	// Get historical atomic swaps
	AtomicSwapHistory(context.Context, *AtomicSwapHistoryRequest) (*AtomicSwapHistoryResponse, error)
	mustEmbedUnimplementedInjectiveSpotExchangeRPCServer()
}

// UnimplementedInjectiveSpotExchangeRPCServer must be embedded to have forward compatible implementations.
type UnimplementedInjectiveSpotExchangeRPCServer struct {
}

func (UnimplementedInjectiveSpotExchangeRPCServer) Markets(context.Context, *MarketsRequest) (*MarketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Markets not implemented")
}
func (UnimplementedInjectiveSpotExchangeRPCServer) Market(context.Context, *MarketRequest) (*MarketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Market not implemented")
}
func (UnimplementedInjectiveSpotExchangeRPCServer) StreamMarkets(*StreamMarketsRequest, InjectiveSpotExchangeRPC_StreamMarketsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamMarkets not implemented")
}
func (UnimplementedInjectiveSpotExchangeRPCServer) OrderbookV2(context.Context, *OrderbookV2Request) (*OrderbookV2Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderbookV2 not implemented")
}
func (UnimplementedInjectiveSpotExchangeRPCServer) OrderbooksV2(context.Context, *OrderbooksV2Request) (*OrderbooksV2Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderbooksV2 not implemented")
}
func (UnimplementedInjectiveSpotExchangeRPCServer) StreamOrderbookV2(*StreamOrderbookV2Request, InjectiveSpotExchangeRPC_StreamOrderbookV2Server) error {
	return status.Errorf(codes.Unimplemented, "method StreamOrderbookV2 not implemented")
}
func (UnimplementedInjectiveSpotExchangeRPCServer) StreamOrderbookUpdate(*StreamOrderbookUpdateRequest, InjectiveSpotExchangeRPC_StreamOrderbookUpdateServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamOrderbookUpdate not implemented")
}
func (UnimplementedInjectiveSpotExchangeRPCServer) Orders(context.Context, *OrdersRequest) (*OrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Orders not implemented")
}
func (UnimplementedInjectiveSpotExchangeRPCServer) StreamOrders(*StreamOrdersRequest, InjectiveSpotExchangeRPC_StreamOrdersServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamOrders not implemented")
}
func (UnimplementedInjectiveSpotExchangeRPCServer) Trades(context.Context, *TradesRequest) (*TradesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Trades not implemented")
}
func (UnimplementedInjectiveSpotExchangeRPCServer) StreamTrades(*StreamTradesRequest, InjectiveSpotExchangeRPC_StreamTradesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamTrades not implemented")
}
func (UnimplementedInjectiveSpotExchangeRPCServer) TradesV2(context.Context, *TradesV2Request) (*TradesV2Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradesV2 not implemented")
}
func (UnimplementedInjectiveSpotExchangeRPCServer) StreamTradesV2(*StreamTradesV2Request, InjectiveSpotExchangeRPC_StreamTradesV2Server) error {
	return status.Errorf(codes.Unimplemented, "method StreamTradesV2 not implemented")
}
func (UnimplementedInjectiveSpotExchangeRPCServer) SubaccountOrdersList(context.Context, *SubaccountOrdersListRequest) (*SubaccountOrdersListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubaccountOrdersList not implemented")
}
func (UnimplementedInjectiveSpotExchangeRPCServer) SubaccountTradesList(context.Context, *SubaccountTradesListRequest) (*SubaccountTradesListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubaccountTradesList not implemented")
}
func (UnimplementedInjectiveSpotExchangeRPCServer) OrdersHistory(context.Context, *OrdersHistoryRequest) (*OrdersHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrdersHistory not implemented")
}
func (UnimplementedInjectiveSpotExchangeRPCServer) StreamOrdersHistory(*StreamOrdersHistoryRequest, InjectiveSpotExchangeRPC_StreamOrdersHistoryServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamOrdersHistory not implemented")
}
func (UnimplementedInjectiveSpotExchangeRPCServer) AtomicSwapHistory(context.Context, *AtomicSwapHistoryRequest) (*AtomicSwapHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AtomicSwapHistory not implemented")
}
func (UnimplementedInjectiveSpotExchangeRPCServer) mustEmbedUnimplementedInjectiveSpotExchangeRPCServer() {
}

// UnsafeInjectiveSpotExchangeRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InjectiveSpotExchangeRPCServer will
// result in compilation errors.
type UnsafeInjectiveSpotExchangeRPCServer interface {
	mustEmbedUnimplementedInjectiveSpotExchangeRPCServer()
}

func RegisterInjectiveSpotExchangeRPCServer(s grpc.ServiceRegistrar, srv InjectiveSpotExchangeRPCServer) {
	s.RegisterService(&InjectiveSpotExchangeRPC_ServiceDesc, srv)
}

func _InjectiveSpotExchangeRPC_Markets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveSpotExchangeRPCServer).Markets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helios_spot_exchange_rpc.InjectiveSpotExchangeRPC/Markets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveSpotExchangeRPCServer).Markets(ctx, req.(*MarketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveSpotExchangeRPC_Market_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveSpotExchangeRPCServer).Market(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helios_spot_exchange_rpc.InjectiveSpotExchangeRPC/Market",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveSpotExchangeRPCServer).Market(ctx, req.(*MarketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveSpotExchangeRPC_StreamMarkets_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamMarketsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InjectiveSpotExchangeRPCServer).StreamMarkets(m, &heliosSpotExchangeRPCStreamMarketsServer{stream})
}

type InjectiveSpotExchangeRPC_StreamMarketsServer interface {
	Send(*StreamMarketsResponse) error
	grpc.ServerStream
}

type heliosSpotExchangeRPCStreamMarketsServer struct {
	grpc.ServerStream
}

func (x *heliosSpotExchangeRPCStreamMarketsServer) Send(m *StreamMarketsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _InjectiveSpotExchangeRPC_OrderbookV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderbookV2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveSpotExchangeRPCServer).OrderbookV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helios_spot_exchange_rpc.InjectiveSpotExchangeRPC/OrderbookV2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveSpotExchangeRPCServer).OrderbookV2(ctx, req.(*OrderbookV2Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveSpotExchangeRPC_OrderbooksV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderbooksV2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveSpotExchangeRPCServer).OrderbooksV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helios_spot_exchange_rpc.InjectiveSpotExchangeRPC/OrderbooksV2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveSpotExchangeRPCServer).OrderbooksV2(ctx, req.(*OrderbooksV2Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveSpotExchangeRPC_StreamOrderbookV2_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamOrderbookV2Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InjectiveSpotExchangeRPCServer).StreamOrderbookV2(m, &heliosSpotExchangeRPCStreamOrderbookV2Server{stream})
}

type InjectiveSpotExchangeRPC_StreamOrderbookV2Server interface {
	Send(*StreamOrderbookV2Response) error
	grpc.ServerStream
}

type heliosSpotExchangeRPCStreamOrderbookV2Server struct {
	grpc.ServerStream
}

func (x *heliosSpotExchangeRPCStreamOrderbookV2Server) Send(m *StreamOrderbookV2Response) error {
	return x.ServerStream.SendMsg(m)
}

func _InjectiveSpotExchangeRPC_StreamOrderbookUpdate_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamOrderbookUpdateRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InjectiveSpotExchangeRPCServer).StreamOrderbookUpdate(m, &heliosSpotExchangeRPCStreamOrderbookUpdateServer{stream})
}

type InjectiveSpotExchangeRPC_StreamOrderbookUpdateServer interface {
	Send(*StreamOrderbookUpdateResponse) error
	grpc.ServerStream
}

type heliosSpotExchangeRPCStreamOrderbookUpdateServer struct {
	grpc.ServerStream
}

func (x *heliosSpotExchangeRPCStreamOrderbookUpdateServer) Send(m *StreamOrderbookUpdateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _InjectiveSpotExchangeRPC_Orders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveSpotExchangeRPCServer).Orders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helios_spot_exchange_rpc.InjectiveSpotExchangeRPC/Orders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveSpotExchangeRPCServer).Orders(ctx, req.(*OrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveSpotExchangeRPC_StreamOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamOrdersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InjectiveSpotExchangeRPCServer).StreamOrders(m, &heliosSpotExchangeRPCStreamOrdersServer{stream})
}

type InjectiveSpotExchangeRPC_StreamOrdersServer interface {
	Send(*StreamOrdersResponse) error
	grpc.ServerStream
}

type heliosSpotExchangeRPCStreamOrdersServer struct {
	grpc.ServerStream
}

func (x *heliosSpotExchangeRPCStreamOrdersServer) Send(m *StreamOrdersResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _InjectiveSpotExchangeRPC_Trades_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveSpotExchangeRPCServer).Trades(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helios_spot_exchange_rpc.InjectiveSpotExchangeRPC/Trades",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveSpotExchangeRPCServer).Trades(ctx, req.(*TradesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveSpotExchangeRPC_StreamTrades_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamTradesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InjectiveSpotExchangeRPCServer).StreamTrades(m, &heliosSpotExchangeRPCStreamTradesServer{stream})
}

type InjectiveSpotExchangeRPC_StreamTradesServer interface {
	Send(*StreamTradesResponse) error
	grpc.ServerStream
}

type heliosSpotExchangeRPCStreamTradesServer struct {
	grpc.ServerStream
}

func (x *heliosSpotExchangeRPCStreamTradesServer) Send(m *StreamTradesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _InjectiveSpotExchangeRPC_TradesV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradesV2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveSpotExchangeRPCServer).TradesV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helios_spot_exchange_rpc.InjectiveSpotExchangeRPC/TradesV2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveSpotExchangeRPCServer).TradesV2(ctx, req.(*TradesV2Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveSpotExchangeRPC_StreamTradesV2_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamTradesV2Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InjectiveSpotExchangeRPCServer).StreamTradesV2(m, &heliosSpotExchangeRPCStreamTradesV2Server{stream})
}

type InjectiveSpotExchangeRPC_StreamTradesV2Server interface {
	Send(*StreamTradesV2Response) error
	grpc.ServerStream
}

type heliosSpotExchangeRPCStreamTradesV2Server struct {
	grpc.ServerStream
}

func (x *heliosSpotExchangeRPCStreamTradesV2Server) Send(m *StreamTradesV2Response) error {
	return x.ServerStream.SendMsg(m)
}

func _InjectiveSpotExchangeRPC_SubaccountOrdersList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubaccountOrdersListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveSpotExchangeRPCServer).SubaccountOrdersList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helios_spot_exchange_rpc.InjectiveSpotExchangeRPC/SubaccountOrdersList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveSpotExchangeRPCServer).SubaccountOrdersList(ctx, req.(*SubaccountOrdersListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveSpotExchangeRPC_SubaccountTradesList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubaccountTradesListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveSpotExchangeRPCServer).SubaccountTradesList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helios_spot_exchange_rpc.InjectiveSpotExchangeRPC/SubaccountTradesList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveSpotExchangeRPCServer).SubaccountTradesList(ctx, req.(*SubaccountTradesListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveSpotExchangeRPC_OrdersHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrdersHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveSpotExchangeRPCServer).OrdersHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helios_spot_exchange_rpc.InjectiveSpotExchangeRPC/OrdersHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveSpotExchangeRPCServer).OrdersHistory(ctx, req.(*OrdersHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveSpotExchangeRPC_StreamOrdersHistory_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamOrdersHistoryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InjectiveSpotExchangeRPCServer).StreamOrdersHistory(m, &heliosSpotExchangeRPCStreamOrdersHistoryServer{stream})
}

type InjectiveSpotExchangeRPC_StreamOrdersHistoryServer interface {
	Send(*StreamOrdersHistoryResponse) error
	grpc.ServerStream
}

type heliosSpotExchangeRPCStreamOrdersHistoryServer struct {
	grpc.ServerStream
}

func (x *heliosSpotExchangeRPCStreamOrdersHistoryServer) Send(m *StreamOrdersHistoryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _InjectiveSpotExchangeRPC_AtomicSwapHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AtomicSwapHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveSpotExchangeRPCServer).AtomicSwapHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helios_spot_exchange_rpc.InjectiveSpotExchangeRPC/AtomicSwapHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveSpotExchangeRPCServer).AtomicSwapHistory(ctx, req.(*AtomicSwapHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InjectiveSpotExchangeRPC_ServiceDesc is the grpc.ServiceDesc for InjectiveSpotExchangeRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InjectiveSpotExchangeRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helios_spot_exchange_rpc.InjectiveSpotExchangeRPC",
	HandlerType: (*InjectiveSpotExchangeRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Markets",
			Handler:    _InjectiveSpotExchangeRPC_Markets_Handler,
		},
		{
			MethodName: "Market",
			Handler:    _InjectiveSpotExchangeRPC_Market_Handler,
		},
		{
			MethodName: "OrderbookV2",
			Handler:    _InjectiveSpotExchangeRPC_OrderbookV2_Handler,
		},
		{
			MethodName: "OrderbooksV2",
			Handler:    _InjectiveSpotExchangeRPC_OrderbooksV2_Handler,
		},
		{
			MethodName: "Orders",
			Handler:    _InjectiveSpotExchangeRPC_Orders_Handler,
		},
		{
			MethodName: "Trades",
			Handler:    _InjectiveSpotExchangeRPC_Trades_Handler,
		},
		{
			MethodName: "TradesV2",
			Handler:    _InjectiveSpotExchangeRPC_TradesV2_Handler,
		},
		{
			MethodName: "SubaccountOrdersList",
			Handler:    _InjectiveSpotExchangeRPC_SubaccountOrdersList_Handler,
		},
		{
			MethodName: "SubaccountTradesList",
			Handler:    _InjectiveSpotExchangeRPC_SubaccountTradesList_Handler,
		},
		{
			MethodName: "OrdersHistory",
			Handler:    _InjectiveSpotExchangeRPC_OrdersHistory_Handler,
		},
		{
			MethodName: "AtomicSwapHistory",
			Handler:    _InjectiveSpotExchangeRPC_AtomicSwapHistory_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamMarkets",
			Handler:       _InjectiveSpotExchangeRPC_StreamMarkets_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamOrderbookV2",
			Handler:       _InjectiveSpotExchangeRPC_StreamOrderbookV2_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamOrderbookUpdate",
			Handler:       _InjectiveSpotExchangeRPC_StreamOrderbookUpdate_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamOrders",
			Handler:       _InjectiveSpotExchangeRPC_StreamOrders_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamTrades",
			Handler:       _InjectiveSpotExchangeRPC_StreamTrades_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamTradesV2",
			Handler:       _InjectiveSpotExchangeRPC_StreamTradesV2_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamOrdersHistory",
			Handler:       _InjectiveSpotExchangeRPC_StreamOrdersHistory_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "goadesign_goagen_helios_spot_exchange_rpc.proto",
}
