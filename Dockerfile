FROM golang:1.10
COPY ./ /go/src/github.com/james-millner/go-lang-web-app/
WORKDIR /go/src/github.com/james-millner/go-lang-web-app/

RUN make test
RUN make build

CMD ["/main"]pwd