package rabbit

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

type EmployeesQueue struct {
	Name string  `json:"name"`
	ID   string  `json:"id"`
	Role string  `json:"role"`
	Wage float64 `json:"wage"`
}

func JsonToCsv() {
	f, _ := os.Open("./input/message.json")
	file, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("func JsonToCsv: ", err)
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

	csvFile, err := os.Create("./input/empQueue.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	w := csv.NewWriter(csvFile)
	header := []string{"name", "id", "role", "wage"}
	if err = w.Write(header); err != nil {
		return
	}

	for _, c := range emps {
		var record []string
		record = append(record, c.Name,
			c.ID,
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
