package assignment

import (
	"net/http"

	"github.com/ProjectWidyaprada/backend/pkg/auth"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// GetApplyStatus godoc
// @Summary      Status pendaftaran user
// @Description  Menunggu verifikasi / lolos / tidak lolos
// @Tags         ujikom
// @Accept       json
// @Produce      json
// @Success      200 {object} entity.ApplyStatusResponse
// @Failure      401 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/ujikom/apply/status [get]
func (h *assignmentHTTPHandler) GetApplyStatus(c *gin.Context) {
	claims := auth.GetClaimsFromContext(c)
	if claims == nil {
		helper.ResponseOutput(c, int32(http.StatusUnauthorized), "Unauthorized", nil)
		return
	}

	resp, err := h.assignmentUsecase.GetApplyStatus(c.Request.Context(), claims.UserID)
	if err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	if resp == nil {
		c.JSON(http.StatusOK, gin.H{"status": "", "message": "Belum pernah apply"})
		return
	}
	c.JSON(http.StatusOK, resp)
}
