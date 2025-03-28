package evmv1

import (
	"math/big"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func DeriveChainID(v *big.Int) *big.Int {
	if v == nil || v.Sign() < 1 {
		return nil
	}

	if v.BitLen() <= 64 {
		v := v.Uint64()
		if v == 27 || v == 28 {
			return new(big.Int)
		}

		if v < 35 {
			return nil
		}

		// V MUST be of the form {0,1} + CHAIN_ID * 2 + 35
		return new(big.Int).SetUint64((v - 35) / 2)
	}
	v = new(big.Int).Sub(v, big.NewInt(35))
	return v.Div(v, big.NewInt(2))
}

// GetChainID returns the chain id field from the derived signature values
func (tx *LegacyTx) GetChainID() *big.Int {
	v, _, _ := tx.GetRawSignatureValues()
	return DeriveChainID(v)
}

// AsEthereumData returns an LegacyTx transaction tx from the proto-formatted
// TxData defined on the Cosmos EVM.
func (tx *LegacyTx) AsEthereumData() ethtypes.TxData {
	v, r, s := tx.GetRawSignatureValues()
	return &ethtypes.LegacyTx{
		Nonce:    tx.GetNonce(),
		GasPrice: stringToBigInt(tx.GetGasPrice()),
		Gas:      tx.GetGas(),
		To:       stringToAddress(tx.GetTo()),
		Value:    stringToBigInt(tx.GetValue()),
		Data:     tx.GetData(),
		V:        v,
		R:        r,
		S:        s,
	}
}

// GetRawSignatureValues returns the V, R, S signature values of the transaction.
// The return values should not be modified by the caller.
func (tx *LegacyTx) GetRawSignatureValues() (v, r, s *big.Int) {
	return RawSignatureValues(tx.V, tx.R, tx.S)
}

// GetAccessList returns nil
func (tx *LegacyTx) GetAccessList() ethtypes.AccessList {
	return nil
}
