package common

import (
	ctypes "github.com/Helios-Chain-Labs/sdk-go/chain/types"
	log "github.com/Helios-Chain-Labs/suplog"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	"google.golang.org/grpc/credentials"
)

func init() {
	// set the address prefixes
	config := sdk.GetConfig()

	// This is specific to Helios chain
	ctypes.SetBech32Prefixes(config)
	ctypes.SetBip44CoinType(config)
}

type ClientOptions struct {
	GasPrices                 string
	Gas                       string
	TLSCert                   credentials.TransportCredentials
	TxFactory                 *tx.Factory
	ShouldFixSequenceMismatch bool
}

type ClientOption func(opts *ClientOptions) error

func DefaultClientOptions() *ClientOptions {
	return &ClientOptions{
		ShouldFixSequenceMismatch: true,
	}
}

func OptionGasPrices(gasPrices string) ClientOption {
	return func(opts *ClientOptions) error {
		_, err := sdk.ParseDecCoins(gasPrices)
		if err != nil {
			err = errors.Wrapf(err, "failed to ParseDecCoins %s", gasPrices)
			return err
		}

		opts.GasPrices = gasPrices
		return nil
	}
}

func OptionGas(gas string) ClientOption {
	return func(opts *ClientOptions) error {
		opts.Gas = gas
		return nil
	}
}

func OptionTLSCert(tlsCert credentials.TransportCredentials) ClientOption {
	return func(opts *ClientOptions) error {
		if tlsCert == nil {
			log.Infoln("client does not use grpc secure transport")
		} else {
			log.Infoln("successfully load server TLS cert")
		}
		opts.TLSCert = tlsCert
		return nil
	}
}

func OptionTxFactory(txFactory *tx.Factory) ClientOption {
	return func(opts *ClientOptions) error {
		opts.TxFactory = txFactory
		return nil
	}
}
