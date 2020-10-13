package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_version(t *testing.T) {
	stdout := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetArgs([]string{"version"})
	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("Error from executing version command: %v", err)
	}

	out, err := ioutil.ReadAll(stdout)
	if err != nil {
		t.Fatalf("Error from reading version command stdout: %v", err)
	}

	expected := version + "\n"
	if d := cmp.Diff(string(out), expected); d != "" {
		t.Fatalf("-got, +want: %v", d)
	}
}
