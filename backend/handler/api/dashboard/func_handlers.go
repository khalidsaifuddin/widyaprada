package dashboard

import (
	"net/http"

	"github.com/ProjectWidyaprada/backend/pkg/auth"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

func (h *dashboardHTTPHandler) GetAssignments(c *gin.Context) {
	claims := auth.GetClaimsFromContext(c)
	if claims == nil {
		helper.ResponseOutput(c, int32(http.StatusUnauthorized), "Unauthorized", nil)
		return
	}

	limit := int64(10)
	if l := c.Query("limit"); l != "" {
		if n, err := parseInt64(l); err == nil && n > 0 && n <= 100 {
			limit = n
		}
	}
	page := int64(1)
	if p := c.Query("page"); p != "" {
		if n, err := parseInt64(p); err == nil && n > 0 {
			page = n
		}
	}

	resp, err := h.dashboardUsecase.GetAssignments(c.Request.Context(), claims.UserID, limit, page)
	if err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *dashboardHTTPHandler) GetMyJournals(c *gin.Context) {
	claims := auth.GetClaimsFromContext(c)
	if claims == nil {
		helper.ResponseOutput(c, int32(http.StatusUnauthorized), "Unauthorized", nil)
		return
	}

	limit := int64(10)
	if l := c.Query("limit"); l != "" {
		if n, err := parseInt64(l); err == nil && n > 0 && n <= 100 {
			limit = n
		}
	}
	page := int64(1)
	if p := c.Query("page"); p != "" {
		if n, err := parseInt64(p); err == nil && n > 0 {
			page = n
		}
	}

	resp, err := h.dashboardUsecase.GetMyJournals(c.Request.Context(), claims.UserID, limit, page)
	if err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}
