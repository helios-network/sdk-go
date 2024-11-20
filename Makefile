all:

copy-exchange-client:
	rm -rf exchange/*
	mkdir -p exchange/health_rpc
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

	cp -r ../helios-indexer/api/gen/grpc/health/pb exchange/health_rpc/pb
	cp -r ../helios-indexer/api/gen/grpc/helios_accounts_rpc/pb exchange/accounts_rpc/pb
	cp -r ../helios-indexer/api/gen/grpc/helios_accounts_rpc/pb exchange/accounts_rpc/pb
	cp -r ../helios-indexer/api/gen/grpc/helios_auction_rpc/pb exchange/auction_rpc/pb
	cp -r ../helios-indexer/api/gen/grpc/helios_campaign_rpc/pb exchange/campaign_rpc/pb
	cp -r ../helios-indexer/api/gen/grpc/helios_derivative_exchange_rpc/pb exchange/derivative_exchange_rpc/pb
	cp -r ../helios-indexer/api/gen/grpc/helios_exchange_rpc/pb exchange/exchange_rpc/pb
	cp -r ../helios-indexer/api/gen/grpc/helios_explorer_rpc/pb exchange/explorer_rpc/pb
	cp -r ../helios-indexer/api/gen/grpc/helios_insurance_rpc/pb exchange/insurance_rpc/pb
	cp -r ../helios-indexer/api/gen/grpc/helios_meta_rpc/pb exchange/meta_rpc/pb
	cp -r ../helios-indexer/api/gen/grpc/helios_oracle_rpc/pb exchange/oracle_rpc/pb
	cp -r ../helios-indexer/api/gen/grpc/helios_portfolio_rpc/pb exchange/portfolio_rpc/pb
	cp -r ../helios-indexer/api/gen/grpc/helios_spot_exchange_rpc/pb exchange/spot_exchange_rpc/pb
	cp -r ../helios-indexer/api/gen/grpc/helios_trading_rpc/pb exchange/trading_rpc/pb

.PHONY: copy-exchange-client tests coverage

copy-chain-types:
	cp ../helios-core/helios-chain/crypto/ethsecp256k1/*.go chain/crypto/ethsecp256k1
	rm -rf chain/crypto/ethsecp256k1/*test.go rm -rf chain/crypto/ethsecp256k1/*gw.go
	cp ../helios-core/helios-chain/codec/types/*.go chain/codec/types
	rm -rf chain/codec/types/*test.go rm -rf chain/codec/types/*gw.go
	cp ../github.com/Helios-Chain-Labs/sdk-go/chain/auction/types/*.go chain/auction/types
	rm -rf chain/auction/types/*test.go  rm -rf chain/auction/types/*gw.go
	cp ../github.com/Helios-Chain-Labs/sdk-go/chain/exchange/types/*.go chain/exchange/types
	rm -rf chain/exchange/types/*test.go  rm -rf chain/exchange/types/*gw.go
	cp ../github.com/Helios-Chain-Labs/sdk-go/chain/insurance/types/*.go chain/insurance/types
	rm -rf chain/insurance/types/*test.go  rm -rf chain/insurance/types/*gw.go
	cp ../github.com/Helios-Chain-Labs/sdk-go/chain/ocr/types/*.go chain/ocr/types
	rm -rf chain/ocr/types/*test.go  rm -rf chain/ocr/types/*gw.go
	cp ../github.com/Helios-Chain-Labs/sdk-go/chain/oracle/types/*.go chain/oracle/types
	cp -r ../github.com/Helios-Chain-Labs/sdk-go/chain/oracle/bandchain chain/oracle
	rm -rf chain/oracle/types/*test.go  rm -rf chain/oracle/types/*gw.go
	cp ../github.com/Helios-Chain-Labs/sdk-go/chain/peggy/types/*.go chain/peggy/types
	rm -rf chain/peggy/types/*test.go  rm -rf chain/peggy/types/*gw.go
	cp ../github.com/Helios-Chain-Labs/sdk-go/chain/permissions/types/*.go chain/permissions/types
	rm -rf chain/permissions/types/*test.go  rm -rf chain/permissions/types/*gw.go
	cp ../github.com/Helios-Chain-Labs/sdk-go/chain/tokenfactory/types/*.go chain/tokenfactory/types
	rm -rf chain/tokenfactory/types/*test.go  rm -rf chain/tokenfactory/types/*gw.go
	cp ../github.com/Helios-Chain-Labs/sdk-go/chain/wasmx/types/*.go chain/wasmx/types
	rm -rf chain/wasmx/types/*test.go  rm -rf chain/wasmx/types/*gw.go
	cp ../helios-core/helios-chain/stream/types/*.go chain/stream/types
	rm -rf chain/stream/types/*test.go  rm -rf chain/stream/types/*gw.go
	cp ../helios-core/helios-chain/types/*.go chain/types
	rm -rf chain/types/*test.go rm -rf chain/types/*gw.go

	@echo "👉 Replace github.com/Helios-Chain-Labs/sdk-go/chain with sdk-go/chain"
	@echo "👉 Replace helios-core/helios-chain/codec with sdk-go/chain/codec"
	@echo "👉 Replace helios-core/helios-chain/codec/types with sdk-go/chain/codec/types"
	@echo "👉 Replace helios-core/helios-chain/types with sdk-go/chain/types"
	@echo "👉 Replace helios-core/helios-chain/crypto with sdk-go/chain/crypto"

tests:
	go test -race ./client/... ./ethereum/...
coverage:
	go test -race -coverprofile=coverage.out -covermode=atomic ./client/... ./ethereum/...
