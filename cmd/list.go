package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

const (
	header = "COMPONENT\tVERSION"
	body   = "%s\t%s\n"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed Tekton components on a Kubernetes cluster",
	Long: `List installed Tekton components on a Kubernetes cluster.

# List available Tekton components on a Kubernetes cluster
tekton-install list`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		componentVersions := make(map[string]string)
		for _, component := range components {
			version, _ := getComponentVersion(component, true)
			if version != "" {
				componentVersions[component] = version
			}
		}

		if len(componentVersions) == 0 {
			fmt.Println("No components installed")
		} else {
			list(componentVersions)
		}
	},
}

func list(componentVersions map[string]string) {
	w := tabwriter.NewWriter(os.Stdout, 0, 5, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, header)

	for component, version := range componentVersions {
		fmt.Fprintf(w, body, component, version)
	}

	w.Flush()
}

func init() {
	rootCmd.AddCommand(listCmd)
}
