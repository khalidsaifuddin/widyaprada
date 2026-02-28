package question

import (
	"net/http"

	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// GetCategories godoc
// @Summary      List question categories
// @Description  Get all question categories for dropdown
// @Tags         questions
// @Accept       json
// @Produce      json
// @Success      200 {array} entity.QuestionCategory
// @Failure      401 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/questions/categories [get]
func (h *questionHTTPHandler) GetCategories(c *gin.Context) {
	resp, err := h.bankSoalUsecase.ListCategories(c.Request.Context())
	if err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusOK, resp)
}
