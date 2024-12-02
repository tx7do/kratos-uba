GOPATH ?= $(shell go env GOPATH)
# GOVERSION is the current go version, e.g. go1.9.2
GOVERSION ?= $(shell go version | awk '{print $$3;}')

# Ensure GOPATH is set before running build process.
ifeq "$(GOPATH)" ""
  $(error Please set the environment variable GOPATH before running `make`)
endif
FAIL_ON_STDOUT := awk '{ print } END { if (NR > 0) { exit 1 } }'

GO_CMD			:= GO111MODULE=on go
GIT_CMD			:= git
DOCKER_CMD		:= docker

ARCH      := "`uname -s`"
LINUX     := "Linux"
MAC       := "Darwin"

DEFAULT_VERSION=0.0.1

ifeq ($(OS),Windows_NT)
    IS_WINDOWS:=TRUE
endif

ifneq (git,)
	GIT_EXIST:=TRUE
endif

ifneq ("$(wildcard .git)", "")
	HAS_DOTGIT:=TRUE
endif

ifeq ($(GIT_EXIST),TRUE)
ifeq ($(HAS_DOTGIT),TRUE)
	# CUR_TAG is the last git tag plus the delta from the current commit to the tag
	# e.g. v1.5.5-<nr of commits since>-g<current git sha>
	CUR_TAG ?= $(shell git describe --tags --first-parent)

	# LAST_TAG is the last git tag
    # e.g. v1.5.5
    LAST_TAG ?= $(shell git describe --match "v*" --abbrev=0 --tags --first-parent)

    # VERSION is the last git tag without the 'v'
    # e.g. 1.5.5
    VERSION ?= $(shell git describe --match "v*" --abbrev=0 --tags --first-parent | cut -c 2-)
endif
endif

CUR_TAG ?= $(DEFAULT_VERSION)
LAST_TAG ?= v$(DEFAULT_VERSION)
VERSION ?= $(DEFAULT_VERSION)

# GOFLAGS is the flags for the go compiler.
GOFLAGS ?= -ldflags "-X main.version=$(VERSION)"

APP_RELATIVE_PATH=$(shell a=`basename $$PWD` && cd .. && b=`basename $$PWD` && echo $$b/$$a)
APP_NAME=$(shell echo $(APP_RELATIVE_PATH) | sed -En "s/\//-/p")
APP_DOCKER_IMAGE=$(shell echo $(APP_NAME) |awk -F '@' '{print "kratos-uba/" $$0 ":0.1.0"}')


.PHONY: init dep vendor build clean docker conf ent wire api openapi run test cover vet lint app

# initialize develop environment
init:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	@go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	@go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	@go install github.com/envoyproxy/protoc-gen-validate@latest
	@go install github.com/bufbuild/buf/cmd/buf@latest
	@go install github.com/google/gnostic@latest
	@go install entgo.io/ent/cmd/ent@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/go-kratos/kratos/cmd/kratos/v2@latest

# download dependencies of module
dep:
	@go mod download

# create vendor
vendor:
	@go mod vendor

# build golang application
build:
ifeq ("$(wildcard ./bin/)","")
	mkdir bin
endif
	@go build -ldflags "-X main.Service.Version=$(APP_VERSION)" -o ./bin/ ./...

# clean build files
clean:
	@go clean
	$(if $(IS_WINDOWS), del "coverage.out", rm -f "coverage.out")

# build docker image
docker:
	@docker build -t $(APP_DOCKER_IMAGE) . \
				  -f ../../../.docker/Dockerfile \
				  --build-arg APP_RELATIVE_PATH=$(APP_RELATIVE_PATH) GRPC_PORT=9000 REST_PORT=8000

# generate config define code
conf:
	protoc --proto_path=./internal/conf/ \
	       --proto_path=../../../api/third_party \
	       --go_out=paths=source_relative:./internal/conf/ \
	       ./internal/conf/conf.proto

# generate ent code
ent:
ifneq ("$(wildcard ./internal/data/ent)","")
	@go run -mod=mod entgo.io/ent/cmd/ent generate \
				--feature privacy \
				--feature sql/modifier \
				--feature entql \
				--feature sql/upsert \
				./internal/data/ent/schema
endif

# generate wire code
wire:
	@go run -mod=mod github.com/google/wire/cmd/wire ./cmd/server

# generate protobuf api go code
api:
	@cd ../../../ && \
	buf generate

# generate OpenAPI v3 doc
openapi:
	@cd ../../../ && \
	buf generate --path api/admin/service/v1 --template api/admin/service/v1/buf.openapi.gen.yaml

# run application
run:
	@go run ./cmd/server -conf ./configs

# run tests
test:
	@go test ./...

# run coverage tests
cover:
	@go test -v ./... -coverprofile=coverage.out

# run static analysis
vet:
	@go vet

# run lint
lint:
	@golangci-lint run

# build service app
app: api wire conf ent build

# show help
help:
	@echo ""
	@echo "Usage:"
	@echo " make [target]"
	@echo ""
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
