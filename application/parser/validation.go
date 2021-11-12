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
	dupList := make(map[string]bool)
	emailMap := make(map[string]bool)
	for _, e := range list {
		if emailMap[e.Email] && e.Email != "" {
			log.Printf("Email duplicated found = %s ", e.Email)
			dupList[e.Email] = true
		}
		emailMap[e.Email] = true
	}
	return dupList
}

func IdDuplicated(list []model.Employee) map[string]bool {
	dupIdList := make(map[string]bool)
	idMap := make(map[string]bool)
	for _, e := range list {
		if idMap[e.Id] && e.Id != "" {
			log.Printf("Id duplicated found = %s ", e.Id)
			dupIdList[e.Id] = true
		}
		idMap[e.Id] = true
	}
	return dupIdList
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
			fmt.Println("Users imported => ", emp.Name)
		}
	}
	//return list of employees imported and create a Json file with the list
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
