package cbt

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/auth"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// GetCBTQuestions godoc
// @Summary      Daftar soal
// @Description  Daftar soal untuk attempt
// @Tags         cbt
// @Accept       json
// @Produce      json
// @Param        attemptId path string true "Attempt ID"
// @Success      200 {object} entity.CBTListQuestionsResponse
// @Failure      401 {object} helper.ErrorResponse
// @Failure      403 {object} helper.ErrorResponse
// @Failure      404 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/cbt/attempts/{attemptId}/questions [get]
func (h *cbtHTTPHandler) GetCBTQuestions(c *gin.Context) {
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

	resp, err := h.cbtUsecase.GetSoal(c.Request.Context(), claims.UserID, attemptID)
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
	c.JSON(http.StatusOK, resp)
}

// GetCBTQuestionByNum godoc
// @Summary      Soal per nomor
// @Description  Get soal by nomor (1-based)
// @Tags         cbt
// @Accept       json
// @Produce      json
// @Param        attemptId path string true "Attempt ID"
// @Param        num path int true "Nomor soal"
// @Success      200 {object} entity.CBTQuestionItem
// @Failure      401 {object} helper.ErrorResponse
// @Failure      403 {object} helper.ErrorResponse
// @Failure      404 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/cbt/attempts/{attemptId}/questions/{num} [get]
func (h *cbtHTTPHandler) GetCBTQuestionByNum(c *gin.Context) {
	claims := auth.GetClaimsFromContext(c)
	if claims == nil {
		helper.ResponseOutput(c, int32(http.StatusUnauthorized), "Unauthorized", nil)
		return
	}

	attemptID := c.Param("attemptId")
	numStr := c.Param("num")
	if attemptID == "" || numStr == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Attempt ID dan nomor wajib diisi", nil)
		return
	}
	num, err := strconv.Atoi(numStr)
	if err != nil || num < 1 {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Nomor soal tidak valid", nil)
		return
	}

	resp, err := h.cbtUsecase.GetSoalByNomor(c.Request.Context(), claims.UserID, attemptID, num)
	if err != nil {
		if errors.Is(err, entity.ErrCBTAttemptNotFound) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Attempt atau soal tidak ditemukan", nil)
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
	c.JSON(http.StatusOK, resp)
}
