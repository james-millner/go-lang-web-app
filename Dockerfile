FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN ls -al
RUN go build .
CMD ["/app/main"]