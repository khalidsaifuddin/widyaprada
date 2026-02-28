package questionpackage

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// CreatePackage godoc
// @Summary      Create question package
// @Tags         question-packages
// @Accept       json
// @Produce      json
// @Param        body body entity.CreatePackageRequest true "Create package request"
// @Success      201 {object} entity.PackageDetailResponse
// @Failure      400 {object} helper.ErrorResponse
// @Failure      401 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/question-packages [post]
func (h *questionPackageHTTPHandler) CreatePackage(c *gin.Context) {
	var req entity.CreatePackageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}
	resp, err := h.paketSoalUsecase.Create(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, entity.ErrPackageCodeExists) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Kode paket sudah digunakan", nil)
			return
		}
		if errors.Is(err, entity.ErrPackageMinOneQuestion) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Minimal 1 soal dalam paket", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusCreated, resp)
}
