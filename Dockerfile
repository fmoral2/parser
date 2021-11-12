FROM golang:1.17 AS golang
RUN mkdir /rain
ADD . /rain
RUN apt-get update
WORKDIR /rain 
ENTRYPOINT [ "" ]
COPY go.mod go.sum ./rain
COPY . .
RUN go mod tidy
CMD ["go", "run", "main.go"]

