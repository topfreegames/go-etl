FROM golang:1.10-alpine

COPY . $GOPATH/src/github.com/topfreegames/go-etl

WORKDIR $GOPATH/src/github.com/topfreegames/go-etl

RUN apk update && apk add git make && \
    go get -u github.com/golang/dep/cmd/dep && \
    dep ensure && \
    make plugins

CMD ["make", "start"]
