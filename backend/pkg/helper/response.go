package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status string      `json:"status,omitempty"`
	Code   int32       `json:"code,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Status  string `json:"status,omitempty"`
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

func ResponseOutput(c *gin.Context, statusCode int32, statusMessage string, data interface{}) {
	c.JSON(int(statusCode), Response{
		Status: statusMessage,
		Code:   statusCode,
		Data:   data,
	})
}

func ErrorResponseOutput(c *gin.Context, statusCode int32, err error, message string) {
	errStr := "internal server error"
	if err != nil {
		errStr = err.Error()
	}
	c.JSON(int(statusCode), ErrorResponse{
		Status:  "error",
		Code:    statusCode,
		Message: message,
		Error:   errStr,
	})
}

func SafeErrorResponse(c *gin.Context, statusCode int32, err error, operation string) {
	if err == nil {
		ErrorResponseOutput(c, statusCode, nil, "Unknown error")
		return
	}
	LogError("Error in %s during %s: %v", c.Request.URL.Path, operation, err)
	obfuscated := ObfuscateErrorWithContext(err, operation)
	ErrorResponseOutput(c, statusCode, obfuscated, "Operation failed")
}

func DatabaseErrorResponse(c *gin.Context, err error, operation string) {
	if err == nil {
		return
	}
	LogError("Database error in %s during %s: %v", c.Request.URL.Path, operation, err)
	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Status:  "error",
		Code:    int32(http.StatusInternalServerError),
		Message: "Database operation failed",
		Error:   "internal server error",
	})
}

func GenerateTotalPage(totalData, limit int64) int64 {
	if limit <= 0 {
		return 0
	}
	n := totalData / limit
	if totalData%limit > 0 {
		n++
	}
	return n
}

func GetOffsetAndLimit(page, pageSize int64) (offset, limit int64) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	return (page - 1) * pageSize, pageSize
}
