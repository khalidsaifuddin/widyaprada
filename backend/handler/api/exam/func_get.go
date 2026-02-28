package exam

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// GetExamDetail godoc
// @Summary      Get exam detail
// @Description  Get exam by ID with contents and participants
// @Tags         exams
// @Accept       json
// @Produce      json
// @Param        id path string true "Exam ID"
// @Success      200 {object} entity.ExamDetailResponse
// @Failure      401 {object} helper.ErrorResponse
// @Failure      404 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/exams/{id} [get]
func (h *examHTTPHandler) GetExamDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "ID wajib diisi", nil)
		return
	}
	resp, err := h.examUsecase.Get(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, entity.ErrExamNotFound) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Ujian tidak ditemukan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusOK, resp)
}
