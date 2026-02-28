package user

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// UpdateUser godoc
// @Summary      Update user
// @Description  Update user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id path string true "User ID"
// @Param        body body entity.UpdateUserRequest true "Update user request"
// @Success      200 {object} entity.UserDetailResponse
// @Failure      400 {object} helper.ErrorResponse
// @Failure      401 {object} helper.ErrorResponse
// @Failure      404 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/users/{id} [put]
func (h *userHTTPHandler) UpdateUser(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "User ID required", nil)
		return
	}

	var req entity.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}

	actor := h.getActorContext(c)
	resp, err := h.userUsecase.Update(c.Request.Context(), userID, req, actor)
	if err != nil {
		if errors.Is(err, entity.ErrUserNotFound) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Pengguna tidak ditemukan", nil)
			return
		}
		if errors.Is(err, entity.ErrEmailExists) {
			helper.ResponseOutput(c, int32(http.StatusConflict), "Email sudah digunakan", nil)
			return
		}
		if errors.Is(err, entity.ErrUsernameExists) {
			helper.ResponseOutput(c, int32(http.StatusConflict), "Username sudah digunakan", nil)
			return
		}
		if errors.Is(err, entity.ErrAtLeastOneRole) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Minimal satu role wajib dipilih", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}
