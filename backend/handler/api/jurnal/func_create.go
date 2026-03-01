package jurnal

import (
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/auth"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

func (h *jurnalHTTPHandler) Create(c *gin.Context) {
	claims := auth.GetClaimsFromContext(c)
	if claims == nil {
		helper.ResponseOutput(c, int32(http.StatusUnauthorized), "Unauthorized", nil)
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
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Format request tidak valid", err.Error())
		return
	}

	j := &entity.Jurnal{
		Title:    body.Title,
		Author:   body.Author,
		Abstract: body.Abstract,
		Content:  body.Content,
		PdfURL:   body.PdfURL,
		Category: body.Category,
		Year:     body.Year,
		Status:   entity.JurnalStatusDraft,
		UserID:   &claims.UserID,
	}
	if err := h.jurnalUsecase.Create(c.Request.Context(), j); err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Gagal menyimpan jurnal", nil)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": j.ID, "message": "Jurnal berhasil dibuat"})
}
