package ports

import (
	"log"

	"github.com/fmoral2/parser/application/parser"
	"github.com/fmoral2/parser/application/repository"
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
