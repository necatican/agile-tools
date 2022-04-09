#! /usr/bin/env bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=source_relative \
    pkg/apis/retro/v1alpha1/board/board.proto

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=source_relative \
    pkg/apis/retro/v1alpha1/task/task.proto
