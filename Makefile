#!/bin/make
GOPATH:=$(shell go env GOPATH)

.PHONY: test

all:
	$(GOPATH)/bin/goimports -w -l .
	go build -v

test:
	go test -v
