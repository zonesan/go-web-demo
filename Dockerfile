FROM alpine

EXPOSE 8080


ADD go-web-demo /


CMD ["/go-web-demo"]
