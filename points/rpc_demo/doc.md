# GRPC

Quick Start: https://grpc.io/docs/languages/go/quickstart

client: Go  
server: Go

## Require

### protoc

download: https://github.com/protocolbuffers/protobuf/releases, remember set PATH

### protoc-gen-go

install script: go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26

### protoc-gen-go-grpc

install script: go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

## Run Demo

1. write proto file(s)
2. generate grpc code
3. start server
4. start client

## FAQ

### windows not allow run powershell script

(admin start)Set-ExecutionPolicy RemoteSigned