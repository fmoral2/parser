package parser_test

import (
	"errors"
	"testing"

	"github.com/morlfm/csv_parser/application/model"
	"github.com/morlfm/csv_parser/application/parser"

	"github.com/stretchr/testify/assert"
)

func TestParseHeader(t *testing.T) {
	expectedId := 1
	expectedName := 0
	expectedEmail := 2
	expectedSalary := 3

	cases := map[string]struct {
		columns        []string
		config         model.HeaderConfig
		expectedHeader model.Header
		expectedErr    error
	}{
		"success": {
			columns: []string{
				"name", "id", "email", "salary",
			},
			config: model.HeaderConfig{
				Id:     []string{"id"},
				Name:   []string{"name"},
				Email:  []string{"email"},
				Salary: []string{"salary"},
			},
			expectedHeader: model.Header{
				Id:     &expectedId,
				Name:   &expectedName,
				Email:  &expectedEmail,
				Salary: &expectedSalary,
			},
			expectedErr: nil,
		},
	}
	for caseTitle, tc := range cases {
		t.Run(caseTitle, func(t *testing.T) {
			result, err := parser.ParseHeader(tc.columns, tc.config)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedHeader.Id, result.Id)
			assert.Equal(t, tc.expectedHeader.Name, result.Name)
			assert.Equal(t, tc.expectedHeader.Email, result.Email)
			assert.Equal(t, tc.expectedHeader.Salary, result.Salary)
		})
	}
}

func TestParseHeaderError(t *testing.T) {

	cases := map[string]struct {
		columns     []string
		config      model.HeaderConfig
		expectedErr error
	}{
		"id_error": {
			columns: []string{
				"name", "id", "email", "salary",
			},
			config: model.HeaderConfig{
				Id:     []string{"phone"},
				Name:   []string{"name"},
				Email:  []string{"email"},
				Salary: []string{"salary"},
			},
			expectedErr: errors.New("missing headers"),
		},
		"name_error": {
			columns: []string{
				"name", "id", "email", "salary",
			},
			config: model.HeaderConfig{
				Id:     []string{"id"},
				Name:   []string{"mobile"},
				Email:  []string{"email"},
				Salary: []string{"salary"},
			},
			expectedErr: errors.New("missing headers"),
		},
	}
	for caseTitle, tc := range cases {
		t.Run(caseTitle, func(t *testing.T) {
			_, err := parser.ParseHeader(tc.columns, tc.config)
			assert.Equal(t, err.Error(), tc.expectedErr.Error())
		})
	}
}
