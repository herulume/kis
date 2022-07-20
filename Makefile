include .bingo/Variables.mk

.DEFAULT_GOAL := build

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-18s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

test:
	go test ./...
	go test -race ./...
.PHONY:test

lint: $(GOLANGCI_LINT)
	$(GOLANGCI_LINT) cache status
	$(GOLANGCI_LINT) run --timeout=5m
.PHONY: lint