FROM golang:1.8

MAINTAINER Lohan Bodevan <lohan.bodevan@gmail.com>

ADD . /go/src/github.com/lohanbodevan/grud
WORKDIR /go/src/github.com/lohanbodevan/grud

RUN cd /go/src/github.com/lohanbodevan/grud; go get
CMD cd /go/src/github.com/lohanbodevan/grud; go run main.go

EXPOSE 8080
