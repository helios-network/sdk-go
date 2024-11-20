package main

import (
	"context"
	"encoding/json"
	"fmt"

	explorerPB "github.com/Helios-Chain-Labs/sdk-go/exchange/explorer_rpc/pb"

	"github.com/Helios-Chain-Labs/sdk-go/client/common"
	explorerclient "github.com/Helios-Chain-Labs/sdk-go/client/explorer"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")
	explorerClient, err := explorerclient.NewExplorerClient(network)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	req := explorerPB.GetCw20BalanceRequest{
		Address: "helios1dc6rrxhfjaxexzdcrec5w3ryl4jn6x5t7t9j3z",
	}
	res, err := explorerClient.GetCW20Balance(ctx, &req)
	if err != nil {
		fmt.Println(err)
	}

	str, _ := json.MarshalIndent(res, "", " ")
	fmt.Print(string(str))
}
