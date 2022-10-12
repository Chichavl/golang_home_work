UID := $(shell id -u)
GID := $(shell id -g)

.DEFAULT_GOAL := help

all: lint test ## Run lint, test

.PHONY: gofumpt
format: ## Run lint formaters (gofumpt)
	@docker run --rm -u "$(UID):$(UID)" -v "$(PWD)":/local -w /local vejed/gofumpt -w .

.PHONY: lint
lint: ## Run linters
	@printf "\n*** Running linters ***\n\n"
	@docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:v1.49.0 golangci-lint run -v -c .mygolangci.yml

.PHONY: test
test: ## Run tests
	@printf "\n*** Running tests ***\n\n"
	@go test -v -count=1 -race -timeout=1m .
	@[ -f ./test.sh ] && ./test.sh

.PHONY: t
t: ## Run test class by name
	@printf "\n*** Running tests ***\n\n"
	@go test -v -count=1 -race -timeout=1m -run TestList

.PHONY: cover
cover: ## Generate coverage.html
	@go test -v -count=1 -race -timeout=1m -coverprofile=c.out .
	@go tool cover -html=c.out -o coverage.html

.PHONY: clear 
clear: ## Remove unnecessary go modules & files
	@ rm coverage.html c.out .sync || true
	@ go mod tidy

help: ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'