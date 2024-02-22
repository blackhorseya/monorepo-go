# Makefile for Your Golang Monorepo Project

# Variables
GO := go
BUILD_DIR := build
BIN_DIR := $(BUILD_DIR)/bin
LDFLAGS := -w -s

# Targets
.PHONY: all help
.PHONY: lint build clean test
.PHONY: gen-pb gen-mocks gen-swagger gazelle gazelle-repos

all: help

help: ## show help
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'

lint: ## run golangci-lint
	@golangci-lint run ./...

build: $(BIN_DIR)/stringx ## build binary

$(BIN_DIR)/stringx: adapter/stringx/main.go
	@mkdir -p $(BIN_DIR)
	$(GO) build -ldflags "$(LDFLAGS)" -o $@ $<

clean: ## clean build directory
	@rm -rf $(BUILD_DIR)

test: ## run test
	@bazel test //...

gen-pb: ## generate protobuf
	@$(GO) get -u google.golang.org/protobuf/proto
	@$(GO) get -u google.golang.org/protobuf/cmd/protoc-gen-go

	## Starting generate pb
	@protoc --proto_path=./pb \
		--go_out=paths=source_relative:./ \
		--go-grpc_out=paths=source_relative,require_unimplemented_servers=false:./ \
		--go-grpc-mock_out=paths=source_relative,require_unimplemented_servers=false:./ \
		./pb/entity/domain/*/model/*.proto
	@protoc --proto_path=./pb \
		--go_out=paths=source_relative:./ \
		--go-grpc_out=paths=source_relative,require_unimplemented_servers=false:./ \
		--go-grpc-mock_out=paths=source_relative,require_unimplemented_servers=false:./ \
		./pb/entity/*/domain/*/model/*.proto
	@echo Successfully generated proto

	## Starting inject tags
	@protoc-go-inject-tag -input="./entity/domain/*/model/*.pb.go"
	@protoc-go-inject-tag -input="./entity/*/domain/*/model/*.pb.go"
	@echo Successfully injected tags

gen-mocks: ## generate mocks
	@$(GO) generate ./...

gen-swagger: ## generate swagger
	## reurl
	@swag init -q -d ./adapter/reurl,./app,./pkg,./entity -o ./adapter/reurl/api/docs

	## ekko
	@swag init -q -d ./adapter/ekko,./app,./pkg,./entity -o ./adapter/ekko/api/docs

	## orianna
	@swag init -q -d ./adapter/orianna,./app,./pkg,./entity -o ./adapter/orianna/api/docs

	## sion
	@swag init -q -d ./adapter/sion,./app,./pkg,./entity -o ./adapter/sion/api/docs

gazelle-repos: ## run gazelle with bazel
	@bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies -prune

gazelle: ## run gazelle with bazel
	@bazel run //:gazelle

## docker
.PHONY: docker-push
docker-push: docker-push-ekko docker-push-orianna docker-push-sion ## push docker image

.PHONY: docker-push-ekko
docker-push-ekko: ## push docker image
	bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //adapter/ekko:push

.PHONY: docker-push-orianna
docker-push-orianna: ## push docker image
	bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //adapter/orianna:push

.PHONY: docker-push-sion
docker-push-sion: ## push docker image
	bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //adapter/sion:push
