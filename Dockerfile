FROM golang:alpine

WORKDIR /app

COPY . .

RUN mkdir /input && cp -r input/* /input
RUN apk add --update alpine-sdk
RUN go mod tidy

RUN go build -o main ./cmd/

USER 1001

CMD ["./main"]



