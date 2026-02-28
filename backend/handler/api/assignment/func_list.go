package assignment

import (
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/auth"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// GetAssignmentList godoc
// @Summary      Tugas Saya
// @Description  Ujian yang user ikuti, filter status, sort batas_waktu
// @Tags         assignments
// @Accept       json
// @Produce      json
// @Param        status query string false "belum_dikerjakan | sudah_dikerjakan"
// @Param        sort_by query string false "jadwal_selesai" default(jadwal_selesai)
// @Param        sort_order query string false "asc | desc" default(asc)
// @Param        page query int false "Page" default(1)
// @Param        page_size query int false "Page size" default(10)
// @Success      200 {object} entity.GetAssignmentListResponse
// @Failure      401 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/assignments [get]
func (h *assignmentHTTPHandler) GetAssignmentList(c *gin.Context) {
	claims := auth.GetClaimsFromContext(c)
	if claims == nil {
		helper.ResponseOutput(c, int32(http.StatusUnauthorized), "Unauthorized", nil)
		return
	}

	var req entity.GetAssignmentListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}

	resp, err := h.assignmentUsecase.ListAssignments(c.Request.Context(), claims.UserID, req)
	if err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusOK, resp)
}
