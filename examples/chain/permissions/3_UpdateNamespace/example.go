package main

import (
	"encoding/json"
	"fmt"
	"os"

	permissionstypes "sdk-go/chain/permissions/types"
	"github.com/Helios-Chain-Labs/sdk-go/client"
	chainclient "github.com/Helios-Chain-Labs/sdk-go/client/chain"
	"github.com/Helios-Chain-Labs/sdk-go/client/common"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
)

func main() {
	network := common.LoadNetwork("devnet", "lb")
	tmClient, err := rpchttp.New(network.TmEndpoint, "/websocket")
	if err != nil {
		panic(err)
	}

	senderAddress, cosmosKeyring, err := chainclient.InitCosmosKeyring(
		os.Getenv("HOME")+"/.heliosd",
		"heliosd",
		"file",
		"inj-user",
		"12345678",
		"f9db9bf330e23cb7839039e944adef6e9df447b90b503d5b4464c90bea9022f3", // keyring will be used if pk not provided
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

	namespaceDenom := "factory/inj1hkhdaj2a2clmq5jq6mspsggqs32vynpk228q3r/inj_test"

	msg := &permissionstypes.MsgUpdateNamespace{
		Sender:         senderAddress.String(),
		NamespaceDenom: namespaceDenom,
		WasmHook:       &permissionstypes.MsgUpdateNamespace_MsgSetWasmHook{NewValue: "inj19ld6swyldyujcn72j7ugnu9twafhs9wxlyye5m"},
		SendsPaused:    &permissionstypes.MsgUpdateNamespace_MsgSetSendsPaused{NewValue: true},
		MintsPaused:    &permissionstypes.MsgUpdateNamespace_MsgSetMintsPaused{NewValue: true},
		BurnsPaused:    &permissionstypes.MsgUpdateNamespace_MsgSetBurnsPaused{NewValue: true},
	}

	//AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	response, err := chainClient.SyncBroadcastMsg(msg)

	if err != nil {
		panic(err)
	}

	str, _ := json.MarshalIndent(response, "", " ")
	fmt.Print(string(str))
}
