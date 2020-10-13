// +build e2e

package e2e

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Version_Command(t *testing.T) {
	t.Run("Run version command", func(t *testing.T) {
		argv := []string{"version"}
		output, errMsg := ExecuteCommandOutput(TektonInstallCmd, argv, false)
		if errMsg != "" {
			t.Log(errMsg)
		}

		version := "v0.0.1\n"
		if d := cmp.Diff(output, version); d != "" {
			t.Fatalf("-got, +want: %v", d)
		}
	})
}
