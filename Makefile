env:
	go env

build: env mod 
	docker build -t rustbot:latest .

mod:
	go mod tidy

.PHONY: env build mod
