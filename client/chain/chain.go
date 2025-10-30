package chain

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	permissionstypes "github.com/Helios-Chain-Labs/sdk-go/chain/permissions/types"

	"github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"

	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"

	"github.com/cosmos/cosmos-sdk/types/query"

	"google.golang.org/grpc/credentials/insecure"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
	tokenfactorytypes "github.com/Helios-Chain-Labs/sdk-go/chain/tokenfactory/types"
	"github.com/Helios-Chain-Labs/sdk-go/client/common"
	log "github.com/Helios-Chain-Labs/suplog"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	ibcconnectiontypes "github.com/cosmos/ibc-go/v8/modules/core/03-connection/types"
	ibcchanneltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	ethcommon "github.com/ethereum/go-ethereum/common"
)

type OrderbookType string

const (
	msgCommitBatchSizeLimit          = 1024
	msgCommitBatchTimeLimit          = 500 * time.Millisecond
	defaultBroadcastStatusPoll       = 100 * time.Millisecond
	defaultBroadcastTimeout          = 40 * time.Second
	defaultTimeoutHeight             = 20
	defaultTimeoutHeightSyncInterval = 20 * time.Second
	SpotOrderbook                    = "helios.exchange.v1beta1.EventOrderbookUpdate.spot_orderbooks"
	DerivativeOrderbook              = "helios.exchange.v1beta1.EventOrderbookUpdate.derivative_orderbooks"
)

var (
	ErrTimedOut       = errors.New("tx timed out")
	ErrQueueClosed    = errors.New("queue is closed")
	ErrEnqueueTimeout = errors.New("enqueue timeout")
	ErrReadOnly       = errors.New("client is in read-only mode")
)

type ChainClient interface {
	CanSignTransactions() bool
	FromAddress() sdk.AccAddress
	QueryClient() *grpc.ClientConn
	ClientContext() client.Context
	// return account number and sequence without increasing sequence
	GetAccNonce() (accNum uint64, accSeq uint64)
	ForceSyncNonce()

	SimulateMsg(clientCtx client.Context, msgs ...sdk.Msg) (*txtypes.SimulateResponse, error)
	AsyncBroadcastMsg(msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error)
	SyncBroadcastMsg(msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error)

	// Build signed tx with given accNum and accSeq, useful for offline siging
	// If simulate is set to false, initialGas will be used
	BuildSignedTx(clientCtx client.Context, accNum, accSeq, initialGas uint64, msg ...sdk.Msg) ([]byte, error)
	SyncBroadcastSignedTx(tyBytes []byte) (*txtypes.BroadcastTxResponse, error)
	AsyncBroadcastSignedTx(txBytes []byte) (*txtypes.BroadcastTxResponse, error)
	QueueBroadcastMsg(msgs ...sdk.Msg) error

	// Bank Module
	GetBankBalances(ctx context.Context, address string) (*banktypes.QueryAllBalancesResponse, error)
	GetBankBalance(ctx context.Context, address string, denom string) (*banktypes.QueryBalanceResponse, error)
	GetBankSpendableBalances(ctx context.Context, address string, pagination *query.PageRequest) (*banktypes.QuerySpendableBalancesResponse, error)
	GetBankSpendableBalancesByDenom(ctx context.Context, address string, denom string) (*banktypes.QuerySpendableBalanceByDenomResponse, error)
	GetBankTotalSupply(ctx context.Context, pagination *query.PageRequest) (*banktypes.QueryTotalSupplyResponse, error)
	GetBankSupplyOf(ctx context.Context, denom string) (*banktypes.QuerySupplyOfResponse, error)
	GetDenomMetadata(ctx context.Context, denom string) (*banktypes.QueryDenomMetadataResponse, error)
	GetDenomsMetadata(ctx context.Context, pagination *query.PageRequest) (*banktypes.QueryDenomsMetadataResponse, error)
	GetDenomOwners(ctx context.Context, denom string, pagination *query.PageRequest) (*banktypes.QueryDenomOwnersResponse, error)
	GetBankSendEnabled(ctx context.Context, denoms []string, pagination *query.PageRequest) (*banktypes.QuerySendEnabledResponse, error)
	GetAccount(ctx context.Context, address string) (*authtypes.QueryAccountResponse, error)

	DefaultSubaccount(acc sdk.AccAddress) ethcommon.Hash
	Subaccount(account sdk.AccAddress, index int) ethcommon.Hash

	GetGasFee() (string, error)

	// get tx from chain node
	GetTx(ctx context.Context, txHash string) (*txtypes.GetTxResponse, error)

	// hyperion module
	GetAttestation(ctx context.Context, eventNonce uint64, claimHash []byte) (*hyperiontypes.QueryAttestationResponse, error)

	// wasm module
	FetchContractInfo(ctx context.Context, address string) (*wasmtypes.QueryContractInfoResponse, error)
	FetchContractHistory(ctx context.Context, address string, pagination *query.PageRequest) (*wasmtypes.QueryContractHistoryResponse, error)
	FetchContractsByCode(ctx context.Context, codeId uint64, pagination *query.PageRequest) (*wasmtypes.QueryContractsByCodeResponse, error)
	FetchAllContractsState(ctx context.Context, address string, pagination *query.PageRequest) (*wasmtypes.QueryAllContractStateResponse, error)
	RawContractState(
		ctx context.Context,
		contractAddress string,
		queryData []byte,
	) (*wasmtypes.QueryRawContractStateResponse, error)
	SmartContractState(
		ctx context.Context,
		contractAddress string,
		queryData []byte,
	) (*wasmtypes.QuerySmartContractStateResponse, error)
	FetchCode(ctx context.Context, codeId uint64) (*wasmtypes.QueryCodeResponse, error)
	FetchCodes(ctx context.Context, pagination *query.PageRequest) (*wasmtypes.QueryCodesResponse, error)
	FetchPinnedCodes(ctx context.Context, pagination *query.PageRequest) (*wasmtypes.QueryPinnedCodesResponse, error)
	FetchContractsByCreator(ctx context.Context, creator string, pagination *query.PageRequest) (*wasmtypes.QueryContractsByCreatorResponse, error)

	// tokenfactory module
	FetchDenomAuthorityMetadata(ctx context.Context, creator string, subDenom string) (*tokenfactorytypes.QueryDenomAuthorityMetadataResponse, error)
	FetchDenomsFromCreator(ctx context.Context, creator string) (*tokenfactorytypes.QueryDenomsFromCreatorResponse, error)
	FetchTokenfactoryModuleState(ctx context.Context) (*tokenfactorytypes.QueryModuleStateResponse, error)

	// distribution module
	FetchValidatorDistributionInfo(ctx context.Context, validatorAddress string) (*distributiontypes.QueryValidatorDistributionInfoResponse, error)
	FetchValidatorOutstandingRewards(ctx context.Context, validatorAddress string) (*distributiontypes.QueryValidatorOutstandingRewardsResponse, error)
	FetchValidatorCommission(ctx context.Context, validatorAddress string) (*distributiontypes.QueryValidatorCommissionResponse, error)
	FetchValidatorSlashes(ctx context.Context, validatorAddress string, startingHeight uint64, endingHeight uint64, pagination *query.PageRequest) (*distributiontypes.QueryValidatorSlashesResponse, error)
	FetchDelegationRewards(ctx context.Context, delegatorAddress string, validatorAddress string) (*distributiontypes.QueryDelegationRewardsResponse, error)
	FetchDelegationTotalRewards(ctx context.Context, delegatorAddress string) (*distributiontypes.QueryDelegationTotalRewardsResponse, error)
	FetchDelegatorValidators(ctx context.Context, delegatorAddress string) (*distributiontypes.QueryDelegatorValidatorsResponse, error)
	FetchDelegatorWithdrawAddress(ctx context.Context, delegatorAddress string) (*distributiontypes.QueryDelegatorWithdrawAddressResponse, error)
	FetchCommunityPool(ctx context.Context) (*distributiontypes.QueryCommunityPoolResponse, error)

	// Tendermint module
	FetchNodeInfo(ctx context.Context) (*cmtservice.GetNodeInfoResponse, error)
	FetchSyncing(ctx context.Context) (*cmtservice.GetSyncingResponse, error)
	FetchLatestBlock(ctx context.Context) (*cmtservice.GetLatestBlockResponse, error)
	FetchBlockByHeight(ctx context.Context, height int64) (*cmtservice.GetBlockByHeightResponse, error)
	FetchLatestValidatorSet(ctx context.Context) (*cmtservice.GetLatestValidatorSetResponse, error)
	FetchValidatorSetByHeight(ctx context.Context, height int64, pagination *query.PageRequest) (*cmtservice.GetValidatorSetByHeightResponse, error)
	ABCIQuery(ctx context.Context, path string, data []byte, height int64, prove bool) (*cmtservice.ABCIQueryResponse, error)

	// IBC Transfer module
	FetchDenomTrace(ctx context.Context, hash string) (*ibctransfertypes.QueryDenomTraceResponse, error)
	FetchDenomTraces(ctx context.Context, pagination *query.PageRequest) (*ibctransfertypes.QueryDenomTracesResponse, error)
	FetchDenomHash(ctx context.Context, trace string) (*ibctransfertypes.QueryDenomHashResponse, error)
	FetchEscrowAddress(ctx context.Context, portId string, channelId string) (*ibctransfertypes.QueryEscrowAddressResponse, error)
	FetchTotalEscrowForDenom(ctx context.Context, denom string) (*ibctransfertypes.QueryTotalEscrowForDenomResponse, error)

	// IBC Core Channel module
	FetchIBCChannel(ctx context.Context, portId string, channelId string) (*ibcchanneltypes.QueryChannelResponse, error)
	FetchIBCChannels(ctx context.Context, pagination *query.PageRequest) (*ibcchanneltypes.QueryChannelsResponse, error)
	FetchIBCConnectionChannels(ctx context.Context, connection string, pagination *query.PageRequest) (*ibcchanneltypes.QueryConnectionChannelsResponse, error)
	FetchIBCChannelClientState(ctx context.Context, portId string, channelId string) (*ibcchanneltypes.QueryChannelClientStateResponse, error)
	FetchIBCChannelConsensusState(ctx context.Context, portId string, channelId string, revisionNumber uint64, revisionHeight uint64) (*ibcchanneltypes.QueryChannelConsensusStateResponse, error)
	FetchIBCPacketCommitment(ctx context.Context, portId string, channelId string, sequence uint64) (*ibcchanneltypes.QueryPacketCommitmentResponse, error)
	FetchIBCPacketCommitments(ctx context.Context, portId string, channelId string, pagination *query.PageRequest) (*ibcchanneltypes.QueryPacketCommitmentsResponse, error)
	FetchIBCPacketReceipt(ctx context.Context, portId string, channelId string, sequence uint64) (*ibcchanneltypes.QueryPacketReceiptResponse, error)
	FetchIBCPacketAcknowledgement(ctx context.Context, portId string, channelId string, sequence uint64) (*ibcchanneltypes.QueryPacketAcknowledgementResponse, error)
	FetchIBCPacketAcknowledgements(ctx context.Context, portId string, channelId string, packetCommitmentSequences []uint64, pagination *query.PageRequest) (*ibcchanneltypes.QueryPacketAcknowledgementsResponse, error)
	FetchIBCUnreceivedPackets(ctx context.Context, portId string, channelId string, packetCommitmentSequences []uint64) (*ibcchanneltypes.QueryUnreceivedPacketsResponse, error)
	FetchIBCUnreceivedAcks(ctx context.Context, portId string, channelId string, packetAckSequences []uint64) (*ibcchanneltypes.QueryUnreceivedAcksResponse, error)
	FetchIBCNextSequenceReceive(ctx context.Context, portId string, channelId string) (*ibcchanneltypes.QueryNextSequenceReceiveResponse, error)

	// IBC Core Chain module
	FetchIBCClientState(ctx context.Context, clientId string) (*ibcclienttypes.QueryClientStateResponse, error)
	FetchIBCClientStates(ctx context.Context, pagination *query.PageRequest) (*ibcclienttypes.QueryClientStatesResponse, error)
	FetchIBCConsensusState(ctx context.Context, clientId string, revisionNumber uint64, revisionHeight uint64, latestHeight bool) (*ibcclienttypes.QueryConsensusStateResponse, error)
	FetchIBCConsensusStates(ctx context.Context, clientId string, pagination *query.PageRequest) (*ibcclienttypes.QueryConsensusStatesResponse, error)
	FetchIBCConsensusStateHeights(ctx context.Context, clientId string, pagination *query.PageRequest) (*ibcclienttypes.QueryConsensusStateHeightsResponse, error)
	FetchIBCClientStatus(ctx context.Context, clientId string) (*ibcclienttypes.QueryClientStatusResponse, error)
	FetchIBCClientParams(ctx context.Context) (*ibcclienttypes.QueryClientParamsResponse, error)
	FetchIBCUpgradedClientState(ctx context.Context) (*ibcclienttypes.QueryUpgradedClientStateResponse, error)
	FetchIBCUpgradedConsensusState(ctx context.Context) (*ibcclienttypes.QueryUpgradedConsensusStateResponse, error)

	// IBC Core Connection module
	FetchIBCConnection(ctx context.Context, connectionId string) (*ibcconnectiontypes.QueryConnectionResponse, error)
	FetchIBCConnections(ctx context.Context, pagination *query.PageRequest) (*ibcconnectiontypes.QueryConnectionsResponse, error)
	FetchIBCClientConnections(ctx context.Context, clientId string) (*ibcconnectiontypes.QueryClientConnectionsResponse, error)
	FetchIBCConnectionClientState(ctx context.Context, connectionId string) (*ibcconnectiontypes.QueryConnectionClientStateResponse, error)
	FetchIBCConnectionConsensusState(ctx context.Context, connectionId string, revisionNumber uint64, revisionHeight uint64) (*ibcconnectiontypes.QueryConnectionConsensusStateResponse, error)
	FetchIBCConnectionParams(ctx context.Context) (*ibcconnectiontypes.QueryConnectionParamsResponse, error)

	// Permissions module
	FetchAllNamespaces(ctx context.Context) (*permissionstypes.QueryAllNamespacesResponse, error)
	FetchNamespaceByDenom(ctx context.Context, denom string, includeRoles bool) (*permissionstypes.QueryNamespaceByDenomResponse, error)
	FetchAddressRoles(ctx context.Context, denom, address string) (*permissionstypes.QueryAddressRolesResponse, error)
	FetchAddressesByRole(ctx context.Context, denom, role string) (*permissionstypes.QueryAddressesByRoleResponse, error)
	FetchVouchersForAddress(ctx context.Context, address string) (*permissionstypes.QueryVouchersForAddressResponse, error)

	GetNetwork() common.Network
	Close()
}

var _ ChainClient = &chainClient{}

type chainClient struct {
	ctx             client.Context
	network         common.Network
	opts            *common.ClientOptions
	logger          log.Logger
	conn            *grpc.ClientConn
	chainStreamConn *grpc.ClientConn
	txFactory       tx.Factory

	doneC   chan bool
	msgC    chan sdk.Msg
	syncMux *sync.Mutex

	cancelCtx context.Context
	cancelFn  func()

	accNum    uint64
	accSeq    uint64
	gasWanted uint64
	gasFee    string

	sessionEnabled bool

	ofacChecker *OfacChecker

	authQueryClient          authtypes.QueryClient
	authzQueryClient         authztypes.QueryClient
	bankQueryClient          banktypes.QueryClient
	distributionQueryClient  distributiontypes.QueryClient
	ibcChannelQueryClient    ibcchanneltypes.QueryClient
	ibcClientQueryClient     ibcclienttypes.QueryClient
	ibcConnectionQueryClient ibcconnectiontypes.QueryClient
	ibcTransferQueryClient   ibctransfertypes.QueryClient
	permissionsQueryClient   permissionstypes.QueryClient
	tendermintQueryClient    cmtservice.ServiceClient
	tokenfactoryQueryClient  tokenfactorytypes.QueryClient
	txClient                 txtypes.ServiceClient
	wasmQueryClient          wasmtypes.QueryClient
	hyperionQueryClient      hyperiontypes.QueryClient
	subaccountToNonce        map[ethcommon.Hash]uint32

	closed  int64
	canSign bool
}

func NewChainClient(
	ctx client.Context,
	network common.Network,
	options ...common.ClientOption,
) (ChainClient, error) {

	// process options
	opts := common.DefaultClientOptions()

	if network.ChainTLSCert != nil {
		options = append(options, common.OptionTLSCert(network.ChainTLSCert))
	}
	for _, opt := range options {
		if err := opt(opts); err != nil {
			err = errors.Wrap(err, "error in client option")
			return nil, err
		}
	}

	// init tx factory
	var txFactory tx.Factory
	if opts.TxFactory == nil {
		txFactory = NewTxFactory(ctx)
		if opts.GasPrices != "" {
			txFactory = txFactory.WithGasPrices(opts.GasPrices)
		}
		if opts.Gas != "" {
			gas, err := strconv.ParseUint(opts.Gas, 10, 64)
			if err != nil {
				return nil, errors.Wrap(err, "error parsing gas limit")
			}
			txFactory = txFactory.WithGas(gas)
		}
	} else {
		txFactory = *opts.TxFactory
	}

	// init grpc connection
	var conn *grpc.ClientConn
	var err error
	stickySessionEnabled := true
	if opts.TLSCert != nil {
		conn, err = grpc.Dial(network.ChainGrpcEndpoint, grpc.WithTransportCredentials(opts.TLSCert), grpc.WithContextDialer(common.DialerFunc))
	} else {
		conn, err = grpc.Dial(network.ChainGrpcEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(common.DialerFunc))
		stickySessionEnabled = false
	}
	if err != nil {
		err = errors.Wrapf(err, "failed to connect to the gRPC: %s", network.ChainGrpcEndpoint)
		return nil, err
	}

	var chainStreamConn *grpc.ClientConn
	if opts.TLSCert != nil {
		chainStreamConn, err = grpc.Dial(network.ChainStreamGrpcEndpoint, grpc.WithTransportCredentials(opts.TLSCert), grpc.WithContextDialer(common.DialerFunc))
	} else {
		chainStreamConn, err = grpc.Dial(network.ChainStreamGrpcEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(common.DialerFunc))
	}
	if err != nil {
		err = errors.Wrapf(err, "failed to connect to the chain stream gRPC: %s", network.ChainStreamGrpcEndpoint)
		return nil, err
	}

	cancelCtx, cancelFn := context.WithCancel(context.Background())
	// build client
	cc := &chainClient{
		ctx:     ctx,
		network: network,
		opts:    opts,

		logger: log.WithFields(log.Fields{
			"module": "sdk-go",
			"svc":    "chainClient",
		}),

		conn:            conn,
		chainStreamConn: chainStreamConn,
		txFactory:       txFactory,
		canSign:         ctx.Keyring != nil,
		syncMux:         new(sync.Mutex),
		msgC:            make(chan sdk.Msg, msgCommitBatchSizeLimit),
		doneC:           make(chan bool, 1),
		cancelCtx:       cancelCtx,
		cancelFn:        cancelFn,

		sessionEnabled: stickySessionEnabled,

		authQueryClient:          authtypes.NewQueryClient(conn),
		authzQueryClient:         authztypes.NewQueryClient(conn),
		bankQueryClient:          banktypes.NewQueryClient(conn),
		distributionQueryClient:  distributiontypes.NewQueryClient(conn),
		ibcChannelQueryClient:    ibcchanneltypes.NewQueryClient(conn),
		ibcClientQueryClient:     ibcclienttypes.NewQueryClient(conn),
		ibcConnectionQueryClient: ibcconnectiontypes.NewQueryClient(conn),
		ibcTransferQueryClient:   ibctransfertypes.NewQueryClient(conn),
		permissionsQueryClient:   permissionstypes.NewQueryClient(conn),
		tendermintQueryClient:    cmtservice.NewServiceClient(conn),
		tokenfactoryQueryClient:  tokenfactorytypes.NewQueryClient(conn),
		txClient:                 txtypes.NewServiceClient(conn),
		wasmQueryClient:          wasmtypes.NewQueryClient(conn),
		subaccountToNonce:        make(map[ethcommon.Hash]uint32),
	}

	cc.ofacChecker, err = NewOfacChecker()
	if err != nil {
		return nil, errors.Wrap(err, "Error creating OFAC checker")
	}
	if cc.canSign {
		var err error
		account, err := cc.txFactory.AccountRetriever().GetAccount(ctx, ctx.GetFromAddress())
		if err != nil {
			err = errors.Wrapf(err, "failed to get account")
			return nil, err
		}
		if cc.ofacChecker.IsBlacklisted(account.GetAddress().String()) {
			return nil, errors.Errorf("Address %s is in the OFAC list", account.GetAddress())
		}
		cc.accNum, cc.accSeq = account.GetAccountNumber(), account.GetSequence()
		go cc.runBatchBroadcast()
		go cc.syncTimeoutHeight()
	}

	return cc, nil
}
func (c *chainClient) syncNonce() {
	num, seq, err := c.txFactory.AccountRetriever().GetAccountNumberSequence(c.ctx, c.ctx.GetFromAddress())
	if err != nil {
		c.logger.WithError(err).Errorln("failed to get account seq")
		return
	} else if num != c.accNum {
		c.logger.WithFields(log.Fields{
			"expected": c.accNum,
			"actual":   num,
		}).Panic("account number changed during nonce sync")
	}

	c.accSeq = seq
}

func (c *chainClient) syncTimeoutHeight() {
	t := time.NewTicker(defaultTimeoutHeightSyncInterval)
	defer t.Stop()

	for {
		block, err := c.ctx.Client.Block(c.cancelCtx, nil)
		if err != nil {
			c.logger.WithError(err).Errorln("failed to get current block")
			return
		}
		c.txFactory.WithTimeoutHeight(uint64(block.Block.Height) + defaultTimeoutHeight)

		select {
		case <-c.cancelCtx.Done():
			return
		case <-t.C:
			continue
		}
	}
}

// PrepareFactory ensures the account defined by ctx.GetFromAddress() exists and
// if the account number and/or the account sequence number are zero (not set),
// they will be queried for and set on the provided Factory. A new Factory with
// the updated fields will be returned.
func PrepareFactory(clientCtx client.Context, txf tx.Factory) (tx.Factory, error) {
	from := clientCtx.GetFromAddress()

	if err := txf.AccountRetriever().EnsureExists(clientCtx, from); err != nil {
		return txf, err
	}

	initNum, initSeq := txf.AccountNumber(), txf.Sequence()
	if initNum == 0 || initSeq == 0 {
		num, seq, err := txf.AccountRetriever().GetAccountNumberSequence(clientCtx, from)
		if err != nil {
			return txf, err
		}

		if initNum == 0 {
			txf = txf.WithAccountNumber(num)
		}

		if initSeq == 0 {
			txf = txf.WithSequence(seq)
		}
	}

	return txf, nil
}

func (c *chainClient) getAccSeq() uint64 {
	defer func() {
		c.accSeq += 1
	}()
	return c.accSeq
}

func (c *chainClient) GetAccNonce() (accNum, accSeq uint64) {
	return c.accNum, c.accSeq
}

func (c *chainClient) ForceSyncNonce() {
	c.syncNonce()
}

func (c *chainClient) QueryClient() *grpc.ClientConn {
	return c.conn
}

func (c *chainClient) ClientContext() client.Context {
	return c.ctx
}

func (c *chainClient) CanSignTransactions() bool {
	return c.canSign
}

func (c *chainClient) FromAddress() sdk.AccAddress {
	if !c.canSign {
		return sdk.AccAddress{}
	}

	return c.ctx.FromAddress
}

func (c *chainClient) Close() {
	if !c.canSign {
		return
	}
	if atomic.CompareAndSwapInt64(&c.closed, 0, 1) {
		close(c.msgC)
	}

	if c.cancelFn != nil {
		c.cancelFn()
	}
	<-c.doneC
	if c.conn != nil {
		c.conn.Close()
	}
	if c.chainStreamConn != nil {
		c.chainStreamConn.Close()
	}
}

// Bank Module

func (c *chainClient) GetBankBalances(ctx context.Context, address string) (*banktypes.QueryAllBalancesResponse, error) {
	req := &banktypes.QueryAllBalancesRequest{
		Address: address,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.AllBalances, req)

	return res, err
}

func (c *chainClient) GetBankBalance(ctx context.Context, address, denom string) (*banktypes.QueryBalanceResponse, error) {
	req := &banktypes.QueryBalanceRequest{
		Address: address,
		Denom:   denom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.Balance, req)

	return res, err
}

func (c *chainClient) GetBankSpendableBalances(ctx context.Context, address string, pagination *query.PageRequest) (*banktypes.QuerySpendableBalancesResponse, error) {
	req := &banktypes.QuerySpendableBalancesRequest{
		Address:    address,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.SpendableBalances, req)

	return res, err
}

func (c *chainClient) GetBankSpendableBalancesByDenom(ctx context.Context, address, denom string) (*banktypes.QuerySpendableBalanceByDenomResponse, error) {
	req := &banktypes.QuerySpendableBalanceByDenomRequest{
		Address: address,
		Denom:   denom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.SpendableBalanceByDenom, req)

	return res, err
}

func (c *chainClient) GetBankTotalSupply(ctx context.Context, pagination *query.PageRequest) (*banktypes.QueryTotalSupplyResponse, error) {
	req := &banktypes.QueryTotalSupplyRequest{Pagination: pagination}
	resp, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.TotalSupply, req)

	return resp, err
}

func (c *chainClient) GetBankSupplyOf(ctx context.Context, denom string) (*banktypes.QuerySupplyOfResponse, error) {
	req := &banktypes.QuerySupplyOfRequest{Denom: denom}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.SupplyOf, req)

	return res, err
}

func (c *chainClient) GetDenomMetadata(ctx context.Context, denom string) (*banktypes.QueryDenomMetadataResponse, error) {
	req := &banktypes.QueryDenomMetadataRequest{Denom: denom}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.DenomMetadata, req)

	return res, err
}

func (c *chainClient) GetDenomsMetadata(ctx context.Context, pagination *query.PageRequest) (*banktypes.QueryDenomsMetadataResponse, error) {
	req := &banktypes.QueryDenomsMetadataRequest{Pagination: pagination}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.DenomsMetadata, req)

	return res, err
}

func (c *chainClient) GetDenomOwners(ctx context.Context, denom string, pagination *query.PageRequest) (*banktypes.QueryDenomOwnersResponse, error) {
	req := &banktypes.QueryDenomOwnersRequest{
		Denom:      denom,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.DenomOwners, req)

	return res, err
}

func (c *chainClient) GetBankSendEnabled(ctx context.Context, denoms []string, pagination *query.PageRequest) (*banktypes.QuerySendEnabledResponse, error) {
	req := &banktypes.QuerySendEnabledRequest{
		Denoms:     denoms,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.SendEnabled, req)

	return res, err
}

// Auth Module

func (c *chainClient) GetAccount(ctx context.Context, address string) (*authtypes.QueryAccountResponse, error) {
	req := &authtypes.QueryAccountRequest{
		Address: address,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.authQueryClient.Account, req)

	return res, err
}

// SyncBroadcastMsg sends Tx to chain and waits until Tx is included in block.
func (c *chainClient) SyncBroadcastMsg(msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error) {
	c.syncMux.Lock()
	defer c.syncMux.Unlock()

	sequence := c.getAccSeq()
	c.txFactory = c.txFactory.WithSequence(sequence)
	c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
	res, err := c.broadcastTx(c.ctx, c.txFactory, true, msgs...)

	if err != nil {
		if c.opts.ShouldFixSequenceMismatch && strings.Contains(err.Error(), "account sequence mismatch") {
			c.syncNonce()
			sequence := c.getAccSeq()
			c.txFactory = c.txFactory.WithSequence(sequence)
			c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
			log.Debugln("retrying broadcastTx with nonce", sequence)
			res, err = c.broadcastTx(c.ctx, c.txFactory, true, msgs...)
		}
		if err != nil {
			resJSON, _ := json.MarshalIndent(res, "", "\t")
			c.logger.WithField("size", len(msgs)).WithError(err).Errorln("failed synchronously broadcast messages:", string(resJSON))
			return nil, err
		}
	}

	return res, nil
}

func (c *chainClient) SimulateMsg(clientCtx client.Context, msgs ...sdk.Msg) (*txtypes.SimulateResponse, error) {
	c.txFactory = c.txFactory.WithSequence(c.accSeq)
	c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
	txf, err := PrepareFactory(clientCtx, c.txFactory)
	if err != nil {
		err = errors.Wrap(err, "failed to prepareFactory")
		return nil, err
	}

	simTxBytes, err := txf.BuildSimTx(msgs...)
	if err != nil {
		err = errors.Wrap(err, "failed to build sim tx bytes")
		return nil, err
	}

	ctx := context.Background()
	req := &txtypes.SimulateRequest{TxBytes: simTxBytes}
	simRes, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.txClient.Simulate, req)

	if err != nil {
		if c.opts.ShouldFixSequenceMismatch && strings.Contains(err.Error(), "account sequence mismatch") {
			c.syncNonce()
			sequence := c.getAccSeq()
			c.txFactory = c.txFactory.WithSequence(sequence)
			c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
			log.Debugln("retrying SimulateMsg with nonce", sequence)
			simRes, err = common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.txClient.Simulate, req)
			if err != nil {
				err = errors.Wrap(err, "failed to SimulateMsg")
				return nil, err
			}
		} else if strings.Contains(err.Error(), "account sequence mismatch") {
			fmt.Println("account sequence mismatch and ShouldFixSequenceMismatch == false")
		}
		if err != nil {
			err = errors.Wrap(err, "failed to SimulateMsg")
			return nil, err
		}
	}

	return simRes, nil
}

// AsyncBroadcastMsg sends Tx to chain and doesn't wait until Tx is included in block. This method
// cannot be used for rapid Tx sending, it is expected that you wait for transaction status with
// external tools. If you want sdk to wait for it, use SyncBroadcastMsg.
func (c *chainClient) AsyncBroadcastMsg(msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error) {
	c.syncMux.Lock()
	defer c.syncMux.Unlock()

	sequence := c.getAccSeq()
	c.txFactory = c.txFactory.WithSequence(sequence)
	c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
	res, err := c.broadcastTx(c.ctx, c.txFactory, false, msgs...)
	if err != nil {
		if c.opts.ShouldFixSequenceMismatch && strings.Contains(err.Error(), "account sequence mismatch") {
			c.syncNonce()
			sequence := c.getAccSeq()
			c.txFactory = c.txFactory.WithSequence(sequence)
			c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
			log.Debugln("retrying broadcastTx with nonce", sequence)
			res, err = c.broadcastTx(c.ctx, c.txFactory, false, msgs...)
		}
		if err != nil {
			resJSON, _ := json.MarshalIndent(res, "", "\t")
			c.logger.WithField("size", len(msgs)).WithError(err).Errorln("failed to asynchronously broadcast messagess:", string(resJSON))
			return nil, err
		}
	}

	return res, nil
}

func (c *chainClient) BuildSignedTx(clientCtx client.Context, accNum, accSeq, initialGas uint64, msgs ...sdk.Msg) ([]byte, error) {
	txf := NewTxFactory(clientCtx).WithSequence(accSeq).WithAccountNumber(accNum).WithGas(initialGas)
	return c.buildSignedTx(clientCtx, txf, msgs...)
}

func (c *chainClient) buildSignedTx(clientCtx client.Context, txf tx.Factory, msgs ...sdk.Msg) ([]byte, error) {
	ctx := context.Background()
	if clientCtx.Simulate {
		simTxBytes, err := txf.BuildSimTx(msgs...)
		if err != nil {
			err = errors.Wrap(err, "failed to build sim tx bytes")
			return nil, err
		}

		req := &txtypes.SimulateRequest{TxBytes: simTxBytes}
		simRes, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.txClient.Simulate, req)

		if err != nil {
			err = errors.Wrap(err, "failed to CalculateGas")
			return nil, err
		}

		adjustedGas := uint64(txf.GasAdjustment() * float64(simRes.GasInfo.GasUsed))
		txf = txf.WithGas(adjustedGas)

		c.gasWanted = adjustedGas
	}

	txf, err := PrepareFactory(clientCtx, txf)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepareFactory")
	}

	txn, err := txf.BuildUnsignedTx(msgs...)
	if err != nil {
		err = errors.Wrap(err, "failed to BuildUnsignedTx")
		return nil, err
	}

	txn.SetFeeGranter(clientCtx.GetFeeGranterAddress())
	err = tx.Sign(ctx, txf, clientCtx.GetFromName(), txn, true)
	if err != nil {
		err = errors.Wrap(err, "failed to Sign Tx")
		return nil, err
	}

	return clientCtx.TxConfig.TxEncoder()(txn.GetTx())
}

func (c *chainClient) SyncBroadcastSignedTx(txBytes []byte) (*txtypes.BroadcastTxResponse, error) {
	req := txtypes.BroadcastTxRequest{
		TxBytes: txBytes,
		Mode:    txtypes.BroadcastMode_BROADCAST_MODE_SYNC,
	}

	ctx := context.Background()

	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.txClient.BroadcastTx, &req)
	if err != nil || res.TxResponse.Code != 0 {
		return res, err
	}

	awaitCtx, cancelFn := context.WithTimeout(context.Background(), defaultBroadcastTimeout)
	defer cancelFn()

	txHash, _ := hex.DecodeString(res.TxResponse.TxHash)
	t := time.NewTimer(defaultBroadcastStatusPoll)

	for {
		select {
		case <-awaitCtx.Done():
			err := errors.Wrapf(ErrTimedOut, "%s", res.TxResponse.TxHash)
			t.Stop()
			return nil, err
		case <-t.C:
			resultTx, err := c.ctx.Client.Tx(awaitCtx, txHash, false)
			if err != nil {
				if errRes := client.CheckCometError(err, txBytes); errRes != nil {
					return &txtypes.BroadcastTxResponse{TxResponse: errRes}, err
				}

				t.Reset(defaultBroadcastStatusPoll)
				continue

			} else if resultTx.Height > 0 {
				resResultTx := sdk.NewResponseResultTx(resultTx, res.TxResponse.Tx, res.TxResponse.Timestamp)
				res = &txtypes.BroadcastTxResponse{TxResponse: resResultTx}
				t.Stop()
				return res, err
			}

			t.Reset(defaultBroadcastStatusPoll)
		}
	}
}

func (c *chainClient) AsyncBroadcastSignedTx(txBytes []byte) (*txtypes.BroadcastTxResponse, error) {
	req := txtypes.BroadcastTxRequest{
		TxBytes: txBytes,
		Mode:    txtypes.BroadcastMode_BROADCAST_MODE_SYNC,
	}

	ctx := context.Background()
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.txClient.BroadcastTx, &req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *chainClient) broadcastTx(
	clientCtx client.Context,
	txf tx.Factory,
	await bool,
	msgs ...sdk.Msg,
) (*txtypes.BroadcastTxResponse, error) {
	txBytes, err := c.buildSignedTx(clientCtx, txf, msgs...)
	if err != nil {
		err = errors.Wrap(err, "failed to build signed Tx")
		return nil, err
	}

	req := txtypes.BroadcastTxRequest{
		TxBytes: txBytes,
		Mode:    txtypes.BroadcastMode_BROADCAST_MODE_SYNC,
	}

	res, err := common.ExecuteCall(context.Background(), c.network.ChainCookieAssistant, c.txClient.BroadcastTx, &req)
	if err != nil || res.TxResponse.Code != 0 || !await {
		return res, err
	}

	awaitCtx, cancelFn := context.WithTimeout(context.Background(), defaultBroadcastTimeout)
	defer cancelFn()

	txHash, _ := hex.DecodeString(res.TxResponse.TxHash)
	t := time.NewTimer(defaultBroadcastStatusPoll)

	for {
		select {
		case <-awaitCtx.Done():
			err := errors.Wrapf(ErrTimedOut, "%s", res.TxResponse.TxHash)
			t.Stop()
			return nil, err
		case <-t.C:
			resultTx, err := clientCtx.Client.Tx(awaitCtx, txHash, false)
			if err != nil {
				if errRes := client.CheckCometError(err, txBytes); errRes != nil {
					return &txtypes.BroadcastTxResponse{TxResponse: errRes}, err
				}

				t.Reset(defaultBroadcastStatusPoll)
				continue

			} else if resultTx.Height > 0 {
				resResultTx := sdk.NewResponseResultTx(resultTx, res.TxResponse.Tx, res.TxResponse.Timestamp)
				res = &txtypes.BroadcastTxResponse{TxResponse: resResultTx}
				t.Stop()
				return res, err
			}

			t.Reset(defaultBroadcastStatusPoll)
		}
	}
}

// QueueBroadcastMsg enqueues a list of messages. Messages will added to the queue
// and grouped into Txns in chunks. Use this method to mass broadcast Txns with efficiency.
func (c *chainClient) QueueBroadcastMsg(msgs ...sdk.Msg) error {
	if !c.canSign {
		return ErrReadOnly
	} else if atomic.LoadInt64(&c.closed) == 1 {
		return ErrQueueClosed
	}

	t := time.NewTimer(10 * time.Second)
	for _, msg := range msgs {
		select {
		case <-t.C:
			return ErrEnqueueTimeout
		case c.msgC <- msg:
		}
	}
	t.Stop()

	return nil
}

func (c *chainClient) runBatchBroadcast() {
	expirationTimer := time.NewTimer(msgCommitBatchTimeLimit)
	msgBatch := make([]sdk.Msg, 0, msgCommitBatchSizeLimit)

	submitBatch := func(toSubmit []sdk.Msg) {
		c.syncMux.Lock()
		defer c.syncMux.Unlock()
		sequence := c.getAccSeq()
		c.txFactory = c.txFactory.WithSequence(sequence)
		c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
		log.Debugln("broadcastTx with nonce", sequence)
		res, err := c.broadcastTx(c.ctx, c.txFactory, true, toSubmit...)
		if err != nil {
			if c.opts.ShouldFixSequenceMismatch && strings.Contains(err.Error(), "account sequence mismatch") {
				c.syncNonce()
				sequence := c.getAccSeq()
				c.txFactory = c.txFactory.WithSequence(sequence)
				c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
				log.Debugln("retrying broadcastTx with nonce", sequence)
				res, err = c.broadcastTx(c.ctx, c.txFactory, true, toSubmit...)
			} else if strings.Contains(err.Error(), "account sequence mismatch") {
				fmt.Println("account sequence mismatch and ShouldFixSequenceMismatch == false")
			}
			if err != nil {
				resJSON, _ := json.MarshalIndent(res, "", "\t")
				c.logger.WithField("size", len(toSubmit)).WithError(err).Errorln("failed to broadcast messages batch:", string(resJSON))
				return
			}
		}

		if res.TxResponse.Code != 0 {
			err = errors.Errorf("error %d (%s): %s", res.TxResponse.Code, res.TxResponse.Codespace, res.TxResponse.RawLog)
			log.WithField("txHash", res.TxResponse.TxHash).WithError(err).Errorln("failed to broadcast messages batch")
		} else {
			log.WithField("txHash", res.TxResponse.TxHash).Debugln("msg batch broadcasted successfully at height", res.TxResponse.Height)
		}

		log.Debugln("gas wanted: ", c.gasWanted)
	}

	for {
		select {
		case msg, ok := <-c.msgC:
			if !ok {
				// exit required
				if len(msgBatch) > 0 {
					submitBatch(msgBatch)
				}

				close(c.doneC)
				return
			}

			msgBatch = append(msgBatch, msg)

			if len(msgBatch) >= msgCommitBatchSizeLimit {
				toSubmit := msgBatch
				msgBatch = msgBatch[:0]
				expirationTimer.Reset(msgCommitBatchTimeLimit)

				submitBatch(toSubmit)
			}
		case <-expirationTimer.C:
			if len(msgBatch) > 0 {
				toSubmit := msgBatch
				msgBatch = msgBatch[:0]
				expirationTimer.Reset(msgCommitBatchTimeLimit)
				submitBatch(toSubmit)
			} else {
				expirationTimer.Reset(msgCommitBatchTimeLimit)
			}
		}
	}
}

func (c *chainClient) GetGasFee() (string, error) {
	gasPrices := strings.Trim(c.opts.GasPrices, "helios")

	gas, err := strconv.ParseFloat(gasPrices, 64)

	if err != nil {
		return "", err
	}

	gasFeeAdjusted := gas * float64(c.gasWanted) / math.Pow(10, 18)
	gasFeeFormatted := strconv.FormatFloat(gasFeeAdjusted, 'f', -1, 64)
	c.gasFee = gasFeeFormatted

	return c.gasFee, err
}

func (c *chainClient) DefaultSubaccount(acc sdk.AccAddress) ethcommon.Hash {
	return c.Subaccount(acc, 0)
}

func (c *chainClient) Subaccount(account sdk.AccAddress, index int) ethcommon.Hash {
	ethAddress := ethcommon.BytesToAddress(account.Bytes())
	ethLowerAddress := strings.ToLower(ethAddress.String())

	subaccountId := fmt.Sprintf("%s%024x", ethLowerAddress, index)
	return ethcommon.HexToHash(subaccountId)
}

func (c *chainClient) GetTx(ctx context.Context, txHash string) (*txtypes.GetTxResponse, error) {
	req := &txtypes.GetTxRequest{
		Hash: txHash,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.txClient.GetTx, req)

	return res, err
}

// hyperion module

func (c *chainClient) GetAttestation(ctx context.Context, eventNonce uint64, claimHash []byte) (*hyperiontypes.QueryAttestationResponse, error) {
	// Ensure the return type matches the interface definition
	req := &hyperiontypes.QueryAttestationRequest{
		Nonce:     eventNonce,
		ClaimHash: claimHash,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.hyperionQueryClient.Attestation, req)

	return res, err
}

// wasm module

func (c *chainClient) FetchContractInfo(ctx context.Context, address string) (*wasmtypes.QueryContractInfoResponse, error) {
	req := &wasmtypes.QueryContractInfoRequest{
		Address: address,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.wasmQueryClient.ContractInfo, req)

	return res, err
}

func (c *chainClient) FetchContractHistory(ctx context.Context, address string, pagination *query.PageRequest) (*wasmtypes.QueryContractHistoryResponse, error) {
	req := &wasmtypes.QueryContractHistoryRequest{
		Address:    address,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.wasmQueryClient.ContractHistory, req)

	return res, err
}

func (c *chainClient) FetchContractsByCode(ctx context.Context, codeId uint64, pagination *query.PageRequest) (*wasmtypes.QueryContractsByCodeResponse, error) {
	req := &wasmtypes.QueryContractsByCodeRequest{
		CodeId:     codeId,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.wasmQueryClient.ContractsByCode, req)

	return res, err
}

func (c *chainClient) FetchAllContractsState(ctx context.Context, address string, pagination *query.PageRequest) (*wasmtypes.QueryAllContractStateResponse, error) {
	req := &wasmtypes.QueryAllContractStateRequest{
		Address:    address,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.wasmQueryClient.AllContractState, req)

	return res, err
}

func (c *chainClient) RawContractState(
	ctx context.Context,
	contractAddress string,
	queryData []byte,
) (*wasmtypes.QueryRawContractStateResponse, error) {
	req := &wasmtypes.QueryRawContractStateRequest{
		Address:   contractAddress,
		QueryData: queryData,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.wasmQueryClient.RawContractState, req)

	return res, err
}

func (c *chainClient) SmartContractState(
	ctx context.Context,
	contractAddress string,
	queryData []byte,
) (*wasmtypes.QuerySmartContractStateResponse, error) {
	req := &wasmtypes.QuerySmartContractStateRequest{
		Address:   contractAddress,
		QueryData: queryData,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.wasmQueryClient.SmartContractState, req)

	return res, err
}

func (c *chainClient) FetchCode(ctx context.Context, codeId uint64) (*wasmtypes.QueryCodeResponse, error) {
	req := &wasmtypes.QueryCodeRequest{
		CodeId: codeId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.wasmQueryClient.Code, req)

	return res, err
}

func (c *chainClient) FetchCodes(ctx context.Context, pagination *query.PageRequest) (*wasmtypes.QueryCodesResponse, error) {
	req := &wasmtypes.QueryCodesRequest{
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.wasmQueryClient.Codes, req)

	return res, err
}

func (c *chainClient) FetchPinnedCodes(ctx context.Context, pagination *query.PageRequest) (*wasmtypes.QueryPinnedCodesResponse, error) {
	req := &wasmtypes.QueryPinnedCodesRequest{
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.wasmQueryClient.PinnedCodes, req)

	return res, err
}

func (c *chainClient) FetchContractsByCreator(ctx context.Context, creator string, pagination *query.PageRequest) (*wasmtypes.QueryContractsByCreatorResponse, error) {
	req := &wasmtypes.QueryContractsByCreatorRequest{
		CreatorAddress: creator,
		Pagination:     pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.wasmQueryClient.ContractsByCreator, req)

	return res, err
}

// Tokenfactory module

func (c *chainClient) FetchDenomAuthorityMetadata(ctx context.Context, creator, subDenom string) (*tokenfactorytypes.QueryDenomAuthorityMetadataResponse, error) {
	req := &tokenfactorytypes.QueryDenomAuthorityMetadataRequest{
		Creator: creator,
	}

	if subDenom != "" {
		req.SubDenom = subDenom
	}

	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tokenfactoryQueryClient.DenomAuthorityMetadata, req)

	return res, err
}

func (c *chainClient) FetchDenomsFromCreator(ctx context.Context, creator string) (*tokenfactorytypes.QueryDenomsFromCreatorResponse, error) {
	req := &tokenfactorytypes.QueryDenomsFromCreatorRequest{
		Creator: creator,
	}

	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tokenfactoryQueryClient.DenomsFromCreator, req)

	return res, err
}

func (c *chainClient) FetchTokenfactoryModuleState(ctx context.Context) (*tokenfactorytypes.QueryModuleStateResponse, error) {
	req := &tokenfactorytypes.QueryModuleStateRequest{}

	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tokenfactoryQueryClient.TokenfactoryModuleState, req)

	return res, err
}

// Distribution module
func (c *chainClient) FetchValidatorDistributionInfo(ctx context.Context, validatorAddress string) (*distributiontypes.QueryValidatorDistributionInfoResponse, error) {
	req := &distributiontypes.QueryValidatorDistributionInfoRequest{
		ValidatorAddress: validatorAddress,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.distributionQueryClient.ValidatorDistributionInfo, req)

	return res, err
}

func (c *chainClient) FetchValidatorOutstandingRewards(ctx context.Context, validatorAddress string) (*distributiontypes.QueryValidatorOutstandingRewardsResponse, error) {
	req := &distributiontypes.QueryValidatorOutstandingRewardsRequest{
		ValidatorAddress: validatorAddress,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.distributionQueryClient.ValidatorOutstandingRewards, req)

	return res, err
}

func (c *chainClient) FetchValidatorCommission(ctx context.Context, validatorAddress string) (*distributiontypes.QueryValidatorCommissionResponse, error) {
	req := &distributiontypes.QueryValidatorCommissionRequest{
		ValidatorAddress: validatorAddress,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.distributionQueryClient.ValidatorCommission, req)

	return res, err
}

func (c *chainClient) FetchValidatorSlashes(ctx context.Context, validatorAddress string, startingHeight, endingHeight uint64, pagination *query.PageRequest) (*distributiontypes.QueryValidatorSlashesResponse, error) {
	req := &distributiontypes.QueryValidatorSlashesRequest{
		ValidatorAddress: validatorAddress,
		StartingHeight:   startingHeight,
		EndingHeight:     endingHeight,
		Pagination:       pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.distributionQueryClient.ValidatorSlashes, req)

	return res, err
}

func (c *chainClient) FetchDelegationRewards(ctx context.Context, delegatorAddress, validatorAddress string) (*distributiontypes.QueryDelegationRewardsResponse, error) {
	req := &distributiontypes.QueryDelegationRewardsRequest{
		DelegatorAddress: delegatorAddress,
		ValidatorAddress: validatorAddress,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.distributionQueryClient.DelegationRewards, req)

	return res, err
}

func (c *chainClient) FetchDelegationTotalRewards(ctx context.Context, delegatorAddress string) (*distributiontypes.QueryDelegationTotalRewardsResponse, error) {
	req := &distributiontypes.QueryDelegationTotalRewardsRequest{
		DelegatorAddress: delegatorAddress,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.distributionQueryClient.DelegationTotalRewards, req)

	return res, err
}

func (c *chainClient) FetchDelegatorValidators(ctx context.Context, delegatorAddress string) (*distributiontypes.QueryDelegatorValidatorsResponse, error) {
	req := &distributiontypes.QueryDelegatorValidatorsRequest{
		DelegatorAddress: delegatorAddress,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.distributionQueryClient.DelegatorValidators, req)

	return res, err
}

func (c *chainClient) FetchDelegatorWithdrawAddress(ctx context.Context, delegatorAddress string) (*distributiontypes.QueryDelegatorWithdrawAddressResponse, error) {
	req := &distributiontypes.QueryDelegatorWithdrawAddressRequest{
		DelegatorAddress: delegatorAddress,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.distributionQueryClient.DelegatorWithdrawAddress, req)

	return res, err
}

func (c *chainClient) FetchCommunityPool(ctx context.Context) (*distributiontypes.QueryCommunityPoolResponse, error) {
	req := &distributiontypes.QueryCommunityPoolRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.distributionQueryClient.CommunityPool, req)

	return res, err
}

// Tendermint module

func (c *chainClient) FetchNodeInfo(ctx context.Context) (*cmtservice.GetNodeInfoResponse, error) {
	req := &cmtservice.GetNodeInfoRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tendermintQueryClient.GetNodeInfo, req)

	return res, err
}

func (c *chainClient) FetchSyncing(ctx context.Context) (*cmtservice.GetSyncingResponse, error) {
	req := &cmtservice.GetSyncingRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tendermintQueryClient.GetSyncing, req)

	return res, err
}

func (c *chainClient) FetchLatestBlock(ctx context.Context) (*cmtservice.GetLatestBlockResponse, error) {
	req := &cmtservice.GetLatestBlockRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tendermintQueryClient.GetLatestBlock, req)

	return res, err
}

func (c *chainClient) FetchBlockByHeight(ctx context.Context, height int64) (*cmtservice.GetBlockByHeightResponse, error) {
	req := &cmtservice.GetBlockByHeightRequest{
		Height: height,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tendermintQueryClient.GetBlockByHeight, req)

	return res, err
}

func (c *chainClient) FetchLatestValidatorSet(ctx context.Context) (*cmtservice.GetLatestValidatorSetResponse, error) {
	req := &cmtservice.GetLatestValidatorSetRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tendermintQueryClient.GetLatestValidatorSet, req)

	return res, err
}

func (c *chainClient) FetchValidatorSetByHeight(ctx context.Context, height int64, pagination *query.PageRequest) (*cmtservice.GetValidatorSetByHeightResponse, error) {
	req := &cmtservice.GetValidatorSetByHeightRequest{
		Height:     height,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tendermintQueryClient.GetValidatorSetByHeight, req)

	return res, err
}

func (c *chainClient) ABCIQuery(ctx context.Context, path string, data []byte, height int64, prove bool) (*cmtservice.ABCIQueryResponse, error) {
	req := &cmtservice.ABCIQueryRequest{
		Path:   path,
		Data:   data,
		Height: height,
		Prove:  prove,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tendermintQueryClient.ABCIQuery, req)

	return res, err
}

// IBC Transfer module
func (c *chainClient) FetchDenomTrace(ctx context.Context, hash string) (*ibctransfertypes.QueryDenomTraceResponse, error) {
	req := &ibctransfertypes.QueryDenomTraceRequest{
		Hash: hash,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcTransferQueryClient.DenomTrace, req)

	return res, err
}

func (c *chainClient) FetchDenomTraces(ctx context.Context, pagination *query.PageRequest) (*ibctransfertypes.QueryDenomTracesResponse, error) {
	req := &ibctransfertypes.QueryDenomTracesRequest{
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcTransferQueryClient.DenomTraces, req)

	return res, err
}

func (c *chainClient) FetchDenomHash(ctx context.Context, trace string) (*ibctransfertypes.QueryDenomHashResponse, error) {
	req := &ibctransfertypes.QueryDenomHashRequest{
		Trace: trace,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcTransferQueryClient.DenomHash, req)

	return res, err
}

func (c *chainClient) FetchEscrowAddress(ctx context.Context, portId, channelId string) (*ibctransfertypes.QueryEscrowAddressResponse, error) {
	req := &ibctransfertypes.QueryEscrowAddressRequest{
		PortId:    portId,
		ChannelId: channelId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcTransferQueryClient.EscrowAddress, req)

	return res, err
}

func (c *chainClient) FetchTotalEscrowForDenom(ctx context.Context, denom string) (*ibctransfertypes.QueryTotalEscrowForDenomResponse, error) {
	req := &ibctransfertypes.QueryTotalEscrowForDenomRequest{
		Denom: denom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcTransferQueryClient.TotalEscrowForDenom, req)

	return res, err
}

// IBC Core Channel module
func (c *chainClient) FetchIBCChannel(ctx context.Context, portId, channelId string) (*ibcchanneltypes.QueryChannelResponse, error) {
	req := &ibcchanneltypes.QueryChannelRequest{
		PortId:    portId,
		ChannelId: channelId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.Channel, req)

	return res, err
}

func (c *chainClient) FetchIBCChannels(ctx context.Context, pagination *query.PageRequest) (*ibcchanneltypes.QueryChannelsResponse, error) {
	req := &ibcchanneltypes.QueryChannelsRequest{
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.Channels, req)

	return res, err
}

func (c *chainClient) FetchIBCConnectionChannels(ctx context.Context, connection string, pagination *query.PageRequest) (*ibcchanneltypes.QueryConnectionChannelsResponse, error) {
	req := &ibcchanneltypes.QueryConnectionChannelsRequest{
		Connection: connection,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.ConnectionChannels, req)

	return res, err
}

func (c *chainClient) FetchIBCChannelClientState(ctx context.Context, portId, channelId string) (*ibcchanneltypes.QueryChannelClientStateResponse, error) {
	req := &ibcchanneltypes.QueryChannelClientStateRequest{
		PortId:    portId,
		ChannelId: channelId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.ChannelClientState, req)

	return res, err
}

func (c *chainClient) FetchIBCChannelConsensusState(ctx context.Context, portId, channelId string, revisionNumber, revisionHeight uint64) (*ibcchanneltypes.QueryChannelConsensusStateResponse, error) {
	req := &ibcchanneltypes.QueryChannelConsensusStateRequest{
		PortId:         portId,
		ChannelId:      channelId,
		RevisionNumber: revisionNumber,
		RevisionHeight: revisionHeight,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.ChannelConsensusState, req)

	return res, err
}

func (c *chainClient) FetchIBCPacketCommitment(ctx context.Context, portId, channelId string, sequence uint64) (*ibcchanneltypes.QueryPacketCommitmentResponse, error) {
	req := &ibcchanneltypes.QueryPacketCommitmentRequest{
		PortId:    portId,
		ChannelId: channelId,
		Sequence:  sequence,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.PacketCommitment, req)

	return res, err
}

func (c *chainClient) FetchIBCPacketCommitments(ctx context.Context, portId, channelId string, pagination *query.PageRequest) (*ibcchanneltypes.QueryPacketCommitmentsResponse, error) {
	req := &ibcchanneltypes.QueryPacketCommitmentsRequest{
		PortId:     portId,
		ChannelId:  channelId,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.PacketCommitments, req)

	return res, err
}

func (c *chainClient) FetchIBCPacketReceipt(ctx context.Context, portId, channelId string, sequence uint64) (*ibcchanneltypes.QueryPacketReceiptResponse, error) {
	req := &ibcchanneltypes.QueryPacketReceiptRequest{
		PortId:    portId,
		ChannelId: channelId,
		Sequence:  sequence,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.PacketReceipt, req)

	return res, err
}

func (c *chainClient) FetchIBCPacketAcknowledgement(ctx context.Context, portId, channelId string, sequence uint64) (*ibcchanneltypes.QueryPacketAcknowledgementResponse, error) {
	req := &ibcchanneltypes.QueryPacketAcknowledgementRequest{
		PortId:    portId,
		ChannelId: channelId,
		Sequence:  sequence,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.PacketAcknowledgement, req)

	return res, err
}

func (c *chainClient) FetchIBCPacketAcknowledgements(ctx context.Context, portId, channelId string, packetCommitmentSequences []uint64, pagination *query.PageRequest) (*ibcchanneltypes.QueryPacketAcknowledgementsResponse, error) {
	req := &ibcchanneltypes.QueryPacketAcknowledgementsRequest{
		PortId:                    portId,
		ChannelId:                 channelId,
		Pagination:                pagination,
		PacketCommitmentSequences: packetCommitmentSequences,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.PacketAcknowledgements, req)

	return res, err
}

func (c *chainClient) FetchIBCUnreceivedPackets(ctx context.Context, portId, channelId string, packetCommitmentSequences []uint64) (*ibcchanneltypes.QueryUnreceivedPacketsResponse, error) {
	req := &ibcchanneltypes.QueryUnreceivedPacketsRequest{
		PortId:                    portId,
		ChannelId:                 channelId,
		PacketCommitmentSequences: packetCommitmentSequences,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.UnreceivedPackets, req)

	return res, err
}

func (c *chainClient) FetchIBCUnreceivedAcks(ctx context.Context, portId, channelId string, packetAckSequences []uint64) (*ibcchanneltypes.QueryUnreceivedAcksResponse, error) {
	req := &ibcchanneltypes.QueryUnreceivedAcksRequest{
		PortId:             portId,
		ChannelId:          channelId,
		PacketAckSequences: packetAckSequences,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.UnreceivedAcks, req)

	return res, err
}

func (c *chainClient) FetchIBCNextSequenceReceive(ctx context.Context, portId, channelId string) (*ibcchanneltypes.QueryNextSequenceReceiveResponse, error) {
	req := &ibcchanneltypes.QueryNextSequenceReceiveRequest{
		PortId:    portId,
		ChannelId: channelId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.NextSequenceReceive, req)

	return res, err
}

// IBC Core Chain module
func (c *chainClient) FetchIBCClientState(ctx context.Context, clientId string) (*ibcclienttypes.QueryClientStateResponse, error) {
	req := &ibcclienttypes.QueryClientStateRequest{
		ClientId: clientId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcClientQueryClient.ClientState, req)

	return res, err
}

func (c *chainClient) FetchIBCClientStates(ctx context.Context, pagination *query.PageRequest) (*ibcclienttypes.QueryClientStatesResponse, error) {
	req := &ibcclienttypes.QueryClientStatesRequest{
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcClientQueryClient.ClientStates, req)

	return res, err
}

func (c *chainClient) FetchIBCConsensusState(ctx context.Context, clientId string, revisionNumber, revisionHeight uint64, latestHeight bool) (*ibcclienttypes.QueryConsensusStateResponse, error) {
	req := &ibcclienttypes.QueryConsensusStateRequest{
		ClientId:       clientId,
		RevisionNumber: revisionNumber,
		RevisionHeight: revisionHeight,
		LatestHeight:   latestHeight,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcClientQueryClient.ConsensusState, req)

	return res, err
}

func (c *chainClient) FetchIBCConsensusStates(ctx context.Context, clientId string, pagination *query.PageRequest) (*ibcclienttypes.QueryConsensusStatesResponse, error) {
	req := &ibcclienttypes.QueryConsensusStatesRequest{
		ClientId:   clientId,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcClientQueryClient.ConsensusStates, req)

	return res, err
}

func (c *chainClient) FetchIBCConsensusStateHeights(ctx context.Context, clientId string, pagination *query.PageRequest) (*ibcclienttypes.QueryConsensusStateHeightsResponse, error) {
	req := &ibcclienttypes.QueryConsensusStateHeightsRequest{
		ClientId:   clientId,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcClientQueryClient.ConsensusStateHeights, req)

	return res, err
}

func (c *chainClient) FetchIBCClientStatus(ctx context.Context, clientId string) (*ibcclienttypes.QueryClientStatusResponse, error) {
	req := &ibcclienttypes.QueryClientStatusRequest{
		ClientId: clientId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcClientQueryClient.ClientStatus, req)

	return res, err
}

func (c *chainClient) FetchIBCClientParams(ctx context.Context) (*ibcclienttypes.QueryClientParamsResponse, error) {
	req := &ibcclienttypes.QueryClientParamsRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcClientQueryClient.ClientParams, req)

	return res, err
}

func (c *chainClient) FetchIBCUpgradedClientState(ctx context.Context) (*ibcclienttypes.QueryUpgradedClientStateResponse, error) {
	req := &ibcclienttypes.QueryUpgradedClientStateRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcClientQueryClient.UpgradedClientState, req)

	return res, err
}

func (c *chainClient) FetchIBCUpgradedConsensusState(ctx context.Context) (*ibcclienttypes.QueryUpgradedConsensusStateResponse, error) {
	req := &ibcclienttypes.QueryUpgradedConsensusStateRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcClientQueryClient.UpgradedConsensusState, req)

	return res, err
}

// IBC Core Connection module
func (c *chainClient) FetchIBCConnection(ctx context.Context, connectionId string) (*ibcconnectiontypes.QueryConnectionResponse, error) {
	req := &ibcconnectiontypes.QueryConnectionRequest{
		ConnectionId: connectionId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcConnectionQueryClient.Connection, req)

	return res, err
}

func (c *chainClient) FetchIBCConnections(ctx context.Context, pagination *query.PageRequest) (*ibcconnectiontypes.QueryConnectionsResponse, error) {
	req := &ibcconnectiontypes.QueryConnectionsRequest{
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcConnectionQueryClient.Connections, req)

	return res, err
}

func (c *chainClient) FetchIBCClientConnections(ctx context.Context, clientId string) (*ibcconnectiontypes.QueryClientConnectionsResponse, error) {
	req := &ibcconnectiontypes.QueryClientConnectionsRequest{
		ClientId: clientId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcConnectionQueryClient.ClientConnections, req)

	return res, err
}

func (c *chainClient) FetchIBCConnectionClientState(ctx context.Context, connectionId string) (*ibcconnectiontypes.QueryConnectionClientStateResponse, error) {
	req := &ibcconnectiontypes.QueryConnectionClientStateRequest{
		ConnectionId: connectionId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcConnectionQueryClient.ConnectionClientState, req)

	return res, err
}

func (c *chainClient) FetchIBCConnectionConsensusState(ctx context.Context, connectionId string, revisionNumber, revisionHeight uint64) (*ibcconnectiontypes.QueryConnectionConsensusStateResponse, error) {
	req := &ibcconnectiontypes.QueryConnectionConsensusStateRequest{
		ConnectionId:   connectionId,
		RevisionNumber: revisionNumber,
		RevisionHeight: revisionHeight,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcConnectionQueryClient.ConnectionConsensusState, req)

	return res, err
}

func (c *chainClient) FetchIBCConnectionParams(ctx context.Context) (*ibcconnectiontypes.QueryConnectionParamsResponse, error) {
	req := &ibcconnectiontypes.QueryConnectionParamsRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcConnectionQueryClient.ConnectionParams, req)

	return res, err
}

// Permissions module

func (c *chainClient) FetchAllNamespaces(ctx context.Context) (*permissionstypes.QueryAllNamespacesResponse, error) {
	req := &permissionstypes.QueryAllNamespacesRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.permissionsQueryClient.AllNamespaces, req)

	return res, err
}

func (c *chainClient) FetchNamespaceByDenom(ctx context.Context, denom string, includeRoles bool) (*permissionstypes.QueryNamespaceByDenomResponse, error) {
	req := &permissionstypes.QueryNamespaceByDenomRequest{
		Denom:        denom,
		IncludeRoles: includeRoles,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.permissionsQueryClient.NamespaceByDenom, req)

	return res, err
}

func (c *chainClient) FetchAddressRoles(ctx context.Context, denom, address string) (*permissionstypes.QueryAddressRolesResponse, error) {
	req := &permissionstypes.QueryAddressRolesRequest{
		Denom:   denom,
		Address: address,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.permissionsQueryClient.AddressRoles, req)

	return res, err
}

func (c *chainClient) FetchAddressesByRole(ctx context.Context, denom, role string) (*permissionstypes.QueryAddressesByRoleResponse, error) {
	req := &permissionstypes.QueryAddressesByRoleRequest{
		Denom: denom,
		Role:  role,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.permissionsQueryClient.AddressesByRole, req)

	return res, err
}

func (c *chainClient) FetchVouchersForAddress(ctx context.Context, address string) (*permissionstypes.QueryVouchersForAddressResponse, error) {
	req := &permissionstypes.QueryVouchersForAddressRequest{
		Address: address,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.permissionsQueryClient.VouchersForAddress, req)

	return res, err
}

func (c *chainClient) GetNetwork() common.Network {
	return c.network
}
