# tekton-install

`tekton-install` is a CLI for installing Tekton components (pipeline, triggers, and dashboard) on a Kubernetes cluster. It currently supports installing and uninstalling individual components of Tekton. By default, it will install the latest version of each component, but it also adds flags to allow users to specify specific versions to install.

There is currently no validation of whether or not certain versions of components are compatible, so please make sure to verify component versions will work with one another before installing certain version combinations. 

### Prerequisites

* `kubectl` must be installed - As of now, `tekton-install` is simply a wrapper around `kubectl` to pass arguments to `kubectl apply -f` and `kubectl delete -f` commands.
* A Kubernetes cluster version 1.15 or higher for Tekton Pipelines v0.11.0 or higher, or a Kubernetes cluster version 1.11 or higher for Tekton releases before v0.11.0.

### Install tekton-install

```
go get -u github.com/danielhelfand/tekton-install
```

### Examples

#### Install

```
# Install latest version of Tekton pipeline component
tekton-install install pipeline

# Install specific version of Tekton pipeline component
tekton-install install pipeline --pipeline-version 0.15.0

# Install latest version of Tekton triggers component
tekton-install install triggers

# Install specific version of Tekton triggers component
tekton-install install triggers --triggers-version 0.6.0

# Install latest version of Tekton dashboard component
tekton-install install dashboard

# Install specific version of Tekton dashboard component
tekton-install install dashboard --dashboard-version 0.6.0

# Install all of latest components
tekton-install install all

# Install all components with specific versions
tekton-install install all --pipeline-version 0.15.0 --triggers-version 0.6.0 --dashboard-version 0.6.0
```

#### Uninstall

```
# Uninstall the Tekton pipeline component
# NOTE: Uninstalling pipeline component will
# also uninstall other installed Tekton components
tekton-install uninstall pipeline

# Uninstall the Tekton triggers component
tekton-install uninstall triggers

# Uninstall the Tekton dashboard component
tekton-install uninstall dashboard

# Uninstall all Tekton components
tekton-install uninstall all

# Uninstall Tekton components without being prompted for approval
tekton-install uninstall triggers dashboard pipeline -f

# Specify a Timeout of 1 minute 30 seconds for uninstalling each component.
# This example will produce a timeout error if the uninstall process lasts 
# longer than 1 minute 30 seconds for the triggers, dashboard, or pipeline component
tekton-install uninstall triggers dashboard pipeline --timeout 1m30s
```

#### List

```
# List available Tekton components on a Kubernetes cluster
tekton-install list
```

### Shell Autocompletion

Similar to [`kubectl`'s shell autocompletion](https://kubernetes.io/docs/tasks/tools/install-kubectl/#enabling-shell-autocompletion) 
approach, `tekton-install` features a `completion` command to enable autocompletion for bash and zsh.

#### Bash

For both Linux and macOS, installing [`bash-completion`](https://github.com/scop/bash-completion#installation) is required.

**Linux**

Use one of the approaches below to enable shell autocompletion for `tekton-install` on a Linux distribution.

To add the script to your `.bashrc` file:
```
echo 'source <(tekton-install completion bash)' >>~/.bashrc
```

Add the completion script to the `/etc/bash_completion.d` directory:
```
tekton-install completion bash >/etc/bash_completion.d/tekton-install
```

After completing the steps above, refresh your terminal to start using autocompletion. 

**Mac**

Add the following to your `~/.bash_profile` file:
```
export BASH_COMPLETION_COMPAT_DIR="/usr/local/etc/bash_completion.d"
[[ -r "/usr/local/etc/profile.d/bash_completion.sh" ]] && . "/usr/local/etc/profile.d/bash_completion.sh"
```

Use one of the approaches below to enable shell autocompletion for `tekton-install` on macOS.

To add the script to your `.bash_profile`:
```
echo 'source <(tekton-install completion bash)' >>~/.bash_profile
```

Add the completion script to the `/usr/local/etc/bash_completion.d` directory:
```
tekton-install completion bash >/usr/local/etc/bash_completion.d/tekton-install
```

After completing the steps above, refresh your terminal to start using autocompletion. 

#### Zsh

Add the following to your `~/.zshrc` file:

```
source <(tekton-install completion zsh)
```

After completing the steps above, refresh your terminal to start using autocompletion.
