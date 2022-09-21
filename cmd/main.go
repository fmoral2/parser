package main

import (
	"github.com/morlfm/csv_parser/application/repository"
	"github.com/morlfm/csv_parser/rabbit"
)

func init() {
	repository.Dynamo = repository.ConnectDynamo()
}

func main() {
	rabbit.Consumer()
}
