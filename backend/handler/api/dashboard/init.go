package dashboard

import (
	dashboard_usecase "github.com/ProjectWidyaprada/backend/core/usecase/dashboard"
	"github.com/gin-gonic/gin"
)

type DashboardHTTPHandler interface {
	GetAssignments(c *gin.Context)
	GetMyJournals(c *gin.Context)
}

type dashboardHTTPHandler struct {
	dashboardUsecase dashboard_usecase.DashboardUsecase
}

func NewDashboardHTTPHandler(dashboardUsecase dashboard_usecase.DashboardUsecase) DashboardHTTPHandler {
	return &dashboardHTTPHandler{dashboardUsecase: dashboardUsecase}
}
