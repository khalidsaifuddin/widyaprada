package beranda

import (
	beranda_usecase "github.com/ProjectWidyaprada/backend/core/usecase/beranda"
	"github.com/gin-gonic/gin"
)

type BerandaHTTPHandler interface {
	GetPengumuman(c *gin.Context)
}

type berandaHTTPHandler struct {
	berandaUsecase beranda_usecase.BerandaUsecase
}

func NewBerandaHTTPHandler(berandaUsecase beranda_usecase.BerandaUsecase) BerandaHTTPHandler {
	return &berandaHTTPHandler{berandaUsecase: berandaUsecase}
}
