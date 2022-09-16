package repository

import (
	"github.com/morlfm/csv_parser/application/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var Dynamo *dynamodb.DynamoDB

func ConnectDynamo() (db *dynamodb.DynamoDB) {
	return dynamodb.New(session.Must(session.NewSession(&aws.Config{
		Region: &model.RegionName,
	})))
}

func CreateTable() error {
	_, err := Dynamo.CreateTable(&dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("ID"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("ID"),
				KeyType:       aws.String("HASH"),
			},
		},
		BillingMode: aws.String(dynamodb.BillingModePayPerRequest),
		TableName:   aws.String(model.TableName),
	})

	return err
}

func PutItem(usersEmp model.Employee) error {
	_, err := Dynamo.PutItem(&dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(usersEmp.Id),
			},
			"Name": {
				S: aws.String(usersEmp.Name),
			},
			"Email": {
				S: aws.String(usersEmp.Email),
			},
			"Salary": {
				S: aws.String(usersEmp.Salary),
			},
		},
		TableName: aws.String(model.TableName),
	})

	return err
}
