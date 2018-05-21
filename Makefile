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

build:
	go build

default: build
