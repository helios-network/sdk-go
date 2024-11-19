# Helios Protocol Golang SDK 🌟

[![codecov](https://codecov.io/gh/HeliosLabs/sdk-go/graph/badge.svg?token=XDGZV265EE)](https://codecov.io/gh/HeliosLabs/sdk-go)

---

## 📚 Getting Started

Clone the repository locally and install needed dependencies

```bash
$ git clone git@github.com:HeliosLabs/sdk-go.git
$ cd sdk-go
$ go install ./...
```

## Run examples
```bash
# import pk into keyring if you use keyring
d keys unsafe-import-eth-key helios-user 5d386fbdbf11f1141010f81a46b40f94887367562bd33b452bbaa6ce1cd1381e

# run chain example
go run examples/chain/0_MsgSend.go

# run exchange example
go run examples/exchange/derivative_exchange_rpc/0_GetOrderbook.go
```

---

## Updating Exchange API proto and client

```bash
$ make copy-exchange-client
```

(you have to clone [this repo](https://github.com/Helios-Chain-Labs/-indexer) into `../-indexer`)

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

- Website at <a href="https://.com" target="_blank">`.com`</a>
- Twitter at <a href="https://twitter.com/HeliosLabs" target="_blank">`@HeliosLabs`</a>

---

## License

Copyright © 2020 - 2022 Helios Labs Inc. (https://labs.org/)

<a href="https://drive.google.com/uc?export=view&id=1-fPQRh_D_dnun2yTtSsPW5MypVBOVYJP"><img src="https://drive.google.com/uc?export=view&id=1-fPQRh_D_dnun2yTtSsPW5MypVBOVYJP" style="width: 300px; max-width: 100%; height: auto" />

Originally released by Helios Labs Inc. under: <br />
Apache License <br />
Version 2.0, January 2004 <br />
http://www.apache.org/licenses/
