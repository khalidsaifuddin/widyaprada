package rbac

import (
	"context"
	"fmt"
	"strings"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *roleUsecase) List(ctx context.Context, req entity.GetRoleListRequest) (entity.GetRoleListResponse, error) {
	return u.rbacRepo.ListRoles(ctx, req)
}

func (u *roleUsecase) Get(ctx context.Context, roleID string) (*entity.RoleDetailResponse, error) {
	return u.rbacRepo.GetRoleByID(ctx, roleID)
}

func (u *roleUsecase) Create(ctx context.Context, req entity.CreateRoleRequest) (*entity.RoleDetailResponse, error) {
	code := strings.TrimSpace(req.Code)
	name := strings.TrimSpace(req.Name)
	if code == "" || name == "" {
		return nil, entity.ErrInvalidData
	}

	roleID, err := u.rbacRepo.CreateRole(ctx, code, name, req.Description)
	if err != nil {
		return nil, err
	}

	if len(req.PermissionIDs) > 0 {
		if err := u.rbacRepo.SyncRolePermissions(ctx, roleID, req.PermissionIDs); err != nil {
			return nil, err
		}
	}

	return u.rbacRepo.GetRoleByID(ctx, roleID)
}

func (u *roleUsecase) Update(ctx context.Context, roleID string, req entity.UpdateRoleRequest) (*entity.RoleDetailResponse, error) {
	role, err := u.rbacRepo.GetRoleByID(ctx, roleID)
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, entity.WrapRecordNotFoundf("role not found: %s", roleID)
	}

	code := strings.TrimSpace(req.Code)
	name := strings.TrimSpace(req.Name)
	if code == "" {
		code = role.Code
	}
	if name == "" {
		name = role.Name
	}

	if err := u.rbacRepo.UpdateRole(ctx, roleID, code, name, req.Description); err != nil {
		return nil, err
	}

	if req.PermissionIDs != nil {
		if err := u.rbacRepo.SyncRolePermissions(ctx, roleID, req.PermissionIDs); err != nil {
			return nil, err
		}
	}

	return u.rbacRepo.GetRoleByID(ctx, roleID)
}

func (u *roleUsecase) Delete(ctx context.Context, roleID string, reason string) error {
	reason = strings.TrimSpace(reason)
	if reason == "" {
		return entity.ErrDeleteReasonRequired
	}

	count, err := u.rbacRepo.CountUsersWithRole(ctx, roleID)
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("%w: %d pengguna", entity.ErrRoleInUse, count)
	}

	return u.rbacRepo.DeleteRole(ctx, roleID, reason)
}
