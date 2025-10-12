.PHONY: run
run:
	@command -v air >/dev/null 2>&1 || { \
		go install github.com/air-verse/air@v1.63.0 \
		exit 1; \
	}
	air

.PHONY: dev
dev:
	docker compose -f docker-compose.yml -f docker-compose.dev.yml up -d
	make run

.PHONY: test
test:
	go test ./...


GOLANGCI_LINT_URL = https://golangci-lint.run/usage/install/

.PHONY: check-linter
require-golangci-lint:
	@command -v golangci-lint >/dev/null 2>&1 || { \
		echo "golangci-lint is not installed. Visit $(GOLANGCI_LINT_URL) to install it."; \
		exit 1; \
	}

.PHONY: lint
lint: require-golangci-lint
	golangci-lint run

lint-fix: require-golangci-lint
	golangci-lint run --fix

format: require-golangci-lint
	golangci-lint fmt
