package auth

import (
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// ForgotPassword godoc
// @Summary      Forgot Password
// @Description  Minta link reset password ke email
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body body entity.ForgotPasswordRequest true "Email"
// @Success      200 {object} entity.ForgotPasswordResponse
// @Failure      400 {object} helper.ErrorResponse
// @Router       /api/v1/auth/forgot-password [post]
func (h *authHTTPHandler) ForgotPassword(c *gin.Context) {
	var req entity.ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}

	resp, err := h.forgotPasswordUsecase.RequestReset(c.Request.Context(), req.Email)
	if err != nil {
		if err == entity.ErrInvalidEmailFormat {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Format email tidak valid", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}

	c.JSON(http.StatusOK, resp)
}
