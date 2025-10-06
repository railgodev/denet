tidy:
	@echo "Organizing modules..."
	@go mod tidy

lint:
	@echo "Linting..."
	@golangci-lint run

prepare-env:
	@if [ ! -f ./.env ]; then cp ./.example.env ./.env; fi

up: prepare-env
	@docker compose  up -d --build

down:
	@docker compose  down

stop:
	@docker compose  stop
