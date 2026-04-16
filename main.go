// Command anthropic-cli provides a command-line interface for interacting with Anthropic's Claude AI models.
// Personal fork: added non-zero exit code logging for easier debugging in scripts.
package main

import (
	"fmt"
	"os"

	"github.com/anthropics/anthropic-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
