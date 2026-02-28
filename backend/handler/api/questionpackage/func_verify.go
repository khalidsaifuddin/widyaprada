package questionpackage

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// VerifyPackage godoc
// @Summary      Verify question package
// @Description  Mark package as verified (Verifikator or Super Admin)
// @Tags         question-packages
// @Accept       json
// @Produce      json
// @Param        id path string true "Package ID"
// @Success      200 {object} map[string]string
// @Failure      401 {object} helper.ErrorResponse
// @Failure      403 {object} helper.ErrorResponse
// @Failure      404 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/question-packages/{id}/verify [post]
func (h *questionPackageHTTPHandler) VerifyPackage(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "ID wajib diisi", nil)
		return
	}
	err := h.paketSoalUsecase.Verify(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, entity.ErrPackageNotFound) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Paket soal tidak ditemukan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Paket soal berhasil diverifikasi"})
}

// UnverifyPackage godoc
// @Summary      Unverify question package
// @Description  Cancel verification of package (Verifikator or Super Admin)
// @Tags         question-packages
// @Accept       json
// @Produce      json
// @Param        id path string true "Package ID"
// @Success      200 {object} map[string]string
// @Failure      401 {object} helper.ErrorResponse
// @Failure      403 {object} helper.ErrorResponse
// @Failure      404 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/question-packages/{id}/unverify [post]
func (h *questionPackageHTTPHandler) UnverifyPackage(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "ID wajib diisi", nil)
		return
	}
	err := h.paketSoalUsecase.Unverify(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, entity.ErrPackageNotFound) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Paket soal tidak ditemukan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Verifikasi paket dibatalkan"})
}
