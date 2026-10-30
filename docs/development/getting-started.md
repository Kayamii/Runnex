# Getting started with development

## Requirements

- **Go** — the version declared in [`go.mod`](../../go.mod). With a recent Go
  installed, the toolchain is downloaded automatically if needed.
- **GNU Make**
- **golangci-lint v2** — installation:
  <https://golangci-lint.run/welcome/install/>. CI pins the version in
  [`.github/workflows/ci.yml`](../../.github/workflows/ci.yml); use the same one
  locally to avoid differing results.
- **Git**

Windows works for building and testing; use Git Bash or WSL so the Makefile's
shell commands behave. `make dist` additionally requires GNU tar and
`sha256sum` (Linux, or macOS with GNU tar installed).

## Building and running

```bash
make build                 # → bin/runnex
./bin/runnex version
make run ARGS="help"       # build and run in one step
```

The version printed by a local build comes from `git describe`.

## Make targets

| Target           | What it does                                                   |
| ---------------- | -------------------------------------------------------------- |
| `make help`      | List targets                                                   |
| `make build`     | Build a static binary into `bin/` with version information     |
| `make run`       | Run from source; pass arguments via `ARGS="..."`               |
| `make test`      | Run unit tests                                                 |
| `make test-race` | Run unit tests with the race detector (needs cgo; used in CI)  |
| `make fmt`       | Format code with `gofmt -s`                                    |
| `make fmt-check` | Fail if any file is not formatted                              |
| `make vet`       | Run `go vet`                                                   |
| `make lint`      | Run `golangci-lint`                                            |
| `make tidy`      | Run `go mod tidy`                                              |
| `make mod-check` | Verify module checksums and that `go.mod` is tidy              |
| `make check`     | `mod-check`, `fmt-check`, `vet`, `lint`, `test`, `build`       |
| `make dist`      | Build release archives into `dist/`                            |
| `make clean`     | Remove `bin/` and `dist/`                                      |

Run `make check` before pushing; CI runs the same steps.

## Code layout

See [Architecture overview → Code layout](../architecture/overview.md#code-layout).
New code goes under `internal/` unless there is a deliberate reason to make it
a public API.

## Further reading

- [Git workflow](git-workflow.md)
- [Releasing](releasing.md)
- [CONTRIBUTING.md](../../CONTRIBUTING.md)
