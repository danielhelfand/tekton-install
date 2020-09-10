package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	components = []string{triggers, dashboard, pipeline}
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall various components of Tekton",
	Long: `Uninstall the pipeline, triggers, and dashboard components. 

# Uninstall the Tekton pipeline component
# NOTE: Uninstalling pipeline component will
# also uninstall other installed Tekton components
tekton-install uninstall pipeline

# Uninstall the Tekton triggers component
tekton-install uninstall triggers

# Uninstall the Tekton dashboard component
tekton-install uninstall dashboard

# Uninstall all Tekton components
tekton-install uninstall all`,
	Args: cobra.RangeArgs(1, 3),
	RunE: func(cmd *cobra.Command, args []string) error {
		return uninstall(args)
	},
}

func uninstall(args []string) error {
	if err := validateArgs(args); err != nil {
		return err
	}

	var allArgs []string
	all := false
	if args[0] == "all" {
		allArgs = components
		all = true
	} else {
		allArgs = args
	}

	componentVersions := make(map[string]string)
	for _, arg := range allArgs {
		version, err := getComponentVersion(arg, all)
		if err != nil {
			return err
		}
		if version != "" {
			componentVersions[arg] = version
		}
	}

	if len(componentVersions) == 0 {
		fmt.Println("No components on cluster to uninstall")
		return nil
	}

	return uninstallComponents(componentVersions)
}

func getComponentVersion(component string, all bool) (string, error) {
	var argv []string
	if component != dashboard {
		// Since deployment for pipeline is named tekton-pipelines-controller, reassign component name to pipelines
		if component == pipeline {
			component = "pipelines"
		}
		argv = []string{"get", "deployment/tekton-" + component + "-controller", "-o", "jsonpath={.metadata.labels['app\\.kubernetes\\.io/version']}", "-n", "tekton-pipelines"}
	} else {
		argv = []string{"get", "deployment/tekton-" + component, "-o", "jsonpath={.metadata.labels.version}", "-n", "tekton-pipelines"}
	}

	kubectlCmd := exec.Command("kubectl", argv...)
	version, err := kubectlCmd.Output()
	if err != nil {
		if all {
			return "", nil
		}

		return "", fmt.Errorf("failed to get version of component %s, check if it is installed", component)
	}

	return string(version), nil
}

func uninstallComponents(componentVersions map[string]string) error {
	for _, component := range components {
		var argv []string
		if _, ok := componentVersions[component]; ok {
			if component != dashboard {
				argv = []string{"delete", "-f", "https://storage.googleapis.com/tekton-releases/" + component + "/previous/" + componentVersions[component] + "/release.yaml"}
			} else {
				argv = []string{"delete", "-f", "https://storage.googleapis.com/tekton-releases/" + component + "/previous/" + componentVersions[component] + "/tekton-dashboard-release.yaml"}
			}
			kubectlCmd := exec.Command("kubectl", argv...)
			kubectlCmd.Env = os.Environ()
			kubectlCmd.Stderr = os.Stderr
			if err := kubectlCmd.Run(); err != nil {
				return fmt.Errorf("uninstall of %s has failed", component)
			}
			fmt.Printf("Component %s has been uninstalled successfully\n", component)
		}
	}
	return nil
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
