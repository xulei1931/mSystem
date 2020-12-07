#!/usr/local/bin bash
protoc --proto_path=./pb --go_out=plugins=grpc:./common/pb ./pb/*.proto
