package assignment

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/auth"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// GetAssignmentResult godoc
// @Summary      Hasil ujian user
// @Description  Hasil ujian untuk exam tersebut (hanya peserta)
// @Tags         assignments
// @Accept       json
// @Produce      json
// @Param        examId path string true "Exam ID"
// @Success      200 {object} entity.AssignmentResultResponse
// @Failure      401 {object} helper.ErrorResponse
// @Failure      403 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/assignments/{examId}/result [get]
func (h *assignmentHTTPHandler) GetAssignmentResult(c *gin.Context) {
	claims := auth.GetClaimsFromContext(c)
	if claims == nil {
		helper.ResponseOutput(c, int32(http.StatusUnauthorized), "Unauthorized", nil)
		return
	}

	examID := c.Param("examId")
	if examID == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "examId wajib diisi", nil)
		return
	}

	resp, err := h.assignmentUsecase.GetAssignmentResult(c.Request.Context(), claims.UserID, examID)
	if err != nil {
		if errors.Is(err, entity.ErrAssignmentForbidden) {
			helper.ResponseOutput(c, int32(http.StatusForbidden), "Akses ditolak", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusOK, resp)
}
