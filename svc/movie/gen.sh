#!/usr/local/bin bash
protoc --proto_path=./proto --go_out=plugins=grpc:./common/pb ./proto/*.proto
