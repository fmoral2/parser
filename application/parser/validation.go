package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/morlfm/csv_parser/application/model"
)

func NotImportedEmployees(employees []model.Employee) []model.Employee {
	listErrors := []model.Employee{}
	MailDup := EmailDuplicated(employees)
	IdDup := IdDuplicated(employees)

	for _, emp := range employees {
		if MailDup[emp.Email] || emp.Id == "" || emp.Email == "" || IdDup[emp.Id] || emp.Salary == "" {
			listErrors = append(listErrors, emp)
		}
	}
	return listErrors
}

func EmailDuplicated(list []model.Employee) map[string]bool {
	dupMap := make(map[string]bool)
	emailMap := make(map[string]bool)
	for _, e := range list {
		if _, ok := emailMap[e.Email]; len(e.Email) > 0 && ok {
			log.Printf("Email duplicated found = %s ", e.Email)
			dupMap[e.Email] = true
		}
		emailMap[e.Email] = true
	}
	return dupMap
}

func IdDuplicated(list []model.Employee) map[string]bool {
	dupMap := make(map[string]bool)
	idMap := make(map[string]bool)
	for _, e := range list {
		if _, ok := idMap[e.Id]; len(e.Id) > 0 && ok {
			log.Printf("Id duplicated found = %s ", e.Id)
			dupMap[e.Id] = true
		}
		idMap[e.Id] = true
	}
	return dupMap
}

func ValidateEmployees(employees []model.Employee) []model.Employee {
	listValidated := []model.Employee{}
	dupMap := EmailDuplicated(employees)
	IdDupMap := IdDuplicated(employees)

	for _, emp := range employees {
		if dupMap[emp.Email] {
			fmt.Println("Duplicated email: user not imported => ", emp.Name)
		} else if emp.Id == "" {
			fmt.Println("Id null: user not imported =>", emp.Name)
		} else if emp.Email == "" {
			fmt.Println("Email null: user not imported =>", emp.Name)
		} else if IdDupMap[emp.Id] {
			fmt.Println("Duplicated id: user not imported =>", emp.Name)
		} else if emp.Salary == "" {
			fmt.Println("Salary null: user not imported =>", emp.Name)
		} else {
			listValidated = append(listValidated, emp)
			fmt.Println("Users imported in DynamoDb => ", emp.Name)
		}
	}

	json_data, err := json.Marshal(listValidated)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	json_file, err := os.Create("outputs/imported.json")
	if err != nil {
		fmt.Println(err)
	}
	defer json_file.Close()

	json_file.Write(json_data)
	json_file.Close()

	return listValidated
}
