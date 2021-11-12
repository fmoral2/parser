package main

import (
	"log"

	"github.com/morlfm/csv_parser/internal/parser"
	"github.com/morlfm/csv_parser/internal/repository"
)

func init() {
	repository.Dynamo = repository.ConnectDynamo()
}

func main() {
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
