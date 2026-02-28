package jurnal

import (
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

func (h *jurnalHTTPHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.jurnalUsecase.GetByID(c.Request.Context(), id)
	if err != nil || resp == nil {
		if entity.IsRecordNotFound(err) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Jurnal tidak ditemukan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}
