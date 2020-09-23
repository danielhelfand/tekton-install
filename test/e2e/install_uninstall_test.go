// +build e2e

package e2e

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"gotest.tools/assert"
)

const (
	noResourcesFound  = "No resources found in tekton-pipelines namespace.\n"
	allCompsInstall   = "Component pipeline has been installed successfully\nComponent dashboard has been installed successfully\nComponent triggers has been installed successfully\n"
	allCompsUninstall = "Component triggers has been uninstalled successfully\nComponent dashboard has been uninstalled successfully\nComponent pipeline has been uninstalled successfully\n"
)

func Test_Install_Uninstall_Commands(t *testing.T) {
	t.Run("Install pipeline component", func(t *testing.T) {
		argv := []string{"install", "pipeline"}
		output, errMsg := ExecuteCommandOutput(TektonInstallCmd, argv)
		if errMsg != "" {
			t.Log(errMsg)
		}

		if d := cmp.Diff(output, installSuccess("pipeline")); d != "" {
			t.Fatalf("-got, +want: %v", d)
		}

		t.Log("Waiting for pods to be available in tekton-pipelines namespace")
		_, errMsg = WaitForAllPodStatus("Ready", "tekton-pipelines", "3m")
		if errMsg != "" {
			t.Log(errMsg)
		}
	})

	t.Run("Install triggers component", func(t *testing.T) {
		argv := []string{"install", "triggers"}
		output, errMsg := ExecuteCommandOutput(TektonInstallCmd, argv)
		if errMsg != "" {
			t.Log(errMsg)
		}

		if d := cmp.Diff(output, installSuccess("triggers")); d != "" {
			t.Fatalf("-got, +want: %v", d)
		}

		t.Log("Waiting for pods to be available in tekton-pipelines namespace")
		_, errMsg = WaitForAllPodStatus("Ready", "tekton-pipelines", "3m")
		if errMsg != "" {
			t.Log(errMsg)
		}
	})

	t.Run("Install dashboard component", func(t *testing.T) {
		argv := []string{"install", "dashboard"}
		output, errMsg := ExecuteCommandOutput(TektonInstallCmd, argv)
		if errMsg != "" {
			t.Log(errMsg)
		}

		if d := cmp.Diff(output, installSuccess("dashboard")); d != "" {
			t.Fatalf("-got, +want: %v", d)
		}

		t.Log("Waiting for pods to be available in tekton-pipelines namespace")
		_, errMsg = WaitForAllPodStatus("Ready", "tekton-pipelines", "3m")
		if errMsg != "" {
			t.Log(errMsg)
		}
	})

	t.Run("Uninstall dashboard component", func(t *testing.T) {
		argv := []string{"uninstall", "dashboard", "-f"}
		output, errMsg := ExecuteCommandOutput(TektonInstallCmd, argv)
		if errMsg != "" {
			t.Log(errMsg)
		}

		if d := cmp.Diff(output, uninstallSuccess("dashboard")); d != "" {
			t.Fatalf("-got, +want: %v", d)
		}
	})

	t.Run("Uninstall triggers component", func(t *testing.T) {
		argv := []string{"uninstall", "triggers", "-f"}
		output, errMsg := ExecuteCommandOutput(TektonInstallCmd, argv)
		if errMsg != "" {
			t.Log(errMsg)
		}

		if d := cmp.Diff(output, uninstallSuccess("triggers")); d != "" {
			t.Fatalf("-got, +want: %v", d)
		}
	})

	t.Run("Uninstall pipeline component", func(t *testing.T) {
		argv := []string{"uninstall", "pipeline", "-f"}
		output, errMsg := ExecuteCommandOutput(TektonInstallCmd, argv)
		if errMsg != "" {
			t.Log(errMsg)
		}

		if d := cmp.Diff(output, uninstallSuccess("pipeline")); d != "" {
			t.Fatalf("-got, +want: %v", d)
		}

		// Make sure no pods remain after uninstall for all components
		argv = []string{"get", "pods", "-n", "tekton-pipelines"}
		output, errMsg = ExecuteCommandOutput(KubectlCmd, argv)
		if errMsg == "" {
			t.Logf("Expected no pods to be found but pods were found\n%s", output)
		}
		if d := cmp.Diff(errMsg, noResourcesFound); d != "" {
			t.Fatalf("-got, +want: %v", d)
		}
	})

	t.Run("Install all components", func(t *testing.T) {
		argv := []string{"install", "all"}
		output, errMsg := ExecuteCommandOutput(TektonInstallCmd, argv)
		if errMsg != "" {
			t.Log(errMsg)
		}

		if d := cmp.Diff(output, allCompsInstall); d != "" {
			t.Fatalf("-got, +want: %v", d)
		}

		t.Log("Waiting for pods to be available in tekton-pipelines namespace")
		_, errMsg = WaitForAllPodStatus("Ready", "tekton-pipelines", "3m")
		if errMsg != "" {
			t.Log(errMsg)
		}
	})

	t.Run("Uninstall all components", func(t *testing.T) {
		argv := []string{"uninstall", "all", "-f"}
		output, errMsg := ExecuteCommandOutput(TektonInstallCmd, argv)
		if errMsg != "" {
			t.Log(errMsg)
		}

		if d := cmp.Diff(output, allCompsUninstall); d != "" {
			t.Fatalf("-got, +want: %v", d)
		}

		// Make sure no pods remain after uninstall for all components
		argv = []string{"get", "pods", "-n", "tekton-pipelines"}
		output, errMsg = ExecuteCommandOutput(KubectlCmd, argv)
		if errMsg == "" {
			t.Logf("Expected no pods to be found but pods were found\n%s", output)
		}
		if d := cmp.Diff(errMsg, noResourcesFound); d != "" {
			t.Fatalf("-got, +want: %v", d)
		}
	})

	t.Run("Install pipeline component with version", func(t *testing.T) {
		argv := []string{"install", "pipeline", "--pipeline-version", "0.16.0"}
		output, errMsg := ExecuteCommandOutput(TektonInstallCmd, argv)
		if errMsg != "" {
			t.Log(errMsg)
		}

		if d := cmp.Diff(output, installSuccess("pipeline")); d != "" {
			t.Fatalf("-got, +want: %v", d)
		}

		t.Log("Waiting for pods to be available in tekton-pipelines namespace")
		_, errMsg = WaitForAllPodStatus("Ready", "tekton-pipelines", "3m")
		if errMsg != "" {
			t.Log(errMsg)
		}

		argv = []string{"list"}
		output, errMsg = ExecuteCommandOutput(TektonInstallCmd, argv)
		if errMsg != "" {
			t.Log(errMsg)
		}

		assert.Assert(t, strings.Contains(output, "pipeline") && strings.Contains(output, "v0.16.0"))
	})

	t.Run("Install triggers component with version", func(t *testing.T) {
		argv := []string{"install", "triggers", "--triggers-version", "0.7.0"}
		output, errMsg := ExecuteCommandOutput(TektonInstallCmd, argv)
		if errMsg != "" {
			t.Log(errMsg)
		}

		if d := cmp.Diff(output, installSuccess("triggers")); d != "" {
			t.Fatalf("-got, +want: %v", d)
		}

		t.Log("Waiting for pods to be available in tekton-pipelines namespace")
		_, errMsg = WaitForAllPodStatus("Ready", "tekton-pipelines", "3m")
		if errMsg != "" {
			t.Log(errMsg)
		}

		argv = []string{"list"}
		output, errMsg = ExecuteCommandOutput(TektonInstallCmd, argv)
		if errMsg != "" {
			t.Log(errMsg)
		}

		assert.Assert(t, strings.Contains(output, "triggers") && strings.Contains(output, "v0.7.0"))
	})

	t.Run("Install dashboard component with version", func(t *testing.T) {
		argv := []string{"install", "dashboard", "--dashboard-version", "0.8.0"}
		output, errMsg := ExecuteCommandOutput(TektonInstallCmd, argv)
		if errMsg != "" {
			t.Log(errMsg)
		}

		if d := cmp.Diff(output, installSuccess("dashboard")); d != "" {
			t.Fatalf("-got, +want: %v", d)
		}

		t.Log("Waiting for pods to be available in tekton-pipelines namespace")
		_, errMsg = WaitForAllPodStatus("Ready", "tekton-pipelines", "3m")
		if errMsg != "" {
			t.Log(errMsg)
		}

		argv = []string{"list"}
		output, errMsg = ExecuteCommandOutput(TektonInstallCmd, argv)
		if errMsg != "" {
			t.Log(errMsg)
		}

		assert.Assert(t, strings.Contains(output, "dashboard") && strings.Contains(output, "v0.8.0"))
	})

	t.Run("Uninstall all installed components with specific versions", func(t *testing.T) {
		argv := []string{"uninstall", "all", "-f"}
		output, errMsg := ExecuteCommandOutput(TektonInstallCmd, argv)
		if errMsg != "" {
			t.Log(errMsg)
		}

		if d := cmp.Diff(output, allCompsUninstall); d != "" {
			t.Fatalf("-got, +want: %v", d)
		}

		// Make sure no pods remain after uninstall for all components
		argv = []string{"get", "pods", "-n", "tekton-pipelines"}
		output, errMsg = ExecuteCommandOutput(KubectlCmd, argv)
		if errMsg == "" {
			t.Logf("Expected no pods to be found but pods were found\n%s", output)
		}
		if d := cmp.Diff(errMsg, noResourcesFound); d != "" {
			t.Fatalf("-got, +want: %v", d)
		}
	})
}

func installSuccess(component string) string {
	return fmt.Sprintf("Component %s has been installed successfully\n", component)
}

func uninstallSuccess(component string) string {
	return fmt.Sprintf("Component %s has been uninstalled successfully\n", component)
}
