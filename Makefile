include ./config.mk

SHELL := /bin/bash
IMAGE_NAME := rain/go-test
CONTAINER_NAME := golang
CONTAINER_BASE_TEST_DIRECTORY := /Users/moral/go/rain

build:
	@docker build --platform linux/arm64 -t ${IMAGE_NAME} .

env-up:
	@docker run -d --name ${CONTAINER_NAME} -v ${CONTAINER_BASE_TEST_DIRECTORY}:/Users/moral/go/rain ${IMAGE_NAME}

	@docker exec -it ${CONTAINER_NAME} /bin/bash 'cd ${CONTAINER_BASE_TEST_DIRECTORY}/${dir} && go run *.go'

env-down:
	@docker stop ${CONTAINER_NAME}
	@docker rm ${CONTAINER_NAME}

	env: