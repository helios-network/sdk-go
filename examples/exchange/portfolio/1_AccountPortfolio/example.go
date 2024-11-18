package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Helios-Chain-Labs/sdk-go/client/common"
	exchangeclient "github.com/Helios-Chain-Labs/sdk-go/client/exchange"
)

func main() {
	// select network: local, testnet, mainnet
	network := common.LoadNetwork("testnet", "lb")
	exchangeClient, err := exchangeclient.NewExchangeClient(network)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	accountAddress := "inj1clw20s2uxeyxtam6f7m84vgae92s9eh7vygagt"
	res, err := exchangeClient.GetAccountPortfolioBalances(ctx, accountAddress)
	if err != nil {
		fmt.Println(err)
	}

	str, _ := json.MarshalIndent(res, "", " ")
	fmt.Print(string(str))
}
