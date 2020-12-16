#!/usr/local/bin bash
protoc --proto_path=./proto --go_out=plugins=grpc:./svc/common/pb ./proto/*.proto
