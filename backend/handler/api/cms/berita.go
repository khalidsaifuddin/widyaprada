package cms

import (
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

func (h *cmsHTTPHandler) GetArticleList(c *gin.Context) {
	var req entity.GetArticleListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}
	resp, err := h.beritaUsecase.List(c.Request.Context(), req, h.getActor(c))
	if err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *cmsHTTPHandler) GetArticleDetail(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.beritaUsecase.Get(c.Request.Context(), id, h.getActor(c))
	if err != nil || resp == nil {
		helper.ResponseOutput(c, int32(http.StatusNotFound), "Berita tidak ditemukan", nil)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *cmsHTTPHandler) CreateArticle(c *gin.Context) {
	var req entity.CreateArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}
	resp, err := h.beritaUsecase.Create(c.Request.Context(), req, h.getActor(c))
	if err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Gagal membuat berita", err.Error())
		return
	}
	c.JSON(http.StatusCreated, resp)
}

func (h *cmsHTTPHandler) UpdateArticle(c *gin.Context) {
	id := c.Param("id")
	var req entity.UpdateArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}
	resp, err := h.beritaUsecase.Update(c.Request.Context(), id, req, h.getActor(c))
	if err != nil {
		if entity.IsRecordNotFound(err) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Berita tidak ditemukan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Gagal memperbarui berita", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *cmsHTTPHandler) DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	if err := h.beritaUsecase.Delete(c.Request.Context(), id, h.getActor(c)); err != nil {
		if entity.IsRecordNotFound(err) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Berita tidak ditemukan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Gagal menghapus berita", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Berita berhasil dihapus"})
}
