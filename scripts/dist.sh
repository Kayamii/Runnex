#!/usr/bin/env bash
# Build reproducible release archives for every supported platform.
#
# Usage: VERSION=v0.1.0 ./scripts/dist.sh
#
# Archives and a SHA256SUMS file are written to ./dist. Binaries are built
# through `make build` so the version ldflags are defined in one place.
# Requires GNU tar (for --sort and --mtime) and sha256sum.
set -euo pipefail

PLATFORMS=(
	linux/amd64
	linux/arm64
	darwin/amd64
	darwin/arm64
)

: "${VERSION:?VERSION must be set, e.g. VERSION=v0.1.0}"

root="$(cd "$(dirname "$0")/.." && pwd)"
dist="$root/dist"

# Use the commit time for every timestamp in the archive so that two builds of
# the same commit produce identical files.
SOURCE_DATE_EPOCH="${SOURCE_DATE_EPOCH:-$(git -C "$root" log -1 --format=%ct)}"
export SOURCE_DATE_EPOCH

rm -rf "$dist"
mkdir -p "$dist"

for platform in "${PLATFORMS[@]}"; do
	os="${platform%/*}"
	arch="${platform#*/}"
	name="runnex_${VERSION#v}_${os}_${arch}"
	stage="$dist/$name"

	echo "==> $os/$arch"
	GOOS="$os" GOARCH="$arch" "${MAKE:-make}" -C "$root" --no-print-directory build BIN_DIR="$stage"
	cp "$root/LICENSE" "$root/README.md" "$stage/"

	tar --sort=name --mtime="@$SOURCE_DATE_EPOCH" \
		--owner=0 --group=0 --numeric-owner --mode='u+rwX,go+rX,go-w' \
		-C "$dist" -cf - "$name" | gzip -n >"$dist/$name.tar.gz"
	rm -rf "$stage"
done

(cd "$dist" && sha256sum -- *.tar.gz >SHA256SUMS)
echo "==> Artifacts in $dist"
