SHELL := /bin/bash
export GOPATH:=$(HOME)/.gopath:$(PWD)
TARGET=/usr/local/bin

## help: this help file
help:
	@( echo "" && echo "Makefile targets..." && echo "" )
	@( cat Makefile | grep '^##' | sed -e 's/##/ -/' | sort && echo "" )

## build: build the cli and tcp services
build: 
	@[ -d bin ] || mkdir bin
	( /bin/rm -f bin/* )
	( go build -o bin/unique src/main.go )
	( go build -o bin/unique-tcp tcp/unique-tcp.go )

## install: install all the cli versions locally, ulid, uuid, txid, xuid, etc
install:
	@make build
	cp -f bin/unique $(TARGET)/unique
	ln -f $(TARGET)/unique $(TARGET)/ulid
	ln -f $(TARGET)/unique $(TARGET)/uuid
	ln -f $(TARGET)/unique $(TARGET)/guid
	ln -f $(TARGET)/unique $(TARGET)/tsid
	ln -f $(TARGET)/unique $(TARGET)/txid
	ln -f $(TARGET)/unique $(TARGET)/cuid
	ln -f $(TARGET)/unique $(TARGET)/xuid

## build-linux: build the targets for linux
build-linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o linux/unique src/main.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o linux/unique-tcp tcp/unique-tcp.go

## docker: build and copy the tcp service 
docker:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o linux/unique-tcp tcp/unique-tcp.go
	( cd linux && ./build.sh )

## install-deps: install all required dependencies
install-deps:
	go get -u github.com/golang/lint/golint
	go get github.com/oklog/ulid
	go get github.com/franela/goblin

## format: format all souce code
format:
	( gofmt -s -w src/*.go src/unique/*.go test/*.go tcp/*.go )

## lint: run lint on the source code
lint:
	@( golint src/... && golint test/... )

## test: run all unit tests
test:
	@( go vet src/unique/*.go && go vet src/unique/*.go && go vet src/*.go && cd test/ && go test -cover )
	@( make lint )

## run: run the service
run:
	go run src/main.go

## run-tcp: run the tcp service
run-tcp:
	go run tcp/unique-tcp.go

## examples: run the example client
examples:
	javac examples/UniqueClient.java

## watch: start the source code watcher
watch:
	./watcher.js

## edit: edit the source code
edit:
	vi -O2 src/*/*.go test/*.go src/*.go

.PHONY: format test watch examples edit run lint format install-deps
