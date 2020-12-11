package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCommand.AddCommand(versionCommand)
}
  
var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Print the version of Stee",
	Long:  `Print the version of Stee`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`Stee v0.0
Development version - very unstable`)
	},
	Aliases: []string{"v"},
}