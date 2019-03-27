FROM golang:1.11

WORKDIR /go/src/app

RUN git clone https://github.com/ma-null/NetInterface
COPY . .

RUN go get github.com/urfave/cli && \
    go get github.com/julienschmidt/httprouter && \
    go get github.com/ma-null/NetInterface && \
    go get github.com/ma-null/net_cli

CMD ["go run server.go"]