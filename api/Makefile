.PHONY: run
run:
	go run ./cmd/api/*

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
