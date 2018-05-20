ifeq ($(GOBIN),)
GOBIN := $(GOPATH)/bin
endif

clean:
	go clean

build:
	go build

default: build
