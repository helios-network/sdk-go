package main

import (
	"context"
	"fmt"

	"sdk-go/client/common"
	tmclient "sdk-go/client/tm"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")
	tmClient := tmclient.NewRPCClient(network.TmEndpoint)
	clientCtx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	res, err := tmClient.GetBlock(clientCtx, 15478013)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res.Block.Txs)

}
