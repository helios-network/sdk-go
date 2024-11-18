package main

import (
	"context"
	"encoding/json"
	"fmt"

	"sdk-go/client/common"
	exchangeclient "sdk-go/client/exchange"
	metaPB "sdk-go/exchange/meta_rpc/pb"
)

func main() {
	// network := common.LoadNetwork("mainnet", "k8s")
	network := common.LoadNetwork("testnet", "lb")
	exchangeClient, err := exchangeclient.NewExchangeClient(network)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	req := metaPB.InfoRequest{}

	res, err := exchangeClient.GetInfo(ctx, &req)
	if err != nil {
		fmt.Println(err)
	}

	str, _ := json.MarshalIndent(res, "", " ")
	fmt.Print(string(str))
}
