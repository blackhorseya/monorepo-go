# Makefile for Your Golang Monorepo Project

# Variables
GO := go
BUILD_DIR := build
BIN_DIR := $(BUILD_DIR)/bin
LDFLAGS := -w -s

# Targets
.PHONY: all build clean

all: build

build: $(BIN_DIR)/user

$(BIN_DIR)/user: adapter/user/main.go
	@mkdir -p $(BIN_DIR)
	$(GO) build -ldflags "$(LDFLAGS)" -o $@ $<

clean:
	@rm -rf $(BUILD_DIR)
