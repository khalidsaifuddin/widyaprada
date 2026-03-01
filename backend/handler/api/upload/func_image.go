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
	imagesSubDir  = "images"
	maxImageSize  = 5 << 20 // 5MB
)

// UploadImage handles image upload for slider & berita (CMS)
func (h *uploadHTTPHandler) UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "File wajib diisi", nil)
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowed := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true, ".webp": true}
	if !allowed[ext] {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Hanya gambar (jpg, png, gif, webp) yang diizinkan", nil)
		return
	}

	ct := file.Header.Get("Content-Type")
	if !strings.HasPrefix(ct, "image/") {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "File harus berupa gambar", nil)
		return
	}

	if file.Size > maxImageSize {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Ukuran file maksimal 5MB", nil)
		return
	}

	cfg := config.Get()
	uploadDir := cfg.UploadDir
	if uploadDir == "" {
		uploadDir = "uploads"
	}
	dir := filepath.Join(uploadDir, imagesSubDir)
	if err := os.MkdirAll(dir, 0755); err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Gagal membuat direktori upload", nil)
		return
	}

	filename := uuid.New().String() + ext
	savePath := filepath.Join(dir, filename)
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Gagal menyimpan file", nil)
		return
	}

	urlPath := "/uploads/" + imagesSubDir + "/" + filename
	c.JSON(http.StatusOK, gin.H{
		"url":     urlPath,
		"path":    urlPath,
		"message": "Gambar berhasil diunggah",
	})
}
