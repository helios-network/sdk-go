package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Helios-Chain-Labs/sdk-go/client/common"
	exchangeclient "github.com/Helios-Chain-Labs/sdk-go/client/exchange"
)

func main() {
	network := common.LoadNetwork("mainnet", "lb")
	exchangeClient, err := exchangeclient.NewExchangeClient(network)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	res, err := exchangeClient.GetOracleList(ctx)
	if err != nil {
		fmt.Println(err)
	}

	str, _ := json.MarshalIndent(res, "", " ")
	fmt.Print(string(str))
}
