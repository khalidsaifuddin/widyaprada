package assignment

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/auth"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// ApplyUjikom godoc
// @Summary      Apply pendaftaran Ujikom
// @Description  Apply pendaftaran dengan dokumen persyaratan
// @Tags         ujikom
// @Accept       json
// @Produce      json
// @Param        body body entity.ApplyUjikomRequest true "Apply request"
// @Success      201 {object} map[string]string
// @Failure      400 {object} helper.ErrorResponse
// @Failure      401 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/ujikom/apply [post]
func (h *assignmentHTTPHandler) ApplyUjikom(c *gin.Context) {
	claims := auth.GetClaimsFromContext(c)
	if claims == nil {
		helper.ResponseOutput(c, int32(http.StatusUnauthorized), "Unauthorized", nil)
		return
	}

	// Expect: { "jenis_ujikom": "...", "documents": [ { "document_type": "...", "file_path": "..." | "portofolio_text": "..." } ] }
	var body struct {
		JenisUjikom string                         `json:"jenis_ujikom" binding:"required"`
		Documents   []entity.ApplyUjikomDocumentInput `json:"documents" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}

	req := entity.ApplyUjikomRequest{JenisUjikom: body.JenisUjikom}
	err := h.assignmentUsecase.ApplyUjikom(c.Request.Context(), claims.UserID, req, body.Documents)
	if err != nil {
		if errors.Is(err, entity.ErrApplyAlreadyExists) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Anda sudah pernah apply untuk jenis ujikom ini", nil)
			return
		}
		if errors.Is(err, entity.ErrDokumenPersyaratanReq) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Dokumen persyaratan wajib belum lengkap", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Pendaftaran berhasil dikirim"})
}
