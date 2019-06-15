.PHONY: build clean test run

build:
	GOOS=linux GOARCH=amd64 go build -o api ./main.go

clean: 
	rm -rf ./api

test:
	go test ./...

run:
	GOOS=linux GOARCH=amd64 go build -o api ./main.go
	open http://localhost:8080/healthcheck
	sam local start-api -p 8080 --template template.json