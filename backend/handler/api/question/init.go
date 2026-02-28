package question

import (
	banksoal_usecase "github.com/ProjectWidyaprada/backend/core/usecase/banksoal"
	"github.com/gin-gonic/gin"
)

type QuestionHTTPHandler interface {
	GetQuestionList(c *gin.Context)
	GetQuestionDetail(c *gin.Context)
	CreateQuestion(c *gin.Context)
	UpdateQuestion(c *gin.Context)
	DeleteQuestion(c *gin.Context)
	VerifyQuestion(c *gin.Context)
	UnverifyQuestion(c *gin.Context)
	GetCategories(c *gin.Context)
}

type questionHTTPHandler struct {
	bankSoalUsecase banksoal_usecase.BankSoalUsecase
}

func NewQuestionHTTPHandler(bankSoalUsecase banksoal_usecase.BankSoalUsecase) QuestionHTTPHandler {
	return &questionHTTPHandler{bankSoalUsecase: bankSoalUsecase}
}
