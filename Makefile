.PHONY: install test

install:
	go build -o ${GOPATH}/bin/y y.go

test:
	go test ./...
