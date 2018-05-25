PACKAGE  = go-lang-web-app

clean:
	go clean

test:
	go get ./...
	go test ./...

build:
	go get github.com/PuerkitoBio/goquery && go get github.com/stretchr/testify/assert && go get github.com/gorilla/mux
	go build cmd/main/main.go

default: build
