package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Completion(t *testing.T) {
	for _, shell := range []string{"bash", "zsh"} {
		stdout := bytes.NewBufferString("")
		stderr := bytes.NewBufferString("")
		rootCmd.SetOut(stdout)
		rootCmd.SetErr(stderr)
		rootCmd.SetArgs([]string{"completion", shell})
		err := rootCmd.Execute()
		require.NoError(t, err, "Expected error to be nil but was not:\n"+stderr.String())
		require.Equal(t, "", stderr.String(), shell+" arg experienced error with tekton-install completion:\n"+stderr.String())
		require.NotEmpty(t, stdout.String(), shell+" arg reported nothing to stdout with tekton-install completion")
	}
}
