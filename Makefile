.PHONY: usage build test staticcheck get-linter lint docker-build integration-test run

OK_COLOR=\033[32;01m
NO_COLOR=\033[0m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

# Build Flags
BUILD_DATE = $(shell date -u --rfc-3339=seconds)
BUILD_HASH ?= $(shell git rev-parse --short HEAD)
APP_VERSION ?= undefined
BUILD_NUMBER ?= dev

NOW = $(shell date -u '+%Y%m%d%I%M%S')
PORT :=8080

GO := go
GO_LINTER := golint
DOCKER := docker
BUILDOS ?= $(shell go env GOHOSTOS)
BUILDARCH ?= amd64
GOFLAGS ?=
ECHOFLAGS ?=
ROOT_DIR := $(realpath .)

BIN := account

PKGS = $(shell $(GO) list ./...)

ENVFLAGS ?= CGO_ENABLED=0
BUILDENV ?= GOOS=$(BUILDOS) GOARCH=$(BUILDARCH)
BUILDFLAGS ?= -a -installsuffix cgo $(GOFLAGS) $(GO_LINKER_FLAGS)
EXTLDFLAGS ?= -extldflags "-lm -lstdc++ -static"

usage: Makefile
	@echo $(ECHOFLAGS) "to use make call:"
	@echo $(ECHOFLAGS) "    make <action>"
	@echo $(ECHOFLAGS) ""
	@echo $(ECHOFLAGS) "list of available actions:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'


## build: build all
##OLD build: lint staticcheck test
build:
	@echo $(ECHOFLAGS) "$(OK_COLOR)==> Building binary ($(BUILDOS)/$(BUILDARCH)/$(BIN))...$(NO_COLOR)"
	@echo $(ECHOFLAGS) $(ENVFLAGS) $(BUILDENV) $(GO) build $(BUILDFLAGS) -o bin/$(BUILDOS)_$(BUILDARCH)/$(BIN) ./cmd/webserver
	@$(ENVFLAGS) $(BUILDENV) $(GO) build $(BUILDFLAGS) -o bin/$(BUILDOS)_$(BUILDARCH)/$(BIN) ./cmd/webserver

## test: run unit tests
test:
	@echo $(ECHOFLAGS) "$(OK_COLOR)==> Running unit tests $(NO_COLOR)"
	@$(LOCAL_VARIABLES) $(ENVFLAGS) $(GO) test $(GOFLAGS) $(PKGS) --cover

## staticcheck: run staticcheck on packages
staticcheck:
	@echo $(ECHOFLAGS) "$(OK_COLOR)==> Running staticcheck...$(NO_COLOR)"
	@$(GO) get -v honnef.co/go/tools/cmd/staticcheck
	@$(ENVFLAGS) staticcheck $(PKGS)

## get-linter: install linter
get-linter:
	@echo $(ECHOFLAGS) "$(OK_COLOR)==> Getting linter...$(NO_COLOR)"
	@go get -u golang.org/x/lint/golint

## lint: lint package
lint: get-linter
	@echo $(ECHOFLAGS) "$(OK_COLOR)==> Running linter...$(NO_COLOR)"
	@$(GO_LINTER) -set_exit_status $(PKGS)


## docker-build: create the docker image
docker-build:
	@echo $(ECHOFLAGS) "$(OK_COLOR)==> build container image...$(NO_COLOR)"
	@ROOT_DIR=$(ROOT_DIR) $(DOCKER) build -t webserver .

## run: runs application	
run: docker-build
	@echo $(ECHOFLAGS) "$(OK_COLOR) ==> running webserver...$(NO_COLOR)"
	@ROOT_DIR=$(ROOT_DIR) $(DOCKER) run  -p $(PORT):8080 --rm -i webserver