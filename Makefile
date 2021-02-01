build:
	@go build -o bin/main main.go

run:
	@go run -race main.go server --file=dev.env

test:							## Run all tests
	@go test ./...

migrate:
	@goose -dir ./migrations mysql "db:db@/db?parseTime=true" up