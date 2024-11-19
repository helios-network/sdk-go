package types

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// HELIOS defines the default coin denomination used in Ethermint in:
	//
	// - Staking parameters: denomination used as stake in the dPoS chain
	// - Mint parameters: denomination minted due to fee distribution rewards
	// - Governance parameters: denomination used for spam prevention in proposal deposits
	// - Crisis parameters: constant fee denomination used for spam prevention to check broken invariant
	// - EVM parameters: denomination used for running EVM state transitions in Ethermint.
	HeliosCoin string = "helios"

	// BaseDenomUnit defines the base denomination unit for Photons.
	// 1 photon = 1x10^{BaseDenomUnit} helios
	BaseDenomUnit = 18
)

// NewHeliosCoin is a utility function that returns an "helios" coin with the given sdkmath.Int amount.
// The function will panic if the provided amount is negative.
func NewHeliosCoin(amount sdkmath.Int) sdk.Coin {
	return sdk.NewCoin(HeliosCoin, amount)
}

// NewHeliosDecCoin is a utility function that returns an "helios" decimal coin with the given sdkmath.Int amount.
// The function will panic if the provided amount is negative.
func NewHeliosDecCoin(amount sdkmath.Int) sdk.DecCoin {
	return sdk.NewDecCoin(HeliosCoin, amount)
}

// NewHeliosCoinInt64 is a utility function that returns an "helios" coin with the given int64 amount.
// The function will panic if the provided amount is negative.
func NewHeliosCoinInt64(amount int64) sdk.Coin {
	return sdk.NewInt64Coin(HeliosCoin, amount)
}
