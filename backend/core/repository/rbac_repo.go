package repository

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

// RBACRepo interface untuk operasi RBAC (SDD_RBAC)
type RBACRepo interface {
	// Roles
	ListRoles(ctx context.Context, req entity.GetRoleListRequest) (entity.GetRoleListResponse, error)
	GetRoleByID(ctx context.Context, roleID string) (*entity.RoleDetailResponse, error)
	CreateRole(ctx context.Context, code, name, description string) (string, error)
	UpdateRole(ctx context.Context, roleID, code, name, description string) error
	DeleteRole(ctx context.Context, roleID, reason string) error
	CountUsersWithRole(ctx context.Context, roleID string) (int64, error)

	// Role permissions
	SyncRolePermissions(ctx context.Context, roleID string, permissionIDs []string) error

	// Permissions
	ListPermissions(ctx context.Context, req entity.GetPermissionListRequest) (entity.GetPermissionListResponse, error)
	GetPermissionByID(ctx context.Context, permID string) (*entity.PermissionDetailResponse, error)
	CreatePermission(ctx context.Context, code, name, group, description string) (string, error)
	UpdatePermission(ctx context.Context, permID, code, name, group, description string) error
	DeletePermission(ctx context.Context, permID, reason string) error
	CountRolesWithPermission(ctx context.Context, permID string) (int64, error)
}
