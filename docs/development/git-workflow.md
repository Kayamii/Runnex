# Git workflow

Runnex uses a simple, GitFlow-inspired branching model. No special tooling is
required — only ordinary Git commands and pull requests.

## Branches

| Branch      | Purpose                                         | Branches from | Merges into          |
| ----------- | ----------------------------------------------- | ------------- | -------------------- |
| `main`      | Stable, releasable code only. Every release tag points here. | —   | —                    |
| `develop`   | Integration branch for upcoming work.           | `main`        | `release/*`          |
| `feature/*` | New functionality.                              | `develop`     | `develop`            |
| `fix/*`     | Non-urgent bug fixes.                           | `develop`     | `develop`            |
| `release/*` | Release preparation (version bump, changelog, final fixes). | `develop` | `main`, then back into `develop` |
| `hotfix/*`  | Critical fixes to a released version.           | `main`        | `main`, then back into `develop` |

Branch names are lowercase and hyphenated:

```text
feature/runner-pools
feature/github-provider
feature/server-registration
fix/version-output
release/0.2.0
hotfix/0.1.1
```

## Flow

```text
feature/* ──▶ develop ──▶ release/* ──▶ main
fix/*     ──▶                 │          │
                              │          ▼
                              │         tag (vX.Y.Z) ──▶ release workflow
                              ▼
                           develop  (release fixes merged back)

hotfix/* (from main) ──▶ main ──▶ tag
                           └────▶ develop
```

### Everyday work

```bash
git switch develop && git pull
git switch -c feature/runner-pools
# ... commit ...
git push -u origin feature/runner-pools
# open a pull request into develop
```

### Releasing

1. Create `release/X.Y.Z` from `develop`.
2. Update `CHANGELOG.md`; only release-blocking fixes go on this branch.
3. Open a pull request into `main` and merge it.
4. Tag the merge commit on `main` and push the tag (see [releasing](releasing.md)).
5. Merge `main` back into `develop` so release fixes are not lost.

### Hotfixes

1. Create `hotfix/X.Y.Z` from `main`.
2. Fix, add a test, update `CHANGELOG.md`.
3. Merge into `main`, tag, then merge `main` into `develop`.

## Pull requests

- All changes reach `main` and `develop` through pull requests; CI must pass.
- Prefer **squash merge** for `feature/*` and `fix/*`, so each change is one
  Conventional Commit on `develop`.
- Use a **merge commit** for `release/*` and `hotfix/*` into `main`, so the
  history of what was released stays visible.

Recommended branch protection for `main` and `develop` (configured in the
GitHub repository settings): require pull requests, require the `CI / Check`
status, disallow force pushes.

## Commit messages

Use [Conventional Commits](https://www.conventionalcommits.org/):

```text
<type>(<optional scope>): <summary>
```

Types: `feat`, `fix`, `docs`, `refactor`, `test`, `build`, `ci`, `chore`,
`perf`, `security`.

```text
feat(runner): add ephemeral runner lifecycle
fix(scheduler): prevent over-allocation
docs: add server registration guide
ci: add ARM64 build
```

- Summary in imperative mood, lowercase, no trailing period, ≤ 72 characters.
- Use the body to explain *why* when it is not obvious.
- Breaking changes: `feat(config)!: rename pool settings` plus a
  `BREAKING CHANGE:` footer.
