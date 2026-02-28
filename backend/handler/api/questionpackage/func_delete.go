package questionpackage

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// DeletePackage godoc
// @Summary      Delete question package
// @Description  Soft delete package by ID, reason required
// @Tags         question-packages
// @Accept       json
// @Produce      json
// @Param        id path string true "Package ID"
// @Param        body body entity.DeletePackageRequest true "Delete reason"
// @Success      200 {object} map[string]string
// @Failure      400 {object} helper.ErrorResponse
// @Failure      401 {object} helper.ErrorResponse
// @Failure      404 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/question-packages/{id} [delete]
func (h *questionPackageHTTPHandler) DeletePackage(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "ID wajib diisi", nil)
		return
	}
	var req entity.DeletePackageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Alasan penghapusan wajib diisi", err.Error())
		return
	}
	err := h.paketSoalUsecase.Delete(c.Request.Context(), id, req.Reason)
	if err != nil {
		if errors.Is(err, entity.ErrPackageNotFound) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Paket soal tidak ditemukan", nil)
			return
		}
		if errors.Is(err, entity.ErrPackageDeleteReason) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Alasan penghapusan wajib diisi", nil)
			return
		}
		if errors.Is(err, entity.ErrPackageInUseByExam) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Paket ini digunakan oleh ujian", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Paket soal berhasil dihapus"})
}
