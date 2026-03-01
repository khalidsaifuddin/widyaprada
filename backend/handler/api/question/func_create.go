package question

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// CreateQuestion godoc
// @Summary      Create question
// @Description  Create new question (Super Admin only)
// @Tags         questions
// @Accept       json
// @Produce      json
// @Param        body body entity.CreateQuestionRequest true "Create question request"
// @Success      201 {object} entity.QuestionDetailResponse
// @Failure      400 {object} helper.ErrorResponse
// @Failure      401 {object} helper.ErrorResponse
// @Failure      403 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/questions [post]
func (h *questionHTTPHandler) CreateQuestion(c *gin.Context) {
	var req entity.CreateQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}
	resp, err := h.bankSoalUsecase.Create(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, entity.ErrQuestionCodeExists) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Kode soal sudah digunakan", nil)
			return
		}
		if errors.Is(err, entity.ErrQuestionOptionsRequired) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Opsi dan kunci jawaban wajib untuk PG, MRA, dan Benar-Salah", nil)
			return
		}
		if errors.Is(err, entity.ErrCategoryNotFound) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Kategori tidak ditemukan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusCreated, resp)
}
