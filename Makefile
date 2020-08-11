env:
	go env

build: env mod test 
	docker build -t rustbot:latest .

mod:
	go mod download

test:
	go test -v ./...

.PHONY: env build mod test
