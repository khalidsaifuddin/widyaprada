// Package docs provides Swagger documentation. Regenerate with: swag init -g main.go -o docs
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/examples": {
            "get": {
                "description": "Returns a paginated list of examples",
                "tags": ["example"],
                "summary": "Get example list",
                "parameters": [
                    {"type": "integer", "name": "page", "in": "query", "default": 1},
                    {"type": "integer", "name": "page_size", "in": "query", "default": 10},
                    {"type": "string", "name": "search", "in": "query"}
                ],
                "responses": {
                    "200": {"description": "OK", "schema": {"$ref": "#/definitions/helper.Response"}},
                    "400": {"description": "Bad Request", "schema": {"$ref": "#/definitions/helper.ErrorResponse"}},
                    "500": {"description": "Internal Server Error", "schema": {"$ref": "#/definitions/helper.ErrorResponse"}}
                }
            }
        },
        "/api/v1/examples/{id}": {
            "get": {
                "description": "Returns a single example by ID",
                "tags": ["example"],
                "summary": "Get example detail",
                "parameters": [{"type": "integer", "name": "id", "in": "path", "required": true}],
                "responses": {
                    "200": {"description": "OK", "schema": {"$ref": "#/definitions/helper.Response"}},
                    "400": {"description": "Bad Request", "schema": {"$ref": "#/definitions/helper.ErrorResponse"}},
                    "404": {"description": "Not Found", "schema": {"$ref": "#/definitions/helper.ErrorResponse"}},
                    "500": {"description": "Internal Server Error", "schema": {"$ref": "#/definitions/helper.ErrorResponse"}}
                }
            }
        }
    },
    "definitions": {
        "helper.Response": {
            "type": "object",
            "properties": {
                "status": {"type": "string"},
                "code": {"type": "integer"},
                "data": {}
            }
        },
        "helper.ErrorResponse": {
            "type": "object",
            "properties": {
                "status": {"type": "string"},
                "code": {"type": "integer"},
                "message": {"type": "string"},
                "error": {"type": "string"}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{"http"},
	Title:            "Widyaprada Backend API",
	Description:      "Backend service (hexagon/clean architecture)",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
