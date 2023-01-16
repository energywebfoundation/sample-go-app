# sample-go-app
Sample application written in Go. Here should be a description of what the purpose of the application is.


Checklist:
- [] Inform DevOps Team
- [] Clear README file describing app purpose and instructions how to run
- [] Defined branching and release strategy
- [] Branch protection rules
- [] CI/test pipelines set up, make sure to use Snyk for security, code quality and Docker IMAGE (not Dockerfile!) scanning
- [] Set up CD pipeline for creating GH releases (tags), changelog and building Docker image
- [] Docker image published with GH packages
- [] Deployment guide instructions (preferably k8s with helm chart, generic microservice should be just enough, custom (public) helm otherwise)
- [] Add license, code owners, issue template, contribution guide etc

## Run locally
Instructions how to run app locally.


```
go run cmd/sample-app/main.go
```

### Build application
```
go build -o ./bin/sample-app cmd/sample-app/main.go
```

Run:
```
./bin/sample-app
```


### Format code
```
gofmt -s -w .
```


## How to use

Hello world:
```
% curl localhost:8000
Hello, world!
```


Get latest block of Volta network:
```
% curl localhost:8000/block
Latest block is: 21133370
```

## Docker
Instructions how to run app with Docker.

```
docker run --rm -p 8000:8000 ghcr.io/energywebfoundation/sample-go-app:latest
```


Build docker image:
```
docker build --tag sample-go-app -f ./docker/Dockerfile .
```

Run local container:
```
docker run --rm -p 8000:8000 sample-go-app
```



## Kubernetes
Instructions how to run app on Kubernetes.

```
helm install sample-app oci://ghcr.io/energywebfoundation/sample-go-app-helm
```

```
export POD_NAME=$(kubectl get pods --namespace default -l "app.kubernetes.io/name=sample-go-app,app.kubernetes.io/instance=sample-app" -o jsonpath="{.items[0].metadata.name}")
export CONTAINER_PORT=$(kubectl get pod --namespace default $POD_NAME -o jsonpath="{.spec.containers[0].ports[0].containerPort}")
kubectl --namespace default port-forward $POD_NAME $CONTAINER_PORT:$CONTAINER_PORT
```

```
% curl localhost:8000/block
Latest block is: 21133370
helm un sample-app
```