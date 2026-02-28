package question

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// DeleteQuestion godoc
// @Summary      Delete question
// @Description  Soft delete question by ID, reason required (Super Admin only)
// @Tags         questions
// @Accept       json
// @Produce      json
// @Param        id path string true "Question ID"
// @Param        body body entity.DeleteQuestionRequest true "Delete reason"
// @Success      200 {object} map[string]string
// @Failure      400 {object} helper.ErrorResponse
// @Failure      401 {object} helper.ErrorResponse
// @Failure      403 {object} helper.ErrorResponse
// @Failure      404 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/questions/{id} [delete]
func (h *questionHTTPHandler) DeleteQuestion(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "ID wajib diisi", nil)
		return
	}
	var req entity.DeleteQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Alasan penghapusan wajib diisi", err.Error())
		return
	}
	err := h.bankSoalUsecase.Delete(c.Request.Context(), id, req.Reason)
	if err != nil {
		if errors.Is(err, entity.ErrQuestionNotFound) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Soal tidak ditemukan", nil)
			return
		}
		if errors.Is(err, entity.ErrQuestionDeleteReason) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Alasan penghapusan wajib diisi", nil)
			return
		}
		if errors.Is(err, entity.ErrQuestionInUseByPaket) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Soal ini digunakan oleh paket ujian", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Soal berhasil dihapus"})
}
