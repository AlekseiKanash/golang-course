#!/bin/bash
go get google.golang.org/grpc
go get google.golang.org/protobuf/reflect/protoreflect@latest
go mod download github.com/golang/protobuf
go get golang.org/x/net/context@latest
go mod download golang.org/x/text
./generate_proto.sh