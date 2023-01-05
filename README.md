# sample-go-app
Sample application written in Go

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
`./bin/sample-app`


### Format code
```
gofmt -s -w .
```

## Docker
Instructions how to run app with Docker.


Build docker image:
`docker build --tag sample-go-app -f ./docker/Dockerfile .`

Run local container:
`docker run --rm -p 8000:8000 sample-go-app`



## Kubernetes
Instructions how to run app on Kubernetes.