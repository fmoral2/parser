include ./config.mk

SHELL := /bin/bash
IMAGE_NAME := golangtest
CONTAINER_NAME := golang-container
CONTAINER_BASE_TEST_DIRECTORY := /app/application/parser
build:
	@docker build -t ${IMAGE_NAME} .

run:
	@docker run --name ${CONTAINER_NAME} ${IMAGE_NAME} 
	
run-tests:
	@docker run ${IMAGE_NAME} sh -c 'cd /app/application/parser/${dir} && go test'

env-down:
	@docker stop ${CONTAINER_NAME}
	@docker rm ${CONTAINER_NAME}