package repository

import (
	"log"

	"github.com/morlfm/csv_parser/application/parser"

	"github.com/morlfm/csv_parser/application/model"
)

func PutItemDynamo(empList []model.Employee) {
	validatedList := parser.ValidateEmployees(empList)
	for _, e := range validatedList {
		err := PutItem(e)
		if err != nil {
			log.Println(err)
		}
	}
}
