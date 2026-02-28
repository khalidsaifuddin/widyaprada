package cbt

import (
	"net/http"

	"github.com/ProjectWidyaprada/backend/pkg/auth"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// GetCBTHistory godoc
// @Summary      Riwayat ujian
// @Description  Riwayat ujian + nilai user
// @Tags         cbt
// @Accept       json
// @Produce      json
// @Success      200 {object} entity.CBTHistoryResponse
// @Failure      401 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/cbt/history [get]
func (h *cbtHTTPHandler) GetCBTHistory(c *gin.Context) {
	claims := auth.GetClaimsFromContext(c)
	if claims == nil {
		helper.ResponseOutput(c, int32(http.StatusUnauthorized), "Unauthorized", nil)
		return
	}

	resp, err := h.cbtUsecase.GetRiwayatHasil(c.Request.Context(), claims.UserID)
	if err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusOK, resp)
}
