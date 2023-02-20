FROM registry.access.redhat.com/ubi9/go-toolset:latest

COPY go.mod ./
COPY go.sum ./
COPY *.go ./
RUN mkdir -p controller
RUN mkdir -p model
RUN mkdir -p tools
COPY controller/*.go controller
COPY model/*.go model
COPY tools/*.go tools

RUN go mod download
RUN go mod tidy

EXPOSE 8081

RUN go run .
