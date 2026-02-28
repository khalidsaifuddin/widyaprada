package user

import (
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// GetUserList godoc
// @Summary      List users
// @Description  Get paginated list of users with filters
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        q query string false "Search"
// @Param        role_id query string false "Filter by role"
// @Param        satker_id query string false "Filter by satker"
// @Param        status query string false "active, inactive, all"
// @Param        page query int false "Page" default(1)
// @Param        page_size query int false "Page size" default(10)
// @Param        sort_by query string false "Sort by"
// @Param        sort_order query string false "asc, desc"
// @Success      200 {object} entity.GetUserListResponse
// @Failure      401 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/users [get]
func (h *userHTTPHandler) GetUserList(c *gin.Context) {
	var req entity.GetUserListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	actor := h.getActorContext(c)
	resp, err := h.userUsecase.List(c.Request.Context(), req, actor)
	if err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}
