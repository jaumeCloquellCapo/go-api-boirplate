build:
	@gogo build -o bin/main main.go

run:
	@go run -race main.go server

test:							## Run all tests
	@go test ./... -v