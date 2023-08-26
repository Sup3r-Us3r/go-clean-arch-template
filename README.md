# Go Clean Arch Template

## Overview

This template implements the Clean Architecture in a REST API made in Golang, it also has unit tests for `Entities`, `Repositories`, `UseCases`, `Handlers`, `Utility Functions` and `Factories`.

```text
.
├── cmd/
│   └── barber/
│       └── main.go
├── config/
│   ├── config.go
│   └── jwt.go
├── docs/
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── internal/
│   ├── domain/
│   │   ├── apperr/
│   │   │   ├── apperr.go
│   │   │   └── barber.go
│   │   ├── entity/
│   │   │   ├── barber_test.go
│   │   │   └── barber.go
│   │   └── gateway/
│   │       └── barber.go
│   ├── infra/
│   │   ├── database/
│   │   │   └── mongodb.go
│   │   ├── repository/
│   │   │   ├── memory/
│   │   │   │   ├── barber_repository_test.go
│   │   │   │   └── barber_repository.go
│   │   │   ├── mongodb/
│   │   │   │   └── barber_repository.go
│   │   │   └── repository.go
│   │   └── web/
│   │       ├── handler/
│   │       │   ├── v1/
│   │       │   │   ├── auth/
│   │       │   │   │   ├── sign_in_test.go
│   │       │   │   │   └── sign_in.go
│   │       │   │   └── barber/
│   │       │   │       ├── create_barber_test.go
│   │       │   │       ├── create_barber_test.go
│   │       │   │       ├── delete_barber_test.go
│   │       │   │       ├── delete_barber.go
│   │       │   │       ├── fetch_barbers_test.go
│   │       │   │       ├── fetch_barbers.go
│   │       │   │       ├── get_barber_by_id_test.go
│   │       │   │       ├── get_barber_by_id.go
│   │       │   │       ├── update_barber_test.go
│   │       │   │       └── update_barber.go
│   │       │   └── handler.go
│   │       ├── middleware/
│   │       │   ├── logger.go
│   │       │   └── verify_token.go
│   │       └── webserver/
│   │           └── webserver.go
│   ├── mapper/
│   │   └── barber.go
│   ├── usecase/
│   │   ├── auth/
│   │   │   ├── sign_in_test.go
│   │   │   └── sign_in.go
│   │   ├── barber/
│   │   │   ├── create_barber_test.go
│   │   │   ├── create_barber_test.go
│   │   │   ├── delete_barber_test.go
│   │   │   ├── delete_barber.go
│   │   │   ├── fetch_barbers_test.go
│   │   │   ├── fetch_barbers.go
│   │   │   ├── get_barber_by_email_test.go
│   │   │   ├── get_barber_by_email.go
│   │   │   ├── get_barber_by_id_test.go
│   │   │   ├── get_barber_by_id.go
│   │   │   ├── update_barber_test.go
│   │   │   └── update_barber.go
│   │   └── usecase.go
│   └── util/
│       ├── do_passwords_match_test.go
│       ├── do_passwords_match.go
│       ├── email_is_valid_test.go
│       ├── email_is_valid.go
│       ├── generate_random_salt_test.go
│       ├── generate_random_salt.go
│       ├── hash_password_test.go
│       ├── hash_password.go
│       ├── password_is_valid_test.go
│       └── password_is_valid.go
├── log/
│   └── log.go
├── test/
│   └── factory/
│       ├── barber_test.go
│       └── barber.go
├── .editorconfig
├── .env.example
├── .gitignore
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── LICENSE
├── Makefile
└── README.md
```

- [x] REST API
- [x] Golang
- [x] Clean Arch
- [x] SOLID
- [x] Go Chi
- [x] JWT
- [x] MongoDB
- [x] Unit test
- [x] Factory
- [x] Container Pattern
- [x] Mapper
- [x] Swagger
- [x] Swag
- [x] Docker + Docker Compose

---

## Setup API

Access the folder:

```sh
$ cd go-clean-arch-template
```

Create `.env.development` file:

```sh
$ cp .env.example .env.development
```

> Update the variable values as needed.

Configure `/etc/hosts`:

```sh
# Mac and Linux
# /etc/hosts

# Windows
# C:\Windows\System32\drivers\etc\hosts

127.0.0.1 host.docker.internal
```

## Run API

> Note: Need Golang 1.21 or higher

Up Container:

```sh
$ docker-compose up -d
```

Access container:

```sh
$ docker exec -it barber-api /bin/bash
```

Install dependencies:

```sh
$ go mod tidy
```

Run API:

```sh
$ make dev
```

## Scripts

Build API:

```sh
$ make build
```

Start API after build:

```sh
$ make start
```

Upgrade dependencies:

```sh
$ make upgrade-dependencies
```

Generate coverage:

```sh
$ make test-coverage
```

Generate docs:

> You need to install [swag](https://github.com/swaggo/swag) to use the CLI.

```sh
$ make generate-docs
```
