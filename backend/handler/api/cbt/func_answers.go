package cbt

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/auth"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// SaveCBTAnswer godoc
// @Summary      Simpan jawaban
// @Description  Simpan jawaban (option_id untuk PG/B-S, answer_text untuk Essay)
// @Tags         cbt
// @Accept       json
// @Produce      json
// @Param        attemptId path string true "Attempt ID"
// @Param        body body entity.CBTSaveAnswerRequest true "Jawaban"
// @Success      200 {object} entity.CBTSaveAnswerResponse
// @Failure      400 {object} helper.ErrorResponse
// @Failure      401 {object} helper.ErrorResponse
// @Failure      403 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/cbt/attempts/{attemptId}/answers [post]
func (h *cbtHTTPHandler) SaveCBTAnswer(c *gin.Context) {
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

	var req entity.CBTSaveAnswerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}

	err := h.cbtUsecase.SimpanJawaban(c.Request.Context(), claims.UserID, attemptID, req)
	if err != nil {
		if errors.Is(err, entity.ErrCBTAttemptNotFound) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Attempt tidak ditemukan", nil)
			return
		}
		if errors.Is(err, entity.ErrCBTAttemptNotOwned) {
			helper.ResponseOutput(c, int32(http.StatusForbidden), "Akses ditolak", nil)
			return
		}
		if errors.Is(err, entity.ErrCBTAlreadySubmitted) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Ujian sudah disubmit", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusOK, entity.CBTSaveAnswerResponse{Message: "Jawaban disimpan"})
}
