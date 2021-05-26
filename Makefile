# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: build

GOBIN = ./build/bin
GORUN = env GO111MODULE=on go run
GOLINT := golangci-lint
ZKSNARK_FILE=zksnark/libzksnark.a

all: clean dep zksnark_generate lint test build

dep:
	go mod tidy

install_libs_ubuntu:
	apt install -y gcc g++ musl-dev gmpc-dev libmpfr-dev

zksnark_generate:
	if [ ! -f ${ZKSNARK_FILE} ]; then go generate zksnark/zksnark.go ; fi

zksnark_clean:
	rm -f zksnark/libzksnark.a

lint: dep check-lint
	 $(GOLINT) run --timeout 1h -c .golangci.yml

cilint: zksnark_generate lint

test: dep zksnark_generate
	 go test -tags unit -cover -count=1 ./bitbon/... ./quorum/...

citest: test

check-lint:
	@which $(GOLINT) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.25.0

check-mockgen:
	@which mockgen || go get github.com/golang/mock/mockgen@v1.5.0

mockgen_generate: dep check-mockgen
	mockgen -package transfer -self_package github.com/simcord-llc/bitbon-system-blockchain/bitbon/transfer -source bitbon/interfaces.go > bitbon/transfer/mock_managers.go
	mockgen -package quorum -self_package github.com/simcord-llc/bitbon-system-blockchain/quorum -source quorum/interfaces.go > quorum/mock_peer_services.go
	mockgen -package quorum -source p2p/message.go > quorum/mock/mock_message_buffer.go
	mockgen -package quorum -source p2p/protocols/protocol.go > quorum/mock_protocol_hook.go

build: dep zksnark_generate
	CGO_ENABLED=1  go build -v -o build/bin/bitbon_node  ./cmd/bitbon_node/
	@echo "Done building."

clean: zksnark_clean
	env GO111MODULE=on go clean -cache
	rm -f build/bin/bitbon_node
	rm -rf vendor

abi-gen: abigen-build
	docker run -v `pwd`:/sources abigen-local --abi /sources/bitbon/contracts/abi/ContractStorageImpl.abi --pkg contracts --type ContractStorageImpl --out /sources/bitbon/contracts/gen_ContractStorageImpl.go
	docker run -v `pwd`:/sources abigen-local --abi /sources/bitbon/contracts/abi/BitbonImpl.abi --pkg contracts --type BitbonImpl --out /sources/bitbon/contracts/gen_BitbonImpl.go
	docker run -v `pwd`:/sources abigen-local --abi /sources/bitbon/contracts/abi/TransferImpl.abi --pkg contracts --type TransferImpl --out /sources/bitbon/contracts/gen_TransferImpl.go
	docker run -v `pwd`:/sources abigen-local --abi /sources/bitbon/contracts/abi/BitbonSupportImpl.abi --pkg contracts --type BitbonSupportImpl --out /sources/bitbon/contracts/gen_BitbonSupportImpl.go
	docker run -v `pwd`:/sources abigen-local --abi /sources/bitbon/contracts/abi/AssetboxImpl.abi --pkg contracts --type AssetboxImpl --out /sources/bitbon/contracts/gen_AssetboxImpl.go
	docker run -v `pwd`:/sources abigen-local --abi /sources/bitbon/contracts/abi/AssetboxStorageImpl.abi --pkg contracts --type AssetboxStorageImpl --out /sources/bitbon/contracts/gen_AssetboxStorageImpl.go
	docker run -v `pwd`:/sources abigen-local --abi /sources/bitbon/contracts/abi/SafeTransferStorageImpl.abi --pkg contracts --type SafeTransferStorageImpl --out /sources/bitbon/contracts/gen_SafeTransferStorageImpl.go
	docker run -v `pwd`:/sources abigen-local --abi /sources/bitbon/contracts/abi/BitbonStorageImpl.abi --pkg contracts --type BitbonStorageImpl --out /sources/bitbon/contracts/gen_BitbonStorageImpl.go
	docker run -v `pwd`:/sources abigen-local --abi /sources/bitbon/contracts/abi/ReservedAliasStorageImpl.abi --pkg contracts --type ReservedAliasStorageImpl --out /sources/bitbon/contracts/gen_ReservedAliasStorageImpl.go
	docker run -v `pwd`:/sources abigen-local --abi /sources/bitbon/contracts/abi/AccessStorageImpl.abi --pkg contracts --type AccessStorageImpl --out /sources/bitbon/contracts/gen_AccessStorageImpl.go
	docker run -v `pwd`:/sources abigen-local --abi /sources/bitbon/contracts/abi/DistributionStorageImpl.abi --pkg contracts --type DistributionStorageImpl --out /sources/bitbon/contracts/gen_DistributionStorageImpl.go
	docker run -v `pwd`:/sources abigen-local --abi /sources/bitbon/contracts/abi/MiningAgentStorageImpl.abi --pkg contracts --type MiningAgentStorageImpl --out /sources/bitbon/contracts/gen_MiningAgentStorageImpl.go

proto-generate:
	docker run --rm -v `pwd`:/defs namely/protoc-all:1.30_0 -i bitbon/dto/external/proto -f assetbox.proto -l go -o bitbon/dto/external
	docker run --rm -v `pwd`:/defs namely/protoc-all:1.30_0 -i bitbon/dto/external/proto -f error.proto -l go -o bitbon/dto/external
	docker run --rm -v `pwd`:/defs namely/protoc-all:1.30_0 -i bitbon/dto/external/proto -f assetbox-create.proto -l go -o bitbon/dto/external
	docker run --rm -v `pwd`:/defs namely/protoc-all:1.30_0 -i bitbon/dto/external/proto -f assetbox-delete.proto -l go -o bitbon/dto/external
	docker run --rm -v `pwd`:/defs namely/protoc-all:1.30_0 -i bitbon/dto/external/proto -f assetbox-info.proto -l go -o bitbon/dto/external
	docker run --rm -v `pwd`:/defs namely/protoc-all:1.30_0 -i bitbon/dto/external/proto -f block.proto -l go -o bitbon/dto/external
	docker run --rm -v `pwd`:/defs namely/protoc-all:1.30_0 -i bitbon/dto/external/proto -f block-info.proto -l go -o bitbon/dto/external
	docker run --rm -v `pwd`:/defs namely/protoc-all:1.30_0 -i bitbon/dto/external/proto -f events.proto -l go -o bitbon/dto/external
	docker run --rm -v `pwd`:/defs namely/protoc-all:1.30_0 -i bitbon/dto/external/proto -f mining-agent.proto -l go -o bitbon/dto/external
	docker run --rm -v `pwd`:/defs namely/protoc-all:1.30_0 -i bitbon/dto/external/proto -f node-transaction.proto -l go -o bitbon/dto/external
	docker run --rm -v `pwd`:/defs namely/protoc-all:1.30_0 -i bitbon/dto/external/proto -f transaction.proto -l go -o bitbon/dto/external
	docker run --rm -v `pwd`:/defs namely/protoc-all:1.30_0 -i bitbon/dto/external/proto -f transfer.proto -l go -o bitbon/dto/external

abigen-build:
	docker build -t abigen-local -f Dockerfile-abigen .

fmt: ## format source files
	go fmt github.com/simcord-llc/bitbon-system-blockchain/...
