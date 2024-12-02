NAME := roster
DESC := Hockey roster web app
VERSION := $(shell git describe --tags --always --dirty --match "[0-9]*.[0-9]*.[0-9]*")
BUILDVERSION := $(shell go version)
BUILDTIME := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
BUILDER := $(shell echo "`git config user.name` <`git config user.email`>")
LDFLAGS := -X 'main.version=$(VERSION)' \
           -X 'main.buildTime=$(BUILDTIME)' \
           -X 'main.builder=$(BUILDER)' \
           -X 'main.goversion=$(BUILDVERSION)' \
           -X 'main.name=$(NAME)'

build: staticcheck lint test clean target/local

init:
	git config core.hooksPath .githooks

staticcheck:
	staticcheck ./...

lint:
	golangci-lint -v run --fix ./...

test:
	go test ./...

target/local: modules
	mkdir -p target/bin && go build -ldflags "$(LDFLAGS)" -o target/bin/server ./cmd/server


modules:
	go mod tidy

clean-target:
	rm -rf target

clean: clean-target

.PHONY: clean lint modules build
