package main

import (
	"context"
	"fmt"
	"os"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/Helios-Chain-Labs/sdk-go/chain/types"
	chainclient "github.com/Helios-Chain-Labs/sdk-go/client/chain"
	"github.com/Helios-Chain-Labs/sdk-go/client/common"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
)

func main() {
	// network := common.LoadNetwork("mainnet", "k8s")
	network := common.LoadNetwork("mainnet", "sentry0")
	tmClient, err := rpchttp.New(network.TmEndpoint, "/websocket")
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

	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint).WithClient(tmClient)

	queryClient := authtypes.NewQueryClient(clientCtx)

	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	address := "helios1akxycslq8cjt0uffw4rjmfm3echchptu52a2dq"

	res, err := queryClient.Account(ctx, &authtypes.QueryAccountRequest{
		Address: address,
	})
	if err != nil {
		fmt.Println(err)
	}

	var account types.EthAccount

	clientCtx.Codec.MustUnmarshal(res.GetAccount().GetValue(), &account)

	fmt.Println(account.String())

}
