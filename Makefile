SHELL := /bin/bash

govendor:
	go mod tidy -compat=1.19
	go mod vendor
	git add vendor

build:
	go build -o main main.go

wire:
	source ence.sh && wire
