MAKEFILE_PATH := $(abspath $(dir $(abspath $(lastword $(MAKEFILE_LIST)))))

BUILD_TARGET=$(MAKEFILE_PATH)/bin/sntt

GREEN_COLOR   = "\033[0;32m"
DEFAULT_COLOR = "\033[m"


.PHONY: env test build docker protoc

test:
	@echo -e $(GREEN_COLOR)[running tests..]$(DEFAULT_COLOR)
	@go generate ./... && go test -v `go list ./... | grep -v integration`
	@echo -e $(GREEN_COLOR)[tests run done]$(DEFAULT_COLOR)

build:
	@echo -e $(GREEN_COLOR)[go build running...]$(DEFAULT_COLOR)
	@echo -e $(GREEN_COLOR)[building sntt to $(BUILD_TARGET)]$(DEFAULT_COLOR)
	@go build -o $(BUILD_TARGET)
	@echo -e $(GREEN_COLOR)[go build done]$(DEFAULT_COLOR)

docker:
	@echo -e $(GREEN_COLOR)[docker build running...]$(DEFAULT_COLOR)
	@docker build -f ./build/Dockerfile ./
	@echo -e $(GREEN_COLOR)[docker build done]$(DEFAULT_COLOR)

protoc:
	@export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN
	@echo -e $(GREEN_COLOR)[protoc generation running...]$(DEFAULT_COLOR)
	@protoc ./api/api.proto \
	--go_out=./pkg \
	--go-grpc_out=./pkg \
	--grpc-gateway_out=./pkg \
	--go_opt=paths=source_relative  \
	--go-grpc_opt=paths=source_relative \
	--grpc-gateway_opt=paths=source_relative \
	--grpc-gateway_opt generate_unbound_methods=true \
	--openapiv2_out ./swaggerui
	@echo -e $(GREEN_COLOR)[protoc generation done]$(DEFAULT_COLOR)