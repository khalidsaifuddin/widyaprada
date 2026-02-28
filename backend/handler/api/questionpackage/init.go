package questionpackage

import (
	paketsoal_usecase "github.com/ProjectWidyaprada/backend/core/usecase/paketsoal"
	"github.com/gin-gonic/gin"
)

type QuestionPackageHTTPHandler interface {
	GetPackageList(c *gin.Context)
	GetPackageDetail(c *gin.Context)
	CreatePackage(c *gin.Context)
	UpdatePackage(c *gin.Context)
	DeletePackage(c *gin.Context)
	VerifyPackage(c *gin.Context)
	UnverifyPackage(c *gin.Context)
}

type questionPackageHTTPHandler struct {
	paketSoalUsecase paketsoal_usecase.PaketSoalUsecase
}

func NewQuestionPackageHTTPHandler(paketSoalUsecase paketsoal_usecase.PaketSoalUsecase) QuestionPackageHTTPHandler {
	return &questionPackageHTTPHandler{paketSoalUsecase: paketSoalUsecase}
}
