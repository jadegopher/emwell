package emwell

import (
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
)

//go:generate  protoc -I . --proto_path=./internal/api/.protobuf --grpc-gateway_out=logtostderr=true:./internal/api/.protobuf --swagger_out=allow_merge=true,merge_file_name=api:./internal/api/.protobuf --go_out=plugins=grpc:./internal/api/.protobuf ./internal/api/.protobuf/api.proto
