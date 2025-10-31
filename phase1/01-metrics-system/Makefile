.PHONY: help test test-race test-coverage lint fmt clean bench install-deps

# Default target
.DEFAULT_GOAL := help

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-20s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

install-deps: ## Install project dependencies
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy
	@echo " Dependencies installed"

test: ## Run all tests
	@echo "Running tests..."
	@go test -v ./...
	@echo " All tests passed"

test-race: ## Run tests with race detector (recommended before commit)
	@echo "Running tests with race detector..."
	@go test -race -v ./...
	@echo " All race tests passed"

test-coverage: ## Generate test coverage report
	@echo "Generating coverage report..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@go tool cover -func=coverage.out
	@echo " Coverage report generated: coverage.html"

test-short: ## Run short tests only
	@go test -short ./...

bench: ## Run benchmarks
	@echo "Running benchmarks..."
	@go test -bench=. -benchmem ./...

bench-compare: ## Run benchmarks and save for comparison
	@echo "Running benchmarks and saving results..."
	@go test -bench=. -benchmem ./... | tee bench_new.txt

lint: ## Run linters (requires golangci-lint)
	@echo "Running linters..."
	@golangci-lint run ./...
	@echo " Linting passed"

fmt: ## Format all Go files
	@echo "Formatting code..."
	@go fmt ./...
	@echo " Code formatted"

vet: ## Run go vet
	@echo "Running go vet..."
	@go vet ./...
	@echo " Vet passed"

clean: ## Clean build artifacts and test cache
	@echo "Cleaning..."
	@go clean -testcache -cache
	@rm -f coverage.out coverage.html
	@echo " Clean complete"

run-example: ## Run example code
	@echo "Running examples..."
	@go test -v -run Example

check: fmt vet lint test-race ## Run all checks (format, vet, lint, test)
	@echo " All checks passed"

ci: install-deps check test-coverage ## Run all CI checks
	@echo " CI pipeline complete"

.PHONY: all
all: clean install-deps check test-coverage bench ## Run everything
	@echo " All tasks complete"
