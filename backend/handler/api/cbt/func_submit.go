package cbt

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/auth"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// SubmitCBTExam godoc
// @Summary      Submit ujian
// @Description  Submit ujian dan hitung skor (PG/B-S otomatis)
// @Tags         cbt
// @Accept       json
// @Produce      json
// @Param        attemptId path string true "Attempt ID"
// @Success      200 {object} entity.CBTSubmitResponse
// @Failure      401 {object} helper.ErrorResponse
// @Failure      403 {object} helper.ErrorResponse
// @Failure      404 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/cbt/attempts/{attemptId}/submit [post]
func (h *cbtHTTPHandler) SubmitCBTExam(c *gin.Context) {
	claims := auth.GetClaimsFromContext(c)
	if claims == nil {
		helper.ResponseOutput(c, int32(http.StatusUnauthorized), "Unauthorized", nil)
		return
	}

	attemptID := c.Param("attemptId")
	if attemptID == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Attempt ID wajib diisi", nil)
		return
	}

	resp, err := h.cbtUsecase.SubmitUjian(c.Request.Context(), claims.UserID, attemptID)
	if err != nil {
		if errors.Is(err, entity.ErrCBTAttemptNotFound) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Attempt tidak ditemukan", nil)
			return
		}
		if errors.Is(err, entity.ErrCBTAttemptNotOwned) {
			helper.ResponseOutput(c, int32(http.StatusForbidden), "Akses ditolak", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusOK, resp)
}
