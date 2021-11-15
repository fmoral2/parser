include ./config.mk

SHELL := /bin/bash
IMAGE_NAME := golangtest
CONTAINER_NAME := golang-container

build:
	@docker build -t ${IMAGE_NAME} .

enter-container:
	@docker run -e AWS_ACCESS_KEY_ID="${AWS_ACCESS_KEY_ID}" -e AWS_SECRET_ACCESS_KEY="${AWS_SECRET_ACCESS_KEY}" -it --name ${CONTAINER_NAME} ${IMAGE_NAME} sh

run:
	@docker run -e AWS_ACCESS_KEY_ID="${AWS_ACCESS_KEY_ID}" -e AWS_SECRET_ACCESS_KEY="${AWS_SECRET_ACCESS_KEY}" --name ${CONTAINER_NAME} ${IMAGE_NAME}
		
run-tests:
	@docker run ${IMAGE_NAME} sh -c 'cd /rain/application/parser/${dir} && go test'

env-down:
	@docker stop ${CONTAINER_NAME}
	@docker rm ${CONTAINER_NAME}