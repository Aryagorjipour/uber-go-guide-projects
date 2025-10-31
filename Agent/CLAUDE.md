# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This is a learning repository containing 10 production-quality Go projects organized into three progressive phases, demonstrating best practices from the [Uber Go Style Guide](https://github.com/uber-go/guide).

**Key Architecture:**
- Multi-module monorepo using Go workspaces (`go.work`)
- Each project is a separate Go module in `phase1/`, `phase2/`, or `phase3/` directories
- Shared utilities live in `shared/` module
- All projects must adhere to Uber Go Style Guide conventions

## Build and Test Commands

### Root-level Commands (via Makefile)

Run these from the repository root:

```bash
# Run all tests across all projects
make test

# Run tests with race detector (always use this before committing)
make test-race

# Generate coverage reports for all projects
make test-coverage

# Run linters on all projects
make lint

# Format all Go files
make fmt

# Tidy all go.mod files
make tidy

# Clean build artifacts and test caches
make clean

# Initialize all projects and install tools
make init-all

# Initialize specific phase
make init-phase1  # or init-phase2, init-phase3
```

### Individual Project Commands

Navigate to a specific project directory (e.g., `cd phase1/01-metrics-system`) and use:

```bash
# Download dependencies
go mod download

# Run tests for current project
go test -v ./...

# Run tests with race detector
go test -race -v ./...

# Run a single test
go test -v -run TestName ./...

# Generate coverage for current project
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# Format code
go fmt ./...

# Tidy dependencies
go mod tidy
```

## Development Workflow

### Working on a Project

1. Navigate to the specific project directory (e.g., `phase1/01-metrics-system`)
2. Implement features following Uber Go Style Guide
3. Run tests with race detector: `go test -race -v ./...`
4. Ensure linting passes: `golangci-lint run ./...`
5. From root, verify all tests still pass: `make test-race`

### Code Quality Requirements

All code must meet these standards:
- 100% compliance with Uber Go Style Guide
- Pass race detector (`go test -race`)
- Minimum 80% test coverage
- Pass all linters (golangci-lint)

### Testing Philosophy

- Use table-driven tests
- Test with race detector enabled
- Mock external dependencies
- Test error paths and edge cases
- Verify zero-value usability where applicable

## Project Organization

### Phase 1: Foundation
- **01-metrics-system**: Thread-safe metrics collection (atomics, mutexes, interfaces)
- **02-config-manager**: Configuration loading & validation (field tags, error types)
- **03-file-watcher**: File system event monitoring (goroutine lifecycle, channels)

### Phase 2: Intermediate
- **01-api-gateway**: HTTP reverse proxy & routing (context, middleware)
- **02-tsdb-client**: Time-series database client (time handling, retries)
- **03-retry-library**: Exponential backoff retry logic (functional options)

### Phase 3: Advanced
- **01-task-scheduler**: Distributed task scheduling (worker pools, graceful shutdown)
- **02-log-aggregator**: Real-time log collection (channel sizing, optimization)
- **03-message-broker**: Pub/sub message broker (lifecycle management)
- **04-connection-pool**: Database connection pooling (resource pooling, health checks)

## Uber Go Style Guide Key Patterns

When implementing features, follow these patterns:

**Zero-Value Usability:**
- Structs should be usable without explicit initialization when possible
- Use sync.Mutex as struct field (not embedded) to preserve zero value

**Concurrency:**
- Start goroutines with defer for cleanup
- Use context for cancellation and timeouts
- Protect shared state with mutexes or atomics
- Always test concurrent code with race detector

**Error Handling:**
- Use typed errors for expected errors
- Wrap errors with context: `fmt.Errorf("context: %w", err)`
- Name error variables with `Err` prefix (e.g., `ErrNotFound`)

**Interfaces:**
- Keep interfaces small and focused
- Define interfaces in consumer packages, not implementer packages
- Verify interface compliance at compile time: `var _ Interface = (*Type)(nil)`

**Naming:**
- Use MixedCaps, not underscores
- Receiver names: use 1-2 letter abbreviations, be consistent
- Package names: lowercase, no underscores, concise
- Avoid stutter: `log.Logger`, not `log.LogLogger`

## Go Workspace Structure

This repo uses Go workspaces (`go.work`). All modules are part of the workspace:
- Changes in `shared/` are immediately visible to all projects
- Run `go work sync` if you encounter module resolution issues
- Each project has its own `go.mod` with independent versioning

## Shared Module

The `shared/` directory contains utilities used across projects:
- Common test helpers
- Shared interfaces
- Reusable patterns

When adding shared code:
1. Ensure it's genuinely reusable across multiple projects
2. Document thoroughly
3. Maintain backward compatibility
4. Test extensively
