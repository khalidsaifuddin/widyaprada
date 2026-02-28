package user

import (
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// GetUserDetail godoc
// @Summary      Get user detail
// @Description  Get user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id path string true "User ID"
// @Success      200 {object} entity.UserDetailResponse
// @Failure      401 {object} helper.ErrorResponse
// @Failure      404 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/users/{id} [get]
func (h *userHTTPHandler) GetUserDetail(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "User ID required", nil)
		return
	}

	actor := h.getActorContext(c)
	resp, err := h.userUsecase.Get(c.Request.Context(), userID, actor)
	if err != nil {
		if err == entity.ErrUserNotFound {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Pengguna tidak ditemukan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}
