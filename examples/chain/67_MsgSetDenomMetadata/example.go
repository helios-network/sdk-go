package main

import (
	"encoding/json"
	"fmt"
	"os"

	tokenfactorytypes "github.com/Helios-Chain-Labs/sdk-go/chain/tokenfactory/types"
	"github.com/Helios-Chain-Labs/sdk-go/client"
	chainclient "github.com/Helios-Chain-Labs/sdk-go/client/chain"
	"github.com/Helios-Chain-Labs/sdk-go/client/common"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")
	tmClient, err := rpchttp.New(network.TmEndpoint, "/websocket")
	if err != nil {
		panic(err)
	}

	senderAddress, cosmosKeyring, err := chainclient.InitCosmosKeyring(
		os.Getenv("HOME")+"/.d",
		"d",
		"file",
		"helios-user",
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
		fmt.Println(err)
		return
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

	denom := "factory/helios1hkhdaj2a2clmq5jq6mspsggqs32vynpk228q3r/helios_test"
	subdenom := "helios_test"
	tokenDecimals := uint32(6)

	microDenomUnit := banktypes.DenomUnit{
		Denom:    denom,
		Exponent: 0,
		Aliases:  []string{fmt.Sprintf("micro%s", subdenom)},
	}
	denomUnit := banktypes.DenomUnit{
		Denom:    subdenom,
		Exponent: tokenDecimals,
		Aliases:  []string{subdenom},
	}

	metadata := banktypes.Metadata{
		Description: "Helios Test Token",
		DenomUnits:  []*banktypes.DenomUnit{&microDenomUnit, &denomUnit},
		Base:        denom,
		Display:     subdenom,
		Name:        "Helios Test",
		Symbol:      "HELIOSTEST",
		URI:         "http://-test.com/icon.jpg",
		URIHash:     "",
	}

	message := new(tokenfactorytypes.MsgSetDenomMetadata)
	message.Sender = senderAddress.String()
	message.Metadata = metadata

	//AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	response, err := chainClient.AsyncBroadcastMsg(message)

	if err != nil {
		panic(err)
	}

	str, _ := json.MarshalIndent(response, "", " ")
	fmt.Print(string(str))
}
