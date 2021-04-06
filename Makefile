PACKAGES := $(go list ./... | grep -v node_modules | grep -v cmd | grep -v examples | grep -v test)

.PHONY: run-local

run-local: ## Starts the server
	go run cmd/tech-challenge-v2-server/main.go --port 8080

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36mmake %-15s\033[0m %s\n", $$1, $$2}'

linter: ## Executes the linter for reviewing the code health
	@echo "Executing linter for golang"
	@golangci-lint run
	@echo "Linter successfully passed"
