# README

Build:

```console
$ make
```

Test `create-cm` and `create-dyn` with:

```console
$ kubectl apply -f po.yaml
pod/run-go created
$ kubectl exec -it run-go -- /bin/sh
```

In console of the pod:

```console
$ create-cm
$ create-dyn
```

Test file uploader:

```console
$ kubectl apply -f uploader.yaml
deployment.apps/uploader created
service/uploader created
ingress.networking.k8s.io/uploader-ingress created
```

Test with `curl`:

```console
curl -F "file=@filename.txt" http://<your-ingress-host>/upload
```
