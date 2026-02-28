package question

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// UpdateQuestion godoc
// @Summary      Update question
// @Description  Update question by ID (Super Admin only)
// @Tags         questions
// @Accept       json
// @Produce      json
// @Param        id path string true "Question ID"
// @Param        body body entity.UpdateQuestionRequest true "Update question request"
// @Success      200 {object} entity.QuestionDetailResponse
// @Failure      400 {object} helper.ErrorResponse
// @Failure      401 {object} helper.ErrorResponse
// @Failure      403 {object} helper.ErrorResponse
// @Failure      404 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/questions/{id} [put]
func (h *questionHTTPHandler) UpdateQuestion(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "ID wajib diisi", nil)
		return
	}
	var req entity.UpdateQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}
	resp, err := h.bankSoalUsecase.Update(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, entity.ErrQuestionNotFound) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Soal tidak ditemukan", nil)
			return
		}
		if errors.Is(err, entity.ErrQuestionCodeExists) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Kode soal sudah digunakan", nil)
			return
		}
		if errors.Is(err, entity.ErrQuestionOptionsRequired) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Opsi dan kunci jawaban wajib untuk PG dan Benar-Salah", nil)
			return
		}
		if errors.Is(err, entity.ErrCategoryNotFound) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Kategori tidak ditemukan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusOK, resp)
}
