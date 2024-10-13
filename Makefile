
SHELL := /bin/bash

GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build


SERVICE=telegram-clicker-game-be

init:
	$(GOCMD) mod init $(SERVICE)

tidy:
	$(GOCMD) mod tidy

run:
	$(GORUN) main.go

build:
	GIN_MODE=release $(GOBUILD) main.go -tags urfave_cli_no_docs	

migration-up:
	$(GORUN) main.go --migrate-up

migration-down:
	$(GORUN) main.go --migrate-down