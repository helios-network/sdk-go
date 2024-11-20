package main

import (
	"context"
	"fmt"

	"github.com/Helios-Chain-Labs/sdk-go/client/common"
	exchangeclient "github.com/Helios-Chain-Labs/sdk-go/client/exchange"
)

func main() {
	network := common.LoadNetwork("devnet-1", "")
	exchangeClient, err := exchangeclient.NewExchangeClient(network)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	marketIds := []string{"0x4ca0f92fc28be0c9761326016b5a1a2177dd6375558365116b5bdda9abc229ce"}
	stream, err := exchangeClient.StreamDerivativeOrderbookV2(ctx, marketIds)
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
				fmt.Println(err)
				return
			}
			fmt.Println(res.MarketId, len(res.Orderbook.Sells), len(res.Orderbook.Buys))
		}
	}
}
