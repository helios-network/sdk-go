package main

import (
	"context"
	"encoding/json"
	"fmt"

	explorerPB "sdk-go/exchange/explorer_rpc/pb"

	"sdk-go/client/common"
	explorerclient "sdk-go/client/explorer"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")
	explorerClient, err := explorerclient.NewExplorerClient(network)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	receiver := "helios1ddcp5ftqmntudn4m6heg2adud6hn58urnwlmkh"

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
