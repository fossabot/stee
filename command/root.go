package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "stee",
	Short: "Stee is a simple URL minifier.",
	Long: `A simple URL minifier with a lot of extra possibilities`,
}

// Execute executes the provided command, taking care of arguments and flags.
func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}