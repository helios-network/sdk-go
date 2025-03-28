package main

import (
	"fmt"
	"os"
	"time"

	"cosmossdk.io/math"

	"github.com/Helios-Chain-Labs/sdk-go/client"
	"github.com/Helios-Chain-Labs/sdk-go/client/common"

	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
	chainclient "github.com/Helios-Chain-Labs/sdk-go/client/chain"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")
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

	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network,
		common.OptionGasPrices(client.DefaultGasPriceWithDenom),
	)

	if err != nil {
		panic(err)
	}

	dest := "0xaf79152ac5df276d9a8e1e2e22822f9713474902"

	amount := sdktypes.Coin{
		Denom: "helios", Amount: math.NewInt(5000000000000000000), // 5 HELIOS
	}
	bridgeFee := sdktypes.Coin{
		Denom: "helios", Amount: math.NewInt(2000000000000000000), // 2 HELIOS
	}

	msg := &hyperiontypes.MsgSendToChain{
		Sender:         senderAddress.String(),
		DestHyperionId: 1,
		Amount:         amount,
		Dest:           dest,
		BridgeFee:      bridgeFee,
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
