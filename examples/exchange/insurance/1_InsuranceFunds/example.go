package main

import (
	"context"
	"encoding/json"
	"fmt"

	"sdk-go/client/common"
	exchangeclient "sdk-go/client/exchange"
	insurancePB "sdk-go/exchange/insurance_rpc/pb"
)

func main() {
	// network := common.LoadNetwork("mainnet", "k8s")
	network := common.LoadNetwork("testnet", "lb")
	exchangeClient, err := exchangeclient.NewExchangeClient(network)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	req := insurancePB.FundsRequest{}

	res, err := exchangeClient.GetInsuranceFunds(ctx, &req)
	if err != nil {
		fmt.Println(err)
	}

	str, _ := json.MarshalIndent(res, "", " ")
	fmt.Print(string(str))
}
