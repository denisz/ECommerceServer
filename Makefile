.PHONY: build run update test

install:
	go get . main.go

credential:
	GOOS=darwin GOARCH=amd64 go build -o bin/credential credential/credential.go

build:
	GOOS=darwin GOARCH=amd64 go build -o store main.go