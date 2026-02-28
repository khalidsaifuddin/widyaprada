package user

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// DeleteUser godoc
// @Summary      Delete user
// @Description  Soft delete user by ID (reason required)
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id path string true "User ID"
// @Param        body body entity.DeleteUserRequest true "Delete reason"
// @Success      200 {object} map[string]string
// @Failure      400 {object} helper.ErrorResponse
// @Failure      401 {object} helper.ErrorResponse
// @Failure      403 {object} helper.ErrorResponse
// @Failure      404 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/users/{id} [delete]
func (h *userHTTPHandler) DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "User ID required", nil)
		return
	}

	var req entity.DeleteUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format - alasan penghapusan wajib", err.Error())
		return
	}

	actor := h.getActorContext(c)
	err := h.userUsecase.Delete(c.Request.Context(), userID, req.Reason, actor)
	if err != nil {
		if errors.Is(err, entity.ErrUserNotFound) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Pengguna tidak ditemukan", nil)
			return
		}
		if errors.Is(err, entity.ErrDeleteReasonRequired) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Alasan penghapusan wajib diisi", nil)
			return
		}
		if errors.Is(err, entity.ErrCannotDeleteSelf) {
			helper.ResponseOutput(c, int32(http.StatusForbidden), "Tidak dapat menghapus akun sendiri", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Pengguna berhasil dihapus"})
}
