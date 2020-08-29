FROM golang:1.15.0-alpine3.12

ENV GO111MODULE=on

RUN apk update && \
    apk --no-cache add git gcc musl-dev make curl

RUN go get github.com/pilu/fresh

WORKDIR /app
