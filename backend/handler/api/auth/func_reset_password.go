package auth

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// ResetPassword godoc
// @Summary      Reset Password
// @Description  Atur ulang kata sandi dengan token dari email
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body body entity.ResetPasswordRequest true "Token, password, confirm"
// @Success      200 {object} entity.ResetPasswordResponse
// @Failure      400 {object} helper.ErrorResponse
// @Router       /api/v1/auth/reset-password [post]
func (h *authHTTPHandler) ResetPassword(c *gin.Context) {
	var req entity.ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}

	resp, err := h.forgotPasswordUsecase.ResetPassword(c.Request.Context(), req.Token, req.Password, req.PasswordConfirm)
	if err != nil {
		if errors.Is(err, entity.ErrResetTokenInvalid) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Token tidak valid atau kadaluarsa", nil)
			return
		}
		if errors.Is(err, entity.ErrPasswordTooShort) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Kata sandi minimal 8 karakter", nil)
			return
		}
		if errors.Is(err, entity.ErrPasswordMismatch) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Konfirmasi kata sandi tidak cocok", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}

	c.JSON(http.StatusOK, resp)
}
