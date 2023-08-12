# Go Clean Arch Template

## Overview

This template implements the Clean Architecture in a REST API made in Golang, it also has unit tests for `Entity`, `Repository`, `UseCase`, `Handler` and `Factory`.

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
│   │       │   │       └── create_barber.go
│   │       │   └── handler.go
│   │       ├── middleware/
│   │       │   ├── logger.go
│   │       │   └── verify_token.go
│   │       └── webserver/
│   │           └── webserver.go
│   ├── mapper/
│   │   └── barber.go
│   └── usecase/
│       ├── auth/
│       │   ├── sign_in_test.go
│       │   └── sign_in.go
│       ├── barber/
│       │   ├── create_barber_test.go
│       │   └── create_barber.go
│       └── usecase.go
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
