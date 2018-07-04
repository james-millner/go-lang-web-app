PACKAGE  = go-lang-web-app

clean:
	go clean

test:
	go get ./...
	go get -u github.com/stretchr/testify/assert
	go test ./...

build:
	go build cmd/main/main.go

default: build