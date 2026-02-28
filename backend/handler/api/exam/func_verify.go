package exam

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// VerifyExam godoc
// @Summary      Verify exam
// @Description  Mark exam as verified
// @Tags         exams
// @Accept       json
// @Produce      json
// @Param        id path string true "Exam ID"
// @Success      200 {object} map[string]string
// @Failure      401 {object} helper.ErrorResponse
// @Failure      404 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/exams/{id}/verify [post]
func (h *examHTTPHandler) VerifyExam(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "ID wajib diisi", nil)
		return
	}
	err := h.examUsecase.Verify(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, entity.ErrExamNotFound) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Ujian tidak ditemukan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Ujian berhasil diverifikasi"})
}

// UnverifyExam godoc
// @Summary      Unverify exam
// @Description  Cancel verification of exam
// @Tags         exams
// @Accept       json
// @Produce      json
// @Param        id path string true "Exam ID"
// @Success      200 {object} map[string]string
// @Failure      401 {object} helper.ErrorResponse
// @Failure      404 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/exams/{id}/unverify [post]
func (h *examHTTPHandler) UnverifyExam(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "ID wajib diisi", nil)
		return
	}
	err := h.examUsecase.Unverify(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, entity.ErrExamNotFound) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Ujian tidak ditemukan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Verifikasi ujian dibatalkan"})
}
