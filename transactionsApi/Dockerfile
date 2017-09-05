FROM golang:1.6
MAINTAINER Gabriel Barros

# Install dependencies
RUN go get gopkg.in/mgo.v2
RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/handlers

ADD . /go/src/github.com/006627/pismo/transactionsApi
WORKDIR /go/src/github.com/006627/pismo/transactionsApi

RUN go install github.com/006627/pismo/transactionsApi
ENTRYPOINT /go/bin/transactionsApi

EXPOSE 8000