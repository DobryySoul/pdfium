APP_NAME=pdfium
BINARY=$(APP_NAME)
DB_URL=postgres://postgres:password@postgres:5432/pdfium?sslmode=disable


migrations-up:
	docker-compose run --rm pdfium migrate -path /app/migrations -database '$(DB_URL)' up

migrations-down:
	docker-compose run --rm pdfium migrate -path /app/migrations -database '$(DB_URL)' down