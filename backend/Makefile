.PHONY: swagger swag docs build vendor test test-cover

# Generate Swagger documentation from annotations
# Requires: go install github.com/swaggo/swag/cmd/swag@latest
swagger swag docs:
	swag init -g main.go -o docs

# Sync vendor with go.mod (fix inconsistent vendoring)
vendor:
	go mod tidy && go mod vendor

# Build binary
build:
	mkdir -p bin && go build -o bin/backend .

# Run unit tests
test:
	go test -v ./...

# Run tests with coverage
test-cover:
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out -o coverage.html
