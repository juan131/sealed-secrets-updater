# Using file inputs encrypted with git-crypt

Sealed Secrets Updater can be configured to retrieve secrets inputs from files. It is very convenient to use this approach when these inputs are committed in the Git repository. This way, everything (including secrets) is stored in the repository and it's feasible to adopt a GitOps approach.

However, it makes little sense to use Sealed Secrets to encrypt the K8s manifests to deploy on a cluster if the secrets inputs used to create them are also committed in the repository in plaintext. This is where [git-crypt](https://github.com/AGWA/git-crypt) comes in handy.

Using `git-crypt` we can easily encrypt the secrets inputs using the GPG keys of the developers we trust and commit them in the repository in a safe way. Then, we can decrypt them locally (or in a CI/CD pipeline) and use Sealed Secrets Updater to generate the Sealed Secrets manifests.

## A full example

### Prerequisites

- [Git](https://git-scm.com/).
- [git-crypt](https://github.com/AGWA/git-crypt/blob/master/INSTALL.md)
- [GPG tools](https://gpgtools.org)
- K8s cluster and `kubectl`  configured to access it.
- [Sealed Secrets controller](https://github.com/bitnami-labs/sealed-secrets#installation) installed in the cluster.

### Scenario

Let's assume we have the folder structure below in our repository:

```console
.
├── manifests
│   ├── bar
│   │   ├── deployment.yaml
│   │   └── sealed-secret.yaml
│   └── foo
│       ├── deployment.yaml
│       └── sealed-secret.yaml
├── secrets
│   ├── bar.json
│   └── foo.json
└── sealed-secrets-updater.json
```

- The `manifests` folder contains the K8s manifests we use to deploy applications `bar` and `foo` in the cluster.
- The `secrets` folder contains the secrets data inputs (key-value) in JSON format.
- The `sealed-secrets-updater.json` file contains the configuration for Sealed Secrets Updater and its content is the one below:

```json
{
  "secrets": [{
    "name": "bar",
    "namespace": "default",
    "input": {
      "type": "file",
      "config": {
        "path": "secrets/bar.json"
      }
    },
    "output": {
      "type": "file",
      "config": {
        "path": "manifests/bar/sealed-secret.yaml"
      }
    }
  }, {
    "name": "foo",
    "namespace": "default",
    "input": {
      "type": "file",
      "config": {
        "path": "secrets/foo.json"
      }
    },
    "output": {
      "type": "file",
      "config": {
        "path": "manifests/foo/sealed-secret.yaml"
      }
    }
  }]
}
```

### Encrypting secrets inputs with git-crypt

We can use `git-crypt` to encrypt the secrets inputs in the `secrets` folder:

- First, we need to initialize `git-crypt` in the repository.

```bash
git-crypt init
```

- Then, we create a `.gitattributes` file in the root of the repository with the following content:

```bash
secrets/** filter=git-crypt diff=git-crypt
```

- After that, we have to import the GPG keys of the developers we trust as it's explained on [this guide](https://gpgtools.tenderapp.com/kb/gpg-keychain-faq/how-to-find-public-keys-of-your-friends-and-import-them).
- Finally, we add the GPG keys of the developers we trust:

```bash
git-crypt add-gpg-user --trusted DEVELOPER@MAIL.com 
```

From this moment on, every developer can encrypt/decrypt the secrets inputs running the commands below:

```bash
# Encrypt secrets inputs
git-crypt lock
# Decrypt secrets inputs
git-crypt unlock
```

### Generating Sealed Secrets manifests

Once we have the secrets inputs encrypted with `git-crypt`, developers can use Sealed Secrets Updater to update the Sealed Secrets manifests in a secure way:

```bash
# Decrypt secrets inputs
git-crypt unlock
# Generate Sealed Secrets manifests
sealed-secrets-updater update --config sealed-secrets-updater.json
# Encrypt secrets inputs
git-crypt lock
# Commit changes
git add . && git commit -m "Update Sealed Secrets manifests"
```

Both the secrets inputs and the Sealed Secrets manifests are now encrypted and committed in the repository and we can now use a GitOps approach to deploy them in the cluster.

### Next steps

- [Using Sealed Secrets Updater in a CI pipeline](ci.md)
