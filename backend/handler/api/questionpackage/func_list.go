package questionpackage

import (
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

// GetPackageList godoc
// @Summary      List question packages
// @Description  Get paginated list of question packages with filters
// @Tags         question-packages
// @Accept       json
// @Produce      json
// @Param        q query string false "Search"
// @Param        status query string false "Draft, Aktif, all"
// @Param        status_verifikasi query string false "Belum, Sudah, all"
// @Param        page query int false "Page" default(1)
// @Param        page_size query int false "Page size" default(10)
// @Success      200 {object} entity.GetPackageListResponse
// @Failure      401 {object} helper.ErrorResponse
// @Failure      500 {object} helper.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/question-packages [get]
func (h *questionPackageHTTPHandler) GetPackageList(c *gin.Context) {
	var req entity.GetPackageListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}
	resp, err := h.paketSoalUsecase.List(c.Request.Context(), req)
	if err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", nil)
		return
	}
	c.JSON(http.StatusOK, resp)
}
