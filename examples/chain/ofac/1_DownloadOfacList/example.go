package main

import (
	chainclient "github.com/Helios-Chain-Labs/sdk-go/client/chain"
)

func main() {
	err := chainclient.DownloadOfacList()
	if err != nil {
		panic(err)
	}
}
