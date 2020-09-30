# tekton-install Development Documentation

This documentation covers how to build the `tekton-install` CLI, run unit tests and end to end tests, and about 
the continuous integration process. 

### Prerequisites

* Go version 1.14
* `kubectl` must be installed
* A Kubernetes cluster version 1.15 or higher for Tekton Pipelines v0.11.0 or higher, or a Kubernetes cluster version 1.11 or higher for Tekton releases before v0.11.0

### Build

To begin, fork this [project's GitHub repository](https://github.com/danielhelfand/tekton-install) and clone it locally. 

Building `tekton-install` can be done with the following command at the root level of this repository:

```
go build
```

After running the command, there should be a `tekton-install` binary created to run locally, which can be done as shown below:

```
./tekton-install
```

If you make changes, simply rebuild the binary with `go build`, and the `tekton-install` binary should have your changes.

### Unit Tests

Running all unit tests for the project can be done with the following command:

```
go test ./...
```

### End to End Tests

Running end to end tests locally can be done against a local Kubernetes cluster like [`Minikube`](https://kubernetes.io/docs/tasks/tools/install-minikube/) or [`KinD`](https://kind.sigs.k8s.io/) (Kubernetes in Docker) or against a remote cluster. Just make sure whatever kubeconfig you are using locally is set up to work with the cluster you would like to run tests against. 

Once a cluster is available, run the following command to execute and see results of end to end tests:

```
go test -v -count=1 -tags=e2e -failfast -timeout=10m ./test/e2e
```

Since the `-failfast` flag is specified, there could be scenarios with test failures that result in needing to clean up the cluster to run tests 
again. You should be able to run `tekton-install uninstall all` to restore the cluster to a clean state.

### Continuous Integration

The CI process for `tekton-install` uses [GitHub Actions](https://docs.github.com/en/free-pro-team@latest/actions) and can be found under the `.github/workflows` folder. When a pull request is opened or pushes are made after the pull request is opened to the tekton-install GitHub repository, the CI process will execute and report build and test successes and failures in the opened pull request. Logs of CI builds can also be viewed in the opened pull request.

Currently, the CI process builds the `tekton-install` binary, runs unit tests for the project, runs a quick round of end to end tests, and then executes 
the suite of end to end tests written in Go. There is also a linting process that uses [`golangci-lint`](https://golangci-lint.run/). Linters used by the project are defined in `.golangci.yml`. 
