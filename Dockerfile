FROM golang

EXPOSE 8080

ADD main.go $GOPATH/src/github.com/zonesan/go-web-demo/

WORKDIR $GOPATH/src/github.com/zonesan/go-web-demo

RUN go build

CMD ["./go-web-demo"]
