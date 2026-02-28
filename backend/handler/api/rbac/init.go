package rbac

import (
	rbac_usecase "github.com/ProjectWidyaprada/backend/core/usecase/rbac"
	"github.com/gin-gonic/gin"
)

type RBACHTTPHandler interface {
	GetRoleList(c *gin.Context)
	GetRoleDetail(c *gin.Context)
	CreateRole(c *gin.Context)
	UpdateRole(c *gin.Context)
	DeleteRole(c *gin.Context)
	GetPermissionList(c *gin.Context)
	GetPermissionDetail(c *gin.Context)
	CreatePermission(c *gin.Context)
	UpdatePermission(c *gin.Context)
	DeletePermission(c *gin.Context)
}

type rbacHTTPHandler struct {
	roleUsecase       rbac_usecase.RoleUsecase
	permissionUsecase rbac_usecase.PermissionUsecase
}

func NewRBACHTTPHandler(roleUsecase rbac_usecase.RoleUsecase, permissionUsecase rbac_usecase.PermissionUsecase) RBACHTTPHandler {
	return &rbacHTTPHandler{
		roleUsecase:       roleUsecase,
		permissionUsecase: permissionUsecase,
	}
}
