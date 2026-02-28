package exam

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// UpdateExam godoc
// @Summary      Update exam
// @Description  Update exam (only Draft status)
// @Tags         exams
// @Accept       json
// @Produce      json
// @Param        id path string true "Exam ID"
// @Param        body body entity.UpdateExamRequest true "Update exam request"
// @Success      200 {object} entity.ExamDetailResponse
// @Failure      400 {object} helper.ErrorResponse
// @Failure      401 {object} helper.ErrorResponse
// @Failure      404 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/exams/{id} [put]
func (h *examHTTPHandler) UpdateExam(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "ID wajib diisi", nil)
		return
	}
	var req entity.UpdateExamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}
	resp, err := h.examUsecase.Update(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, entity.ErrExamNotFound) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Ujian tidak ditemukan", nil)
			return
		}
		if errors.Is(err, entity.ErrExamNotDraft) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Hanya ujian status Draft yang dapat diedit", nil)
			return
		}
		if errors.Is(err, entity.ErrExamCodeExists) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Kode ujian sudah digunakan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusOK, resp)
}
