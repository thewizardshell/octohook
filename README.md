<h1 align="center"> Octohook </h1>
<div align="center">
  <img src="https://github.com/user-attachments/assets/3d589a9e-ed13-401c-8ab9-663bf43b8499" alt="OctoHook Logo" width="600" />
  
  <p><strong>Intelligent git hooks with dependency tracking</strong></p>
  <p>Built-in caching. Dependency-aware testing. Multi-language support.</p>

  [![Release](https://img.shields.io/github/v/release/thewizardshell/octohook?label=Release&color=8b5cf6)](https://github.com/thewizardshell/octohook/releases)
  [![Platforms](https://img.shields.io/badge/Platforms-Windows%20%7C%20Linux%20%7C%20macOS-a78bfa?style=flat)](#)
  [![Go Version](https://img.shields.io/badge/Go-1.20%2B-7c3aed?style=flat)](#)

</div>

## Why OctoHook?

- **Dependency-aware** — Only runs tests affected by your changes
- **Smart caching** — Skips unchanged tests automatically
- **Multi-language** — Supports Go, TypeScript, JavaScript, and Python
- **Zero dependencies** — Single binary, no runtime required
- **Beautiful TUI** — Real-time test execution with visual feedback

## How It Works

OctoHook builds a dependency graph of your project by analyzing imports across your codebase. When you commit changes, it:

1. Identifies which files were modified
2. Traces their dependencies to find affected code
3. Runs only the tests that matter
4. Caches results to skip redundant test runs

This means faster commits and more confidence that you're testing what actually changed.

## Installation

```bash
# Download the latest release
# Coming soon

# Or build from source
go build -o octohook cmd/main.go

# Install hooks in your project
./octohook install
```

## Configuration

Create an `octohook.yml` file in your project root:

```yaml
pre-commit:
  command: npm
  arg: ["test"]
  path:
    services:
      - "src/**/*.ts"
    test:
      - "src/**/*.spec.ts"
  cache: true
  use_directory: false
```

### Configuration Options

- `command`: The executable to run (e.g., `npm`, `go`, `pytest`)
- `arg`: Arguments passed to the command
- `path.services`: Glob patterns for source files
- `path.test`: Glob patterns for test files
- `cache`: Enable/disable result caching (default: true)
- `use_directory`: Pass directory instead of file path (useful for Go)

### Multi-Language Examples

**TypeScript/JavaScript:**
```yaml
pre-commit:
  command: npm
  arg: ["test"]
  path:
    services: ["**/*.ts"]
    test: ["**/*.spec.ts"]
```

**Go:**
```yaml
pre-commit:
  command: go
  arg: ["test"]
  path:
    services: ["**/*.go"]
    test: ["**/*_test.go"]
  use_directory: true
```

**Python:**
```yaml
pre-commit:
  command: pytest
  arg: []
  path:
    services: ["**/*.py"]
    test: ["**/test_*.py"]
```

## Available Hooks

- `pre-commit` — Runs before commit is created
- `post-commit` — Runs after commit is created
- `pre-push` — Runs before pushing to remote
- `post-push` — Runs after pushing to remote

## Commands

```bash
octohook install           # Install all configured hooks
octohook uninstall         # Remove all hooks
octohook uninstall-hook    # Remove a specific hook
octohook pre-commit        # Run pre-commit hook manually
octohook post-commit       # Run post-commit hook manually
octohook pre-push          # Run pre-push hook manually
octohook post-push         # Run post-push hook manually
```

## Features

### Dependency Graph Analysis
OctoHook parses your codebase to understand import relationships, building a reverse dependency graph that tracks which files depend on others.

### Intelligent Test Selection
Instead of running all tests, OctoHook identifies:
- Tests directly affected by changed files
- Tests in the same directory as changed files
- Tests for files that import your changes

### File-based Caching
Results are cached using SHA-256 hashes of file contents. If neither the test nor its related source files have changed, the test is skipped.

### Real-time TUI
Built with Bubble Tea, the terminal interface shows test progress with spinners, status indicators, and immediate feedback on failures.

## Roadmap

- [x] Core hook execution engine
- [x] Dependency graph analysis
- [x] Intelligent cache system
- [x] Basic multi-language support (Go, TS, JS, Python)
- [x] TUI with real-time feedback
- [x] Hook installation/uninstallation
- [ ] Installation script
- [ ] Improve dependency graph system
- [ ] Interactive init command for configuration
- [ ] npx-based installation
- [ ] Extended language support
- [ ] Comprehensive testing suite
- [ ] Test in production projects
- [ ] Pre-built binaries for all platforms

## License

MIT
