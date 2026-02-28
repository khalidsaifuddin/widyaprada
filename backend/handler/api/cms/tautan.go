package cms

import (
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

func (h *cmsHTTPHandler) GetLinkList(c *gin.Context) {
	var req entity.GetLinkListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}
	resp, err := h.tautanUsecase.List(c.Request.Context(), req, h.getActor(c))
	if err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *cmsHTTPHandler) GetLinkDetail(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.tautanUsecase.Get(c.Request.Context(), id, h.getActor(c))
	if err != nil || resp == nil {
		helper.ResponseOutput(c, int32(http.StatusNotFound), "Tautan tidak ditemukan", nil)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *cmsHTTPHandler) CreateLink(c *gin.Context) {
	var req entity.CreateLinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}
	resp, err := h.tautanUsecase.Create(c.Request.Context(), req, h.getActor(c))
	if err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Gagal membuat tautan", err.Error())
		return
	}
	c.JSON(http.StatusCreated, resp)
}

func (h *cmsHTTPHandler) UpdateLink(c *gin.Context) {
	id := c.Param("id")
	var req entity.UpdateLinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}
	resp, err := h.tautanUsecase.Update(c.Request.Context(), id, req, h.getActor(c))
	if err != nil {
		if entity.IsRecordNotFound(err) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Tautan tidak ditemukan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Gagal memperbarui tautan", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *cmsHTTPHandler) DeleteLink(c *gin.Context) {
	id := c.Param("id")
	if err := h.tautanUsecase.Delete(c.Request.Context(), id, h.getActor(c)); err != nil {
		if entity.IsRecordNotFound(err) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Tautan tidak ditemukan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Gagal menghapus tautan", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tautan berhasil dihapus"})
}
