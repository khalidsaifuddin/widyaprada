package upload

import (
	"github.com/gin-gonic/gin"
)

type UploadHTTPHandler interface {
	UploadJournalPDF(c *gin.Context)
	UploadImage(c *gin.Context)
}

type uploadHTTPHandler struct{}

func NewUploadHTTPHandler() UploadHTTPHandler {
	return &uploadHTTPHandler{}
}
