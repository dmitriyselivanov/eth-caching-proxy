FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/eth-caching-proxy
COPY . $GOPATH/src/eth-caching-proxy
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./eth-caching-proxy"]
