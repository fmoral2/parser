package ports

import (
	"github.com/morlfm/csv_parser/application/parser"
	"github.com/morlfm/csv_parser/application/repository"
	"log"
)

func Entry() {
	messagesEntry := "input/empsRabbit.csv"
	_ = repository.CreateTable()
	empList := parser.ReadFiles(messagesEntry)
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
