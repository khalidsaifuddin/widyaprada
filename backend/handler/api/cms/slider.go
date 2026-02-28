package cms

import (
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

func (h *cmsHTTPHandler) GetSlideList(c *gin.Context) {
	var req entity.GetSlideListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}
	resp, err := h.sliderUsecase.List(c.Request.Context(), req, h.getActor(c))
	if err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *cmsHTTPHandler) GetSlideDetail(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.sliderUsecase.Get(c.Request.Context(), id, h.getActor(c))
	if err != nil || resp == nil {
		helper.ResponseOutput(c, int32(http.StatusNotFound), "Slider tidak ditemukan", nil)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *cmsHTTPHandler) CreateSlide(c *gin.Context) {
	var req entity.CreateSlideRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}
	resp, err := h.sliderUsecase.Create(c.Request.Context(), req, h.getActor(c))
	if err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Gagal membuat slider", err.Error())
		return
	}
	c.JSON(http.StatusCreated, resp)
}

func (h *cmsHTTPHandler) UpdateSlide(c *gin.Context) {
	id := c.Param("id")
	var req entity.UpdateSlideRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}
	resp, err := h.sliderUsecase.Update(c.Request.Context(), id, req, h.getActor(c))
	if err != nil {
		if entity.IsRecordNotFound(err) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Slider tidak ditemukan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Gagal memperbarui slider", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *cmsHTTPHandler) DeleteSlide(c *gin.Context) {
	id := c.Param("id")
	if err := h.sliderUsecase.Delete(c.Request.Context(), id, h.getActor(c)); err != nil {
		if entity.IsRecordNotFound(err) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Slider tidak ditemukan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Gagal menghapus slider", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Slider berhasil dihapus"})
}
