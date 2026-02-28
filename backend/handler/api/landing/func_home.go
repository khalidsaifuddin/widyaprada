package landing

import (
	"net/http"

	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

func (h *landingHTTPHandler) GetHome(c *gin.Context) {
	resp, err := h.landingUsecase.GetHome(c.Request.Context())
	if err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}
