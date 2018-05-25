PACKAGE  = go-lang-web-app

GO         = go
DEP        = dep
GODOC      = godoc
GOFMT      = gofmt
GO2XUNIT   = go2xunit
GOCOVMERGE = gocovmerge
GOCOV      = gocov
GOCOVXML   = gocov-xml
TIMEOUT    = 15

clean:
	go clean

test:
	go get github.com/PuerkitoBio/goquery && go get github.com/stretchr/testify/assert && go get github.com/gorilla/mux
	go test ./...

build:
	go build

default: build
