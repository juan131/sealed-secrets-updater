# Using Sealed Secrets Updater in your CI pipeline

> Note: This tutorial assumes that you are familiar with encrypted inputs using `git-crypt`. If you are not, please read [this tutorial](./git-crypt.md) first.

Sealed Secrets Updater can be used in your CI pipeline to update the Sealed Secrets manifests when a change is made in the secrets inputs. This way, you don't need to worry anymore about Sealed Secrets manifests and exclusively focus on the secrets inputs.

## A full example using GitHub Actions

### Prerequisites

- [Git](https://git-scm.com/) and a GitHub repository.
- [git-crypt](https://github.com/AGWA/git-crypt/blob/master/INSTALL.md)
- [GPG tools](https://gpgtools.org)
- K8s cluster.
- [Sealed Secrets controller](https://github.com/bitnami-labs/sealed-secrets#installation) installed in the cluster.

### Scenario

This tutorial starts from the scenario we left after completing the [previous tutorial](./git-crypt.md).

### Adding a GPG key to GitHub Encrypted Secrets

In order to run Sealed Secrets Updater in your CI pipeline, we need to previously decrypt the secrets inputs. For that, we need to export a GPG key and add it to GitHub Encrypted Secrets.

- First, export your GPG key using the command below and copy it to your clipboard:

```bash
git-crypt export-key ./tmp-key && cat ./tmp-key | base64 | pbcopy && rm ./tmp-key
```

- Then, follow the steps described in [this guide](https://docs.github.com/en/actions/security-guides/encrypted-secrets#creating-encrypted-secrets-for-a-repository) to create a new encrypted secret in your repository called `GIT_CRYPT_KEY` with the value of the key you just imported in your clipboard.

### Adding a GitHub workflow to update the Sealed Secrets manifests

Now, we are ready to add a GitHub workflow to update the Sealed Secrets manifests when a change is made in the secrets inputs. To do so, we just need create a new GitHub workflow file (e.g `.github/workflows/update-sealed-secrets.yml`) in the repository with the content below:

```yaml
name: Update Sealed Secrets manifests

on:
  pull_request:
    branches:
      - main
    paths:
      - 'secrets/**'

jobs:
  update-sealed-secrets:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Setup git-crypt
        run: sudo apt-get install -y git-crypt

      - name: Setup Sealed Secrets Updater
        run: |
          latest_release_name="$(curl -sH "Accept: application/vnd.github.v3+json" https://api.github.com/repos/juan131/sealed-secrets-updater/releases | jq -r "map(select(.prerelease == false)) | .[0].name")"
          latest_version="${latest_release_name#"sealed-secrets-updater-v"}"
          curl -sL "https://github.com/juan131/sealed-secrets-updater/releases/download/v${latest_version}/sealed-secrets-updater-${latest_version}-linux-amd64.tar.gz" | tar -xz sealed-secrets-updater
          mv sealed-secrets-updater /usr/local/bin/sealed-secrets-updater
          chmod +x /usr/local/bin/sealed-secrets-updater
    
      - name: Unlock secrets inputs
        run: |
          echo ${{ secrets.GIT_CRYPT_KEY }} | base64 -d > ./git-crypt-key
          git-crypt unlock ./git-crypt-key
          rm ./git-crypt-key

      - name: Authenticate with GCP
        uses: google-github-actions/auth@v1
        with:
          credentials_json: '${{ secrets.gcp_credentials }}'

      - name: Get GKE credentials
        uses: google-github-actions/get-gke-credentials@v1
        with:
          cluster_name: my-cluster
          location: us-central1-a

      - name: Update the Sealed Secrets manifests
        run: |
          sealed-secrets-updater update --config sealed-secrets-updater.json

      - name: Commit
        uses: EndBug/add-and-commit@v7.2.0
        with:
          add: 'manifests'
          message: 'chore: update Sealed Secrets manifests'
```

> Note: the above workflow assumes you're using GKE as your K8s cluster and authenticate to it as described in [this guide](https://github.com/google-github-actions/get-gke-credentials#authenticating-via-service-account-key-json). Please adapt it to your needs if you're using a different cloud provider or authentication method.

Now, every time a PR is created attempting to change the secrets inputs, a GitHub workflow will be launched to update the Sealed Secrets manifests and commit the changes.

### Next steps

- [Using Sealed Secrets Updater with ArgoCD](./argocd.md)
