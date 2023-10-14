# Makefile for Your Golang Monorepo Project

# Variables
GO := go
BUILD_DIR := build
BIN_DIR := $(BUILD_DIR)/bin
LDFLAGS := -w -s

# Targets
.PHONY: all help build clean lint gen-pb gen-mocks

all: help

help: ## show help
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'

lint: ## run golangci-lint
	@golangci-lint run ./...

build: $(BIN_DIR)/stringx ## build binary

$(BIN_DIR)/user: adapter/stringx/main.go
	@mkdir -p $(BIN_DIR)
	$(GO) build -ldflags "$(LDFLAGS)" -o $@ $<

clean: ## clean build directory
	@rm -rf $(BUILD_DIR)

gen-pb: ## generate protobuf
	@$(GO) get -u google.golang.org/protobuf/proto
	@$(GO) get -u google.golang.org/protobuf/cmd/protoc-gen-go

	## Starting generate pb
	@protoc --proto_path=./pb \
		--go_out=paths=source_relative:./ \
		--go-grpc_out=paths=source_relative,require_unimplemented_servers=false:./ \
		--go-grpc-mock_out=paths=source_relative,require_unimplemented_servers=false:./ \
		./pb/entity/domain/*/*/*.proto
	@echo Successfully generated proto

	## Starting inject tags
	@protoc-go-inject-tag -input="./entity/domain/*/model/*.pb.go"
	@echo Successfully injected tags

gen-mocks: ## generate mocks
	@$(GO) generate ./...
