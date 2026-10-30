// Command runnex is the entry point for the Runnex CLI.
package main

import (
	"os"

	"github.com/kayamii/runnex/internal/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:], os.Stdout, os.Stderr))
}
