all:

.PHONY: tests coverage

copy-chain-types:
	cp ../helios-core/helios-chain/eip712/*.go chain/eip712
	rm -rf chain/eip712/*test.go rm -rf chain/eip712/*gw.go
	cp ../helios-core/helios-chain/crypto/ethsecp256k1/*.go chain/crypto/ethsecp256k1
	rm -rf chain/crypto/ethsecp256k1/*test.go rm -rf chain/crypto/ethsecp256k1/*gw.go
	cp ../helios-core/helios-chain/codec/types/*.go chain/codec/types
	rm -rf chain/codec/types/*test.go rm -rf chain/codec/types/*gw.go
	cp ../helios-core/helios-chain/modules/ocr/types/*.go chain/ocr/types
	rm -rf chain/ocr/types/*test.go  rm -rf chain/ocr/types/*gw.go
	cp ../helios-core/helios-chain/modules/hyperion/types/*.go chain/hyperion/types
	rm -rf chain/hyperion/types/*test.go  rm -rf chain/hyperion/types/*gw.go
	cp ../helios-core/helios-chain/modules/tokenfactory/types/*.go chain/tokenfactory/types
	rm -rf chain/tokenfactory/types/*test.go  rm -rf chain/tokenfactory/types/*gw.go
	cp ../helios-core/helios-chain/stream/types/*.go chain/stream/types
	rm -rf chain/stream/types/*test.go  rm -rf chain/stream/types/*gw.go
	cp ../helios-core/helios-chain/types/*.go chain/types
	rm -rf chain/types/*test.go rm -rf chain/types/*gw.go

	@echo "👉 Replace helios-core/helios-chain/modules with sdk-go/chain"
	@echo "👉 Replace helios-core/helios-chain/codec with sdk-go/chain/codec"
	@echo "👉 Replace helios-core/helios-chain/codec/types with sdk-go/chain/codec/types"
	@echo "👉 Replace helios-core/helios-chain/types with sdk-go/chain/types"
	@echo "👉 Replace helios-core/helios-chain/crypto with sdk-go/chain/crypto"
	@echo "👉 Replace helios-core/helios-chain/eip712 with sdk-go/chain/eip712"

tests:
	go test -race ./client/... ./ethereum/...
coverage:
	go test -race -coverprofile=coverage.out -covermode=atomic ./client/... ./ethereum/...
