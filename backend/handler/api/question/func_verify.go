package question

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// VerifyQuestion godoc
// @Summary      Verify question
// @Description  Mark question as verified (Verifikator or Super Admin)
// @Tags         questions
// @Accept       json
// @Produce      json
// @Param        id path string true "Question ID"
// @Success      200 {object} map[string]string
// @Failure      401 {object} helper.ErrorResponse
// @Failure      403 {object} helper.ErrorResponse
// @Failure      404 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/questions/{id}/verify [post]
func (h *questionHTTPHandler) VerifyQuestion(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "ID wajib diisi", nil)
		return
	}
	err := h.bankSoalUsecase.Verify(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, entity.ErrQuestionNotFound) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Soal tidak ditemukan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Soal berhasil diverifikasi"})
}

// UnverifyQuestion godoc
// @Summary      Unverify question
// @Description  Cancel verification of question (Verifikator or Super Admin)
// @Tags         questions
// @Accept       json
// @Produce      json
// @Param        id path string true "Question ID"
// @Success      200 {object} map[string]string
// @Failure      401 {object} helper.ErrorResponse
// @Failure      403 {object} helper.ErrorResponse
// @Failure      404 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/questions/{id}/unverify [post]
func (h *questionHTTPHandler) UnverifyQuestion(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "ID wajib diisi", nil)
		return
	}
	err := h.bankSoalUsecase.Unverify(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, entity.ErrQuestionNotFound) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Soal tidak ditemukan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Verifikasi soal dibatalkan"})
}
