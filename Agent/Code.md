---
globs: *.go
alwaysApply: false
---

# Go Development Rules

## Code Quality & Style
- Follow idiomatic Go style as described in Effective Go and the official Go style guide.
- Use `gofmt`/`go fmt` and `goimports` to enforce formatting before committing.
- Prefer short, clear names; exported identifiers start with uppercase, unexported with lowercase.
- Avoid redundant naming: package name is part of the symbol context.
- Keep function length manageable; each function should perform one logical task.
- Avoid deep nesting; handle error cases early and return.
- Use composition over inheritance; prefer small types and interfaces.

## Error Handling
- Use `func(...) (T, error)` for functions that may fail; handle errors explicitly.
- Never ignore errors; avoid `_ = err` unless justified.
- Add context using `fmt.Errorf("...: %w", err)` or use `errors.Join`/`errors.Is` for wrapping.
- Avoid `panic` in library code; reserve for truly unrecoverable conditions.
- Propagate `context.Context` for cancellation and deadlines.
- Document error conditions in function comments.

## Ownership & Memory / Data Handling
- Avoid unnecessary allocations; use pointers vs values appropriately.
- Reuse buffers and slices; be mindful of `len` and `cap`.
- Avoid `unsafe` unless absolutely necessary and document its use.
- Use value receivers for small immutable types, pointer receivers otherwise.
- Design types with meaningful zero-value defaults.

## Testing
- Write unit tests in `_test.go` files with `package foo` or `package foo_test`.
- Use table-driven tests for clarity and coverage.
- Use the `testing` package; avoid heavy third-party assertion libraries.
- Run `go test -race` to detect data races.
- Write benchmarks in `*_test.go` using `Benchmark...`.
- Include edge cases, errors, and boundary tests.
- Keep tests maintainable and clean.

## Documentation
- Use GoDoc comments for exported identifiers starting with the identifier’s name.
- Include examples via `Example...` functions for doctests.
- Document concurrency, side effects, and cancellation behavior.
- Verify documentation using `pkg.go.dev`.
- Maintain README and module-level documentation.

## Dependencies & Build
- Use Go Modules (`go.mod`, `go.sum`) for dependency management.
- Pin module versions with semantic versioning.
- Keep external dependencies minimal and well-justified.
- Run `go vet`, `go fmt`, and `go mod tidy` in CI.
- Ensure reproducible builds; consider `go mod vendor` if needed.
- Use `-ldflags "-X"` for version embedding in binaries.

## Performance & Safety
- Profile before optimizing using `pprof`; measure, don’t guess.
- Use concurrency judiciously; avoid overcomplicating with goroutines.
- Reuse buffers, use `sync.Pool` for high-frequency allocations.
- Manage goroutines; ensure they exit with `context.Done()`.
- Use `defer` responsibly; avoid hiding errors.
- Avoid global mutable state; pass dependencies explicitly.
- Document invariants and unsafe assumptions.

## Async / Concurrency Code
- Prefer structured concurrency: worker pools, pipelines, and `select` loops.
- Use `context.Context` for cancellation, timeouts, and deadlines.
- Close channels properly; avoid goroutine leaks.
- Use `errgroup` from `golang.org/x/sync/errgroup` for structured error handling.
- Avoid blocking operations inside goroutines unless explicitly managed.

## Project Structure
- Keep modules and packages small and cohesive.
- Use `internal/` for private packages.
- Use `cmd/<appname>/` for executables.
- Use `pkg/` for reusable libraries and shared code.
- Separate concerns: domain logic, I/O, persistence, transport.
- Use build tags or feature flags for optional dependencies.

## CLI-Specific
- Use `cobra` or `urfave/cli/v2` for complex CLI tools.
- Provide `--help` and `--version` commands.
- Return proper exit codes (0 = success, non-0 = failure).
- Handle `EPIPE` gracefully for broken pipes.
- Use standard `flag` package for simple cases.

## Data Processing
- Stream data using channels or iterators instead of loading all into memory.
- Use goroutines and worker pools for parallel workloads when beneficial.
- Use standard encoders (`encoding/json`, `encoding/xml`, `encoding/gob`) or optimized ones thoughtfully.
- Monitor memory and GC performance via `pprof`.
- Use `bufio.Reader`/`Writer` or `io` abstractions for efficient I/O pipelines.
