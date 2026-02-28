package cbt

import (
	"net/http"

	"github.com/ProjectWidyaprada/backend/pkg/auth"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// GetCBTExams godoc
// @Summary      Daftar ujian tersedia
// @Description  Ujian yang bisa dikerjakan (Diterbitkan, jadwal aktif, user peserta)
// @Tags         cbt
// @Accept       json
// @Produce      json
// @Success      200 {object} entity.CBTListExamsResponse
// @Failure      401 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/cbt/exams [get]
func (h *cbtHTTPHandler) GetCBTExams(c *gin.Context) {
	claims := auth.GetClaimsFromContext(c)
	if claims == nil {
		helper.ResponseOutput(c, int32(http.StatusUnauthorized), "Unauthorized", nil)
		return
	}

	resp, err := h.cbtUsecase.ListUjianTersedia(c.Request.Context(), claims.UserID)
	if err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusOK, resp)
}
