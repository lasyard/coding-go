# README

Build and publish the image:

```sh
make docker-build
make push-to-registry
```

Be sure to change the visibility to `Public` on GitHub in "Package settings", or credentials will be needed.

Create service account if not exists:

```console
$ kubectl apply -f internal-create-sa.yaml
serviceaccount/internal-create created
role.rbac.authorization.k8s.io/configmap-creator created
rolebinding.rbac.authorization.k8s.io/configmap-creator-binding created
```

Run the pod:

```console
$ kubectl apply -f internal-create-po.yaml
pod/internal-create created
```

The command need to be run in the pod mannually:

```console
$ kubectl exec internal-create -- /usr/local/bin/create-cm
Created ConfigMap `test` in `default` namespace.
```

```console
$ kubectl exec internal-create -- /usr/local/bin/create-cm-dyn
Created ConfigMap `test-dyn` in `default` namespace.
```
