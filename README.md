# tekton-install

`tekton-install` is a CLI for installing Tekton components (pipeline, triggers, and dashboard) on a Kubernetes cluster. This is still a WIP, but it currently supports installing and uninstalling individual components of Tekton. By default, it will install the latest version of each component, but it also adds flags to allow users to specify specific versions to install.

There is currently no validation of whether or not certain versions of components are compatible, so please make sure to verify component versions will work with one another.

### Prerequisites

* `kubectl` must be installed - As of now, `tekton-install` is simply a wrapper around `kubectl` to pass arguments to `kubectl apply -f` and `kubectl delete -f` commands.
* A Kubernetes cluster version 1.15 or higher for Tekton Pipelines v0.11.0 or higher, or a Kubernetes cluster version 1.11 or higher for Tekton releases before v0.11.0.

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
```

### Install tekton-install

```
go get -u github.com/danielhelfand/tekton-install
```
