package example

import (
	"github.com/gin-gonic/gin"
	example_usecase "github.com/ProjectWidyaprada/backend/core/usecase/example"
)

type ExampleHTTPHandler interface {
	GetExampleList(c *gin.Context)
	GetExampleDetail(c *gin.Context)
}

type exampleHTTPHandler struct {
	exampleUsecase example_usecase.ExampleUsecase
}

func NewExampleHTTPHandler(exampleUsecase example_usecase.ExampleUsecase) ExampleHTTPHandler {
	return &exampleHTTPHandler{exampleUsecase: exampleUsecase}
}
