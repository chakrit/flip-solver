#!/usr/bin/make

GO  := go
PKG := .
BIN := $(shell basename `pwd`)

ALL_PUZ := $(wildcard puzzles/*-*.table)
ALL_SOL := $(patsubst %.table,%.sol,$(ALL_PUZ))

.PHONY: default vet test deps lint install

default: all

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

# puzzles
%.sol: install
%.sol: %.table
	@echo
	@echo solving $@...
	@time $(BIN) $< | tee $@
	@echo
	@echo ---

all: $(ALL_SOL)

