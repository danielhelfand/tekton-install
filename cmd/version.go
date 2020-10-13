package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	version = "v0.0.1"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version of tekton-install",
	Long: `Show the version of tekton-install.

# Show version of tekton-install
tekton-install version`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintln(cmd.OutOrStdout(), version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
