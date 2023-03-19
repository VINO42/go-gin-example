FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/vino42/go-gin-example
COPY . $GOPATH/src/github.com/vino42/go-gin-example
RUN go build .

EXPOSE 8001
ENTRYPOINT ["./go-gin-example"]