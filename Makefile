.PHONY:build
build: 
	go build -v .

.PHONY:test
test:
	go test -v -race -timeout 30s ./ ...

.DEFAULD_GOAL := build