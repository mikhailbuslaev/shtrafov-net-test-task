MAKEFILE_PATH := $(abspath $(dir $(abspath $(lastword $(MAKEFILE_LIST)))))

export GOBIN = $(MAKEFILE_PATH)/bin

BUILD_TARGET=$(MAKEFILE_PATH)/bin/sntt

# protoc
PROTOC = /usr/bin/protoc

GREEN_COLOR   = "\033[0;32m"
DEFAULT_COLOR = "\033[m"

.PHONY: test build docker

test:
	@echo -e $(GREEN_COLOR)[running tests]$(DEFAULT_COLOR)
	@go generate ./... && go test -v `go list ./... | grep -v integration`

build:
	@echo -e $(GREEN_COLOR)[building sntt to $(BUILD_TARGET)]$(DEFAULT_COLOR)
	@go build -o $(BUILD_TARGET)

docker:
	@docker build -f ./build/Dockerfile ./

protoc:
	@echo -e $(GREEN_COLOR)[protoc]$(DEFAULT_COLOR)
	@$(PROTOC) protoc $(MAKEFILE_PATH)/api/api.proto \
	--go_out=$(MAKEFILE_PATH)/pkg --go-grpc_out=$(MAKEFILE_PATH)/pkg \
	--go_opt=paths=source_absolute  \
	--go-grpc_opt=paths=source_absolute \