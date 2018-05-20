PACKAGE  = go-lang-web-app
DATE    ?= $(shell date +%FT%T%z)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)
PKGS     = $(or $(PKG),$(shell $(GO) list ./... | grep -v "^$(PACKAGE)/vendor/"))
TESTPKGS = $(shell $(GO) list -f '{{ if or .TestGoFiles .XTestGoFiles }}{{ .ImportPath }}{{ end }}' $(PKGS))

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
