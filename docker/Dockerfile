##
## Build
##

FROM golang:1.19-buster AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

ADD cmd cmd
ADD internal internal

RUN go build -o /sample-app cmd/sample-app/main.go

##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /sample-app /sample-app

EXPOSE 8000

USER nonroot:nonroot

ENTRYPOINT ["/sample-app"]