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

Currently **only local secrets are supported**, but we plan to add support for secrets managers in the future such as Vault, AWS Secrets Manager, etc.

> Note: It is highly recommended to encrypt your local secrets using [git-crypt](https://github.com/AGWA/git-crypt) or similar tools.

## Installation

TODO

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
          "path": "path/to/my-secret.json"
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

You can find some basic examples in the [examples](./examples) directory.

> Note: Refer to the [JSON Schema](./api/secrets.schema.json) for the full list of available options.

## Tutorials

Please refer to the [tutorials](./docs/tutorials) directory for some tutorials on how to use Sealed Secrets Updater with other tools such as ArgoCD, GitHub Actions, etc.
