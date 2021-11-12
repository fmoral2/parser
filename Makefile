include ./config.mk

SHELL := /bin/bash
IMAGE_NAME := rain/go-test
CONTAINER_NAME := golang
CONTAINER_BASE_TEST_DIRECTORY := /Users/moral/go/rain

build:
	@docker build --platform linux/arm64 -t ${IMAGE_NAME} .

env-up:
	@docker exec -it ${CONTAINER_NAME} /bin/bash 'cd ${CONTAINER_BASE_TEST_DIRECTORY}/${dir} && go run main.go'

env-down:
	@docker stop ${CONTAINER_NAME}
	@docker rm ${CONTAINER_NAME}
