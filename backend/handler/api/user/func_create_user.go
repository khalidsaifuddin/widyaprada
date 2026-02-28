package user

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Summary      Create user
// @Description  Create new user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        body body entity.CreateUserRequest true "Create user request"
// @Success      201 {object} entity.UserDetailResponse
// @Failure      400 {object} helper.ErrorResponse
// @Failure      401 {object} helper.ErrorResponse
// @Failure      409 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/users [post]
func (h *userHTTPHandler) CreateUser(c *gin.Context) {
	var req entity.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}

	actor := h.getActorContext(c)
	resp, err := h.userUsecase.Create(c.Request.Context(), req, actor)
	if err != nil {
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
		if errors.Is(err, entity.ErrUserNotFound) {
			helper.ResponseOutput(c, int32(http.StatusForbidden), "Tidak memiliki wewenang", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", err.Error())
		return
	}
	c.JSON(http.StatusCreated, resp)
}
