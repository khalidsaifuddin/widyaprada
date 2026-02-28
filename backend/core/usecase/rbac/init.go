package rbac

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/core/repository"
)

// RoleUsecase interface untuk manajemen role (SDD_RBAC)
type RoleUsecase interface {
	List(ctx context.Context, req entity.GetRoleListRequest) (entity.GetRoleListResponse, error)
	Get(ctx context.Context, roleID string) (*entity.RoleDetailResponse, error)
	Create(ctx context.Context, req entity.CreateRoleRequest) (*entity.RoleDetailResponse, error)
	Update(ctx context.Context, roleID string, req entity.UpdateRoleRequest) (*entity.RoleDetailResponse, error)
	Delete(ctx context.Context, roleID string, reason string) error
}

// PermissionUsecase interface untuk manajemen permission (SDD_RBAC)
type PermissionUsecase interface {
	List(ctx context.Context, req entity.GetPermissionListRequest) (entity.GetPermissionListResponse, error)
	Get(ctx context.Context, permID string) (*entity.PermissionDetailResponse, error)
	Create(ctx context.Context, req entity.CreatePermissionRequest) (*entity.PermissionDetailResponse, error)
	Update(ctx context.Context, permID string, req entity.UpdatePermissionRequest) (*entity.PermissionDetailResponse, error)
	Delete(ctx context.Context, permID string, reason string) error
}

type roleUsecase struct {
	rbacRepo repository.RBACRepo
}

type permissionUsecase struct {
	rbacRepo repository.RBACRepo
}

func NewRoleUsecase(rbacRepo repository.RBACRepo) RoleUsecase {
	return &roleUsecase{rbacRepo: rbacRepo}
}

func NewPermissionUsecase(rbacRepo repository.RBACRepo) PermissionUsecase {
	return &permissionUsecase{rbacRepo: rbacRepo}
}
