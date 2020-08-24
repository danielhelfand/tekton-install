# tekton-install

`tekton-install` is a CLI for installing Tekton components (pipeline, triggers, and dashboard) on a Kubernetes cluster. This is still a WIP, but it currently supports installing individual components of Tekton. By default, it will install the latest version of each component, but it also adds flags to allow users to specify specific versions to install.

This will also eventually have an uninstall command to uninstall Tekton components from a Kubernetes cluster. 

### Examples

```
# Install latest version of Tekton pipeline component
tekton-install install pipeline

# Install specific version of Tekton pipeline component
tekton-install install pipeline --pipeline-version v0.15.0

# Install latest version of Tekton triggers component
tekton-install install triggers

# Install specific version of Tekton triggers component
tekton-install install triggers --triggers-version v0.6.0

# Install latest version of Tekton dashboard component
tekton-install install dashboard

# Install specific version of Tekton dashboard component
tekton-install install dashboard --dashboard-version v0.6.0

# Install all of latest components
tekton-install install all

# Install all components with specific versions
tekton-install install all --pipeline-version v0.15.0 --triggers-version v0.6.0 --dashboard-version v0.6.0
```

### Install

```
go get github.com/danielhelfand/tekton-install
```
