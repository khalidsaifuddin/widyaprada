package exam

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// PublishExam godoc
// @Summary      Publish exam
// @Description  Publish exam (Draft -> Diterbitkan)
// @Tags         exams
// @Accept       json
// @Produce      json
// @Param        id path string true "Exam ID"
// @Success      200 {object} map[string]string
// @Failure      401 {object} helper.ErrorResponse
// @Failure      404 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/exams/{id}/publish [post]
func (h *examHTTPHandler) PublishExam(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "ID wajib diisi", nil)
		return
	}
	err := h.examUsecase.Publish(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, entity.ErrExamNotFound) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Ujian tidak ditemukan", nil)
			return
		}
		if errors.Is(err, entity.ErrExamAlreadyPublished) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Ujian sudah diterbitkan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Ujian berhasil diterbitkan"})
}
