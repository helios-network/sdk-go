package main

import (
	"context"
	"encoding/json"
	"fmt"

	chainStreamModule "github.com/Helios-Chain-Labs/sdk-go/chain/stream/types"
	"github.com/Helios-Chain-Labs/sdk-go/client"
	chainclient "github.com/Helios-Chain-Labs/sdk-go/client/chain"
	"github.com/Helios-Chain-Labs/sdk-go/client/common"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")

	clientCtx, err := chainclient.NewClientContext(
		network.ChainId,
		"",
		nil,
	)
	if err != nil {
		panic(err)
	}
	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint)

	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network,
		common.OptionGasPrices(client.DefaultGasPriceWithDenom),
	)

	if err != nil {
		panic(err)
	}

	subaccountId := "0xbdaedec95d563fb05240d6e01821008454c24c36000000000000000000000000"

	heliosUsdtMarket := "0x0611780ba69656949525013d947713300f56c37b6175e02f26bffa495c3208fe"
	heliosUsdtPerpMarket := "0x17ef48032cb24375ba7c2e39f384e56433bcab20cbee9a7357e4cba2eb00abe6"

	req := chainStreamModule.StreamRequest{
		BankBalancesFilter: &chainStreamModule.BankBalancesFilter{
			Accounts: []string{"*"},
		},
		SpotOrdersFilter: &chainStreamModule.OrdersFilter{
			MarketIds:     []string{heliosUsdtMarket},
			SubaccountIds: []string{subaccountId},
		},
		DerivativeOrdersFilter: &chainStreamModule.OrdersFilter{
			MarketIds:     []string{heliosUsdtPerpMarket},
			SubaccountIds: []string{subaccountId},
		},
		SpotTradesFilter: &chainStreamModule.TradesFilter{
			MarketIds:     []string{heliosUsdtMarket},
			SubaccountIds: []string{"*"},
		},
		SubaccountDepositsFilter: &chainStreamModule.SubaccountDepositsFilter{
			SubaccountIds: []string{subaccountId},
		},
		DerivativeOrderbooksFilter: &chainStreamModule.OrderbookFilter{
			MarketIds: []string{heliosUsdtPerpMarket},
		},
		SpotOrderbooksFilter: &chainStreamModule.OrderbookFilter{
			MarketIds: []string{heliosUsdtMarket},
		},
		PositionsFilter: &chainStreamModule.PositionsFilter{
			SubaccountIds: []string{subaccountId},
			MarketIds:     []string{heliosUsdtPerpMarket},
		},
		DerivativeTradesFilter: &chainStreamModule.TradesFilter{
			SubaccountIds: []string{"*"},
			MarketIds:     []string{heliosUsdtPerpMarket},
		},
		OraclePriceFilter: &chainStreamModule.OraclePriceFilter{
			Symbol: []string{"HELIOS", "USDT"},
		},
	}

	ctx := context.Background()

	stream, err := chainClient.ChainStream(ctx, req)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case <-ctx.Done():
			return
		default:
			res, err := stream.Recv()
			if err != nil {
				panic(err)
				return
			}
			str, _ := json.MarshalIndent(res, "", " ")
			fmt.Print(string(str))
		}
	}
}
