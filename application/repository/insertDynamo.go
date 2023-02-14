package repository

import (
	"log"

	"github.com/fmoral2/parser/application/parser"

	"github.com/fmoral2/parser/application/model"
)

func CreateTableDynamoDb() {
	err := CreateTable()
	if err != nil {
		log.Println(err)
	}
}

func PutItemDynamoDb(empList []model.Employee) {
	validatedList := parser.ValidateEmployees(empList)
	for _, e := range validatedList {
		err := PutItem(e)
		if err != nil {
			log.Println(err)
		}
	}
}
