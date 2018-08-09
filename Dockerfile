FROM golang:1.10.0 as builder

COPY ./ /go/src/github.com/iqblade/casestudies/
WORKDIR /go/src/github.com/iqblade/casestudies/

RUN make init && make build

FROM gcr.io/distroless/java

EXPOSE 8092

COPY --from=builder /go/src/github.com/iqblade/casestudies/main /main
COPY --from=builder /go/src/github.com/iqblade/casestudies/tika-server-1.14.jar tika-server-1.14.jar

ENTRYPOINT ["/main"]
