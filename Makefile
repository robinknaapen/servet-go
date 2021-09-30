.PHONY: build

build: generate build-broker build-server

build-broker:
	gotip build -o ./bin/broker ./cmd/broker

build-server:
	gotip build -o ./bin/server ./cmd/server

generate:
	gotip generate ./...
