# tekton-install tkn plugins

This folder contains scripts that can be used as plugins with the [Tekton CLI (`tkn`)](https://github.com/tektoncd/cli). Version v0.15.0 of `tkn` features a plugin system that allows users to call scripts or other binaries directly via `tkn`. 

To use `tekton install` via `tkn` instead of directly calling the `tekton-install` binary, add the scripts in this folder to the following path on your local machine: `~/.config/tkn/plugins/`.

After adding the scripts to the path, `tekton-install install` can be run through `tkn` using the following example commands: 

```
# Install latest version of Tekton pipeline component
tkn install pipeline

# Install specific version of Tekton pipeline component
tkn install pipeline --pipeline-version 0.15.0

# Install latest version of Tekton triggers component
tkn install triggers

# Install specific version of Tekton triggers component
tkn install triggers --triggers-version 0.6.0

# Install latest version of Tekton dashboard component
tkn install dashboard

# Install specific version of Tekton dashboard component
tkn install dashboard --dashboard-version 0.6.0

# Install all of latest components
tkn install all

# Install all components with specific versions
tkn install all --pipeline-version 0.15.0 --triggers-version 0.6.0 --dashboard-version 0.6.0
```

The same approach with `tekton-install` as above applies for `tekton-uninstall` as shown below:

```
# Uninstall the Tekton pipeline component
# NOTE: Uninstalling pipeline component will
# also uninstall other installed Tekton components
tkn uninstall pipeline

# Uninstall the Tekton triggers component
tkn uninstall triggers

# Uninstall the Tekton dashboard component
tkn uninstall dashboard

# Uninstall all Tekton components
tkn uninstall all

# Uninstall Tekton components without being prompted for approval
tkn uninstall triggers dashboard pipeline -f

# Specify a Timeout of 1 minute 30 seconds for uninstalling each component.
# This example will produce a timeout error if the uninstall process lasts 
# longer than 1 minute 30 seconds for the triggers, dashboard, or pipeline component
tkn uninstall triggers dashboard pipeline --timeout 1m30s
```