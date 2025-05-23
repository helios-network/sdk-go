package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"

	authzcdc "github.com/cosmos/cosmos-sdk/x/authz/codec"
)

// Legacy Amino Codec for JSON Serialization
var amino = codec.NewLegacyAmino()

// Protobuf Codec for Message Serialization
var ModuleCdc = codec.NewProtoCodec(codectypes.NewInterfaceRegistry())

// Amino JSON codec for backwards compatibility
var AminoCdc = codec.NewAminoCodec(amino) //nolint:staticcheck

func init() {
	RegisterLegacyAminoCodec(amino)
	RegisterLegacyAminoCodec(authzcdc.Amino)
	amino.Seal()
}

// RegisterInterfaces registers the interfaces for the proto stuff
func RegisterInterfaces(registry types.InterfaceRegistry) {
	// Register Hyperion Messages
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgValsetConfirm{},
		&MsgSendToChain{},
		&MsgRequestBatch{},
		&MsgConfirmBatch{},
		&MsgDepositClaim{},
		&MsgWithdrawClaim{},
		&MsgERC20DeployedClaim{},
		&MsgSetOrchestratorAddresses{},
		&MsgValsetUpdatedClaim{},
		&MsgCancelSendToChain{},
		&MsgSubmitBadSignatureEvidence{},
		&MsgUpdateParams{},
		&MsgBlacklistEthereumAddresses{},
		&MsgRevokeEthereumBlacklist{},
	)

	registry.RegisterInterface(
		"hyperion.v1beta1.EthereumClaim",
		(*EthereumClaim)(nil),
		&MsgDepositClaim{},
		&MsgWithdrawClaim{},
		&MsgERC20DeployedClaim{},
		&MsgValsetUpdatedClaim{},
	)

	// Register MsgService Descriptor
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// RegisterLegacyAminoCodec registers the necessary Hyperion messages for JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSetOrchestratorAddresses{}, "hyperion/MsgSetOrchestratorAddresses", nil)
	cdc.RegisterConcrete(&MsgValsetConfirm{}, "hyperion/MsgValsetConfirm", nil)
	cdc.RegisterConcrete(&MsgSendToChain{}, "hyperion/MsgSendToChain", nil)
	cdc.RegisterConcrete(&MsgCancelSendToChain{}, "hyperion/MsgCancelSendToChain", nil)
	cdc.RegisterConcrete(&MsgRequestBatch{}, "hyperion/MsgRequestBatch", nil)
	cdc.RegisterConcrete(&MsgConfirmBatch{}, "hyperion/MsgConfirmBatch", nil)
	cdc.RegisterConcrete(&Valset{}, "hyperion/Valset", nil)
	cdc.RegisterConcrete(&MsgDepositClaim{}, "hyperion/MsgDepositClaim", nil)
	cdc.RegisterConcrete(&MsgWithdrawClaim{}, "hyperion/MsgWithdrawClaim", nil)
	cdc.RegisterConcrete(&MsgERC20DeployedClaim{}, "hyperion/MsgERC20DeployedClaim", nil)
	cdc.RegisterConcrete(&MsgValsetUpdatedClaim{}, "hyperion/MsgValsetUpdatedClaim", nil)
	cdc.RegisterConcrete(&OutgoingTxBatch{}, "hyperion/OutgoingTxBatch", nil)
	cdc.RegisterConcrete(&OutgoingTransferTx{}, "hyperion/OutgoingTransferTx", nil)
	cdc.RegisterConcrete(&Token{}, "hyperion/ERC20Token", nil)
	cdc.RegisterConcrete(&IDSet{}, "hyperion/IDSet", nil)
	cdc.RegisterConcrete(&Attestation{}, "hyperion/Attestation", nil)
	cdc.RegisterConcrete(&MsgSubmitBadSignatureEvidence{}, "hyperion/MsgSubmitBadSignatureEvidence", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "hyperion/MsgUpdateParams", nil)
	cdc.RegisterConcrete(&MsgBlacklistEthereumAddresses{}, "hyperion/MsgBlacklistEthereumAddresses", nil)
	cdc.RegisterConcrete(&MsgRevokeEthereumBlacklist{}, "hyperion/MsgRevokeEthereumBlacklist", nil)
	cdc.RegisterConcrete(&Params{}, "hyperion/Params", nil)
}
