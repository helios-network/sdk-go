package main

import (
	"context"
	"fmt"
	"os"
	"time"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	exchangetypes "sdk-go/chain/exchange/types"
	"github.com/Helios-Chain-Labs/sdk-go/client"
	chainclient "github.com/Helios-Chain-Labs/sdk-go/client/chain"
	"github.com/Helios-Chain-Labs/sdk-go/client/common"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")
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
		fmt.Println(err)
		return
	}

	ctx := context.Background()
	marketsAssistant, err := chainclient.NewMarketsAssistant(ctx, chainClient)
	if err != nil {
		panic(err)
	}

	defaultSubaccountID := chainClient.DefaultSubaccount(senderAddress)

	marketId := "0x0611780ba69656949525013d947713300f56c37b6175e02f26bffa495c3208fe"
	amount := decimal.NewFromFloat(2)
	price := decimal.NewFromFloat(22.5)

	order := chainClient.CreateSpotOrder(
		defaultSubaccountID,
		&chainclient.SpotOrderData{
			OrderType:    exchangetypes.OrderType_BUY, //BUY SELL BUY_PO SELL_PO
			Quantity:     amount,
			Price:        price,
			FeeRecipient: senderAddress.String(),
			MarketId:     marketId,
			Cid:          uuid.NewString(),
		},
		marketsAssistant)
	msg := new(exchangetypes.MsgBatchCreateSpotLimitOrders)
	msg.Sender = senderAddress.String()
	msg.Orders = []exchangetypes.SpotOrder{*order}

	simRes, err := chainClient.SimulateMsg(clientCtx, msg)

	if err != nil {
		fmt.Println(err)
		return
	}

	msgBatchCreateSpotLimitOrdersResponse := exchangetypes.MsgBatchCreateSpotLimitOrdersResponse{}
	err = msgBatchCreateSpotLimitOrdersResponse.Unmarshal(simRes.Result.MsgResponses[0].Value)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("simulated order hashes", msgBatchCreateSpotLimitOrdersResponse.OrderHashes)

	// AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	err = chainClient.QueueBroadcastMsg(msg)

	if err != nil {
		fmt.Println(err)
		return
	}

	time.Sleep(time.Second * 5)

	gasFee, err := chainClient.GetGasFee()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("gas fee:", gasFee, "HELIOS")
}
