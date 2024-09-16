.PHONY: run
run:
	@go run ./cmd/tender-service/main.go
	
.PHONY: up
up:
	@cd ./deploy && docker compose up -d

.PHONY: down
down:
	@cd ./deploy && docker compose down

.PHONY: migrate
migrate:
	@go run ./cmd/migration/main.go

.PHONY: test
test:
	@go test ./... -coverprofile cover.out && go tool cover -func cover.out