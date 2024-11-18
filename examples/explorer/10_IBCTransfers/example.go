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
	receiver := "inj1ddcp5ftqmntudn4m6heg2adud6hn58urnwlmkh"

	req := explorerPB.GetIBCTransferTxsRequest{
		Receiver: receiver,
	}

	res, err := explorerClient.GetIBCTransfers(ctx, &req)
	if err != nil {
		fmt.Println(err)
	}

	str, _ := json.MarshalIndent(res, "", " ")
	fmt.Print(string(str))
}
