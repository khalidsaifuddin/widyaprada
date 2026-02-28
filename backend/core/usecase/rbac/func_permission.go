package rbac

import (
	"context"
	"fmt"
	"strings"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *permissionUsecase) List(ctx context.Context, req entity.GetPermissionListRequest) (entity.GetPermissionListResponse, error) {
	return u.rbacRepo.ListPermissions(ctx, req)
}

func (u *permissionUsecase) Get(ctx context.Context, permID string) (*entity.PermissionDetailResponse, error) {
	return u.rbacRepo.GetPermissionByID(ctx, permID)
}

func (u *permissionUsecase) Create(ctx context.Context, req entity.CreatePermissionRequest) (*entity.PermissionDetailResponse, error) {
	code := strings.TrimSpace(req.Code)
	name := strings.TrimSpace(req.Name)
	if code == "" || name == "" {
		return nil, entity.ErrInvalidData
	}

	permID, err := u.rbacRepo.CreatePermission(ctx, code, name, req.Group, req.Description)
	if err != nil {
		return nil, err
	}

	return u.rbacRepo.GetPermissionByID(ctx, permID)
}

func (u *permissionUsecase) Update(ctx context.Context, permID string, req entity.UpdatePermissionRequest) (*entity.PermissionDetailResponse, error) {
	perm, err := u.rbacRepo.GetPermissionByID(ctx, permID)
	if err != nil {
		return nil, err
	}
	if perm == nil {
		return nil, entity.WrapRecordNotFoundf("permission not found: %s", permID)
	}

	code := strings.TrimSpace(req.Code)
	name := strings.TrimSpace(req.Name)
	if code == "" {
		code = perm.Code
	}
	if name == "" {
		name = perm.Name
	}

	if err := u.rbacRepo.UpdatePermission(ctx, permID, code, name, req.Group, req.Description); err != nil {
		return nil, err
	}

	return u.rbacRepo.GetPermissionByID(ctx, permID)
}

func (u *permissionUsecase) Delete(ctx context.Context, permID string, reason string) error {
	reason = strings.TrimSpace(reason)
	if reason == "" {
		return entity.ErrDeleteReasonRequired
	}

	count, err := u.rbacRepo.CountRolesWithPermission(ctx, permID)
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("%w: %d role", entity.ErrPermissionInUse, count)
	}

	return u.rbacRepo.DeletePermission(ctx, permID, reason)
}
