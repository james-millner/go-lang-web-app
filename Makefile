ifeq ($(GOBIN),)
GOBIN := $(GOPATH)/bin
endif

clean:
	go clean

build:
	go build cmd/go-app/main.go


default: build

.PHONY: test