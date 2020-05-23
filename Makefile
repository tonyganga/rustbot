CURRENT_DIR=$(shell pwd)

clean:
	rm -rf ${CURRENT_DIR}/bin

env:
	go env

build: env
	docker build -t rustbot:latest .

.PHONY: clean env test build
