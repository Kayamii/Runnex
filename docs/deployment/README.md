# Deployment

> **Status:** Runnex has no long-running service yet, so there is nothing to
> deploy in production. This page describes the intended deployment model and
> the artifacts that exist today.

## Intended model

Runnex is meant to run as a single, self-contained binary on a host you
control, with no required external services. Two supported ways to run it
are planned:

- **systemd service** on a Linux host — the primary target.
- **Container** — for environments that already run containers.

## Artifacts available today

| Artifact                                                | State |
| ------------------------------------------------------- | ----- |
| Release binaries (Linux/macOS, amd64/arm64)             | Built by the release workflow on version tags |
| [`Dockerfile`](../../Dockerfile)                        | Builds a minimal, non-root image containing the CLI. No image is published. |
| [`deployments/systemd/runnex.service`](../../deployments/systemd/runnex.service) | Example unit with sandboxing defaults. It cannot be used until a server command exists. |

Build the container image locally:

```bash
docker build -t runnex:dev .
docker run --rm runnex:dev version
```

This page will document installation, configuration and upgrades once the
control plane exists.
