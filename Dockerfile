# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main github.com/morlfm/csv_parser/cmd
CMD ["./main"]

