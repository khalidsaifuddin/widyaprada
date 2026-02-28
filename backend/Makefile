.PHONY: swagger swag docs build vendor

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
