package berita

import (
	berita_usecase "github.com/ProjectWidyaprada/backend/core/usecase/berita"
	"github.com/gin-gonic/gin"
)

type BeritaHTTPHandler interface {
	GetList(c *gin.Context)
	GetBySlug(c *gin.Context)
}

type beritaHTTPHandler struct {
	beritaUsecase berita_usecase.BeritaUsecase
}

func NewBeritaHTTPHandler(beritaUsecase berita_usecase.BeritaUsecase) BeritaHTTPHandler {
	return &beritaHTTPHandler{beritaUsecase: beritaUsecase}
}
