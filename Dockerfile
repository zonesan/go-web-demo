FROM alpine

EXPOSE 8080

ADD go-web-demo /
RUN chmod +x /go-web-demo && \
    mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2 

CMD ["/go-web-demo"]
