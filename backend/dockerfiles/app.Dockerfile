FROM golang:1.22.7-alpine

RUN mkdir -p /usr/local/go/app

WORKDIR /usr/local/go/app

RUN apk update && apk add git