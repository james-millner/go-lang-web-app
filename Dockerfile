FROM golang:latest
COPY ./ /go/src/github.com/james-millner/go-lang-web-app/
WORKDIR /go/src/github.com/james-millner/go-lang-web-app/
RUN go get -u all

RUN go test github.com/james-millner/go-lang-web-app/pkg/web
RUN go build cmd/main/main.go
CMD ["/main"]