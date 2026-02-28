package cbt

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/auth"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// StartCBTExam godoc
// @Summary      Mulai ujian
// @Description  Create attempt dan return attempt_id
// @Tags         cbt
// @Accept       json
// @Produce      json
// @Param        id path string true "Exam ID"
// @Success      200 {object} entity.CBTStartResponse
// @Failure      400 {object} helper.ErrorResponse
// @Failure      401 {object} helper.ErrorResponse
// @Failure      403 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/cbt/exams/{id}/start [post]
func (h *cbtHTTPHandler) StartCBTExam(c *gin.Context) {
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

	resp, err := h.cbtUsecase.MulaiUjian(c.Request.Context(), claims.UserID, examID)
	if err != nil {
		if errors.Is(err, entity.ErrCBTNotParticipant) {
			helper.ResponseOutput(c, int32(http.StatusForbidden), "Anda bukan peserta ujian ini", nil)
			return
		}
		if errors.Is(err, entity.ErrCBTAlreadyStarted) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Sudah pernah mulai ujian ini", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusOK, resp)
}
