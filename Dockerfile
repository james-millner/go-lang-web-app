FROM golang:latest
COPY ./ /go/src/github.com/james-millner/go-lang-web-app/
WORKDIR /go/src/github.com/james-millner/go-lang-web-app/
RUN ls -al && cd .. && ls -al && cd .. && ls -al && cd .. && ls -al && cd .. && ls -al && cd .. && ls -al
RUN go get github.com/PuerkitoBio/goquery
RUN go build cmd/main/main.go
CMD ["/app/main"]