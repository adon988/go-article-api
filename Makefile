OUT        = ../pkg/grpc_client
SWAGGER    = ./api
INCLUDEDIR = $(CURDIR)/includes
PROTOS     = $(shell find . -type f -name '*.proto')

all: base gen

.PHONY: gen

gen: gen/proto

base: gen/base

gen/base:
	go mod tidy
	go get github.com/grpc-ecosystem/grpc-gateway/v2
	go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	buf format -w -d
	buf mod update
gen/proto:
	buf generate

