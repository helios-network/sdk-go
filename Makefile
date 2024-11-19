all:

copy-exchange-client:
	rm -rf exchange/*
	mkdir -p exchange/accounts_rpc
	mkdir -p exchange/auction_rpc
	mkdir -p exchange/campaign_rpc
	mkdir -p exchange/derivative_exchange_rpc
	mkdir -p exchange/exchange_rpc
	mkdir -p exchange/explorer_rpc
	mkdir -p exchange/insurance_rpc
	mkdir -p exchange/meta_rpc
	mkdir -p exchange/oracle_rpc
	mkdir -p exchange/portfolio_rpc
	mkdir -p exchange/spot_exchange_rpc
	mkdir -p exchange/trading_rpc

	cp -r ../-indexer/api/gen/grpc/_accounts_rpc/pb exchange/accounts_rpc/pb
	cp -r ../-indexer/api/gen/grpc/_auction_rpc/pb exchange/auction_rpc/pb
	cp -r ../-indexer/api/gen/grpc/_campaign_rpc/pb exchange/campaign_rpc/pb
	cp -r ../-indexer/api/gen/grpc/_derivative_exchange_rpc/pb exchange/derivative_exchange_rpc/pb
	cp -r ../-indexer/api/gen/grpc/_exchange_rpc/pb exchange/exchange_rpc/pb
	cp -r ../-indexer/api/gen/grpc/_explorer_rpc/pb exchange/explorer_rpc/pb
	cp -r ../-indexer/api/gen/grpc/_insurance_rpc/pb exchange/insurance_rpc/pb
	cp -r ../-indexer/api/gen/grpc/_meta_rpc/pb exchange/meta_rpc/pb
	cp -r ../-indexer/api/gen/grpc/_oracle_rpc/pb exchange/oracle_rpc/pb
	cp -r ../-indexer/api/gen/grpc/_portfolio_rpc/pb exchange/portfolio_rpc/pb
	cp -r ../-indexer/api/gen/grpc/_spot_exchange_rpc/pb exchange/spot_exchange_rpc/pb
	cp -r ../-indexer/api/gen/grpc/_trading_rpc/pb exchange/trading_rpc/pb

.PHONY: copy-exchange-client tests coverage

copy-chain-types:
	cp ../-core/-chain/types/*.go chain/types
	rm -rf chain/types/*test.go rm -rf chain/types/*gw.go
	cp ../-core/-chain/crypto/ethsecp256k1/*.go chain/crypto/ethsecp256k1
	rm -rf chain/crypto/ethsecp256k1/*test.go rm -rf chain/crypto/ethsecp256k1/*gw.go
	cp ../-core/-chain/modules/auction/types/*.go chain/auction/types
	rm -rf chain/auction/types/*test.go  rm -rf chain/auction/types/*gw.go
	cp ../-core/-chain/modules/exchange/types/*.go chain/exchange/types
	rm -rf chain/exchange/types/*test.go  rm -rf chain/exchange/types/*gw.go
	cp ../-core/-chain/modules/insurance/types/*.go chain/insurance/types
	rm -rf chain/insurance/types/*test.go  rm -rf chain/insurance/types/*gw.go
	cp ../-core/-chain/modules/ocr/types/*.go chain/ocr/types
	rm -rf chain/ocr/types/*test.go  rm -rf chain/ocr/types/*gw.go
	cp ../-core/-chain/modules/oracle/types/*.go chain/oracle/types
	cp -r ../-core/-chain/modules/oracle/bandchain chain/oracle
	rm -rf chain/oracle/types/*test.go  rm -rf chain/oracle/types/*gw.go
	cp ../-core/-chain/modules/peggy/types/*.go chain/peggy/types
	rm -rf chain/peggy/types/*test.go  rm -rf chain/peggy/types/*gw.go
	cp ../-core/-chain/modules/wasmx/types/*.go chain/wasmx/types
	rm -rf chain/wasmx/types/*test.go  rm -rf chain/wasmx/types/*gw.go
	cp ../-core/-chain/modules/tokenfactory/types/*.go chain/tokenfactory/types
	rm -rf chain/tokenfactory/types/*test.go  rm -rf chain/tokenfactory/types/*gw.go
	cp ../-core/-chain/stream/types/*.go chain/stream/types

	echo "👉 Replace -core/-chain/modules with sdk-go/chain"
	echo "👉 Replace -core/-chain/types with sdk-go/chain/types"
	echo "👉 Replace -core/-chain/crypto with sdk-go/chain/crypto"

tests:
	go test -race ./client/... ./ethereum/...
coverage:
	go test -race -coverprofile=coverage.out -covermode=atomic ./client/... ./ethereum/...
