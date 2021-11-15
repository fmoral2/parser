## Rain Test

## Dependencies

- Docker
- Go 1.17 +
- Make

## Summary

- This project is to create a parser for incoming third part csv and standardize the file creating 2 outputs and importing to DynamoDb the correct file
import.json ( correct data )
errors.json ( bad data )

## Environment variables (direnv)

- Create a .env file in root and add your credentials like that changing ${aws-access-key-id} and ${aws-secret-access-key} to your credentials or if you do not have DynamoDb just ask me for credentials.

```
export AWS_ACCESS_KEY_ID=${aws-access-key-id}
export AWS_SECRET_ACCESS_KEY=${aws-secret-access-key}
```
## Running the project locally

- export your AWS credentials in your terminal accordingly to your OS.

- To run the project please especify which roster to you want to run: 
- roster , roster1 , roster2 , roster3 , roster4 

```
$ make run-local roster={yourChosenRoster}
```

- To run tests

```
$ make run-test
```

## Running the project with Docker

- To run interactively and see outputs folder:

```
$ make build
```
```
$ make enter-container 
```
```
$ chmod +x /rain 
```
```
$ ./main go run main.go
```

- To see outputs folder:

```
$ cat ./outputs/imported.json && cd ./application
```
```
$ cat ./outputs/errors.json && cd ./application
```

- To run directly:

```
$ make build
```
```
$ make run-docker
```

- Run the unit tests

```
$ make run-tests-docker
```

- To stop and delete container after execution:

```
$ make env-down
```

## Database used to store correct data (employees)

DynamoDb

## Architeture and packages in golang format of 

- Application/internal
- Repository to deal with Db
- Cmd to gather and run main
- Input for data comming and config file
- Output to store imported employees and error file generated
## Bussiness rules applied and why

- All duplicated and null data will not be imported and it will have a console message explaining why and who

- The correct data is imported to DynamoDb to be processed

- In the console will be a summary also explaining the steps of the system

- When you run differente csv files the outputs will be new per run to not cause confusion and to keep it simple


