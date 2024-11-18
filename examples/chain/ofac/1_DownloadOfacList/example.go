package main

import (
	chainclient "sdk-go/client/chain"
)

func main() {
	err := chainclient.DownloadOfacList()
	if err != nil {
		panic(err)
	}
}
