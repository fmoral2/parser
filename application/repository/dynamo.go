package repository

import (
	"strconv"

	"github.com/morlfm/csv_parser/application/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var Dynamo *dynamodb.DynamoDB

// connectDynamo...
func ConnectDynamo() (db *dynamodb.DynamoDB) {
	return dynamodb.New(session.Must(session.NewSession(&aws.Config{
		Region: &model.RegionName,
	})))
}

// CreateTable...
func CreateTable() error {
	_, err := Dynamo.CreateTable(&dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("ID"),
				AttributeType: aws.String("S"),
			},
			// {
			// 	AttributeName: aws.String("Email"),
			// 	AttributeType: aws.String("S"),
			// },
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("ID"),
				KeyType:       aws.String("HASH"),
			},
			// {
			// 	AttributeName: aws.String("Email"),
			// 	KeyType:       aws.String("RANGE"),
			// },
		},
		BillingMode: aws.String(dynamodb.BillingModePayPerRequest),
		TableName:   aws.String(model.TableName),
	})

	return err
}

// PutItem...
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

// UpdateItem...
func UpdateItem(usersEmp model.Employee) error {
	_, err := Dynamo.UpdateItem(&dynamodb.UpdateItemInput{
		ExpressionAttributeNames: map[string]*string{
			"#N": aws.String("Name"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":Name": {
				S: aws.String(usersEmp.Name),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(usersEmp.Id),
			},
		},
		TableName:        aws.String(model.TableName),
		UpdateExpression: aws.String("SET #N = :Name"),
	})

	return err
}

// DeleteItem...
func DeleteItem(usersEmp model.Employee) error {
	_, err := Dynamo.DeleteItem(&dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(usersEmp.Id),
			},
		},
		TableName: aws.String(model.TableName),
	})

	return err
}

// GetItem...
func GetItem(id int) (usersEmp model.Employee, err error) {
	result, err := Dynamo.GetItem(&dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(strconv.Itoa(id)),
			},
		},
		TableName: aws.String(model.TableName),
	})

	if err != nil {
		return usersEmp, err
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &usersEmp)

	return usersEmp, err

}
