FROM golang:latest
COPY ./ /go/src/github.com/james-millner/go-lang-web-app/
WORKDIR /go/src/github.com/james-millner/go-lang-web-app/

RUN go get ./...
RUN go get -u github.com/stretchr/testify/assert
RUN go build cmd/main/main.go

CMD ["/main"]