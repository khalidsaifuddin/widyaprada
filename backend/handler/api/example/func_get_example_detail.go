package example

import (
	"net/http"
	"strconv"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// GetExampleDetail godoc
// @Summary      Get example detail
// @Description  Returns a single example by ID
// @Tags         example
// @Accept       json
// @Produce      json
// @Param        id path int true "Example ID"
// @Success      200 {object} helper.Response
// @Failure      400 {object} helper.ErrorResponse
// @Failure      404 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Router       /api/v1/examples/{id} [get]
func (h *exampleHTTPHandler) GetExampleDetail(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", "ID is required")
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", "ID must be a valid integer")
		return
	}

	resp, err := h.exampleUsecase.GetExampleDetail(c, id)
	if err != nil {
		if entity.IsRecordNotFound(err) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Not found", err.Error())
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", err.Error())
		return
	}
	helper.ResponseOutput(c, int32(http.StatusOK), "Success", resp)
}
