DB_URL=postgres://api_user:qwerty@postgres:5432/api_db?sslmode=disable

.SILENT:

migrate-up:
	docker compose run --rm migrations -path=/migrations -database $(DB_URL) up

migrate-down:
	docker compose run --rm migrations -path=/migrations -database $(DB_URL) down
