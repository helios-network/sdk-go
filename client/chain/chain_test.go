package chain_test

import (
	"os"
	"testing"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	eth "github.com/ethereum/go-ethereum/common"

	"github.com/Helios-Chain-Labs/sdk-go/client"
	"github.com/Helios-Chain-Labs/sdk-go/client/chain"
	"github.com/Helios-Chain-Labs/sdk-go/client/common"
)

func accountForTests() (cosmtypes.AccAddress, keyring.Keyring, error) {
	senderAddress, cosmosKeyring, err := chain.InitCosmosKeyring(
		os.Getenv("HOME")+"/.heliades",
		"heliades",
		"file",
		"helios-user",
		"12345678",
		"5d386fbdbf11f1141010f81a46b40f94887367562bd33b452bbaa6ce1cd1381e", // keyring will be used if pk not provided
		false,
	)

	return senderAddress, cosmosKeyring, err
}

func createClient(senderAddress cosmtypes.AccAddress, cosmosKeyring keyring.Keyring, network common.Network) (chain.ChainClient, error) {
	tmClient, _ := rpchttp.New(network.TmEndpoint, "/websocket")
	clientCtx, err := chain.NewClientContext(
		network.ChainId,
		senderAddress.String(),
		cosmosKeyring,
	)

	if err != nil {
		return nil, err
	}

	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint).WithClient(tmClient)
	// configure Keyring as nil to avoid the account initialization request when running unit tests
	clientCtx.Keyring = nil

	chainClient, err := chain.NewChainClient(
		clientCtx,
		network,
		common.OptionGasPrices(client.DefaultGasPriceWithDenom),
	)

	return chainClient, err
}

func TestDefaultSubaccount(t *testing.T) {
	network := common.LoadNetwork("devnet", "lb")
	senderAddress, cosmosKeyring, err := accountForTests()

	if err != nil {
		t.Errorf("Error creating the address %v", err)
	}

	chainClient, err := createClient(senderAddress, cosmosKeyring, network)

	if err != nil {
		t.Errorf("Error creating the client %v", err)
	}

	defaultSubaccountID := chainClient.DefaultSubaccount(senderAddress)

	expectedSubaccountId := "0xaf79152ac5df276d9a8e1e2e22822f9713474902000000000000000000000000"
	expectedSubaccountIdHash := eth.HexToHash(expectedSubaccountId)
	if defaultSubaccountID != expectedSubaccountIdHash {
		t.Error("The default subaccount is calculated incorrectly")
	}
}

func TestGetSubaccountWithIndex(t *testing.T) {
	network := common.LoadNetwork("devnet", "lb")
	senderAddress, cosmosKeyring, err := accountForTests()

	if err != nil {
		t.Errorf("Error creating the address %v", err)
	}

	chainClient, err := createClient(senderAddress, cosmosKeyring, network)

	if err != nil {
		t.Errorf("Error creating the client %v", err)
	}

	subaccountOne := chainClient.Subaccount(senderAddress, 1)
	subaccountThirty := chainClient.Subaccount(senderAddress, 30)

	expectedSubaccounOnetId := "0xaf79152ac5df276d9a8e1e2e22822f9713474902000000000000000000000001"
	expectedSubaccountOneIdHash := eth.HexToHash(expectedSubaccounOnetId)

	expectedSubaccounThirtytId := "0xaf79152ac5df276d9a8e1e2e22822f971347490200000000000000000000001e"
	expectedSubaccountThirtyIdHash := eth.HexToHash(expectedSubaccounThirtytId)

	if subaccountOne != expectedSubaccountOneIdHash {
		t.Error("The subaccount with index 1 was calculated incorrectly")
	}
	if subaccountThirty != expectedSubaccountThirtyIdHash {
		t.Error("The subaccount with index 30 was calculated incorrectly")
	}
}
