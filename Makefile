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

.PHONY: gen_error_code
gen_error_code: ## generate error code
	${GO} generate github.com/TensoRaws/NuxBT-Backend/module/code/gen

.PHONY: test
test: tidy ## go test
	${GO} test ./...

.PHONY: lint
lint: gen_error_code
	golangci-lint run
	pre-commit install # pip install pre-commit
	pre-commit run --all-files
