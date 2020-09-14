package cmd

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_quotedList_All_Valid_Components(t *testing.T) {
	expected := `"triggers", "dashboard", "pipeline"`
	got := quotedList(components)

	if d := cmp.Diff(got, expected); d != "" {
		t.Fatalf("-got, +want: %v", d)
	}
}

func Test_quotedList_Some_Valid_Components(t *testing.T) {
	expected := `"triggers", "dashboard"`
	got := quotedList([]string{triggers, dashboard})

	if d := cmp.Diff(got, expected); d != "" {
		t.Fatalf("-got, +want: %v", d)
	}
}

func Test_quotedList_One_Valid_Components(t *testing.T) {
	expected := `"pipeline"`
	got := quotedList([]string{pipeline})

	if d := cmp.Diff(got, expected); d != "" {
		t.Fatalf("-got, +want: %v", d)
	}
}

func Test_confirmUninstall_Respond_y(t *testing.T) {
	var stdin bytes.Buffer
	stdin.Write([]byte("y\n"))

	err := confirmUninstall(components, &stdin)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if d := cmp.Diff(err, nil); d != "" {
		t.Fatalf("-got, +want: %v", d)
	}
}

func Test_confirmUninstall_Respond_y_After_Error(t *testing.T) {
	var stdin bytes.Buffer
	stdin.Write([]byte("Y\ny\n"))

	err := confirmUninstall(components, &stdin)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if d := cmp.Diff(err, nil); d != "" {
		t.Fatalf("-got, +want: %v", d)
	}
}

func Test_confirmUninstall_Respond_n(t *testing.T) {
	var stdin bytes.Buffer
	stdin.Write([]byte("n\n"))

	err := confirmUninstall(components, &stdin)
	if err == nil {
		t.Error("expected error but error was nil")
	}

	if d := cmp.Diff(err.Error(), errors.New("cancelling uninstall").Error()); d != "" {
		t.Fatalf("-got, +want: %v", d)
	}
}

func Test_confirmUninstall_Respond_n_After_Error(t *testing.T) {
	var stdin bytes.Buffer
	stdin.Write([]byte("N\nn\n"))

	err := confirmUninstall(components, &stdin)
	if err == nil {
		t.Error("expected error but error was nil")
	}

	if d := cmp.Diff(err.Error(), errors.New("cancelling uninstall").Error()); d != "" {
		t.Fatalf("-got, +want: %v", d)
	}
}

func Test_confirmUninstall_Respond_Invalid(t *testing.T) {
	expected := `Are you sure you want to uninstall "triggers", "dashboard", "pipeline" components (y/n): Please enter y or n: `
	var stdin bytes.Buffer
	stdin.Write([]byte("invalid\n"))

	original := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	os.Stdout = w

	err = confirmUninstall(components, &stdin)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		w.Close()
		os.Stdout = original
	}

	w.Close()
	out, err := ioutil.ReadAll(r)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		os.Stdout = original
	}
	os.Stdout = original

	if d := cmp.Diff(string(out), expected); d != "" {
		t.Fatalf("-got, +want: %v", d)
	}
}
