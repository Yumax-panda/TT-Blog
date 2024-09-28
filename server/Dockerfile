FROM golang:1.23.1 as server-build

WORKDIR  /go/src/TT-Blog

COPY . .

RUN apk upgrade --update && \
    apk --no-cache add git