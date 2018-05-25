PACKAGE  = go-lang-web-app

clean:
	go clean

test:
	go get github.com/PuerkitoBio/goquery && go get github.com/stretchr/testify/assert && go get github.com/gorilla/mux
	go test ./...

build:
	go build

default: build
