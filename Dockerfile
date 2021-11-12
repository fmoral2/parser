FROM golang:1.17 AS golang
RUN mkdir /rain
ADD . /rain
RUN apt-get update
WORKDIR /rain 
ENTRYPOINT [ "" ]
COPY . .
RUN go mod tidy
CMD ["go", "run", "main.go"]

