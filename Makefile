env:
	go env

build: env mod lint
	docker build -t rustbot:latest .

lint:
	golangci-lint run

mod:
	go mod tidy

.PHONY: env build lint mod
