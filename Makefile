SHA := $(shell git rev-parse --short HEAD)
VERSION := $(shell cat VERSION)
ITTERATION := $(shell date +%s)
LOCALPKGS :=  $(shell go list ./... | grep -v /vendor/)

deps:
	go get -u -f $(LOCALPKGS)

build:
	mkdir -p packaging/root/usr/local/bin
	go build -i -v -o packaging/root/usr/local/bin/oohhc-cli github.com/letterj/oohhc/oohhc-cli
	go build -i -v -o packaging/root/usr/local/bin/oohhc-acctd github.com/letterj/oohhc/oohhc-acctd
	go build -i -v -o packaging/root/usr/local/bin/oohhc-filesysd github.com/letterj/oohhc/oohhc-filesysd

clean:
	rm -f packaging/root/usr/local/bin/oohh-acctd
	rm -f packaging/root/usr/local/bin/oohh-cli
	rm -f packaging/root/usr/local/bin/oohh-filesysd

install: build
	cp -av packaging/root/usr/local/bin/* $(GOPATH)/bin

test:
	go test ./...

packages: clean deps build deb
