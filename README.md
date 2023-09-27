# Sealed Secrets Updater

**Problem:** "I follow GitOps using Sealed Secrets, but I need to manually recreate my manifests whenever my secrets need to be updated."

**Solution:** Use this tool to automatically track changes in your secrets manager and update your Sealed Secrets manifests.

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [Supported Secrets Managers](#supported-secrets-managers)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Tutorials](#tutorials)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Supported Secrets Managers

Currently **only input secrets files are supported**, but we plan to add support for secrets managers in the future such as Vault, AWS Secrets Manager, etc.

> Note: It is highly recommended to encrypt your input secrets files using [git-crypt](https://github.com/AGWA/git-crypt) or similar tools.

## Installation

You can download the corresponding binary for every supported version from [releases section](https://github.com/juan131/sealed-secrets-updater/releases). Alternatively, you can use the following commands to install the latest version (assuming linux/amd64):

```bash
latest_release_name="$(curl -sH "Accept: application/vnd.github.v3+json" https://api.github.com/repos/juan131/sealed-secrets-updater/releases | jq -r "map(select(.prerelease == false)) | .[0].name")"
latest_version="${latest_release_name#"sealed-secrets-updater-v"}"
curl -sL "https://github.comjuan131/sealed-secrets-updater/releases/download/v${latest_version}/sealed-secrets-updater-${latest_version}-linux-amd64.tar.gz" | tar -xz sealed-secrets-updater
mv sealed-secrets-updater /usr/local/bin/sealed-secrets-updater
chmod +x /usr/local/bin/sealed-secrets-updater
```

## Usage

Basic usage:

```bash
sealed-secrets-updater update --config config.json
```

Run the command below to see the rest available commands:

```bash
sealed-secrets-updater help
```

## Configuration

Sealed Secrets Updater uses a configuration file (JSON format) to determine how to update your manifests such as the ones below:

```json
{
  "kubesealConfig": {
    "controllerNamespace": "kube-system",
    "controllerName": "sealed-secrets-controller"
  },
  "secrets": [
    {
      "name": "my-secret",
      "namespace": "default",
      "input": {
        "type": "file",
        "config": {
          "path": "path/to/my-secret-inputs.json"
        }
      },
      "output": {
        "type": "file",
        "config": {
          "path": "path/to/my-sealed-secret.json"
        }
      }
    }
  ]
}
```

You can find some basic examples in the [examples](./examples) directory to learn how to configure Sealed Secrets Updater to update your manifests using different output types. Please note only two output types are supported at the moment:

- `apply`: Directly apply the new Sealed Secrets to your cluster.
- `file`: Save the new Sealed Secrets to a file.

> Note: Refer to the [JSON Schema](./api/secrets.schema.json) for the full list of available options.

## Tutorials

Please refer to the [tutorials](./docs/tutorials/index.md) directory for some tutorials on how to use Sealed Secrets Updater with other tools such as ArgoCD, GitHub Actions, etc.
