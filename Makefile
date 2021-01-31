build:
	@go build -o bin/main main.go

run:
	@go run -race main.go server --file=pro.env

test:							## Run all tests
	@go test ./... -v