# Monitoring LMS [Backend Application]

See More [here](https://upserv.su)

## Build & Run (Locally)
### Prerequisites
- go 1.24
- docker & docker-compose
- [golangci-lint](https://github.com/golangci/golangci-lint) (<i>optional</i>, used to run code checks)
- [swag](https://github.com/swaggo/swag) (<i>optional</i>, used to re-generate swagger documentation)

Create .env file in root directory and add following values:
```dotenv
APP_ENV=local

POSTGRESQL_HOST=
POSTGRESQL_PORT=
POSTGRESQL_USER=
POSTGRESQL_PASSWORD=
POSTGRESQL_NAME=

JWT_SIGNING_KEY=

HTTP_HOST=
```

Use `make run` to build&run project, `make lint` to check code with linter.