# Provider Vault

`provider-vault` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Vault API.

## Prerequisites

This provider interacts with HashiCorp Vault. 
To test it you will need Vault installed. 
Vault has many options and various ways for how
it can be installed. We will use a very simple
installation approach to show how to test the
provider-vault manually per below. To test this
provider-vault using automated tests, please
run `make e2e`. This will create, initialize
and unseal a copy of Vault against which the
automated tests are running. The list of tests
is specified in the `Makefile` variable
`UPTEST_EXAMPLE_LIST`. The tests it runs are
specified in the `examples` directory.

Create a vault namespace.
```
kubectl create namespace vault
```

Add hashicorp to your helm repo.
```
helm repo add hashicorp https://helm.releases.hashicorp.com
```

Install vault. Note that the vault-0 pod
will not be ready until vault is unsealed.
This is expected behavior.
```
helm install vault hashicorp/vault -n vault
```

Initialize and unseal vault.
```
kubectl exec -it vault-0 -n vault -- sh
```

Inside of the vault pod, initialize it.
```
vault operator init
```
Make note of the root key and the unseal keys.

Unseal vault 3 times with a different key.
```
vault operator unseal
```

Verify that vault is unsealed.
```
vault status
```

From the commandline, forward the vault pod port.
```
kubectl port-forward vault-0 -n vault 8200:8200
```

Update the examples/providerconfig/secret.yaml.tmpl
with your root token or an access token that was
created in Vault for your provider. Note that 
this token should not be accessible 
by cluster operators, only by Vault admins.

Apply provider-vault package/crds.
```
kubectl apply -f package/crds
```

Apply the secret.
```
kubectl apply -f providerconfig/secret.yaml.tmpl
```

Apply the provider config.
```
kubectl apply -f providerconfig/providerconfig.yaml
```

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/upbound/provider-vault):
```
up ctp provider install upbound/provider-vault:v0.1.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-vault
spec:
  package: xpkg.upbound.io/upbound/provider-vault:v0.1.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/upbound/provider-vault).

## Developing

Initialize the repository with
```bash
make submodules
```

Run code-generation pipeline:
```bash
make generate
```

Run against a Kubernetes cluster:

```bash
make run
```

Build, push, and install:

```bash
make all
```

Build binary, image and Crossplane package (xpkg):

```bash
make build
# to build all architectures / platforms, use:
make -j2 build.all
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/upbound/provider-vault/issues).
