package parser

import (
	"errors"
	"log"
	"regexp"

	"github.com/fmoral2/parser/application/model"
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

	// regex to treat special characters
	re := regexp.MustCompile(`[^a-zA-Z0-9\\.\\,]`)
	emp.Salary = re.ReplaceAllString(emp.Salary, " ")

	re = regexp.MustCompile("[a-zA-Z]")
	emp.Id = re.ReplaceAllString(emp.Id, "")

	return emp
}

func ParseHeader(columns []string, config model.HeaderConfig) (model.Header, error) {
	var header model.Header

	for i, c := range columns {
		identifier, err := config.GetHeaderIdentifier(c)
		if err != nil {
			log.Println("Header " + c + " will be ignored")
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
