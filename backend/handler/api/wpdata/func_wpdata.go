package wpdata

import (
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

func (h *wpdataHTTPHandler) GetWPDataList(c *gin.Context) {
	var req entity.GetWPDataListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}
	resp, err := h.wpdataUsecase.List(c.Request.Context(), req, h.getActor(c))
	if err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *wpdataHTTPHandler) GetWPDataDetail(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.wpdataUsecase.Get(c.Request.Context(), id, h.getActor(c))
	if err != nil || resp == nil {
		if entity.IsRecordNotFound(err) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Data tidak ditemukan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *wpdataHTTPHandler) CreateWPData(c *gin.Context) {
	var req entity.CreateWPDataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}
	resp, err := h.wpdataUsecase.Create(c.Request.Context(), req, h.getActor(c))
	if err != nil {
		if err == entity.ErrNIPExists {
			helper.ResponseOutput(c, int32(http.StatusConflict), "NIP sudah digunakan", err.Error())
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Gagal membuat data", err.Error())
		return
	}
	c.JSON(http.StatusCreated, resp)
}

func (h *wpdataHTTPHandler) UpdateWPData(c *gin.Context) {
	id := c.Param("id")
	var req entity.UpdateWPDataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}
	resp, err := h.wpdataUsecase.Update(c.Request.Context(), id, req, h.getActor(c))
	if err != nil {
		if entity.IsRecordNotFound(err) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Data tidak ditemukan", nil)
			return
		}
		if err == entity.ErrNIPExists {
			helper.ResponseOutput(c, int32(http.StatusConflict), "NIP sudah digunakan", err.Error())
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Gagal memperbarui data", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *wpdataHTTPHandler) DeleteWPData(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Reason string `json:"reason" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Alasan penghapusan wajib diisi", err.Error())
		return
	}
	if err := h.wpdataUsecase.Delete(c.Request.Context(), id, body.Reason, h.getActor(c)); err != nil {
		if entity.IsRecordNotFound(err) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Data tidak ditemukan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Gagal menghapus data", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
