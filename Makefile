.PHONY: build clean test package package-deb ui/build api ui-requirements serve cloc update-vendor internal/statics internal/migrations
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

api:
	@echo "Generating API code from .proto files"
	go generate pb/api.go
