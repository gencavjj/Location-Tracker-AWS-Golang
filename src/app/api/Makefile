.PHONY: build clean test test-api

build:
	GOOS=linux GOARCH=amd64 go build -o api ./main.go

clean: 
	rm -rf ./api

test:
	go test ./...

test-api:
	GOOS=linux GOARCH=amd64 go build -o api ./main.go
	open http://127.0.0.1:3000/healthcheck
	sam local start-api --template template.json