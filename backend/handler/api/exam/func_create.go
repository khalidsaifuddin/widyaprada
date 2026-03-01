package exam

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// CreateExam godoc
// @Summary      Create exam
// @Description  Create new exam (Draft status)
// @Tags         exams
// @Accept       json
// @Produce      json
// @Param        body body entity.CreateExamRequest true "Create exam request"
// @Success      201 {object} entity.ExamDetailResponse
// @Failure      400 {object} helper.ErrorResponse
// @Failure      401 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/exams [post]
func (h *examHTTPHandler) CreateExam(c *gin.Context) {
	var req entity.CreateExamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}
	resp, err := h.examUsecase.Create(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, entity.ErrExamCodeExists) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Kode ujian sudah digunakan", nil)
			return
		}
		if errors.Is(err, entity.ErrExamMinContent) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Minimal 1 soal atau paket dalam ujian", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusCreated, resp)
}
