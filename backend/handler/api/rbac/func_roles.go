package rbac

import (
	"errors"
	"net/http"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

func (h *rbacHTTPHandler) GetRoleList(c *gin.Context) {
	var req entity.GetRoleListRequest
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

	resp, err := h.roleUsecase.List(c.Request.Context(), req)
	if err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *rbacHTTPHandler) GetRoleDetail(c *gin.Context) {
	roleID := c.Param("id")
	if roleID == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Role ID required", nil)
		return
	}

	resp, err := h.roleUsecase.Get(c.Request.Context(), roleID)
	if err != nil {
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", err.Error())
		return
	}
	if resp == nil {
		helper.ResponseOutput(c, int32(http.StatusNotFound), "Role tidak ditemukan", nil)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *rbacHTTPHandler) CreateRole(c *gin.Context) {
	var req entity.CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}

	resp, err := h.roleUsecase.Create(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, entity.ErrInvalidData) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Data tidak valid", nil)
			return
		}
		if errors.Is(err, entity.ErrDuplicateKey) {
			helper.ResponseOutput(c, int32(http.StatusConflict), "Kode role sudah digunakan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", err.Error())
		return
	}
	c.JSON(http.StatusCreated, resp)
}

func (h *rbacHTTPHandler) UpdateRole(c *gin.Context) {
	roleID := c.Param("id")
	if roleID == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Role ID required", nil)
		return
	}

	var req entity.UpdateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format", err.Error())
		return
	}

	resp, err := h.roleUsecase.Update(c.Request.Context(), roleID, req)
	if err != nil {
		if entity.IsRecordNotFound(err) {
			helper.ResponseOutput(c, int32(http.StatusNotFound), "Role tidak ditemukan", nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *rbacHTTPHandler) DeleteRole(c *gin.Context) {
	roleID := c.Param("id")
	if roleID == "" {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Role ID required", nil)
		return
	}

	var req entity.DeleteRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseOutput(c, int32(http.StatusBadRequest), "Invalid request format - alasan penghapusan wajib", err.Error())
		return
	}

	err := h.roleUsecase.Delete(c.Request.Context(), roleID, req.Reason)
	if err != nil {
		if errors.Is(err, entity.ErrDeleteReasonRequired) {
			helper.ResponseOutput(c, int32(http.StatusBadRequest), "Alasan penghapusan wajib diisi", nil)
			return
		}
		if errors.Is(err, entity.ErrRoleInUse) {
			helper.ResponseOutput(c, int32(http.StatusConflict), err.Error(), nil)
			return
		}
		helper.ResponseOutput(c, int32(http.StatusInternalServerError), "Internal server error", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Role berhasil dihapus"})
}
