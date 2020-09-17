package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var (
	force bool
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
tekton-install uninstall all

# Uninstall Tekton components without being prompted for approval
tekton-install uninstall triggers dashboard pipeline -f`,
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

	if !force {
		err := confirmUninstall(allArgs, os.Stdin)
		if err != nil {
			// Return nil for err since err
			// only occurs when uninstall is cancelled
			return nil
		}
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

func confirmUninstall(components []string, reader io.Reader) error {
	qList := quotedList(components)
	fmt.Fprintf(os.Stdout, "Are you sure you want to uninstall %s components (y/n): ", qList)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		t := strings.TrimSpace(scanner.Text())
		if t == "y" {
			break
		} else if t == "n" {
			fmt.Fprintf(os.Stdout, "Cancelling uninstall of %s\n", qList)
			return fmt.Errorf("cancelling uninstall")
		}
		fmt.Fprint(os.Stdout, "Please enter y or n: ")
	}
	return nil
}

func quotedList(components []string) string {
	quoted := make([]string, len(components))
	for i := range components {
		quoted[i] = fmt.Sprintf("%q", components[i])
	}
	return strings.Join(quoted, ", ")
}

func init() {
	uninstallCmd.Flags().BoolVarP(&force, "force", "f", false, "Skip confirmation prompt for uninstalling components.")
	rootCmd.AddCommand(uninstallCmd)
}
