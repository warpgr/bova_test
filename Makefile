.PHONY: test build run-exec lint dc

test:
	go test ./... -race

build:
	go build -o price-service -ldflags="-s -w" cmd/main.go

lint:
	golangci-lint run -v

run: build
	./price-service

dc:
	docker-compose up