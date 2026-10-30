# Contributing to Runnex

Thank you for considering a contribution. Runnex is early-stage, so
discussion of design and direction is as valuable as code.

By participating you agree to follow the [Code of Conduct](CODE_OF_CONDUCT.md).

## Before you start

- **Small fixes** (typos, obvious bugs, docs): open a pull request directly.
- **Anything larger** (new features, new packages, behavior changes): open an
  issue first so the approach can be agreed before you invest time.
- **Security issues**: do not open a public issue. Follow [SECURITY.md](SECURITY.md).

## Workflow

```text
Fork the repository
  ↓
Create a branch from develop   (feature/…, fix/…)
  ↓
Implement the change, with tests
  ↓
Run make check
  ↓
Commit using Conventional Commits
  ↓
Open a pull request against develop
```

1. Fork the repository and clone your fork.
2. Create a branch from `develop`:
   ```bash
   git switch develop
   git switch -c feature/runner-pools
   ```
3. Make your change. Keep each pull request focused on one thing.
4. Run the full validation suite locally:
   ```bash
   make check
   ```
5. Commit using [Conventional Commits](#commit-messages).
6. Push and open a pull request against `develop`. Fill in the template.

The branching model is described in
[docs/development/git-workflow.md](docs/development/git-workflow.md), and
local setup in [docs/development/getting-started.md](docs/development/getting-started.md).

## Commit messages

Runnex uses [Conventional Commits](https://www.conventionalcommits.org/):

```text
<type>(<optional scope>): <short summary in imperative mood>
```

| Type       | Use for                                          |
| ---------- | ------------------------------------------------ |
| `feat`     | New functionality                                |
| `fix`      | Bug fixes                                        |
| `docs`     | Documentation only                               |
| `refactor` | Code changes that neither fix a bug nor add a feature |
| `test`     | Adding or correcting tests                       |
| `build`    | Build system, Makefile, Dockerfile, dependencies |
| `ci`       | GitHub Actions workflows                         |
| `chore`    | Maintenance that fits nowhere else               |
| `perf`     | Performance improvements                         |
| `security` | Security hardening or fixes                      |

Examples:

```text
feat(runner): add ephemeral runner lifecycle
fix(scheduler): prevent over-allocation
docs: add server registration guide
ci: add ARM64 build
```

Mark breaking changes with `!` after the type (`feat(api)!: …`) and explain
the change in a `BREAKING CHANGE:` footer.

## Expectations

### Tests

- New behavior comes with unit tests; bug fixes come with a test that fails
  without the fix.
- Prefer table-driven tests and the standard `testing` package.
- Tests must be deterministic and must not require network access or real
  provider credentials.

### Documentation

- Update the README, `docs/` or the CHANGELOG when behavior visible to users
  changes.
- Explain *why* a design exists, not only what the code does.
- Do not document features as available before they are implemented.

### Code quality

- Idiomatic Go, formatted with `gofmt`, passing `go vet` and `golangci-lint`.
- Prefer the standard library. New dependencies need a justification in the
  pull request.
- Keep packages under `internal/` unless there is a deliberate reason to make
  an API public.
- Handle errors explicitly; wrap them with context using `fmt.Errorf("…: %w", err)`.

### Pull request descriptions

Explain what changed, why, and how you tested it. Link the related issue.
Reviewers should be able to understand the change without reading every line
first.

### Backwards compatibility

Runnex is pre-1.0, so breaking changes are allowed, but they must be
intentional: called out in the PR, marked in the commit message, and recorded
in the [CHANGELOG](CHANGELOG.md). Once configuration formats and APIs exist,
changes to them need a migration note.

### Security-sensitive changes

Changes touching authentication, credentials, provider tokens, runner
isolation or network exposure get extra scrutiny:

- Never log secrets or tokens, including at debug level.
- Request the minimum permissions a provider integration needs.
- Never commit real credentials, even in tests. Use obviously fake values.
- Describe the security impact in the pull request.

## Getting help

Open an issue with your question. There is no dedicated chat channel yet.
