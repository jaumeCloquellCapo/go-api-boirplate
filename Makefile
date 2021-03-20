
build:
	@go build -o bin/main main.go

run:
	@go run -race main.go server --file=dev.env

test:							## Run all tests
	@go test ./...

migrate:
	@goose -dir ./migrations postgres "postgresql://db:db@localhost?sslmode=disable" up

dev_up:
	@docker-compose -f docker-compose.dev.yml up

dev_down:
	@docker-compose -f docker-compose.dev.yml down

pro_up:
	@docker-compose -f docker-compose.yml up

pro_down:
	@docker-compose -f docker-compose.yml down