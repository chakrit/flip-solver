#!/usr/bin/make

GO  := go
PKG := .
BIN := $(shell basename `pwd`)

.PHONY: default all vet test deps lint

default: run

all: install
install: deps
	$(GO) install $(PKG)
build: deps
	$(GO) build $(PKG)
lint: vet
vet: deps
	$(GO) vet $(PKG)
fmt:
	$(GO) fmt $(PKG)
test: deps
	$(GO) test $(PKG)
clean:
	$(GO) clean $(PKG)
deps:
	$(GO) get -d $(PKG)
install:
run: install
	$(BIN)

