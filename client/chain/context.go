package chain

import (
	"os"

	"cosmossdk.io/x/tx/signing"
	"github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/gogoproto/proto"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/pkg/errors"

	keyscodec "github.com/Helios-Chain-Labs/sdk-go/chain/crypto/codec"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"

	hyperion "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
	tokenfactory "github.com/Helios-Chain-Labs/sdk-go/chain/tokenfactory/types"
	chaintypes "github.com/Helios-Chain-Labs/sdk-go/chain/types"

	evidencetypes "cosmossdk.io/x/evidence/types"
	feegranttypes "cosmossdk.io/x/feegrant"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govv1types "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	paramproposaltypes "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	icatypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/types"
	ibcapplicationtypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibccoretypes "github.com/cosmos/ibc-go/v8/modules/core/types"
	ibclightclienttypes "github.com/cosmos/ibc-go/v8/modules/light-clients/06-solomachine"
	ibctenderminttypes "github.com/cosmos/ibc-go/v8/modules/light-clients/07-tendermint"
)

// NewInterfaceRegistry returns a new InterfaceRegistry
func NewInterfaceRegistry() types.InterfaceRegistry {
	registry, err := types.NewInterfaceRegistryWithOptions(types.InterfaceRegistryOptions{
		ProtoFiles: proto.HybridResolver,
		SigningOptions: signing.Options{
			AddressCodec: address.Bech32Codec{
				Bech32Prefix: cosmostypes.GetConfig().GetBech32AccountAddrPrefix(),
			},
			ValidatorAddressCodec: address.Bech32Codec{
				Bech32Prefix: cosmostypes.GetConfig().GetBech32ValidatorAddrPrefix(),
			},
		},
	})
	if err != nil {
		panic(err)
	}
	return registry
}

// NewTxConfig initializes new Cosmos TxConfig with certain signModes enabled.
func NewTxConfig(signModes []signingtypes.SignMode) client.TxConfig {
	interfaceRegistry := NewInterfaceRegistry()
	keyscodec.RegisterInterfaces(interfaceRegistry)
	std.RegisterInterfaces(interfaceRegistry)
	hyperion.RegisterInterfaces(interfaceRegistry)
	chaintypes.RegisterInterfaces(interfaceRegistry)
	tokenfactory.RegisterInterfaces(interfaceRegistry)

	// more cosmos types
	authtypes.RegisterInterfaces(interfaceRegistry)
	authztypes.RegisterInterfaces(interfaceRegistry)
	vestingtypes.RegisterInterfaces(interfaceRegistry)
	banktypes.RegisterInterfaces(interfaceRegistry)
	crisistypes.RegisterInterfaces(interfaceRegistry)
	distributiontypes.RegisterInterfaces(interfaceRegistry)
	evidencetypes.RegisterInterfaces(interfaceRegistry)
	govtypes.RegisterInterfaces(interfaceRegistry)
	govv1types.RegisterInterfaces(interfaceRegistry)
	paramproposaltypes.RegisterInterfaces(interfaceRegistry)
	ibcapplicationtypes.RegisterInterfaces(interfaceRegistry)
	ibccoretypes.RegisterInterfaces(interfaceRegistry)
	ibclightclienttypes.RegisterInterfaces(interfaceRegistry)
	ibctenderminttypes.RegisterInterfaces(interfaceRegistry)
	slashingtypes.RegisterInterfaces(interfaceRegistry)
	stakingtypes.RegisterInterfaces(interfaceRegistry)
	upgradetypes.RegisterInterfaces(interfaceRegistry)
	feegranttypes.RegisterInterfaces(interfaceRegistry)
	wasmtypes.RegisterInterfaces(interfaceRegistry)
	icatypes.RegisterInterfaces(interfaceRegistry)

	marshaler := codec.NewProtoCodec(interfaceRegistry)
	return tx.NewTxConfig(marshaler, signModes)
}

// NewClientContext creates a new Cosmos Client context, where chainID
// corresponds to Cosmos chain ID, fromSpec is either name of the key, or bech32-address
// of the Cosmos account. Keyring is required to contain the specified key.
func NewClientContext(
	chainId, fromSpec string, kb keyring.Keyring,
) (client.Context, error) {
	clientCtx := client.Context{}

	interfaceRegistry := NewInterfaceRegistry()
	keyscodec.RegisterInterfaces(interfaceRegistry)
	std.RegisterInterfaces(interfaceRegistry)
	hyperion.RegisterInterfaces(interfaceRegistry)
	chaintypes.RegisterInterfaces(interfaceRegistry)
	tokenfactory.RegisterInterfaces(interfaceRegistry)

	// more cosmos types
	authtypes.RegisterInterfaces(interfaceRegistry)
	authztypes.RegisterInterfaces(interfaceRegistry)
	vestingtypes.RegisterInterfaces(interfaceRegistry)
	banktypes.RegisterInterfaces(interfaceRegistry)
	crisistypes.RegisterInterfaces(interfaceRegistry)
	distributiontypes.RegisterInterfaces(interfaceRegistry)
	evidencetypes.RegisterInterfaces(interfaceRegistry)
	govtypes.RegisterInterfaces(interfaceRegistry)
	govv1types.RegisterInterfaces(interfaceRegistry)
	paramproposaltypes.RegisterInterfaces(interfaceRegistry)
	ibcapplicationtypes.RegisterInterfaces(interfaceRegistry)
	ibccoretypes.RegisterInterfaces(interfaceRegistry)
	ibclightclienttypes.RegisterInterfaces(interfaceRegistry)
	ibctenderminttypes.RegisterInterfaces(interfaceRegistry)
	slashingtypes.RegisterInterfaces(interfaceRegistry)
	stakingtypes.RegisterInterfaces(interfaceRegistry)
	upgradetypes.RegisterInterfaces(interfaceRegistry)
	feegranttypes.RegisterInterfaces(interfaceRegistry)
	wasmtypes.RegisterInterfaces(interfaceRegistry)
	icatypes.RegisterInterfaces(interfaceRegistry)
	ibcfeetypes.RegisterInterfaces(interfaceRegistry)

	marshaler := codec.NewProtoCodec(interfaceRegistry)
	encodingConfig := EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig: NewTxConfig([]signingtypes.SignMode{
			signingtypes.SignMode_SIGN_MODE_DIRECT,
		}),
	}

	var keyInfo keyring.Record

	if kb != nil {
		addr, err := cosmostypes.AccAddressFromBech32(fromSpec)
		if err == nil {
			record, err := kb.KeyByAddress(addr)
			if err != nil {
				err = errors.Wrapf(err, "failed to load key info by address %s", addr.String())
				return clientCtx, err
			}
			keyInfo = *record
		} else {
			// failed to parse Bech32, is it a name?
			record, err := kb.Key(fromSpec)
			if err != nil {
				err = errors.Wrapf(err, "no key in keyring for name: %s", fromSpec)
				return clientCtx, err
			}
			keyInfo = *record
		}
	}

	clientCtx = newContext(
		chainId,
		encodingConfig,
		kb,
		keyInfo,
	)

	return clientCtx, nil
}

type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Marshaler         codec.Codec
	TxConfig          client.TxConfig
}

func newContext(
	chainId string,
	encodingConfig EncodingConfig,
	kb keyring.Keyring,
	keyInfo keyring.Record,

) client.Context {
	clientCtx := client.Context{
		ChainID:           chainId,
		Codec:             encodingConfig.Marshaler,
		InterfaceRegistry: encodingConfig.InterfaceRegistry,
		Output:            os.Stderr,
		OutputFormat:      "json",
		BroadcastMode:     "block",
		UseLedger:         false,
		Simulate:          true,
		GenerateOnly:      false,
		Offline:           false,
		SkipConfirm:       true,
		TxConfig:          encodingConfig.TxConfig,
		AccountRetriever:  authtypes.AccountRetriever{},
	}

	if keyInfo.PubKey != nil {
		keyInfoAddress, err := keyInfo.GetAddress()
		if err != nil {
			panic(err)
		}
		clientCtx = clientCtx.WithKeyring(kb)
		clientCtx = clientCtx.WithFromAddress(keyInfoAddress)
		clientCtx = clientCtx.WithFromName(keyInfo.Name)
		clientCtx = clientCtx.WithFrom(keyInfo.Name)
	}

	return clientCtx
}
