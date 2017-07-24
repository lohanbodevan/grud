FROM golang:1.8

MAINTAINER Lohan Bodevan <lohan.bodevan@gmail.com>

ADD . /go/src/github.com/lohanbodevan/grud

RUN cd /go/src/github.com/lohanbodevan/grud; go get

EXPOSE 8080

CMD go run main.go
