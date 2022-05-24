FROM golang:1-alpine
ENV GOBIN /go/bin

# gcc is required to support cgo;
# add git for go dependencies
RUN apk add --no-cache git gcc musl-dev make zip protobuf-dev=3.18.1-r1

# Install grpc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Install grpcurl
RUN go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

ENV PATH=$PATH:$GOPATH/bin:/opt/protoc/bin

WORKDIR root/app/
