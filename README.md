# Helios Protocol Golang SDK 🌟

[![codecov](https://codecov.io/gh/Helios-Chain-Labs/sdk-go/graph/badge.svg?token=XDGZV265EE)](https://codecov.io/gh/Helios-Chain-Labs/sdk-go)

---

## 📚 Getting Started

Clone the repository locally and install needed dependencies

```bash
$ git clone git@github.com:Helios-Chain-Labs/sdk-go.git
$ cd sdk-go
$ go mod download
```

## Run examples
```bash
# import pk into keyring if you use keyring
heliades keys unsafe-import-eth-key helios-user 5d386fbdbf11f1141010f81a46b40f94887367562bd33b452bbaa6ce1cd1381e

# run chain example
go run examples/chain/bank/1_MsgSend/example.go

# run exchange example
go run examples/exchange/derivatives/4_Orderbook/example.go
```

---

## Updating Exchange API proto and client

```bash
$ make copy-exchange-client
```

(you have to clone [this repo](https://github.com/Helios-Chain-Labs/helios-indexer) into `../helios-indexer`)

---

## Publishing Tagged Release

```bash
$ git add .
$ git commit -m "bugfix"
$ git tag -a v1.1.1
$ git push origin master --tags
```

---

## ⛑ Support

Reach out to us at one of the following places!

- Website at <a href="https://helioschain.network/" target="_blank">`https://helioschain.network/`</a>
- Twitter at <a href="https://twitter.com/helios_layer1" target="_blank">`@helios_layer1`</a>

---

## License

Copyright © 2025 Helios Foundation (https://helioschain.network/)

<a href="https://drive.google.com/uc?export=view&id=1-fPQRh_D_dnun2yTtSsPW5MypVBOVYJP"><img src="https://drive.google.com/uc?export=view&id=1-fPQRh_D_dnun2yTtSsPW5MypVBOVYJP" style="width: 300px; max-width: 100%; height: auto" />

Originally released by Helios Foundation, Based on Injective Labs sdk-go under: <br />
Apache License <br />
Version 2.0, January 2004 <br />
http://www.apache.org/licenses/
