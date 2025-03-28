all:

.PHONY: copy-exchange-client tests coverage

tests:
	go test -race ./client/... ./ethereum/...
coverage:
	go test -race -coverprofile=coverage.out -covermode=atomic ./client/... ./ethereum/...
