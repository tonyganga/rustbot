image: build
	docker build -t rustbot:latest .

build: mod test 
	go build .

build-rpi: env mod test
	docker build -t rustbot:arm --build-arg GOARCH=arm .

mod:
	go mod tidy && go mod download

test:
	go test -v ./...

clean:
	rm -r ./rustbot

up: build
	./rustbot

.PHONY: build mod test build-rpi image 
