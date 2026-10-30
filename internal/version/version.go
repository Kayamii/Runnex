// Package version exposes build metadata for the Runnex binary.
//
// Release builds inject Version, Commit and Date through -ldflags (see the
// Makefile). Builds without ldflags, such as `go install`, fall back to the
// module and VCS information embedded by the Go toolchain.
package version

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

// Set at build time via -ldflags "-X github.com/kayamii/runnex/internal/version.Version=...".
var (
	// Version is the semantic version of the build, e.g. "v0.1.0".
	Version = "dev"
	// Commit is the full Git commit hash the binary was built from.
	Commit = "unknown"
	// Date is the commit timestamp in RFC 3339 format. The commit time is used
	// instead of the build time so that builds are reproducible.
	Date = "unknown"
)

// Info describes the running binary.
type Info struct {
	Version   string
	Commit    string
	Date      string
	GoVersion string
	Platform  string
}

// Get returns build information for the running binary.
func Get() Info {
	info := Info{
		Version:   Version,
		Commit:    Commit,
		Date:      Date,
		GoVersion: runtime.Version(),
		Platform:  runtime.GOOS + "/" + runtime.GOARCH,
	}
	if bi, ok := debug.ReadBuildInfo(); ok {
		applyBuildInfo(&info, bi)
	}
	return info
}

// applyBuildInfo fills fields that were not set through ldflags from the
// information embedded by the Go toolchain.
func applyBuildInfo(info *Info, bi *debug.BuildInfo) {
	if info.Version == "dev" && bi.Main.Version != "" && bi.Main.Version != "(devel)" {
		info.Version = bi.Main.Version
	}
	for _, s := range bi.Settings {
		switch s.Key {
		case "vcs.revision":
			if info.Commit == "unknown" {
				info.Commit = s.Value
			}
		case "vcs.time":
			if info.Date == "unknown" {
				info.Date = s.Value
			}
		}
	}
}

// String renders the build information in a human-readable form.
func (i Info) String() string {
	return fmt.Sprintf("runnex %s\n  commit:   %s\n  date:     %s\n  go:       %s\n  platform: %s\n",
		i.Version, i.Commit, i.Date, i.GoVersion, i.Platform)
}
