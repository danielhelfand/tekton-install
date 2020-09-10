package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	pipelineVersion  string
	triggersVersion  string
	dashboardVersion string
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install various components of Tekton",
	Long: `Install the pipeline, triggers, and dashboard components. Use flags to specify specific versions of components.

# Install latest version of Tekton pipeline component
tekton-install install pipeline

# Install specific version of Tekton pipeline component
tekton-install install pipeline --pipeline-version 0.15.0

# Install latest version of Tekton triggers component
tekton-install install triggers

# Install specific version of Tekton triggers component
tekton-install install triggers --triggers-version 0.6.0

# Install latest version of Tekton dashboard component
tekton-install install dashboard

# Install specific version of Tekton dashboard component
tekton-install install dashboard --dashboard-version 0.6.0

# Install all of latest components
tekton-install install all

# Install all components with specific versions
tekton-install install all --pipeline-version 0.15.0 --triggers-version 0.6.0 --dashboard-version 0.6.0`,
	Args: cobra.RangeArgs(1, 3),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := validateArgs(args); err != nil {
			return err
		}

		return install(args)
	},
}

func install(args []string) error {
	var allArgs []string
	if args[0] == "all" {
		allArgs = []string{pipeline, dashboard, triggers}
	} else {
		allArgs = args
	}

	componentVersions := make(map[string]string)
	for _, arg := range allArgs {
		switch arg {
		case pipeline:
			componentVersions[arg] = pipelineVersion
		case dashboard:
			componentVersions[arg] = dashboardVersion
		case triggers:
			componentVersions[arg] = triggersVersion
		}
	}

	for _, arg := range allArgs {
		var argv []string
		if arg != dashboard && componentVersions[arg] == "" {
			argv = []string{"apply", "-f", "https://storage.googleapis.com/tekton-releases/" + arg + "/latest/release.yaml"}
		} else if arg != dashboard {
			argv = []string{"apply", "-f", "https://storage.googleapis.com/tekton-releases/" + arg + "/previous/v" + componentVersions[arg] + "/release.yaml"}
		}
		if arg == dashboard && componentVersions[arg] == "" {
			argv = []string{"apply", "-f", "https://storage.googleapis.com/tekton-releases/dashboard/latest/tekton-dashboard-release.yaml"}
		} else if arg == dashboard {
			argv = []string{"apply", "-f", "https://storage.googleapis.com/tekton-releases/" + arg + "/previous/v" + componentVersions[arg] + "/tekton-dashboard-release.yaml"}
		}
		kubectlCmd := exec.Command("kubectl", argv...)
		// Report command errors to stderr
		kubectlCmd.Env = os.Environ()
		kubectlCmd.Stderr = os.Stderr

		if err := kubectlCmd.Run(); err != nil {
			return fmt.Errorf("installation of %s has failed", arg)
		}
		fmt.Printf("Component %s has been installed successfully\n", arg)
	}

	return nil
}

func init() {
	installCmd.Flags().StringVar(&pipelineVersion, "pipeline-version", "", "The version of pipeline to install.")
	installCmd.Flags().StringVar(&triggersVersion, "triggers-version", "", "The version of triggers to install.")
	installCmd.Flags().StringVar(&dashboardVersion, "dashboard-version", "", "The version of dashboard to install.")
	rootCmd.AddCommand(installCmd)
}
