package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		wantCode   int
		wantStdout string
		wantStderr string
	}{
		{name: "no args shows help", args: nil, wantCode: ExitOK, wantStdout: "Usage:"},
		{name: "help", args: []string{"help"}, wantCode: ExitOK, wantStdout: "Usage:"},
		{name: "-h", args: []string{"-h"}, wantCode: ExitOK, wantStdout: "Usage:"},
		{name: "--help", args: []string{"--help"}, wantCode: ExitOK, wantStdout: "Usage:"},
		{name: "version", args: []string{"version"}, wantCode: ExitOK, wantStdout: "runnex "},
		{name: "--version", args: []string{"--version"}, wantCode: ExitOK, wantStdout: "runnex "},
		{name: "unknown command", args: []string{"bogus"}, wantCode: ExitUsage, wantStderr: `unknown command "bogus"`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			code := Run(tt.args, &stdout, &stderr)

			if code != tt.wantCode {
				t.Errorf("exit code = %d, want %d", code, tt.wantCode)
			}
			if !strings.Contains(stdout.String(), tt.wantStdout) {
				t.Errorf("stdout = %q, want it to contain %q", stdout.String(), tt.wantStdout)
			}
			if !strings.Contains(stderr.String(), tt.wantStderr) {
				t.Errorf("stderr = %q, want it to contain %q", stderr.String(), tt.wantStderr)
			}
			if tt.wantStderr == "" && stderr.Len() > 0 {
				t.Errorf("unexpected stderr output: %q", stderr.String())
			}
		})
	}
}
