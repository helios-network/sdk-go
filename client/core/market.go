package core

import (
	"github.com/Helios-Chain-Labs/sdk-go/client/common"
	cosmossdk_io_math "cosmossdk.io/math"
	"github.com/shopspring/decimal"
)

const AdditionalChainFormatDecimals = 18

type SpotMarket struct {
	Id                  string
	Status              string
	Ticker              string
	BaseToken           Token
	QuoteToken          Token
	MakerFeeRate        decimal.Decimal
	TakerFeeRate        decimal.Decimal
	ServiceProviderFee  decimal.Decimal
	MinPriceTickSize    decimal.Decimal
	MinQuantityTickSize decimal.Decimal
}

func (spotMarket SpotMarket) QuantityToChainFormat(humanReadableValue decimal.Decimal) cosmossdk_io_math.LegacyDec {
	chainFormattedValue := humanReadableValue.Mul(decimal.New(1, spotMarket.BaseToken.Decimals))
	quantizedValue := chainFormattedValue.DivRound(spotMarket.MinQuantityTickSize, 0).Mul(spotMarket.MinQuantityTickSize)
	valueInChainFormat, _ := cosmossdk_io_math.NewDecFromStr(quantizedValue.String())

	return valueInChainFormat
}

func (spotMarket SpotMarket) PriceToChainFormat(humanReadableValue decimal.Decimal) cosmossdk_io_math.LegacyDec {
	decimals := spotMarket.QuoteToken.Decimals - spotMarket.BaseToken.Decimals
	chainFormattedValue := humanReadableValue.Mul(decimal.New(1, decimals))
	quantizedValue := chainFormattedValue.DivRound(spotMarket.MinPriceTickSize, 0).Mul(spotMarket.MinPriceTickSize)
	valueInChainFormat, _ := cosmossdk_io_math.NewDecFromStr(quantizedValue.String())

	return valueInChainFormat
}

func (spotMarket SpotMarket) QuantityFromChainFormat(chainValue cosmossdk_io_math.LegacyDec) decimal.Decimal {
	return decimal.RequireFromString(chainValue.String()).Div(decimal.New(1, spotMarket.BaseToken.Decimals))
}

func (spotMarket SpotMarket) PriceFromChainFormat(chainValue cosmossdk_io_math.LegacyDec) decimal.Decimal {
	decimals := spotMarket.BaseToken.Decimals - spotMarket.QuoteToken.Decimals
	return decimal.RequireFromString(chainValue.String()).Mul(decimal.New(1, decimals))
}

func (spotMarket SpotMarket) QuantityFromExtendedChainFormat(chainValue cosmossdk_io_math.LegacyDec) decimal.Decimal {
	return common.RemoveExtraDecimals(spotMarket.QuantityFromChainFormat(chainValue), AdditionalChainFormatDecimals)
}

func (spotMarket SpotMarket) PriceFromExtendedChainFormat(chainValue cosmossdk_io_math.LegacyDec) decimal.Decimal {
	return common.RemoveExtraDecimals(spotMarket.PriceFromChainFormat(chainValue), AdditionalChainFormatDecimals)
}

type DerivativeMarket struct {
	Id                     string
	Status                 string
	Ticker                 string
	OracleBase             string
	OracleQuote            string
	OracleType             string
	OracleScaleFactor      uint32
	InitialMarginRatio     decimal.Decimal
	MaintenanceMarginRatio decimal.Decimal
	QuoteToken             Token
	MakerFeeRate           decimal.Decimal
	TakerFeeRate           decimal.Decimal
	ServiceProviderFee     decimal.Decimal
	MinPriceTickSize       decimal.Decimal
	MinQuantityTickSize    decimal.Decimal
}

func (derivativeMarket DerivativeMarket) QuantityToChainFormat(humanReadableValue decimal.Decimal) cosmossdk_io_math.LegacyDec {
	chainFormattedValue := humanReadableValue
	quantizedValue := chainFormattedValue.DivRound(derivativeMarket.MinQuantityTickSize, 0).Mul(derivativeMarket.MinQuantityTickSize)
	valueInChainFormat, _ := cosmossdk_io_math.NewDecFromStr(quantizedValue.String())

	return valueInChainFormat
}

func (derivativeMarket DerivativeMarket) PriceToChainFormat(humanReadableValue decimal.Decimal) cosmossdk_io_math.LegacyDec {
	decimals := derivativeMarket.QuoteToken.Decimals
	chainFormattedValue := humanReadableValue.Mul(decimal.New(1, decimals))
	quantizedValue := chainFormattedValue.DivRound(derivativeMarket.MinPriceTickSize, 0).Mul(derivativeMarket.MinPriceTickSize)
	valueInChainFormat, _ := cosmossdk_io_math.NewDecFromStr(quantizedValue.String())

	return valueInChainFormat
}

func (derivativeMarket DerivativeMarket) MarginToChainFormat(humanReadableValue decimal.Decimal) cosmossdk_io_math.LegacyDec {
	decimals := derivativeMarket.QuoteToken.Decimals
	chainFormattedValue := humanReadableValue.Mul(decimal.New(1, decimals))
	quantizedValue := chainFormattedValue.DivRound(derivativeMarket.MinQuantityTickSize, 0).Mul(derivativeMarket.MinQuantityTickSize)
	valueInChainFormat, _ := cosmossdk_io_math.NewDecFromStr(quantizedValue.String())

	return valueInChainFormat
}

func (derivativeMarket DerivativeMarket) CalculateMarginInChainFormat(humanReadableQuantity decimal.Decimal, humanReadablePrice decimal.Decimal, leverage decimal.Decimal) cosmossdk_io_math.LegacyDec {
	chainFormattedQuantity := humanReadableQuantity
	chainFormattedPrice := humanReadablePrice.Mul(decimal.New(1, derivativeMarket.QuoteToken.Decimals))

	margin := chainFormattedQuantity.Mul(chainFormattedPrice).Div(leverage)
	// We are using the min_quantity_tick_size to quantize the margin because that is the way margin is validated
	// in the chain (it might be changed to a min_notional in the future)
	quantizedMargin := margin.DivRound(derivativeMarket.MinQuantityTickSize, 0).Mul(derivativeMarket.MinQuantityTickSize)
	valueInChainFormat, _ := cosmossdk_io_math.NewDecFromStr(quantizedMargin.String())

	return valueInChainFormat
}

func (derivativeMarket DerivativeMarket) QuantityFromChainFormat(chainValue cosmossdk_io_math.LegacyDec) decimal.Decimal {
	return decimal.RequireFromString(chainValue.String())
}

func (derivativeMarket DerivativeMarket) PriceFromChainFormat(chainValue cosmossdk_io_math.LegacyDec) decimal.Decimal {
	decimals := -derivativeMarket.QuoteToken.Decimals
	return decimal.RequireFromString(chainValue.String()).Mul(decimal.New(1, decimals))
}

func (derivativeMarket DerivativeMarket) MarginFromChainFormat(chainValue cosmossdk_io_math.LegacyDec) decimal.Decimal {
	decimals := -derivativeMarket.QuoteToken.Decimals
	return decimal.RequireFromString(chainValue.String()).Mul(decimal.New(1, decimals))
}

func (derivativeMarket DerivativeMarket) QuantityFromExtendedChainFormat(chainValue cosmossdk_io_math.LegacyDec) decimal.Decimal {
	return common.RemoveExtraDecimals(derivativeMarket.QuantityFromChainFormat(chainValue), AdditionalChainFormatDecimals)
}

func (derivativeMarket DerivativeMarket) PriceFromExtendedChainFormat(chainValue cosmossdk_io_math.LegacyDec) decimal.Decimal {
	return common.RemoveExtraDecimals(derivativeMarket.PriceFromChainFormat(chainValue), AdditionalChainFormatDecimals)
}

func (derivativeMarket DerivativeMarket) MarginFromExtendedChainFormat(chainValue cosmossdk_io_math.LegacyDec) decimal.Decimal {
	return common.RemoveExtraDecimals(derivativeMarket.MarginFromChainFormat(chainValue), AdditionalChainFormatDecimals)
}
