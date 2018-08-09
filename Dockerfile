FROM golang:1.10.0 as builder

COPY ./ /go/src/github.com/james-millner/go-lang-web-app/
WORKDIR /go/src/github.com/james-millner/go-lang-web-app/

RUN make init && make build

FROM gcr.io/distroless/java

EXPOSE 8092

COPY --from=builder /go/src/github.com/james-millner/go-lang-web-app/main /main
COPY --from=builder /go/src/github.com/james-millner/go-lang-web-app/tika-server-1.14.jar tika-server-1.14.jar

ENTRYPOINT ["/main"]
