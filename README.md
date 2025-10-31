# Uber Go Style Guide - Learning Projects

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

> A comprehensive collection of projects demonstrating best practices from the [Uber Go Style Guide](https://github.com/uber-go/guide).

## 📚 Overview

This repository contains 10 production-quality Go projects organized into three progressive phases. Each project demonstrates specific patterns, conventions, and best practices from the Uber Go Style Guide.

**Purpose:**
- ✅ Learn idiomatic Go through practical implementation
- ✅ Master concurrency patterns and thread-safety
- ✅ Build production-ready, maintainable code
- ✅ Practice comprehensive testing strategies

## 🎯 Learning Path

### Phase 1: Foundation (Start Here)
**Focus:** Core patterns, style conventions, basic concurrency

| Project | Description | Key Concepts |
|---------|-------------|--------------|
| [01-metrics-system](./phase1/01-metrics-system) | Thread-safe metrics collection | Atomics, zero-value mutexes, interfaces |
| [02-config-manager](./phase1/02-config-manager) | Configuration loading & validation | Field tags, error types, struct initialization |
| [03-file-watcher](./phase1/03-file-watcher) | File system event monitoring | Goroutine lifecycle, defer cleanup, channels |

**Estimated Time:** 1-2 weeks
**Prerequisites:** Basic Go syntax and understanding

### Phase 2: Intermediate
**Focus:** API design, advanced error handling, complex concurrency

| Project | Description | Key Concepts |
|---------|-------------|--------------|
| [01-api-gateway](./phase2/01-api-gateway) | HTTP reverse proxy & routing | Exit patterns, middleware, context |
| [02-tsdb-client](./phase2/02-tsdb-client) | Time-series database client | Time handling, boundary protection, retries |
| [03-retry-library](./phase2/03-retry-library) | Exponential backoff retry logic | Functional options, error wrapping |

**Estimated Time:** 2-3 weeks
**Prerequisites:** Phase 1 completion

### Phase 3: Advanced
**Focus:** Production systems, resource management, distributed patterns

| Project | Description | Key Concepts |
|---------|-------------|--------------|
| [01-task-scheduler](./phase3/01-task-scheduler) | Distributed task scheduling | Worker pools, graceful shutdown |
| [02-log-aggregator](./phase3/02-log-aggregator) | Real-time log collection | Channel sizing, performance optimization |
| [03-message-broker](./phase3/03-message-broker) | Pub/sub message broker | Complex lifecycle management |
| [04-connection-pool](./phase3/04-connection-pool) | Database connection pooling | Resource pooling, health checks |

**Estimated Time:** 3-4 weeks
**Prerequisites:** Phase 2 completion

## 🚀 Quick Start

### Clone Repository
```bash
git clone https://github.com/AryaGorjipour/uber-go-guide-projects.git
cd uber-go-guide-projects
```

### Start with Phase 1
```bash
cd phase1/01-metrics-system
go mod download
go test -v -race ./...
```

### Run All Tests (from root)
```bash
make test-all
```

## 📖 Documentation

- [Style Guide Quick Reference](./docs/STYLE_GUIDE.md)
- [Phase 1 Details](./docs/phase1.md)
- [Phase 2 Details](./docs/phase2.md)
- [Phase 3 Details](./docs/phase3.md)

## 🛠️ Development Tools

### Required
- Go 1.21 or higher
- Make (optional, for convenience commands)

### Recommended
- [golangci-lint](https://golangci-lint.run/) for linting
- [go-junit-report](https://github.com/jstemmer/go-junit-report) for CI/CD
- [gotestsum](https://github.com/gotestyourself/gotestsum) for better test output

### Installation
```bash
# Install linters
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Install test tools
go install gotest.tools/gotestsum@latest
```

## 📝 Project Status

| Project | Status | Tests | Coverage |
|---------|--------|-------|----------|
| Phase 1 - Metrics System | ✅ Complete | Passing | 95% |
| Phase 1 - Config Manager | 🚧 In Progress | - | - |
| Phase 1 - File Watcher | 📋 Planned | - | - |
| Phase 2 - API Gateway | 📋 Planned | - | - |
| Phase 2 - TSDB Client | 📋 Planned | - | - |
| Phase 2 - Retry Library | 📋 Planned | - | - |
| Phase 3 - Task Scheduler | 📋 Planned | - | - |
| Phase 3 - Log Aggregator | 📋 Planned | - | - |
| Phase 3 - Message Broker | 📋 Planned | - | - |
| Phase 3 - Connection Pool | 📋 Planned | - | - |

## 🎓 Learning Objectives by Phase

### After Phase 1, you will understand:
- [ ] Zero-value initialization and usability
- [ ] Proper mutex usage (embedded vs field)
- [ ] Interface design and verification
- [ ] Basic goroutine lifecycle management
- [ ] Error types and naming conventions
- [ ] Table-driven testing
- [ ] Style conventions (naming, grouping, imports)

### After Phase 2, you will understand:
- [ ] API design patterns
- [ ] Context usage and cancellation
- [ ] Time handling complexities
- [ ] Functional options pattern
- [ ] Advanced error handling strategies
- [ ] Middleware patterns
- [ ] Performance optimization techniques

### After Phase 3, you will understand:
- [ ] Production-grade concurrency patterns
- [ ] Resource pool management
- [ ] Graceful shutdown sequences
- [ ] Health checking strategies
- [ ] Distributed system patterns
- [ ] Performance profiling
- [ ] Production deployment considerations

## 🤝 Contributing

Contributions are welcome! Please:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/improvement`)
3. Follow the [Uber Go Style Guide](https://github.com/uber-go/guide)
4. Add tests for new functionality
5. Ensure all tests pass (`go test -race ./...`)
6. Submit a pull request

See [CONTRIBUTING.md](./CONTRIBUTING.md) for detailed guidelines.

## 📊 Code Quality Standards

All projects adhere to:
- ✅ 100% compliance with Uber Go Style Guide
- ✅ Race detector passes (`go test -race`)
- ✅ Minimum 80% test coverage
- ✅ All linters pass (golangci-lint)
- ✅ Comprehensive documentation

## 🔗 Resources

### Official Documentation
- [Uber Go Style Guide](https://github.com/uber-go/guide)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

### Community Resources
- [Go by Example](https://gobyexample.com/)
- [Go Concurrency Patterns](https://go.dev/blog/pipelines)
- [Advanced Testing in Go](https://about.sourcegraph.com/go/advanced-testing-in-go)

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Uber Go Style Guide](https://github.com/uber-go/guide) authors
- Go community for best practices
- All contributors to this learning repository

## 📬 Contact

- **GitHub Issues:** [Create an issue](https://github.com/AryaGorjipour/uber-go-guide-projects/issues)
- **Discussions:** [Join discussions](https://github.com/AryaGorjipour/uber-go-guide-projects/discussions)

---

<p align="center">
  <b>⭐ Star this repository if you find it helpful!</b>
</p>

<p align="center">
  Made with ❤️ by developers learning Go the right way
</p>
