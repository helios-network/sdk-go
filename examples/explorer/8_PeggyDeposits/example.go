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

	receiver := "inj14au322k9munkmx5wrchz9q30juf5wjgz2cfqku"

	req := explorerPB.GetPeggyDepositTxsRequest{
		Receiver: receiver,
	}

	res, err := explorerClient.GetPeggyDeposits(ctx, &req)
	if err != nil {
		fmt.Println(err)
	}

	str, _ := json.MarshalIndent(res, "", " ")
	fmt.Print(string(str))
}
