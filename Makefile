include ./config.mk

SHELL := /bin/bash
IMAGE_NAME := golangtest
CONTAINER_NAME := golang-container

## Docker
build:
	@docker build -t ${IMAGE_NAME} .

enter-container:
	@docker run -e AWS_ACCESS_KEY_ID="${AWS_ACCESS_KEY_ID}" -e AWS_SECRET_ACCESS_KEY="${AWS_SECRET_ACCESS_KEY}" -it --name ${CONTAINER_NAME} ${IMAGE_NAME} sh

run-docker:
	@docker run -e AWS_ACCESS_KEY_ID="${AWS_ACCESS_KEY_ID}" -e AWS_SECRET_ACCESS_KEY="${AWS_SECRET_ACCESS_KEY}" --name ${CONTAINER_NAME} ${IMAGE_NAME}
		
run-tests-docker:
	@docker run ${IMAGE_NAME} sh -c ''

env-down:
	@docker stop ${CONTAINER_NAME}
	@docker rm ${CONTAINER_NAME}

env-up:
	@docker compose up -d

## Local
run-test:
	@export AWS_ACCESS_KEY_ID="${AWS_ACCESS_KEY_ID}"
	@export AWS_SECRET_ACCESS_KEY="${AWS_SECRET_ACCESS_KEY}"
	@cd application/parser && go test

run-local:
	@export AWS_ACCESS_KEY_ID="${AWS_ACCESS_KEY_ID}"
	@export AWS_SECRET_ACCESS_KEY="${AWS_SECRET_ACCESS_KEY}"
	@cd cmd && go run main.go ../input/${roster}.csv
