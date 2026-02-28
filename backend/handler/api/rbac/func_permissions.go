package rbac

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

func (h *rbacHTTPHandler) GetPermissionList(c *gin.Context) {
	var req entity.GetPermissionListRequest
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

	resp, err := h.permissionUsecase.List(c.Request.Context(), req)
	if err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *rbacHTTPHandler) GetPermissionDetail(c *gin.Context) {
	permID := c.Param("id")
	if permID == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Permission ID required", nil)
		return
	}

	resp, err := h.permissionUsecase.Get(c.Request.Context(), permID)
	if err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", err.Error())
		return
	}
	if resp == nil {
		helper.ResponseOutput(c, int32(http.StatusNotFound), "Permission tidak ditemukan", nil)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *rbacHTTPHandler) CreatePermission(c *gin.Context) {
	var req entity.CreatePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}

	resp, err := h.permissionUsecase.Create(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, entity.ErrInvalidData) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Data tidak valid", nil)
			return
		}
		if errors.Is(err, entity.ErrDuplicateKey) {
			helper.ResponseOutput(c, int32(http.StatusConflict), "Kode permission sudah digunakan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", err.Error())
		return
	}
	c.JSON(http.StatusCreated, resp)
}

func (h *rbacHTTPHandler) UpdatePermission(c *gin.Context) {
	permID := c.Param("id")
	if permID == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Permission ID required", nil)
		return
	}

	var req entity.UpdatePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}

	resp, err := h.permissionUsecase.Update(c.Request.Context(), permID, req)
	if err != nil {
		if entity.IsRecordNotFound(err) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Permission tidak ditemukan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *rbacHTTPHandler) DeletePermission(c *gin.Context) {
	permID := c.Param("id")
	if permID == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Permission ID required", nil)
		return
	}

	var req entity.DeletePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format - alasan penghapusan wajib", err.Error())
		return
	}

	err := h.permissionUsecase.Delete(c.Request.Context(), permID, req.Reason)
	if err != nil {
		if errors.Is(err, entity.ErrDeleteReasonRequired) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Alasan penghapusan wajib diisi", nil)
			return
		}
		if errors.Is(err, entity.ErrPermissionInUse) {
			helper.ResponseOutput(c, int32(http.StatusConflict), err.Error(), nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Permission berhasil dihapus"})
}
