package cmd

import (
	"fmt"
	"os/exec"
)

func getComponentVersion(component string, all bool) (string, error) {
	var argv []string
	if component != dashboard {
		// Since deployment for pipeline is named tekton-pipelines-controller, reassign component name to pipelines
		if component == pipeline {
			component = "pipelines"
		}
		argv = []string{"get", "deployment/tekton-" + component + "-controller", "-o", "jsonpath={.metadata.labels['app\\.kubernetes\\.io/version']}", "-n", "tekton-pipelines"}
	} else {
		argv = []string{"get", "deployment/tekton-" + component, "-o", "jsonpath={.metadata.labels.version}", "-n", "tekton-pipelines"}
	}

	kubectlCmd := exec.Command("kubectl", argv...)
	version, err := kubectlCmd.Output()
	if err != nil {
		if all {
			return "", nil
		}

		return "", fmt.Errorf("failed to get version of component %s, check if it is installed", component)
	}

	return string(version), nil
}
