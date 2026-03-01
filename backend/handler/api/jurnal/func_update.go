package jurnal

import (
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/auth"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

func (h *jurnalHTTPHandler) Update(c *gin.Context) {
	claims := auth.GetClaimsFromContext(c)
	if claims == nil {
		helper.ResponseOutput(c, int32(http.StatusUnauthorized), "Unauthorized", nil)
		return
	}

	id := c.Param("id")
	if id == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "ID jurnal wajib diisi", nil)
		return
	}

	// Get existing journal and verify ownership
	existing, err := h.jurnalUsecase.GetByIDForOwner(c.Request.Context(), id, claims.UserID)
	if err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Gagal memuat jurnal", nil)
		return
	}
	if existing == nil {
		helper.ResponseOutput(c, int32(http.StatusNotFound), "Jurnal tidak ditemukan", nil)
		return
	}
	if existing.UserID == nil || *existing.UserID != claims.UserID {
		helper.ResponseOutput(c, int32(http.StatusForbidden), "Anda tidak memiliki akses untuk mengubah jurnal ini", nil)
		return
	}

	var body struct {
		Title    string `json:"title" binding:"required"`
		Author   string `json:"author"`
		Abstract string `json:"abstract"`
		Content  string `json:"content"`
		PdfURL   string `json:"pdf_url"`
		Category string `json:"category"`
		Year     int    `json:"year"`
		Status   string `json:"status"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Format request tidak valid", err.Error())
		return
	}

	j := &entity.Jurnal{
		ID:       id,
		Title:    body.Title,
		Author:   body.Author,
		Abstract: body.Abstract,
		Content:  body.Content,
		PdfURL:   body.PdfURL,
		Category: body.Category,
		Year:     body.Year,
		Status:   body.Status,
		UserID:   existing.UserID,
	}
	if j.Status == "" {
		j.Status = existing.Status
	}
	if err := h.jurnalUsecase.Update(c.Request.Context(), j); err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Gagal menyimpan jurnal", nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Jurnal berhasil diperbarui"})
}
