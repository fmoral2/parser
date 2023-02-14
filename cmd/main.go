package main

import (
	"github.com/fmoral2/parser/application/ports"
	"github.com/fmoral2/parser/application/repository"
	"github.com/fmoral2/parser/rabbit"
)

func main() {
	repository.Dynamo = repository.ConnectDynamo()
	rabbit.Consumer()
	ports.Entry()
}
