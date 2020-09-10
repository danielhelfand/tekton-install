package cmd

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_validateArgs_Invalid_Argument(t *testing.T) {
	expectedErr := "invalid argument provided to command: invalid"
	err := validateArgs([]string{"invalid", "pipeline"})
	if err == nil {
		t.Error("expecting error from invalid argument passed to validateArgsInstall but no error returned")
	}

	if d := cmp.Diff(err.Error(), expectedErr); d != "" {
		t.Fatalf("-got, +want: %v", d)
	}
}

func Test_validateArgs_Invalid_AllNotFirst(t *testing.T) {
	expectedErr := "all should be only argument provided when used"
	err := validateArgs([]string{"pipeline", "all"})
	if err == nil {
		t.Error("expecting error from all argument used with additional arguments but no error returned")
	}

	if d := cmp.Diff(err.Error(), expectedErr); d != "" {
		t.Fatalf("-got, +want: %v", d)
	}
}

func Test_validateArgs_Valid_AllValidArgsPassed(t *testing.T) {
	err := validateArgs([]string{"pipeline", "triggers", "dashboard"})
	if err != nil {
		t.Errorf("no error expected but error was returned: %v", err)
	}
}

func Test_validateArgs_Valid_SingleArgPassed(t *testing.T) {
	err := validateArgs([]string{"pipeline"})
	if err != nil {
		t.Errorf("no error expected but error was returned: %v", err)
	}
}
