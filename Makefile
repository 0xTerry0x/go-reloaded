GOCMD := go
GOBUILD := $(GOCMD) build
GOTEST := $(GOCMD) test
BIN_DIR := bin
BIN := $(BIN_DIR)/textfmt
PKG := ./cmd/textfmt

.PHONY: setup fmt lint test coverage build ci clean

setup:
	$(GOCMD) mod tidy

fmt:
	$(GOCMD) fmt ./...

lint:
	golangci-lint run

test:
	$(GOTEST) ./... -race

coverage:
	$(GOTEST) ./... -race -coverprofile=coverage.out
	$(GOCMD) tool cover -func=coverage.out

build:
	mkdir -p $(BIN_DIR)
	$(GOBUILD) -o $(BIN) $(PKG)

ci: fmt lint test coverage build

clean:
	rm -rf $(BIN_DIR) coverage.out
