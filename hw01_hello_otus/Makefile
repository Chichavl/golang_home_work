.DEFAULT_GOAL := help

all: lint test ## Run lint, test

.PHONY: lint
lint: ## Run linters
	@printf "\n*** Running linters ***\n\n"
	@docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:v1.49.0 golangci-lint run -v -c ./.mygolangci.yml

test: ## Run tests
	@printf "\n*** Running tests ***\n\n"
	@go test -v -count=1 -race -timeout=1m .
	@[ -f ./test.sh ] && ./test.sh

help: ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'