SHELL := /bin/bash
export GOPATH:=$(HOME)/.gopath:$(PWD)

build: 
	@[ -d bin ] || mkdir bin
	( go build -o bin/unique src/main.go )

install-deps:
	go get -u github.com/golang/lint/golint
	go get github.com/oklog/ulid
	go get github.com/franela/goblin
	go get github.com/hashicorp/go-uuid

format:
	( gofmt -s -w src/*.go src/unique/*.go test/*.go )

lint:
	@( golint src/... && golint test/... )

test:
	@( go vet src/unique/*.go && go vet src/unique/*.go && go vet src/*.go && cd test/ && go test -cover )
	@( make lint )

run:
	go run src/main.go

watch:
	./watcher.js

edit:
	vi -O3 src/*/*.go test/*.go src/*.go

.PHONY: format
.PHONY: test
.PHONY: watch
