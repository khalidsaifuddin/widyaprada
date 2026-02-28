package questionpackage

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// GetPackageDetail godoc
// @Summary      Get package detail
// @Description  Get question package by ID with list of questions
// @Tags         question-packages
// @Accept       json
// @Produce      json
// @Param        id path string true "Package ID"
// @Success      200 {object} entity.PackageDetailResponse
// @Failure      401 {object} helper.ErrorResponse
// @Failure      404 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/question-packages/{id} [get]
func (h *questionPackageHTTPHandler) GetPackageDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "ID wajib diisi", nil)
		return
	}
	resp, err := h.paketSoalUsecase.Get(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, entity.ErrPackageNotFound) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Paket soal tidak ditemukan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusOK, resp)
}
