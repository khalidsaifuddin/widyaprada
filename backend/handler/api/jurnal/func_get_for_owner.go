package jurnal

import (
	"net/http"

	"github.com/ProjectWidyaprada/backend/pkg/auth"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

func (h *jurnalHTTPHandler) GetByIDForOwner(c *gin.Context) {
	claims := auth.GetClaimsFromContext(c)
	if claims == nil {
		helper.ResponseOutput(c, int32(http.StatusUnauthorized), "Unauthorized", nil)
		return
	}

	id := c.Param("id")
	if id == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "ID jurnal wajib diisi", nil)
		return
	}

	j, err := h.jurnalUsecase.GetByIDForOwner(c.Request.Context(), id, claims.UserID)
	if err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Gagal memuat jurnal", nil)
		return
	}
	if j == nil {
		helper.ResponseOutput(c, int32(http.StatusNotFound), "Jurnal tidak ditemukan", nil)
		return
	}
	c.JSON(http.StatusOK, j)
}
