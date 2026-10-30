BINARY  := runnex
PKG     := github.com/kayamii/runnex
BIN_DIR ?= bin
EXE     := $(shell go env GOEXE)

VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
COMMIT  ?= $(shell git rev-parse HEAD 2>/dev/null || echo unknown)
DATE    ?= $(shell git log -1 --format=%cI 2>/dev/null || echo unknown)

LDFLAGS := -s -w -buildid= \
	-X $(PKG)/internal/version.Version=$(VERSION) \
	-X $(PKG)/internal/version.Commit=$(COMMIT) \
	-X $(PKG)/internal/version.Date=$(DATE)

GOFLAGS_BUILD := -trimpath -ldflags "$(LDFLAGS)"

.DEFAULT_GOAL := help

.PHONY: help
help: ## Show available targets
	@grep -hE '^[a-zA-Z_-]+:.*?## ' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-12s %s\n", $$1, $$2}'

.PHONY: build
build: ## Build the runnex binary into ./bin
	CGO_ENABLED=0 go build $(GOFLAGS_BUILD) -o $(BIN_DIR)/$(BINARY)$(EXE) ./cmd/runnex

.PHONY: run
run: ## Run runnex (pass arguments with ARGS="...")
	go run ./cmd/runnex $(ARGS)

.PHONY: test
test: ## Run unit tests
	go test -count=1 ./...

.PHONY: test-race
test-race: ## Run unit tests with the race detector (requires cgo)
	go test -race -count=1 ./...

.PHONY: fmt
fmt: ## Format Go source files
	gofmt -s -w .

.PHONY: fmt-check
fmt-check: ## Fail if any Go source file is not formatted
	@out="$$(gofmt -s -l .)"; if [ -n "$$out" ]; then echo "Unformatted files:"; echo "$$out"; exit 1; fi

.PHONY: vet
vet: ## Run go vet
	go vet ./...

.PHONY: lint
lint: ## Run golangci-lint
	golangci-lint run

.PHONY: tidy
tidy: ## Tidy go.mod and go.sum
	go mod tidy

.PHONY: mod-check
mod-check: ## Verify module dependencies and that go.mod is tidy
	go mod verify
	go mod tidy -diff

.PHONY: check
check: mod-check fmt-check vet lint test build ## Run all validation steps

.PHONY: dist
dist: ## Build release archives into ./dist (requires GNU tar)
	MAKE=$(MAKE) VERSION=$(VERSION) COMMIT=$(COMMIT) DATE=$(DATE) bash scripts/dist.sh

.PHONY: clean
clean: ## Remove build artifacts
	rm -rf $(BIN_DIR) dist
