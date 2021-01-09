.PHONY: build clean test package package-deb ui/build api ui-requirements serve cloc update-vendor
PKGS := $(shell go list ./... | grep -v /vendor | grep -v /migrations | grep -v /static | grep -v /ui)
VERSION := $(shell git describe --always)

GOARCH ?= amd64

ifeq ($(OS), Windows_NT)
	GOOS ?= windows
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S), Linux)
		GOOS ?= linux
	endif
	ifeq ($(UNAME_S), Darwin)
		GOOS ?= darwin
	endif
endif

build:
	@echo "Compiling source for $(GOOS) $(GOARCH)"
	@mkdir -p build
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-X main.version=$(VERSION)" -o build/movie$(BINEXT)

requirements:
	@echo "Installing development tools"
	go install golang.org/x/lint/golint
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go install github.com/golang/protobuf/protoc-gen-go
	go install github.com/elazarl/go-bindata-assetfs/...
	go install github.com/jteeuwen/go-bindata/...
	go install github.com/goreleaser/goreleaser
	go install github.com/goreleaser/nfpm

ui-requirements:
	@echo "Installing UI requirements"
	@cd ui && npm install 

serve: build
	@echo "Starting Lora Custom Server"
	./build/movie

