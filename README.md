
# Crossplane Provider for Microsoft Azure DevOps

<div style="text-align: center;">

![CI](https://github.com/glalanne/provider-azuredevops/workflows/CI/badge.svg)
[![GitHub release](https://img.shields.io/github/release/glalanne/provider-azuredevops/all.svg)](https://github.com/glalanne/provider-azuredevops/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/glalanne/provider-azuredevops)](https://goreportcard.com/report/github.com/glalanne/provider-azuredevops)
[![Contributors](https://img.shields.io/github/contributors/glalanne/provider-azuredevops)](https://github.com/glalanne/provider-azuredevops/graphs/contributors)


</div>

`provider-azuredevops` is a community [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
[azuredevops](https://registry.terraform.io/providers/azuredevops/azuredevops/latest/docs).

Resources tested:
- Project
- Service connection (Github)
- Build defintion
- Dashboard
- Feed
- Repository
- Workitem
- Variable Group



## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/glalanne/provider-azuredevops):
```
crossplane xpkg install provider xpkg.crossplane.io/glalanne/provider-azuredevops:v2.0.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: upjet-provider-azuredevops
spec:
  package: pkg.crossplane.io/glalanne/provider-azuredevops:v2.0.0
EOF
```

You can see the API reference [here](https://doc.crds.dev/github.com/glalanne/provider-azuredevops).

## Exemples

A few exemples are provided [here](./examples/) to show case how to configure the provider, and how to use some resources


## Developing

Run code-generation pipeline:
```console
make generate
```

Run against a Kubernetes cluster:

```console
make run
```
Run unit tests:

```console
make test
```


Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

End to end tests:

```console
make e2e
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/glalanne/provider-azuredevops/issues).
