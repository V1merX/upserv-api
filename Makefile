.PHONY:
.SILENT:
.DEFAULT_GOAL := run

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/api ./cmd/api/main.go

run: build
	docker-compose up --remove-orphans app

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