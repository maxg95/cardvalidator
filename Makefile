.PHONY: run/api build/api

## run/api: run the cmd/api application
run/api:
	go run ./cmd/api

## build/api: build the cmd/api application
build/api:
	@echo 'Building cmd/api...'
	go build -ldflags='-s -w' -o=./bin/api ./cmd/api
	GOOS=linux GOARCH=amd64 go build -ldflags='-s -w' -o=./bin/linux_amd64/api ./cmd/api