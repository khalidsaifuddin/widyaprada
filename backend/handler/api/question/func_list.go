package question

import (
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// GetQuestionList godoc
// @Summary      List questions
// @Description  Get paginated list of questions with filters
// @Tags         questions
// @Accept       json
// @Produce      json
// @Param        q query string false "Search"
// @Param        tipe query string false "Filter by type (PG, BENAR_SALAH, ESSAY)"
// @Param        kategori_id query string false "Filter by category"
// @Param        status query string false "Draft, Aktif, all"
// @Param        status_verifikasi query string false "Belum, Sudah, all"
// @Param        page query int false "Page" default(1)
// @Param        page_size query int false "Page size" default(10)
// @Param        sort_by query string false "Sort by"
// @Param        sort_order query string false "asc, desc"
// @Success      200 {object} entity.GetQuestionListResponse
// @Failure      401 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/questions [get]
func (h *questionHTTPHandler) GetQuestionList(c *gin.Context) {
	var req entity.GetQuestionListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}
	resp, err := h.bankSoalUsecase.List(c.Request.Context(), req)
	if err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusOK, resp)
}
