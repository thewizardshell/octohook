<h1 align="center"> Octohook </h1>
<div align="center">
  <img src="https://github.com/user-attachments/assets/3d589a9e-ed13-401c-8ab9-663bf43b8499" alt="OctoHook Logo" width="600" />
  
  <p><strong>Git hooks made fast and simple</strong></p>
  <p>Zero dependencies. Intelligent caching. Any project.</p>

  [![Release](https://img.shields.io/github/v/release/thewizardshell/octohook?label=Release&color=8b5cf6)](https://github.com/thewizardshell/octohook/releases)
  [![Platforms](https://img.shields.io/badge/Platforms-Windows%20%7C%20Linux%20%7C%20macOS-a78bfa?style=flat)](#)
  [![Go Version](https://img.shields.io/badge/Go-1.20%2B-7c3aed?style=flat)](#)

</div>

## Why OctoHook?

- **One binary** — No Node, no Python, just works
- **Smart caching** — Skip hooks when nothing changed
- **Any project** — Go, Node, Python, Rust, whatever
- **Beautiful TUI** — See your hooks at a glance

## Quick Start

```bash
# Coming soon
```

## Configuration

```yaml
# octohook.yml
pre-commit:
  - name: lint
    command: golangci-lint run
    files: ["*.go"]
```

## Roadmap

- [ ] Core hook execution
- [ ] Intelligent cache system
- [ ] TUI dashboard
- [ ] Parallel hook execution

## License
MIT
