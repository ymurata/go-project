FROM golang:1.15.0-alpine3.12

ENV GO111MODULE=on

RUN apk update && \
    apk --no-cache add git

RUN go get github.com/pilu/fresh && \
    go get -tags 'postgres' -u github.com/golang-migrate/migrate/v4/cmd/migrate@master

COPY ./docker-entrypoint.sh ./docker-entrypoint.sh

ENTRYPOINT ["./docker-entrypoint.sh"]
