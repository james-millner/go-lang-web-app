FROM golang:1.10
COPY ./ /go/src/codecommit/go-land-web-app/
WORKDIR /go/src/codecommit/go-land-web-app/

RUN make test
RUN make build

CMD ["/main"]pwd