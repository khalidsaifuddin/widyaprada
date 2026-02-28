package auth

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// Register godoc
// @Summary      Register
// @Description  Registrasi mandiri calon peserta (nama, email, nip opsional). Password dikirim ke email.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body body entity.RegisterRequest true "Register request"
// @Success      201 {object} entity.RegisterResponse
// @Failure      400 {object} helper.ErrorResponse
// @Router       /api/v1/auth/register [post]
func (h *authHTTPHandler) Register(c *gin.Context) {
	var req entity.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}

	resp, err := h.registrationUsecase.Register(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, entity.ErrEmailAlreadyRegistered) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Email ini sudah terdaftar. Gunakan Lupa Password jika Anda lupa kata sandi.", nil)
			return
		}
		if errors.Is(err, entity.ErrInvalidEmailFormat) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Format email tidak valid", nil)
			return
		}
		if errors.Is(err, entity.ErrValidation) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Nama dan email wajib diisi", nil)
			return
		}
		if errors.Is(err, entity.ErrRoleNotFound) {
			helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Konfigurasi role tidak ditemukan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}

	c.JSON(http.StatusCreated, resp)
}
