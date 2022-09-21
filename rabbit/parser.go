package rabbit

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type EmployeesQueue struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	//Location string  `json:"location"`
	Role string  `json:"role"`
	Wage float64 `json:"wage"`
}

func JsonToCsv() {
	file, err := ioutil.ReadFile("input/message.json")
	if err != nil {
		return
	}

	var emps []EmployeesQueue
	err = json.Unmarshal(
		file,
		&emps,
	)
	if err != nil {
		return
	}

	csvFile, err := os.Create("input/empsRabbit.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	w := csv.NewWriter(csvFile)
	// creating headers
	header := []string{"name", "id", "role", "wage"}
	if err = w.Write(header); err != nil {
		return
	}

	// creating csv records
	for _, c := range emps {
		var record []string
		record = append(record, c.Name,
			c.ID,
			//c.Location,
			c.Role,
			strconv.FormatFloat(c.Wage, 'f', 2, 64))

		err = w.Write(record)
		if err != nil {
			return
		}
		record = nil
	}
	defer w.Flush()

}
