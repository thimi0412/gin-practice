FROM golang:latest

RUN apt-get update -qq && \
    apt-get install -y mysql-client vim

WORKDIR /go/src/

COPY . .

ENV GO111MODULE=on

RUN go build

ENV PATH /go/bin:$PATH