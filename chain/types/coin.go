package types

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// INJ defines the default coin denomination used in Ethermint in:
	//
	// - Staking parameters: denomination used as stake in the dPoS chain
	// - Mint parameters: denomination minted due to fee distribution rewards
	// - Governance parameters: denomination used for spam prevention in proposal deposits
	// - Crisis parameters: constant fee denomination used for spam prevention to check broken invariant
	// - EVM parameters: denomination used for running EVM state transitions in Ethermint.
	HeliosCoin string = "helios"

	// BaseDenomUnit defines the base denomination unit for Photons.
	// 1 photon = 1x10^{BaseDenomUnit} inj
	BaseDenomUnit = 18
)

// NewHeliosCoin is a utility function that returns an "helios" coin with the given math.Int amount.
// The function will panic if the provided amount is negative.
func NewHeliosCoin(amount math.Int) sdk.Coin {
	return sdk.NewCoin(HeliosCoin, amount)
}

// NewHeliosCoinInt64 is a utility function that returns an "helios" coin with the given int64 amount.
// The function will panic if the provided amount is negative.
func NewHeliosCoinInt64(amount int64) sdk.Coin {
	return sdk.NewInt64Coin(HeliosCoin, amount)
}
