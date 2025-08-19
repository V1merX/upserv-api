.PHONY: build run test test.coverage swag lint gen clean env-load
.SILENT:
.DEFAULT_GOAL := run

# Определение ОС
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
    OS = linux
    BINARY_EXT =
else ifeq ($(UNAME_S),Darwin)
    OS = darwin
    BINARY_EXT =
else
    OS = windows
    BINARY_EXT = .exe
endif

# Определение архитектуры
UNAME_M := $(shell uname -m)
ifeq ($(UNAME_M),x86_64)
    ARCH = amd64
else ifeq ($(UNAME_M),aarch64)
    ARCH = arm64
else ifeq ($(UNAME_M),arm64)
    ARCH = arm64
else
    ARCH = amd64
endif

BIN_DIR = .bin
BINARY_NAME = api$(BINARY_EXT)
BINARY_PATH = $(BIN_DIR)/$(BINARY_NAME)

# Автоматическая загрузка .env файла если существует
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

build:
	@echo "Building for $(OS)/$(ARCH)..."
	go mod download
	CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) go build -o $(BINARY_PATH) ./cmd/api/main.go
	@echo "Build completed: $(BINARY_PATH)"

build-linux:
	@echo "Building for linux/amd64..."
	go mod download
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/api-linux ./cmd/api/main.go
	@echo "Build completed: $(BIN_DIR)/api-linux"

build-darwin:
	@echo "Building for darwin/arm64..."
	go mod download
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o $(BIN_DIR)/api-darwin ./cmd/api/main.go
	@echo "Build completed: $(BIN_DIR)/api-darwin"

run: build
	@echo "Starting application..."
	$(BINARY_PATH)

run-linux: build-linux
	@echo "Starting Linux binary..."
	$(BIN_DIR)/api-linux

run-darwin: build-darwin
	@echo "Starting Darwin binary..."
	$(BIN_DIR)/api-darwin

test:
	go test --short -coverprofile=cover.out -v ./...
	make test.coverage

test.coverage:
	go tool cover -func=cover.out | grep "total"

swag:
#	swag init -g internal/app/app.go

lint:
	golangci-lint run

gen:
#	mockgen -source=internal/repository/repository.go -destination=internal/repository/mocks/mock.go

clean:
	rm -rf $(BIN_DIR)
	rm -f cover.out

env-load:
	@if [ -f .env ]; then \
		export $$(grep -v '^#' .env | xargs); \
		echo "Environment variables loaded from .env"; \
	else \
		echo ".env file not found"; \
	fi

run-with-env: env-load build
	@echo "Starting application with environment variables..."
	$(BINARY_PATH)

docker-build:
	docker build -t my-app .

docker-run:
	docker run -p 8080:8080 my-app

build-all: build-linux build-darwin
	@echo "All binaries built in $(BIN_DIR)/"

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build          - Build for current OS ($(OS)/$(ARCH))"
	@echo "  build-linux    - Build for Linux/amd64"
	@echo "  build-darwin   - Build for macOS/arm64"
	@echo "  build-all      - Build for all platforms"
	@echo "  run            - Build and run for current OS (auto-load .env)"
	@echo "  run-with-env   - Force load .env and run"
	@echo "  run-linux      - Build and run Linux binary"
	@echo "  run-darwin     - Build and run macOS binary"
	@echo "  test           - Run tests with coverage"
	@echo "  lint           - Run linter"
	@echo "  env-load       - Load environment variables from .env"
	@echo "  clean          - Clean build artifacts"
	@echo "  docker-build   - Build Docker image"
	@echo "  docker-run     - Run Docker container"