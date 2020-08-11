env:
	go env

build: env mod test 
	docker build -t rustbot:latest .

build-rpi: env mod test
	docker build -t rustbot:arm --build-arg GOARCH=arm .

mod:
	go mod download

test:
	go test -v ./...

.PHONY: env build mod test build-rpi
