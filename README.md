# tekton-install

`tekton-install` is a CLI for installing Tekton components (pipeline, triggers, and dashboard) on a Kubernetes cluster. This is still a WIP, but it currently supports installing individual components of Tekton. By default it will install the latest version of each component, but it also adds flags to allow users to specify specific versions to install.

This will also eventually have an uninstall command to uninstall Tekton components from a Kubernetes cluster. 

### Install

```
go get github.com/danielhelfand/tekton-install
```
