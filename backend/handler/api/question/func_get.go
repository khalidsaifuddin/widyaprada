package question

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// GetQuestionDetail godoc
// @Summary      Get question detail
// @Description  Get question by ID with options
// @Tags         questions
// @Accept       json
// @Produce      json
// @Param        id path string true "Question ID"
// @Success      200 {object} entity.QuestionDetailResponse
// @Failure      401 {object} helper.ErrorResponse
// @Failure      404 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/questions/{id} [get]
func (h *questionHTTPHandler) GetQuestionDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "ID wajib diisi", nil)
		return
	}
	resp, err := h.bankSoalUsecase.Get(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, entity.ErrQuestionNotFound) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Soal tidak ditemukan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusOK, resp)
}
