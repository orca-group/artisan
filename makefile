build:
	go mod download
	go build -o bin/artisan --ldflags "-s -w" ./

.PHONY: format
format:
	go fmt ./...
