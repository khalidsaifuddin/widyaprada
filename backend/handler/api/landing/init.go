package landing

import (
	landing_usecase "github.com/ProjectWidyaprada/backend/core/usecase/landing"
	"github.com/gin-gonic/gin"
)

type LandingHTTPHandler interface {
	GetHome(c *gin.Context)
}

type landingHTTPHandler struct {
	landingUsecase landing_usecase.LandingUsecase
}

func NewLandingHTTPHandler(landingUsecase landing_usecase.LandingUsecase) LandingHTTPHandler {
	return &landingHTTPHandler{landingUsecase: landingUsecase}
}
