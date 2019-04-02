FROM golang:1.11

WORKDIR /go/src/app

RUN go get github.com/urfave/cli && \
    go get github.com/julienschmidt/httprouter && \
    go get github.com/ma-null/NetInterface && \
    go get github.com/ma-null/net_cli

