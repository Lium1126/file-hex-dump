GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOLANGCI=golangci-lint

all: run

run: ## Run the example case
	$(GORUN) ./main.go example.txt

run-debug: ## Run the example case with debug mode
	$(GORUN) ./main.go --debug example.txt

cmd-help: ## Display help during command execution
	$(GORUN) ./main.go --help

version: ## Display version
	$(GORUN) ./main.go --version

go-lint: ## Run static checks
	$(GOLANGCI) run --config=.golangci.yml ./...

go-test: ## Run go test
	$(GOTEST) -v ./...

go-bench: ## Run go benchmark
	$(GOTEST) -bench -v ./...

go-coverage: ## Check go test coverage
	$(GOCMD) test -coverprofile=cover.out ./...
	$(GOCMD) tool cover -html=cover.out -o cover.html

go-build: ## Build
	$(GOBUILD) -o myhexdump ./main.go

help: ## Display this help screen
	@grep -E '^[a-zA-Z/_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
