package parser

import (
	"errors"
	"log"

	"github.com/morlfm/csv_parser/internal/model"
)

func ParseEmployee(s []string, h model.Header) model.Employee {

	var emp model.Employee

	idIndex := *h.Id
	nameIndex := *h.Name
	emailIndex := *h.Email
	salaryIndex := *h.Salary

	emp.Id = s[idIndex]
	emp.Name = s[nameIndex]
	emp.Email = s[emailIndex]
	emp.Salary = s[salaryIndex]

	return emp
}

func ParseHeader(columns []string, config model.HeaderConfig) (model.Header, error) {

	var header model.Header

	for i, c := range columns {
		identifier, err := config.GetHeaderIdentifier(c)
		if err != nil {
			log.Println("Header " + c + " not identified")
		}

		pos := i
		switch identifier {
		case "Id":
			header.Id = &pos

		case "Name":
			header.Name = &pos

		case "Email":
			header.Email = &pos

		case "Salary":
			header.Salary = &pos
		}
	}
	if header.Id == nil || header.Name == nil || header.Email == nil || header.Salary == nil {
		return header, errors.New("missing headers")
	}
	return header, nil
}
