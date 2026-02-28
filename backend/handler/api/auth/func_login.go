package auth

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary      Login
// @Description  Login dengan identifier (email/username) dan password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body body entity.LoginRequest true "Login request"
// @Success      200 {object} entity.LoginResponse
// @Failure      400 {object} helper.ErrorResponse
// @Failure      401 {object} helper.ErrorResponse
// @Router       /api/v1/auth/login [post]
func (h *authHTTPHandler) Login(c *gin.Context) {
	var req entity.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}

	resp, err := h.authUsecase.Login(c.Request.Context(), req)
	if err != nil {
		// Pesan aman: tidak bocorkan apakah email terdaftar
		if errors.Is(err, entity.ErrInvalidCredentials) {
			helper.ResponseOutput(c, int32(http.StatusUnauthorized), "Email/username atau kata sandi salah", nil)
			return
		}
		if errors.Is(err, entity.ErrAccountInactive) {
			helper.ResponseOutput(c, int32(http.StatusUnauthorized), "Akun ini tidak aktif. Silakan hubungi administrator.", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}

	c.JSON(http.StatusOK, resp)
}
