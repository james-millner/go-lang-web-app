FROM golang:latest
RUN mkdir /app
ADD . /go/src/app
WORKDIR /go/src/app
RUN ls -al
RUN go build .
CMD ["/app/main"]