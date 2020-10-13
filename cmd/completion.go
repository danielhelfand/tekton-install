package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh]",
	Short: "Print completion scripts for tekton-install commands",
	Long: `Detailed instructions for enabling shell autocompletion 
with tekton-install are available at the following link:

https://github.com/danielhelfand/tekton-install#shell-autocompletion`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh"},
	Args:                  cobra.ExactValidArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		switch args[0] {
		case "bash":
			err = cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			err = cmd.Root().GenZshCompletion(os.Stdout)
		}

		return err
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
