package e2e

import (
	"bytes"
	"log"
	"os"
	"os/exec"
)

const (
	TektonInstallCmd = "./../../tekton-install"
	KubectlCmd       = "kubectl"
)

func ExecuteCommand(command string, argv []string) {
	cmd := exec.Command(command, argv...)
	setCommandStdOutErr(cmd, nil, nil)
	runCommand(cmd)
}

func ExecuteCommandOutput(command string, argv []string, expectErr bool) (string, string) {
	cmd := exec.Command(command, argv...)
	var stdout, stderr bytes.Buffer
	setCommandStdOutErr(cmd, &stdout, &stderr)
	output, errMsg := runCommandOutput(cmd, &stdout, &stderr, expectErr)
	return output, errMsg
}

func setCommandStdOutErr(command *exec.Cmd, stdout, stderr *bytes.Buffer) {
	if stdout == nil && stderr == nil {
		command.Stderr = os.Stderr
	} else {
		command.Stdout = stdout
		command.Stderr = stderr
	}
}

func runCommand(command *exec.Cmd) {
	if err := command.Run(); err != nil {
		log.Fatalf("command %s failed", command.Args)
	}
}

func runCommandOutput(command *exec.Cmd, stdout, stderr *bytes.Buffer, expectErr bool) (string, string) {
	err := command.Run()
	if err != nil && !expectErr {
		log.Fatalf("Error occurred from command %s: %v", command.Args, err)
	}

	return stdout.String(), stderr.String()
}

func WaitFor(condition, namespace, resource, timeout string, all bool) (string, string) {
	argv := []string{"wait", "--for=" + condition, resource, "-n", namespace, "--timeout=" + timeout}
	if all {
		argv = append(argv, "--all")
	}
	return ExecuteCommandOutput(KubectlCmd, argv, false)
}
