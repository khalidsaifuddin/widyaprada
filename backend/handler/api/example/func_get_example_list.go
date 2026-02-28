package example

import (
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// GetExampleList godoc
// @Summary      Get example list
// @Description  Returns a paginated list of examples
// @Tags         example
// @Accept       json
// @Produce      json
// @Param        page query int false "Page number" default(1)
// @Param        page_size query int false "Page size" default(10)
// @Param        search query string false "Search by name"
// @Success      200 {object} helper.Response
// @Failure      400 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Router       /api/v1/examples [get]
func (h *exampleHTTPHandler) GetExampleList(c *gin.Context) {
	var req entity.GetExampleListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	resp, err := h.exampleUsecase.GetExampleList(c, req)
	if err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", err.Error())
		return
	}
	helper.ResponseOutput(c, int32(http.StatusOK), "Success", resp)
}
