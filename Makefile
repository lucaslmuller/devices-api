.PHONY: test mocks unit_test docker_build docker_build_test docker_run

test: unit_test e2e_test

mocks:
	@echo "► Generating mocks..."
	@mockgen -destination internal/app/device/domain/mocks/mock_repository.go -package mocks github.com/lucaslmuller/technical-test/internal/app/device/repository IRepository
	@mockgen -destination internal/app/device/domain/mocks/mock_service.go -package mocks github.com/lucaslmuller/technical-test/internal/app/device/domain/service IService
	@mockgen -destination internal/app/device/domain/mocks/mock_cache.go -package mocks github.com/lucaslmuller/technical-test/internal/infrastructure/adapter/redis/cache Cache
	@echo "✔ Mocks generated successfully."

e2e_test: docker_build_test
	@echo "⏹ Stopping any running containers..."
	@docker compose down --remove-orphans
	@echo "► Starting test containers..."
	@docker compose up -d
	@echo "► Running tests..."
	@docker compose exec -T http go test -v ./...
	@echo "-> Cleaning up test environment..."
	@docker compose down

unit_test: mocks
	@echo "► Running unit tests..."
	@go list ./... | grep -v e2e_test | xargs go test -v -race
	@echo "✔ Unit tests completed."

docker_build:
	@echo "► Building Docker image..."
	@docker build . -t service
	@echo "✔ Docker build completed."

docker_build_test:
	@echo "► Building test Docker image..."
	@docker build . -t service_test --target=test
	@echo "✔ Test Docker build completed."

stop:
	@echo "⏹ Stopping any running containers..."
	@docker compose down --remove-orphans

run: stop docker_build
	@echo "► Starting containers..."
	@docker compose up -d
	@echo "► Running service..."
