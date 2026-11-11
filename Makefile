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

BINARY := bin/server

# The sentinel lives in the *common* git dir's hooks directory, resolved
# through git: in a linked worktree `.git` is a file, so a literal
# `.git/hooks` path would not exist there. The common dir (and the
# core.hooksPath config `make init` writes) is shared by every worktree, so
# one init initializes them all. Outside a repository we fall back to the
# plain path, where the guard below fires with its own message instead.
GIT_COMMON_DIR := $(shell git rev-parse --git-common-dir 2>/dev/null)
HOOKS_SENTINEL := $(if $(GIT_COMMON_DIR),$(GIT_COMMON_DIR)/hooks,.git/hooks)/.githooks-installed

.PHONY: all init link checks staticcheck lint lint-check fmt test build verify modules clean

# Everything runs serially, even under `make -j`: golangci-lint --fix can
# rewrite files while the other checks read them, and ordering-sensitive
# sequences below must not be reordered.
.NOTPARALLEL:

# Plain `make` builds: checks, then a clean compile.
all: build

# Guard target — aborts with a helpful message if init was never run.
$(HOOKS_SENTINEL):
	@echo ""
	@echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
	@echo "⚠️  Run 'make init' first to configure git hooks."
	@echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
	@echo ""
	@exit 1

init:
	git config core.hooksPath .githooks
	@mkdir -p $(if $(GIT_COMMON_DIR),$(GIT_COMMON_DIR)/hooks,.git/hooks)
	@touch $(HOOKS_SENTINEL)
	@echo "✓ Git hooks configured. You're ready to develop."

# Create or refresh the ~/bin symlink to the built binary (the convention
# here: ~/bin holds one symlink per project, pointing at that project's
# built binary).
link: build
	@mkdir -p $(HOME)/bin
	@ln -sf $(CURDIR)/$(BINARY) $(HOME)/bin/$(NAME)
	@echo "✓ Linked ~/bin/$(NAME) -> $(CURDIR)/$(BINARY)"

# checks gate the build; the recipe then wipes the output tree and compiles
# from scratch, so every build is clean. The wipe and the compile are recipe
# lines rather than prerequisites so their order can never be scrambled —
# recipe lines always run in order, unlike prerequisites under `make -j`.
build: checks
	rm -rf bin
	mkdir -p bin && go build -ldflags "$(LDFLAGS)" -o $(BINARY) ./cmd/server

checks: $(HOOKS_SENTINEL) staticcheck lint-check test

staticcheck:
	staticcheck ./...

lint:
	golangci-lint -v run --fix ./...

# Verify lint/formatting without rewriting files (CI uses this).
lint-check:
	golangci-lint run ./...

# Apply formatter fixes (goimports) without running the full check suite.
fmt:
	golangci-lint fmt ./...

test:
	go test -race ./...

# Guard: refuse to trust a binary that does not match the current source.
verify: $(HOOKS_SENTINEL)
	bash scripts/verify-build.sh $(BINARY)

modules:
	go mod tidy

clean:
	rm -rf bin
