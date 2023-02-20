FROM registry.access.redhat.com/ubi9/go-toolset:latest

COPY go.mod ./
COPY go.sum ./
COPY *.go ./

RUN go mod download
RUN go mod tidy

EXPOSE 8081

RUN go run .
