package cbt

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/auth"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// GetCBTExam godoc
// @Summary      Detail ujian untuk petunjuk
// @Description  Dapatkan detail ujian untuk halaman petunjuk/mulai. Mengembalikan exam info dan apakah user dapat mulai.
// @Tags         cbt
// @Accept       json
// @Produce      json
// @Param        id path string true "Exam ID"
// @Success      200 {object} entity.CBTExamDetailResponse
// @Failure      403 {object} helper.ErrorResponse
// @Failure      404 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/cbt/exams/{id} [get]
func (h *cbtHTTPHandler) GetCBTExam(c *gin.Context) {
	claims := auth.GetClaimsFromContext(c)
	if claims == nil {
		helper.ResponseOutput(c, int32(http.StatusUnauthorized), "Unauthorized", nil)
		return
	}

	examID := c.Param("id")
	if examID == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Exam ID wajib diisi", nil)
		return
	}

	resp, err := h.cbtUsecase.GetExamDetail(c.Request.Context(), claims.UserID, examID)
	if err != nil {
		if errors.Is(err, entity.ErrExamNotFound) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Ujian tidak ditemukan.", nil)
			return
		}
		if errors.Is(err, entity.ErrCBTNotParticipant) {
			helper.ResponseOutput(c, int32(http.StatusForbidden), "Anda bukan peserta ujian ini.", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusOK, resp)
}
