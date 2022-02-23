package main

import (
	"log"

	"github.com/morlfm/csv_parser/application/parser"
	"github.com/morlfm/csv_parser/application/repository"
	"github.com/morlfm/csv_parser/rabbit"
)

func init() {
	repository.Dynamo = repository.ConnectDynamo()
}

func main() {
	rabbit.Consumer()
	rabbit.JsonToCsv()
	repository.CreateTable()
	empList := parser.ReadFiles()
	errors := parser.NotImportedEmployees(empList)
	parser.ErrorsToJson(errors)

	validatedList := parser.ValidateEmployees(empList)
	for _, e := range validatedList {
		err := repository.PutItem(e)
		if err != nil {
			log.Println(err)
		}
	}
}
