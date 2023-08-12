OS = $(shell uname | tr A-Z a-z)
BINARY_NAME = "barber-api"
EXECUTABLE_PATH = "cmd/barber/main.go"

.PHONY: dev
dev:
	@if [ ! -f .env.development ]; then \
		echo "Error: .env.development file not found. Please create the .env.development file."; \
		exit 1; \
	fi
	@cp .env.development .env
	@go run ${EXECUTABLE_PATH}

.PHONY: build
build:
	@if [ ! -f .env.production ]; then \
		echo "Error: .env.production file not found. Please create the .env.production file."; \
		exit 1; \
	fi
	@cp .env.production .env
	@GOARCH=amd64 GOOS=${OS} CGO_ENABLED=0 go build -ldflags="-w -s" -o ${BINARY_NAME} ${EXECUTABLE_PATH}

.PHONY: start
start:
	@./${BINARY_NAME}

.PHONY: test-coverage
test-coverage:
	@go test ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out

.PHONY: upgrade-dependencies
upgrade-dependencies:
	@go get -u ./...

.PHONY: generate-docs
generate-docs:
	@swag init -g ${EXECUTABLE_PATH}

.PHONY: format-docs
format-docs:
	@swag fmt -g ${EXECUTABLE_PATH}
