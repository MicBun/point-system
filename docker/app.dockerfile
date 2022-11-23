FROM golang:alpine

WORKDIR /point-system

ADD . .

RUN go mod download

ENTRYPOINT go build  && ./point-system