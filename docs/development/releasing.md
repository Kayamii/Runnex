# Releasing

Releases are driven by Git tags. Pushing to a branch never creates a release.

## Versioning

Runnex follows [Semantic Versioning](https://semver.org/). Tags have the form
`vMAJOR.MINOR.PATCH`, optionally with a pre-release suffix:

```text
v0.1.0
v0.1.1
v0.2.0
v0.2.0-rc.1   # published as a pre-release
```

Before 1.0, minor versions may contain breaking changes; they are recorded in
the changelog.

## Cutting a release

From an up-to-date `main` that contains the release (see
[Git workflow](git-workflow.md#releasing)):

```bash
git switch main && git pull
git tag -a v0.1.0 -m "v0.1.0"
git push origin v0.1.0
```

## What the release workflow does

[`.github/workflows/release.yml`](../../.github/workflows/release.yml):

1. Rejects tags that are not valid `vX.Y.Z[-pre]` versions.
2. Runs the unit tests.
3. Runs `make dist`, which builds static binaries for:
   - `linux/amd64`, `linux/arm64`
   - `darwin/amd64`, `darwin/arm64`
4. Creates a GitHub Release for the tag with generated notes, the
   `runnex_<version>_<os>_<arch>.tar.gz` archives and a `SHA256SUMS` file.
   Tags with a pre-release suffix are marked as pre-releases.

## Version information

`make build` injects the version (from the tag), full commit hash and commit
date into `internal/version` via `-ldflags`. `runnex version` prints them.

## Reproducibility

Two builds of the same commit with the same Go version produce identical
binaries and archives:

- `-trimpath` removes local file paths; `-buildid=` removes the build ID.
- `CGO_ENABLED=0` avoids dependence on the host C toolchain.
- The embedded date is the **commit** date, not the build time.
- Archives use the commit time for every file, fixed ownership and modes,
  sorted entries, and `gzip -n`.

The Go version is pinned by `go.mod`, so local and CI builds use the same
toolchain. Verify a release by rebuilding the tag and comparing against
`SHA256SUMS`:

```bash
git checkout v0.1.0
make dist VERSION=v0.1.0
cat dist/SHA256SUMS
```

## Not yet included

Signing (e.g. Sigstore/cosign), SBOMs, and container image publishing are not
set up. They are worth adding once there is a release people depend on.
