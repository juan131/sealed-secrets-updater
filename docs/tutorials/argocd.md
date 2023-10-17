# Using Sealed Secrets Updater with ArgoCD

> Note: This tutorial assumes that you have previously completed the [previous tutorial](ci.md).

Once we have a CI pipeline that updates the Sealed Secrets manifests when a change is made in the secrets inputs, we can use ArgoCD to deploy them in the cluster and completing our GitOps CI/CD flow.

## A full example using ArgoCD

### Scenario

This tutorial starts from the scenario we left after completing the [previous tutorial](ci.md).

### Installing ArgoCD in the cluster

If you already have ArgoCD installed in your cluster, you can skip this step. Otherwise, refer to [ArgoCD installation guide](https://argo-cd.readthedocs.io/en/stable/operator-manual/installation/).

### Configuring ArgoCD

Finally, we just need to add new directory-type applications to deploy `foo` and `bar` applications in the cluster. Please refer to [ArgoCD documentation](https://argo-cd.readthedocs.io/en/stable/user-guide/directory/) for more information about directory-type applications.

We must ensure the `spec.source.path` field should point to the `manifests` folder in the repository.

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: bar
spec:
  destination:
    namespace: default
    server: https://kubernetes.default.svc
  project: default
  source:
    path: manifests/bar
    repoURL: https://github.com/your-user/your-repo.git
    targetRevision: HEAD
---
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: foo
spec:
  destination:
    namespace: default
    server: https://kubernetes.default.svc
  project: default
  source:
    path: manifests/foo
    repoURL: https://github.com/your-user/your-repo.git
    targetRevision: HEAD
```

> Note: remember to update the `spec.source.repoURL` field with your repository URL.

With this configuration, ArgoCD will synchronize the applications `bar` and `foo` in the cluster whenever a new change in the `manifests/bar` or `manifests/foo` folders is detected.
