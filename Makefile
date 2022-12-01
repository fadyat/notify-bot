tg:
	go run ./cmd/tg/main.go

db:
	docker-compose -f build/docker-compose.yml up -d

lint:
	golangci-lint run
