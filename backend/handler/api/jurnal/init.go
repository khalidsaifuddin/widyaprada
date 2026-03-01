package jurnal

import (
	jurnal_usecase "github.com/ProjectWidyaprada/backend/core/usecase/jurnal"
	"github.com/gin-gonic/gin"
)

type JurnalHTTPHandler interface {
	GetList(c *gin.Context)
	GetByID(c *gin.Context)
	GetByIDForOwner(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
}

type jurnalHTTPHandler struct {
	jurnalUsecase jurnal_usecase.JurnalUsecase
}

func NewJurnalHTTPHandler(jurnalUsecase jurnal_usecase.JurnalUsecase) JurnalHTTPHandler {
	return &jurnalHTTPHandler{jurnalUsecase: jurnalUsecase}
}
