# README

Build and publish the image:

```sh
make docker-build
make push-to-registry
```

Be sure to change the visibility to `Public` on GitHub in "Package settings", or credentials will be needed.

Deploy to the cluster:

```console
$ kubectl apply -f upload-svc.yaml 
deployment.apps/upload-svc created
service/upload-svc created
ingress.networking.k8s.io/upload-svc-ingress created
```

Test with `curl`:

```console
$ curl -F "file=@README.md" http://<your-ingress-host>/upload
File uploaded successfully: README.md
```

Check the file is uploaded:

```console
$ kubectl exec "$(kubectl get po -lapp=upload-svc -o jsonpath='{.items[0].metadata.name}')" -- ls /var/uploads
README.md
```
