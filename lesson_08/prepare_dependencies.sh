#!/bin/bash
go get google.golang.org/grpc
go get google.golang.org/protobuf/reflect/protoreflect@latest
go mod download github.com/golang/protobuf
go get golang.org/x/net/context@latest
go mod download golang.org/x/text
go get -u github.com/golang/protobuf/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest