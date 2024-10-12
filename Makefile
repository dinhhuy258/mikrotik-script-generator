.PHONY: lint
lint:
	@echo "==> Running lint check..."
	@golangci-lint run

.PHONY: up
up:
	@go run main.go service
