// Command anthropic-cli provides a command-line interface for interacting with Anthropic's Claude AI models.
package main

import (
	"os"

	"github.com/anthropics/anthropic-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}