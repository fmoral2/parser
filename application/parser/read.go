package parser

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/morlfm/csv_parser/application/model"
)

func ReadFiles() []model.Employee {

	r, _ := os.OpenFile("../input/config2.json", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer r.Close()
	conf, _ := io.ReadAll(r)

	var config model.HeaderConfig

	if err := json.Unmarshal([]byte(conf), &config); err != nil {
		panic(err)
	}

	// read csv file path from command line using os.Args

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	rf := csv.NewReader(file)

	headers, err := rf.Read()
	if err != nil {
		panic(err)
	}
	header, err := ParseHeader(headers, config)
	if err != nil {
		panic(err)
	}
	var empList []model.Employee

	for {
		columns, err := rf.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
		}
		emp := ParseEmployee(columns, header)
		empList = append(empList, emp)
	}

	fmt.Println("Users processed: \n", empList)

	return empList
}
