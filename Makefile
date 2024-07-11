GO ?= go

.DEFAULT_GOAL := default

TAGS ?=

.PHONY: tidy
tidy: ## go mod tidy
	${GO} mod tidy

.PHONY: build
build: ## build binary file
	${GO} build -o nuxbt .

.PNONY: gen
gen: ## generate CURD code
	${GO} run ./cmd/gen/main.go

.PHONY: test
test: tidy ## go test
	${GO} test ./...

.PHONY: lint
lint: ## go lint
	golangci-lint run
