package exam

import (
	exam_usecase "github.com/ProjectWidyaprada/backend/core/usecase/exam"
	"github.com/gin-gonic/gin"
)

type ExamHTTPHandler interface {
	GetExamList(c *gin.Context)
	GetExamDetail(c *gin.Context)
	CreateExam(c *gin.Context)
	UpdateExam(c *gin.Context)
	DeleteExam(c *gin.Context)
	PublishExam(c *gin.Context)
	VerifyExam(c *gin.Context)
	UnverifyExam(c *gin.Context)
}

type examHTTPHandler struct {
	examUsecase exam_usecase.ExamUsecase
}

func NewExamHTTPHandler(examUsecase exam_usecase.ExamUsecase) ExamHTTPHandler {
	return &examHTTPHandler{examUsecase: examUsecase}
}
