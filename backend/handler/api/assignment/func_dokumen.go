package assignment

import (
	"net/http"

	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// GetDokumenPersyaratan godoc
// @Summary      List dokumen persyaratan
// @Description  Daftar dokumen persyaratan sesuai jenis ujikom
// @Tags         ujikom
// @Accept       json
// @Produce      json
// @Param        jenis_ujikom query string true "perpindahan_jabatan | kenaikan_tingkat"
// @Success      200 {object} entity.ListDokumenPersyaratanResponse
// @Failure      401 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/ujikom/dokumen-persyaratan [get]
func (h *assignmentHTTPHandler) GetDokumenPersyaratan(c *gin.Context) {
	jenisUjikom := c.Query("jenis_ujikom")
	if jenisUjikom == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "jenis_ujikom wajib diisi", nil)
		return
	}

	resp, err := h.assignmentUsecase.ListDokumenPersyaratan(c.Request.Context(), jenisUjikom)
	if err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusOK, resp)
}
