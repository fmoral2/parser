package main

import (
	"github.com/morlfm/csv_parser/application/ports"
	"github.com/morlfm/csv_parser/application/repository"
	"github.com/morlfm/csv_parser/rabbit"
)

func main() {
	repository.Dynamo = repository.ConnectDynamo()
	rabbit.Consumer()
	ports.Entry()

}
