BINARY = rustbot
GOARCH = amd64
CURRENT_DIR=$(shell pwd)

all: env linux darwin windows

clean:
	rm -rf ${CURRENT_DIR}/bin

env:
	go env

linux:
	GOOS=linux GOARCH=${GOARCH} go build -v -o ${CURRENT_DIR}/bin/linux-${GOARCH}/${BINARY} . ; \

darwin:
	GOOS=darwin GOARCH=${GOARCH} go build -v -o ${CURRENT_DIR}/bin/darwin-${GOARCH}/${BINARY} . ;\

windows:
	GOOS=windows GOARCH=${GOARCH} go build -v -o ${CURRENT_DIR}/bin/windows-${GOARCH}/${BINARY}.exe . ;\

test:
	go test -v

build:
	docker build -t rustbot:latest .

.PHONY: all setup clean env linux darwin windows
