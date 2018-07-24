FROM golang:1.10-alpine

ENV GOOS=linux 
ENV GOARCH=amd64 
ENV CGO_ENABLED=0

WORKDIR $GOPATH/src/github.com/topfreegames/go-etl

RUN apk update 
RUN apk add git make gcc
RUN go get -u github.com/golang/dep/cmd/dep

COPY . $GOPATH/src/github.com/topfreegames/go-etl

RUN dep ensure 
RUN make plugins

CMD ["make", "start"]
