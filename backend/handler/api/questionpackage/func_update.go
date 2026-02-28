package questionpackage

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// UpdatePackage godoc
// @Summary      Update question package
// @Description  Update question package by ID
// @Tags         question-packages
// @Accept       json
// @Produce      json
// @Param        id path string true "Package ID"
// @Param        body body entity.UpdatePackageRequest true "Update package request"
// @Success      200 {object} entity.PackageDetailResponse
// @Failure      400 {object} helper.ErrorResponse
// @Failure      401 {object} helper.ErrorResponse
// @Failure      404 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/question-packages/{id} [put]
func (h *questionPackageHTTPHandler) UpdatePackage(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "ID wajib diisi", nil)
		return
	}
	var req entity.UpdatePackageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}
	resp, err := h.paketSoalUsecase.Update(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, entity.ErrPackageNotFound) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Paket soal tidak ditemukan", nil)
			return
		}
		if errors.Is(err, entity.ErrPackageCodeExists) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Kode paket sudah digunakan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusOK, resp)
}
