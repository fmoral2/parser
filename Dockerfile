FROM golang:1.17 AS golang

RUN apt-get update
WORKDIR /Users/moral/go/rain
COPY . .
RUN go mod tidy

