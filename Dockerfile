FROM golang:1.8

MAINTAINER Lohan Bodevan <lohan.bodevan@gmail.com>

ADD . /go/src/github.com/lohanbodevan/grud
WORKDIR /go/src/github.com/lohanbodevan/grud

RUN go get
CMD go run main.go

EXPOSE 8080
