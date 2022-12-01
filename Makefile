tg:
	go run ./cmd/tg/main.go

db:
	docker-compose -f build/docker-compose.yml up -d

lint:
	golangci-lint run

migrations-up:
	migrate -path ./internal/migrations/psql \
		-database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" \
		up

migrations-down:
	migrate -path ./internal/migrations/psql \
		-database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" \
		down
