package wpdata

import (
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

func (h *wpdataHTTPHandler) GetCalonPesertaList(c *gin.Context) {
	var req entity.GetCalonPesertaListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}
	resp, err := h.calonPesertaUsecase.List(c.Request.Context(), req, h.getActor(c))
	if err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *wpdataHTTPHandler) GetCalonPesertaDetail(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.calonPesertaUsecase.Get(c.Request.Context(), id, h.getActor(c))
	if err != nil || resp == nil {
		if entity.IsRecordNotFound(err) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Calon peserta tidak ditemukan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *wpdataHTTPHandler) VerifyCalonPeserta(c *gin.Context) {
	id := c.Param("id")
	var req entity.VerifyCalonPesertaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}
	if err := h.calonPesertaUsecase.Verify(c.Request.Context(), id, req, h.getActor(c)); err != nil {
		if entity.IsRecordNotFound(err) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Calon peserta tidak ditemukan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Gagal memverifikasi", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Verifikasi berhasil"})
}
