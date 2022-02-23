package model

import "errors"

type Employee struct {
	Id     string
	Name   string
	Email  string
	Salary string
}

type Header struct {
	Id     *int
	Name   *int
	Email  *int
	Salary *int
}

type HeaderConfig struct {
	Id     []string `json:"Id"`
	Name   []string `json:"Name"`
	Email  []string `json:"Email"`
	Salary []string `json:"Salary"`
}

var (
	TableName  = "EMPLOYEES"
	RegionName = "us-east-1"
)

func (c HeaderConfig) GetHeaderIdentifier(name string) (string, error) {

	for _, i := range c.Id {
		if name == i {
			return "Id", nil
		}
	}
	for _, i := range c.Name {
		if name == i {
			return "Name", nil
		}
	}
	for _, i := range c.Email {
		if name == i {
			return "Email", nil
		}
	}
	for _, i := range c.Salary {
		if name == i {
			return "Salary", nil
		}
	}
	return "", errors.New("Header not found")
}
