// Package cli implements the runnex command-line interface.
//
// The CLI is intentionally small. Commands are added here as the
// corresponding subsystems are implemented.
package cli

import (
	"fmt"
	"io"

	"github.com/kayamii/runnex/internal/version"
)

// Exit codes returned by Run.
const (
	ExitOK    = 0
	ExitUsage = 2
)

const usage = `Runnex manages self-hosted CI/CD runners.

Usage:
  runnex <command>

Commands:
  version     Print version and build information
  help        Show this help

Runnex is in early development. See https://github.com/Kayamii/Runnex
for the roadmap.
`

// Run executes the CLI with the given arguments (excluding the program name)
// and returns the process exit code.
func Run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprint(stdout, usage)
		return ExitOK
	}

	switch args[0] {
	case "help", "-h", "--help":
		fmt.Fprint(stdout, usage)
		return ExitOK
	case "version", "--version":
		fmt.Fprint(stdout, version.Get().String())
		return ExitOK
	default:
		fmt.Fprintf(stderr, "runnex: unknown command %q\n\n", args[0])
		fmt.Fprint(stderr, usage)
		return ExitUsage
	}
}
