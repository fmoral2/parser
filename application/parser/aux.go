package parser

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/morlfm/csv_parser/application/model"
)

func ErrorsToJson(listErrors []model.Employee) {
	json, err := json.Marshal(listErrors)
	if err != nil {
		fmt.Println(err)
	}
	file, err := os.Create("outputs/errors.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.Write(json)
	file.Close()
}
