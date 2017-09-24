SHELL := /bin/bash
export GOPATH:=$(HOME)/.gopath:$(PWD)
TARGET=/usr/local/bin

build: 
	@[ -d bin ] || mkdir bin
	( /bin/rm -f bin/* )
	( go build -o bin/unique src/main.go )

install:
	@make build
	cp -f bin/unique $(TARGET)/unique
	ln -f $(TARGET)/unique $(TARGET)/ulid
	ln -f $(TARGET)/unique $(TARGET)/uuid
	ln -f $(TARGET)/unique $(TARGET)/guid
	ln -f $(TARGET)/unique $(TARGET)/tsid
	ln -f $(TARGET)/unique $(TARGET)/txid

install-deps:
	go get -u github.com/golang/lint/golint
	go get github.com/oklog/ulid
	go get github.com/franela/goblin

format:
	( gofmt -s -w src/*.go src/unique/*.go test/*.go )

lint:
	@( golint src/... && golint test/... )

test:
	@( go vet src/unique/*.go && go vet src/unique/*.go && go vet src/*.go && cd test/ && go test -cover )
	@( make lint )

qtest:
	@( cd test && go test -cover )

run:
	go run src/main.go

watch:
	./watcher.js

edit:
	vi -O3 src/*/*.go test/*.go src/*.go

.PHONY: format
.PHONY: test
.PHONY: watch
