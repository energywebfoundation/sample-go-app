# sample-go-app
Sample application written in Go. Here should be a description of what the purpose of the application is.

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
21118165
```

## Docker
Instructions how to run app with Docker.

```
docker run --rm -p 8000:8000 ghcr.io/energywebfoundation/sample-go-app:master
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