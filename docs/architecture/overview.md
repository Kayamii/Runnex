# Architecture overview

> **Status:** design intent. None of the components below are implemented
> yet. This document records the principles that implementation must follow,
> so that early decisions do not close off later ones.

## What Runnex is

Runnex is a control plane for self-hosted CI/CD runners. It receives demand
from CI providers ("a job needs a runner"), decides where that runner can
safely run, provisions it, and removes it when the job ends.

## Planned components

```text
Runnex control plane
  ├── API                    management interface for the CLI and future UI
  ├── Scheduler              decides which server gets which runner
  ├── Resource manager       tracks declared and used capacity per server
  ├── Provider integrations  GitHub, later GitLab and others
  ├── Runner manager         provisions, observes and destroys runners
  └── Authentication         operator access and provider credentials
```

These are logical boundaries, not a commitment to separate services. The
initial implementation is expected to be a single binary.

## Principles

### Provider agnostic

GitHub is an integration, not the core abstraction.

```text
Provider
   ├── GitHub
   ├── GitLab
   └── Future providers
```

The core reasons about provider-neutral concepts — a *job* needs a *runner*
with certain *labels* and *resources*. Each provider integration translates
between those concepts and the provider's API: receiving job events, issuing
registration tokens, reporting status.

**Why:** GitHub and GitLab differ in registration, token lifetime and job
discovery. If GitHub's model leaks into the scheduler, adding GitLab later
means rewriting the scheduler. Keeping the boundary from the start costs
little.

**Rule:** packages outside a provider's own package must not import that
provider's SDK or types.

### Infrastructure agnostic

Runnex should eventually manage runners on:

```text
Local server
VPS
Cloud VM
Multiple servers
```

The core must not depend on AWS, Kubernetes, Docker, or any single
infrastructure vendor. How a runner is isolated on a server (process,
container, VM) is a detail of the runner manager on that server, chosen
behind an interface.

**Why:** the target users run their own heterogeneous machines. A design
that assumes a particular cloud or orchestrator excludes most of them.

### Resource aware

Runnex tracks, per server:

```text
CPU
Memory
Disk
Concurrent jobs
Runner capacity
```

**A server must never be assigned more runner capacity than it can safely
provide.** Capacity is declared explicitly (with optional detection as a
starting point), reserved before a runner is provisioned, and released only
when the runner is gone. When no server has capacity, the job waits; Runnex
does not overcommit.

**Why:** overcommitted CI servers fail in slow, confusing ways — OOM kills,
timeouts, flaky tests. Refusing work is easier to diagnose than degrading it.

### Ephemeral by default

Each runner executes exactly one job and is then destroyed:

```text
Job
 ↓
Provision runner
 ↓
Execute
 ↓
Collect status/logs
 ↓
Destroy runner
```

**Why:** long-lived runners accumulate files, caches, credentials and
processes from earlier jobs. That makes builds non-reproducible and lets one
job affect the next — a security problem when jobs come from different
repositories. Disposable runners remove that class of problem. Caching for
speed should be added deliberately, not inherited from leftovers.

Teardown is part of the lifecycle, not cleanup after it: a runner that cannot
be confirmed destroyed keeps its resources reserved and is reported.

### Secure by default

- Credentials and provider tokens are never logged, at any log level.
- Secrets are never committed to the repository; configuration references
  them (files, environment) rather than embedding them.
- Provider integrations request the minimum permissions they need — for
  GitHub, a GitHub App with narrowly scoped permissions rather than a personal
  access token.
- Runner registration tokens are short-lived and scoped to a single runner.
- Defaults are the safe choice; less safe behavior requires explicit opt-in.

## Code layout

| Path           | Purpose                                                       |
| -------------- | ------------------------------------------------------------- |
| `cmd/runnex`   | Entry point. Parses arguments and delegates; contains no logic. |
| `internal/`    | All application code. Not importable by other modules, so internal APIs can change freely before 1.0. |
| `pkg/`         | Not created yet. Only for code deliberately offered as a stable public API (for example, a provider SDK). |
| `deployments/` | Deployment artifacts (systemd). The `Dockerfile` stays at the root, where tooling expects it. |
| `scripts/`     | Build and release helpers used by the Makefile.               |

Expected future packages under `internal/` (names indicative, not final):
`config`, `api`, `scheduler`, `resource`, `provider` (with `provider/github`),
`runner`, `auth`. They will be created when there is code to put in them.
