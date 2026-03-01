package upload

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/ProjectWidyaprada/backend/config"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	jurnalSubDir   = "jurnal"
	maxPDFSize     = 15 << 20 // 15MB
	allowedPDFType = "application/pdf"
)

// UploadJournalPDF handles PDF file upload for journal
func (h *uploadHTTPHandler) UploadJournalPDF(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "File wajib diisi", nil)
		return
	}

	// Validate file type
	ct := file.Header.Get("Content-Type")
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ct != allowedPDFType && ext != ".pdf" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Hanya file PDF yang diizinkan", nil)
		return
	}

	if file.Size > maxPDFSize {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Ukuran file maksimal 15MB", nil)
		return
	}

	// Ensure upload directory exists (path relatif dari working dir backend)
	cfg := config.Get()
	uploadDir := cfg.UploadDir
	if uploadDir == "" {
		uploadDir = "uploads"
	}
	dir := filepath.Join(uploadDir, jurnalSubDir)
	if err := os.MkdirAll(dir, 0755); err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Gagal membuat direktori upload", nil)
		return
	}

	// Save with unique filename
	filename := uuid.New().String() + ".pdf"
	savePath := filepath.Join(dir, filename)
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Gagal menyimpan file", nil)
		return
	}

	// Return path relatif (untuk akses via Static /uploads)
	urlPath := "/uploads/" + jurnalSubDir + "/" + filename
	c.JSON(http.StatusOK, gin.H{
		"url":     urlPath,
		"path":    urlPath,
		"message": "File berhasil diunggah",
	})
}
