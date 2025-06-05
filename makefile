APP_NAME = ./cmd/main.go
BUILD_OUTPUT = ./bin/app
DB_URL=postgres://api_user:qwerty@postgres:5432/api_db?sslmode=disable

.PHONY: docker-up docker-down migrate-up migrate-down build run clean
.SILENT:

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

migrate-up:
	docker compose run --rm migrations -path=/migrations -database $(DB_URL) up

migrate-down:
	docker compose run --rm migrations -path=/migrations -database $(DB_URL) down

build:
	go build -o $(BUILD_OUTPUT) $(APP_NAME)

run: build docker-up wait-for-db migrate-up
	./$(BUILD_OUTPUT)

clean:
	rm -f $(BUILD_OUTPUT)

wait-for-db:
	@echo
	sleep 5