FROM golang:1.6
MAINTAINER Gabriel Barros

# Install dependencies
RUN go get gopkg.in/mgo.v2
RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/handlers

ADD . /go/src/github.com/006627/pismo/accountsApi
WORKDIR /go/src/github.com/006627/pismo/accountsApi

RUN go install github.com/006627/pismo/accountsApi
ENTRYPOINT /go/bin/accountsApi

EXPOSE 8080