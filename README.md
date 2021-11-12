## Rain Test

## Main feature

- Ingest data from third part csv file

## Dependencies

- Docker
- Go 1.16 +
- Make

## Summary

- This project is to create a parser for incoming csv and standardize the file creating 2 outputs and importng to DynamoDb the correct file
import.json ( correct data )
errors.json ( bad data )

## Environment variables (direnv)

```
export AWS_ACCESS_KEY_ID=${given-aws-access-key-id}
export AWS_SECRET_ACCESS_KEY=${given-aws-secret-key}
```
## Running the project locally

- go to cmd folder and run 

```
$ go run main.go
```

## Running the project with Docker


Run the unit tests

```
$ make run-tests
```

Run the project

```
$ make run
```

## Database used to store correct data (employees)

DynamoDb

## Architeture and packages in golang format of 

- Application/internal
- Reposioty to deal wuth Db
- Cmd to gather and run main
- Input for data comming and config file
- Output to store imported employees and error file generated
## Bussiness rules applied and why

- All duplicated and null data will not be imported and it will have a console message explaining why and who

- The correct data is imported to DynamoDb to be processed

- In the console will be a summary also explaining the steps of the system



