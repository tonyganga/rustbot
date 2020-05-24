CURRENT_DIR=$(shell pwd)

clean:
	rm -rf ${CURRENT_DIR}/bin

env:
	go env

build: env lint
	docker build -t rustbot:latest .

lint:
	golangci-lint run

.PHONY: clean env test build
