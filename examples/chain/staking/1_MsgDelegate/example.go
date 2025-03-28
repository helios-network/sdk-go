package main

import (
	"fmt"
	"os"
	"time"

	"cosmossdk.io/math"

	"github.com/Helios-Chain-Labs/sdk-go/client"
	"github.com/Helios-Chain-Labs/sdk-go/client/common"

	chainclient "github.com/Helios-Chain-Labs/sdk-go/client/chain"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func main() {
	network := common.LoadNetwork("local", "lb")
	tmClient, err := rpchttp.New(network.TmEndpoint, "/websocket")
	if err != nil {
		panic(err)
	}

	senderAddress, cosmosKeyring, err := chainclient.InitCosmosKeyring(
		os.Getenv("HOME")+"/.heliades",
		"heliades",
		"test",
		"helios-user",
		"",
		"cec6921f279a6f6c598b3406a9dc17875db23d36ae6809e703d1eb75ba9835e6", // keyring will be used if pk not provided
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

	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network,
		common.OptionGasPrices(client.DefaultGasPriceWithDenom),
	)

	if err != nil {
		panic(err)
	}

	msg := new(stakingtypes.MsgDelegate)
	msg.DelegatorAddress = senderAddress.String()
	msg.ValidatorAddress = "heliosvaloper1f9ddqdwc2lj9x7vkhrerdeuemtxg5csk6wrjjx"
	msg.Amount = sdktypes.Coin{
		Denom: "ahelios", Amount: math.NewInt(1000000000000000000), // 1 HELIOS
	}

	// AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	err = chainClient.QueueBroadcastMsg(msg)

	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(time.Second * 5)

	gasFee, err := chainClient.GetGasFee()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("gas fee:", gasFee, "HELIOS")
}
