package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Helios-Chain-Labs/sdk-go/client/common"
	exchangeclient "github.com/Helios-Chain-Labs/sdk-go/client/exchange"
	insurancePB "github.com/Helios-Chain-Labs/sdk-go/exchange/insurance_rpc/pb"
)

func main() {
	// network := common.LoadNetwork("mainnet", "k8s")
	network := common.LoadNetwork("testnet", "lb")
	exchangeClient, err := exchangeclient.NewExchangeClient(network)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	req := insurancePB.RedemptionsRequest{}

	res, err := exchangeClient.GetRedemptions(ctx, &req)
	if err != nil {
		fmt.Println(err)
	}

	str, _ := json.MarshalIndent(res, "", " ")
	fmt.Print(string(str))
}
