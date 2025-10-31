.PHONY: help test test-all test-race lint fmt clean install-tools

# Default target
.DEFAULT_GOAL := help

# Project configuration
PROJECTS_P1 := $(wildcard phase1/*/.)
PROJECTS_P2 := $(wildcard phase2/*/.)
PROJECTS_P3 := $(wildcard phase3/*/.)
ALL_PROJECTS := $(PROJECTS_P1) $(PROJECTS_P2) $(PROJECTS_P3)

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

install-tools: ## Install development tools
	@echo "Installing development tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install gotest.tools/gotestsum@latest
	@echo "✅ Tools installed successfully"

test: ## Run tests for all projects
	@echo "Running tests for all projects..."
	@for project in $(ALL_PROJECTS); do \
		echo "\n📦 Testing $$project"; \
		cd $$project && go test -v ./... || exit 1; \
		cd ../..; \
	done
	@echo "\n✅ All tests passed"

test-race: ## Run tests with race detector
	@echo "Running tests with race detector..."
	@for project in $(ALL_PROJECTS); do \
		echo "\n📦 Testing $$project"; \
		cd $$project && go test -race -v ./... || exit 1; \
		cd ../..; \
	done
	@echo "\n✅ All race tests passed"

test-coverage: ## Generate test coverage report
	@echo "Generating coverage reports..."
	@for project in $(ALL_PROJECTS); do \
		echo "\n📦 Coverage for $$project"; \
		cd $$project && go test -coverprofile=coverage.out ./... && \
		go tool cover -html=coverage.out -o coverage.html; \
		cd ../..; \
	done
	@echo "\n✅ Coverage reports generated"

lint: ## Run linters on all projects
	@echo "Running linters..."
	@for project in $(ALL_PROJECTS); do \
		echo "\n🔍 Linting $$project"; \
		cd $$project && golangci-lint run ./... || exit 1; \
		cd ../..; \
	done
	@echo "\n✅ All linting passed"

fmt: ## Format all Go files
	@echo "Formatting Go files..."
	@for project in $(ALL_PROJECTS); do \
		cd $$project && go fmt ./...; \
		cd ../..; \
	done
	@echo "✅ Formatting complete"

tidy: ## Run go mod tidy on all projects
	@echo "Tidying go modules..."
	@for project in $(ALL_PROJECTS); do \
		echo "📦 Tidying $$project"; \
		cd $$project && go mod tidy; \
		cd ../..; \
	done
	@echo "✅ Modules tidied"

clean: ## Clean build artifacts and test caches
	@echo "Cleaning..."
	@for project in $(ALL_PROJECTS); do \
		cd $$project && go clean -testcache -cache; \
		cd ../..; \
	done
	@find . -name "coverage.out" -delete
	@find . -name "coverage.html" -delete
	@echo "✅ Cleaned successfully"

init-phase1: ## Initialize Phase 1 project dependencies
	@echo "Initializing Phase 1 projects..."
	@for project in $(PROJECTS_P1); do \
		cd $$project && go mod download; \
		cd ../..; \
	done
	@echo "✅ Phase 1 initialized"

init-phase2: ## Initialize Phase 2 project dependencies
	@echo "Initializing Phase 2 projects..."
	@for project in $(PROJECTS_P2); do \
		cd $$project && go mod download; \
		cd ../..; \
	done
	@echo "✅ Phase 2 initialized"

init-phase3: ## Initialize Phase 3 project dependencies
	@echo "Initializing Phase 3 projects..."
	@for project in $(PROJECTS_P3); do \
		cd $$project && go mod download; \
		cd ../..; \
	done
	@echo "✅ Phase 3 initialized"

init-all: install-tools init-phase1 init-phase2 init-phase3 ## Initialize all projects
	@echo "✅ All projects initialized"
