package main

import (
	"encoding/json"
	"fmt"
	"os"

	"cosmossdk.io/math"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"

	exchangetypes "github.com/Helios-Chain-Labs/sdk-go/chain/exchange/types"
	oracletypes "github.com/Helios-Chain-Labs/sdk-go/chain/oracle/types"
	"github.com/Helios-Chain-Labs/sdk-go/client"
	chainclient "github.com/Helios-Chain-Labs/sdk-go/client/chain"
	"github.com/Helios-Chain-Labs/sdk-go/client/common"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")
	tmClient, err := rpchttp.New(network.TmEndpoint, "/websocket")
	if err != nil {
		panic(err)
	}

	senderAddress, cosmosKeyring, err := chainclient.InitCosmosKeyring(
		os.Getenv("HOME")+"/.heliades",
		"heliades",
		"file",
		"helios-user",
		"12345678",
		"5d386fbdbf11f1141010f81a46b40f94887367562bd33b452bbaa6ce1cd1381e", // keyring will be used if pk not provided
		false,
	)

	if err != nil {
		panic(err)
	}

	clientCtx, err := chainclient.NewClientContext(
		network.ChainId,
		senderAddress.String(),
		cosmosKeyring,
	)

	if err != nil {
		panic(err)
	}

	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint).WithClient(tmClient)

	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network,
		common.OptionGasPrices(client.DefaultGasPriceWithDenom),
	)
	if err != nil {
		panic(err)
	}

	minPriceTickSize := math.LegacyMustNewDecFromStr("0.01")
	minQuantityTickSize := math.LegacyMustNewDecFromStr("0.001")

	chainMinPriceTickSize := minPriceTickSize.Mul(math.LegacyNewDecFromIntWithPrec(math.NewInt(1), int64(6)))
	chainMinQuantityTickSize := minQuantityTickSize

	msg := &exchangetypes.MsgInstantExpiryFuturesMarketLaunch{
		Sender:                 senderAddress.String(),
		Ticker:                 "HELIOS/USDC FUT",
		QuoteDenom:             "factory/helios17vytdwqczqz72j65saukplrktd4gyfme5agf6c/usdc",
		OracleBase:             "HELIOS",
		OracleQuote:            "USDC",
		OracleScaleFactor:      6,
		OracleType:             oracletypes.OracleType_Band,
		Expiry:                 2000000000,
		MakerFeeRate:           math.LegacyMustNewDecFromStr("-0.0001"),
		TakerFeeRate:           math.LegacyMustNewDecFromStr("0.001"),
		InitialMarginRatio:     math.LegacyMustNewDecFromStr("0.33"),
		MaintenanceMarginRatio: math.LegacyMustNewDecFromStr("0.095"),
		MinPriceTickSize:       chainMinPriceTickSize,
		MinQuantityTickSize:    chainMinQuantityTickSize,
	}

	// AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	response, err := chainClient.AsyncBroadcastMsg(msg)

	if err != nil {
		panic(err)
	}

	str, _ := json.MarshalIndent(response, "", " ")
	fmt.Print(string(str))
}
