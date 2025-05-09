package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Helios-Chain-Labs/sdk-go/client"
	"github.com/Helios-Chain-Labs/sdk-go/client/common"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"

	chainclient "github.com/Helios-Chain-Labs/sdk-go/client/chain"
)

func main() {
	// network := common.LoadNetwork("mainnet", "k8s")
	network := common.LoadNetwork("mainnet", "lb")
	tmRPC, err := rpchttp.New(network.TmEndpoint, "/websocket")

	if err != nil {
		panic(err)
	}

	senderAddress, cosmosKeyring, err := chainclient.InitCosmosKeyring(
		os.Getenv("HOME")+"/.heliades",
		"heliades",
		"file",
		"helios-user",
		"12345678",
		"5d386fbdbf11f1141010f81a46b40f94887367562bd33b452bbaa6ce1cd1381e", // keyring will be used if pk not provided
		false,
	)

	if err != nil {
		panic(err)
	}

	clientCtx, err := chainclient.NewClientContext(
		network.ChainId,
		senderAddress.String(),
		cosmosKeyring,
	)

	if err != nil {
		panic(err)
	}

	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint).WithClient(tmRPC)

	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network,
		common.OptionGasPrices(client.DefaultGasPriceWithDenom),
	)

	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	timeOutCtx, cancelFn := context.WithTimeout(ctx, 30*time.Second)
	defer cancelFn()

	resp, err := chainClient.GetTx(timeOutCtx, "A2B2B971C690AE7977451D24D6F450AECE6BCCB271E91E32C2563342DDA5254B")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.TxResponse)
}
