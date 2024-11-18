package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Helios-Chain-Labs/sdk-go/client/common"
	explorerclient "github.com/Helios-Chain-Labs/sdk-go/client/explorer"
)

func main() {
	network := common.LoadNetwork("mainnet", "lb")
	explorerClient, err := explorerclient.NewExplorerClient(network)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	res, err := explorerClient.GetBlocks(ctx)
	if err != nil {
		fmt.Println(err)
	}

	str, _ := json.MarshalIndent(res, "", " ")
	fmt.Print(string(str))
}
