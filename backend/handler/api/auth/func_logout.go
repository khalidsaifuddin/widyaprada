package auth

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/auth"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// Logout godoc
// @Summary      Logout
// @Description  Keluar dari aplikasi, invalidasi token. Requires Authorization: Bearer &lt;token&gt;
// @Tags         auth
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} entity.LogoutResponse
// @Failure      401 {object} helper.ErrorResponse
// @Router       /api/v1/auth/logout [post]
func (h *authHTTPHandler) Logout(c *gin.Context) {
	tokenString := auth.GetTokenFromContext(c)
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Token tidak ditemukan"})
		return
	}

	resp, err := h.authUsecase.Logout(c.Request.Context(), tokenString)
	if err != nil {
		if errors.Is(err, entity.ErrInvalidToken) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token tidak valid"})
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}

	c.JSON(http.StatusOK, resp)
}
