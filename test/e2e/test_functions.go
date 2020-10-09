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

func ExecuteCommandOutput(command string, argv []string) (string, string) {
	cmd := exec.Command(command, argv...)
	var stdout, stderr bytes.Buffer
	setCommandStdOutErr(cmd, &stdout, &stderr)
	output, errMsg := runCommandOutput(cmd, &stdout, &stderr)
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
		log.Printf("command %s failed", command.Args)
	}
}

func runCommandOutput(command *exec.Cmd, stdout, stderr *bytes.Buffer) (string, string) {
	err := command.Run()
	if err != nil {
		log.Printf("command %s failed", command.Args)
	}

	return stdout.String(), stderr.String()
}

func WaitForAllPodStatus(condition, namespace, timeout string) (string, string) {
	argv := []string{"wait", "--for=condition=" + condition, "pod", "-n", namespace, "--timeout=" + timeout, "--all"}
	return ExecuteCommandOutput(KubectlCmd, argv)
}
