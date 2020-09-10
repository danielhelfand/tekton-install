package cmd

import "fmt"

func validateArgs(args []string) error {
	validComponents := make(map[string]bool)
	validComponents[pipeline] = true
	validComponents[triggers] = true
	validComponents[dashboard] = true
	validComponents["all"] = true

	for _, arg := range args {
		if !validComponents[arg] {
			return fmt.Errorf("invalid argument provided to command: %s", arg)
		}

		if arg == "all" && args[0] != "all" {
			return fmt.Errorf("all should be only argument provided when used")
		}
	}
	return nil
}
