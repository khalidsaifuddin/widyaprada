package helper

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestGenerateTotalPage(t *testing.T) {
	tests := []struct {
		name      string
		totalData int64
		limit     int64
		want      int64
	}{
		{"exact division", 100, 10, 10},
		{"with remainder", 95, 10, 10},
		{"single page", 5, 10, 1},
		{"zero total", 0, 10, 0},
		{"zero limit", 100, 0, 0},
		{"one item", 1, 10, 1},
		{"limit one", 5, 1, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateTotalPage(tt.totalData, tt.limit)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetOffsetAndLimit(t *testing.T) {
	tests := []struct {
		name     string
		page     int64
		pageSize int64
		wantOff  int64
		wantLim  int64
	}{
		{"first page", 1, 10, 0, 10},
		{"second page", 2, 10, 10, 10},
		{"third page", 3, 20, 40, 20},
		{"page zero defaults to 1", 0, 10, 0, 10},
		{"pageSize zero defaults to 10", 1, 0, 0, 10},
		{"both zero", 0, 0, 0, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			off, lim := GetOffsetAndLimit(tt.page, tt.pageSize)
			assert.Equal(t, tt.wantOff, off)
			assert.Equal(t, tt.wantLim, lim)
		})
	}
}

func TestResponseOutput(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	ResponseOutput(c, 200, "ok", map[string]string{"key": "value"})

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "ok")
	assert.Contains(t, w.Body.String(), "value")
}

func TestErrorResponseOutput(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	err := errors.New("test error")
	ErrorResponseOutput(c, 400, err, "Bad request")

	assert.Equal(t, 400, w.Code)
	assert.Contains(t, w.Body.String(), "error")
	assert.Contains(t, w.Body.String(), "Bad request")
}
