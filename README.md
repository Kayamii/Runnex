# Runnex

Lightweight, self-hosted management for ephemeral CI/CD runners.

[![CI](https://github.com/Kayamii/Runnex/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/Kayamii/Runnex/actions/workflows/ci.yml)
[![Go version](https://img.shields.io/github/go-mod/go-version/Kayamii/Runnex)](go.mod)
[![License](https://img.shields.io/github/license/Kayamii/Runnex)](LICENSE)
[![Release](https://img.shields.io/github/v/release/Kayamii/Runnex?include_prereleases&sort=semver)](https://github.com/Kayamii/Runnex/releases)
[![Issues](https://img.shields.io/github/issues/Kayamii/Runnex)](https://github.com/Kayamii/Runnex/issues)

> [!WARNING]
> **Runnex is in early development and is not usable yet.** The repository
> currently contains the project foundation: build tooling, CI, release
> automation and a minimal CLI. No runner orchestration exists yet. Do not
> deploy it expecting it to manage runners.

## The problem

Self-hosted runners are powerful: they give you control over hardware, network
access, caches and cost. But operating them across many repositories and
several servers quickly becomes difficult:

- Runners are registered by hand and drift out of date.
- Long-lived runners accumulate state from previous jobs, which hurts
  reproducibility and security.
- Nothing stops a server from being given more concurrent jobs than its CPU
  and memory can handle.
- Each CI provider has its own registration model, so every new provider means
  new tooling.

Runnex aims to make runner infrastructure simple, lightweight and
controllable: you bring the servers, Runnex provisions disposable runners on
them as jobs arrive, within the resource limits you define.

## Features

| Status          | Feature                                                         |
| --------------- | --------------------------------------------------------------- |
| **Implemented** | `runnex` CLI with `version` and `help` commands                 |
| **Implemented** | Reproducible, versioned release builds for Linux and macOS      |
| **Planned**     | Configuration system                                            |
| **Planned**     | GitHub Actions integration via a GitHub App                     |
| **Planned**     | Ephemeral runner lifecycle (provision → run one job → destroy)  |
| **Planned**     | Server registration and resource-aware scheduling               |
| **Future**      | Web dashboard                                                   |
| **Future**      | GitLab CI and other providers                                   |

"Planned" means it is part of the current roadmap. "Future" means it is an
intended direction without committed design work yet.

## Architecture

Runnex is designed as a single control plane that coordinates CI providers and
the servers that run jobs. The core is deliberately independent of any single
CI provider or infrastructure vendor.

```text
                 ┌───────────────────────────────────────┐
                 │          Runnex control plane         │
                 │                                       │
 CI providers    │  API ─── Scheduler ─── Resource mgr.  │    Servers
 ────────────    │   │          │              │         │    ───────
 GitHub   ◀────▶ │  Provider    Runner manager           │ ◀─▶ local / VPS /
 GitLab*  ◀────▶ │  integrations                         │     cloud VM
                 │                                       │
                 │  Authentication                       │
                 └───────────────────────────────────────┘
                                         * future
```

None of these components exist yet. The guiding principles — provider
agnostic, infrastructure agnostic, resource aware, ephemeral by default,
secure by default — are described in
[docs/architecture/overview.md](docs/architecture/overview.md).

## Why Runnex?

- **Small footprint.** A single static Go binary with minimal dependencies.
- **Self-hosted.** Your servers, your network, your credentials. No hosted
  service is required.
- **Explicit resources.** Capacity is declared and enforced, not guessed.
- **Ephemeral by default.** Every job gets a clean runner.
- **Not tied to one CI system.** GitHub is the first integration, not the core
  abstraction.

## Quick Start

There is nothing to operate yet. You can build and run the CLI:

```bash
git clone https://github.com/Kayamii/Runnex.git
cd Runnex
make build
./bin/runnex version
```

Once releases are published, prebuilt binaries for Linux and macOS
(amd64/arm64) will be available on the
[Releases page](https://github.com/Kayamii/Runnex/releases).

## Development

Requirements: Go (version in [`go.mod`](go.mod)), GNU Make, and
[golangci-lint](https://golangci-lint.run/) v2.

```bash
make help     # list all targets
make check    # everything CI runs: modules, formatting, vet, lint, tests, build
make run ARGS="version"
```

See [docs/development/getting-started.md](docs/development/getting-started.md)
for the full setup and [docs/development/git-workflow.md](docs/development/git-workflow.md)
for branching and commit conventions.

## Configuration

Runnex has no configuration yet. The configuration system is the next item on
the roadmap; this section will document it once it exists.

## Supported Providers

| Provider       | Status  |
| -------------- | ------- |
| GitHub Actions | Planned |
| GitLab CI      | Future  |

See [docs/providers/](docs/providers/README.md) for how providers fit into the
design.

## Roadmap

No dates are promised. The order reflects current priorities and will be
adjusted as development reveals better boundaries.

### Phase 1 — Foundation

- [x] Project structure
- [x] CLI skeleton
- [x] CI pipeline
- [x] Release pipeline
- [ ] Configuration system
- [ ] Control plane API
- [ ] Authentication

### Phase 2 — GitHub

- [ ] GitHub App integration
- [ ] Repository and organization management
- [ ] Runner registration
- [ ] Ephemeral runner lifecycle
- [ ] Runner pools
- [ ] Job tracking

### Phase 3 — Infrastructure

- [ ] Server registration
- [ ] Resource detection
- [ ] CPU and memory allocation
- [ ] Resource limits
- [ ] Multi-server scheduling
- [ ] Worker installation

### Phase 4 — Dashboard

- [ ] Web UI
- [ ] Server management
- [ ] Runner pools
- [ ] Jobs and logs
- [ ] Metrics

### Phase 5 — Providers

- [ ] Stable provider abstraction
- [ ] GitLab CI
- [ ] Additional CI providers

## Contributing

Contributions are welcome. Please read [CONTRIBUTING.md](CONTRIBUTING.md) and
the [Code of Conduct](CODE_OF_CONDUCT.md) first. For anything beyond a small
fix, open an issue to discuss the approach before writing code.

## Security

Please do not report vulnerabilities in public issues. See
[SECURITY.md](SECURITY.md) for how to report them privately.

## License

Runnex is licensed under the [Apache License 2.0](LICENSE).
