package exam

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// DeleteExam godoc
// @Summary      Delete exam
// @Description  Soft delete exam by ID, reason required
// @Tags         exams
// @Accept       json
// @Produce      json
// @Param        id path string true "Exam ID"
// @Param        body body entity.DeleteExamRequest true "Delete reason"
// @Success      200 {object} map[string]string
// @Failure      400 {object} helper.ErrorResponse
// @Failure      401 {object} helper.ErrorResponse
// @Failure      404 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/exams/{id} [delete]
func (h *examHTTPHandler) DeleteExam(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "ID wajib diisi", nil)
		return
	}
	var req entity.DeleteExamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Alasan penghapusan wajib diisi", err.Error())
		return
	}
	err := h.examUsecase.Delete(c.Request.Context(), id, req.Reason)
	if err != nil {
		if errors.Is(err, entity.ErrExamNotFound) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Ujian tidak ditemukan", nil)
			return
		}
		if errors.Is(err, entity.ErrExamDeleteReason) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Alasan penghapusan wajib diisi", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Ujian berhasil dihapus"})
}
