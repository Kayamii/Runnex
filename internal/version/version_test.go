package version

import (
	"runtime/debug"
	"strings"
	"testing"
)

func TestApplyBuildInfo(t *testing.T) {
	bi := &debug.BuildInfo{
		Main: debug.Module{Version: "v0.3.0"},
		Settings: []debug.BuildSetting{
			{Key: "vcs.revision", Value: "abc123"},
			{Key: "vcs.time", Value: "2026-01-02T03:04:05Z"},
		},
	}

	t.Run("fills unset fields", func(t *testing.T) {
		info := Info{Version: "dev", Commit: "unknown", Date: "unknown"}
		applyBuildInfo(&info, bi)

		want := Info{Version: "v0.3.0", Commit: "abc123", Date: "2026-01-02T03:04:05Z"}
		if info != want {
			t.Errorf("got %+v, want %+v", info, want)
		}
	})

	t.Run("ldflags take precedence", func(t *testing.T) {
		info := Info{Version: "v1.0.0", Commit: "def456", Date: "2026-06-01T00:00:00Z"}
		want := info
		applyBuildInfo(&info, bi)

		if info != want {
			t.Errorf("got %+v, want %+v", info, want)
		}
	})

	t.Run("ignores devel module version", func(t *testing.T) {
		info := Info{Version: "dev"}
		applyBuildInfo(&info, &debug.BuildInfo{Main: debug.Module{Version: "(devel)"}})

		if info.Version != "dev" {
			t.Errorf("Version = %q, want %q", info.Version, "dev")
		}
	})
}

func TestInfoString(t *testing.T) {
	s := Info{Version: "v0.1.0", Commit: "abc", Date: "d", GoVersion: "go1.27.0", Platform: "linux/amd64"}.String()

	for _, want := range []string{"runnex v0.1.0", "abc", "go1.27.0", "linux/amd64"} {
		if !strings.Contains(s, want) {
			t.Errorf("String() = %q, missing %q", s, want)
		}
	}
}
