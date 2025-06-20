.PHONY: all

all:
	docker build . --target=runner -t coding-go/k8s/run-go:latest
	docker tag coding-go/k8s/run-go:latest las3:443/coding-go/k8s/run-go:latest
	docker push las3:443/coding-go/k8s/run-go:latest
