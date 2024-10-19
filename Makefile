DOCKER_RUN = docker run --rm -v $$PWD:/app -w /app
GO_LINT_RUN = golangci-lint run

.PHONY: lint
lint:
	$(DOCKER_RUN) golangci/golangci-lint:v1.61.0-alpine \
		$(GO_LINT_RUN)

.PHONY: up
up:
	@go run main.go service

.PHONY: build-docker
build-docker:
	@docker build -f build/Dockerfile -t mikrotik-script-generator . --platform=linux/amd64
