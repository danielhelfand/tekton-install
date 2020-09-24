// +build e2e

package e2e

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"gotest.tools/assert"
)

const (
	noComponentsMsg = "No components installed\n"
)

func Test_List_Command(t *testing.T) {

	t.Run("Run list command against empty cluster", func(t *testing.T) {
		argv := []string{"list"}
		output, errMsg := ExecuteCommandOutput(TektonInstallCmd, argv)
		if errMsg != "" {
			t.Log(errMsg)
		}

		if d := cmp.Diff(output, noComponentsMsg); d != "" {
			t.Fatalf("-got, +want: %v", d)
		}
	})

	t.Run("Run list command against cluster with everything installed", func(t *testing.T) {
		t.Log("Installing all Tekton components")
		argv := []string{"install", "pipeline", "--pipeline-version", "0.16.3", "triggers", "--triggers-version", "0.8.1", "dashboard", "--dashboard-version", "0.9.0"}
		ExecuteCommand(TektonInstallCmd, argv)

		t.Log("Waiting for pods to be available in tekton-pipelines namespace")
		_, errMsg := WaitForAllPodStatus("Ready", "tekton-pipelines", "3m")
		if errMsg != "" {
			t.Log(errMsg)
		}

		argv = []string{"list"}
		output, errMsg := ExecuteCommandOutput(TektonInstallCmd, argv)
		if errMsg != "" {
			t.Log(errMsg)
		}

		assert.Assert(t, strings.Contains(output, "dashboard") && strings.Contains(output, "v0.9.0"))
		assert.Assert(t, strings.Contains(output, "pipeline") && strings.Contains(output, "v0.16.3"))
		assert.Assert(t, strings.Contains(output, "triggers") && strings.Contains(output, "v0.8.1"))
	})

	t.Run("Run list command against cluster with only pipeline installed", func(t *testing.T) {
		t.Log("Uninstalling dashboard and triggers components")
		argv := []string{"uninstall", "dashboard", "triggers", "-f"}
		ExecuteCommand(TektonInstallCmd, argv)

		argv = []string{"list"}
		output, errMsg := ExecuteCommandOutput(TektonInstallCmd, argv)
		if errMsg != "" {
			t.Log(errMsg)
		}

		// Assert components do not show up after uninstall
		assert.Assert(t, !strings.Contains(output, "dashboard") && !strings.Contains(output, "v0.9.0"))
		assert.Assert(t, !strings.Contains(output, "triggers") && !strings.Contains(output, "v0.8.1"))

		// Assert pipeline component still is displayed
		assert.Assert(t, strings.Contains(output, "pipeline") && strings.Contains(output, "v0.16.3"))
	})

	t.Run("Run list command against empty cluster after everything uninstalled", func(t *testing.T) {
		t.Log("Uninstalling pipeline component")
		argv := []string{"uninstall", "all", "-f"}
		ExecuteCommand(TektonInstallCmd, argv)

		argv = []string{"list"}
		output, errMsg := ExecuteCommandOutput(TektonInstallCmd, argv)
		if errMsg != "" {
			t.Log(errMsg)
		}

		if d := cmp.Diff(output, noComponentsMsg); d != "" {
			t.Fatalf("-got, +want: %v", d)
		}
	})
}
