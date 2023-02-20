FROM registry.access.redhat.com/ubi9/go-toolset:latest

COPY go.mod ./
COPY go.sum ./
COPY *.go ./

RUN go mod download

EXPOSE 8081

RUN go run .
