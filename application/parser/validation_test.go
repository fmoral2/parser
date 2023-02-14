package parser_test

import (
	"testing"

	"github.com/fmoral2/parser/application/model"
	"github.com/fmoral2/parser/application/parser"
	"github.com/stretchr/testify/assert"
)

func TestEmailDuplicated(t *testing.T) {

	cases := map[string]struct {
		list     []model.Employee
		expected map[string]bool
	}{
		"not_duplicated": {
			list: []model.Employee{
				{
					Id:     "1",
					Name:   "Moral",
					Email:  "moral@test.com",
					Salary: "10",
				},
			},
			expected: map[string]bool{},
		},
		"duplicated": {
			list: []model.Employee{
				{
					Id:     "1",
					Name:   "Moral",
					Email:  "moral@test.com",
					Salary: "10",
				},
				{
					Id:     "3",
					Name:   "Francisco",
					Email:  "moral@test.com",
					Salary: "100",
				},
			},
			expected: map[string]bool{
				"moral@test.com": true,
			},
		},
	}

	for caseTitle, c := range cases {
		t.Run(caseTitle, func(t *testing.T) {
			result := parser.EmailDuplicated(c.list)
			assert.Equal(t, c.expected, result)
		})
	}
}

func TestIdDuplicated(t *testing.T) {

	cases := map[string]struct {
		list     []model.Employee
		expected map[string]bool
	}{
		"not_duplicated": {
			list: []model.Employee{
				{
					Id:     "1",
					Name:   "Moral",
					Email:  "moral@test.com",
					Salary: "10",
				},
			},
			expected: map[string]bool{},
		},
		"duplicated": {
			list: []model.Employee{
				{
					Id:     "3",
					Name:   "Moral",
					Email:  "moral@test.com",
					Salary: "10",
				},
				{
					Id:     "3",
					Name:   "John",
					Email:  "john@test.com",
					Salary: "100",
				},
			},
			expected: map[string]bool{
				"3": true,
			},
		},
	}

	for caseTitle, c := range cases {
		t.Run(caseTitle, func(t *testing.T) {
			result := parser.IdDuplicated(c.list)
			assert.Equal(t, c.expected, result)
		})
	}
}
