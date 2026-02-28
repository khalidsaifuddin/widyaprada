package rbacrepo

import (
	"context"
	"strings"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	userrepo "github.com/ProjectWidyaprada/backend/repository/user-repo"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *rbacRepo) ListRoles(ctx context.Context, req entity.GetRoleListRequest) (entity.GetRoleListResponse, error) {
	var resp entity.GetRoleListResponse
	db := r.db.WithContext(ctx).Model(&userrepo.Role{})

	if req.Q != "" {
		q := "%" + strings.TrimSpace(req.Q) + "%"
		db = db.Where("LOWER(code) LIKE LOWER(?) OR LOWER(name) LIKE LOWER(?)", q, q)
	}

	if err := db.Count(&resp.TotalData).Error; err != nil {
		return resp, err
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	offset, limit := helper.GetOffsetAndLimit(req.Page, req.PageSize)
	db = db.Order("name ASC").Offset(int(offset)).Limit(int(limit))

	var roles []userrepo.Role
	if err := db.Find(&roles).Error; err != nil {
		return resp, err
	}

	for i := range roles {
		perms, _ := r.getPermissionCodesByRoleID(ctx, roles[i].ID)
		resp.Items = append(resp.Items, entity.RoleListItem{
			ID:          roles[i].ID,
			Code:        roles[i].Code,
			Name:        roles[i].Name,
			Description: roles[i].Description,
			Permissions: perms,
		})
	}

	resp.TotalPage = helper.GenerateTotalPage(resp.TotalData, req.PageSize)
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	return resp, nil
}

func (r *rbacRepo) GetRoleByID(ctx context.Context, roleID string) (*entity.RoleDetailResponse, error) {
	var role userrepo.Role
	err := r.db.WithContext(ctx).Where("id = ?", roleID).First(&role).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	perms, err := r.getPermissionsByRoleID(ctx, roleID)
	if err != nil {
		return nil, err
	}

	createdAt, updatedAt := "", ""
	if role.CreatedAt != nil {
		createdAt = role.CreatedAt.Format(time.RFC3339)
	}
	if role.UpdatedAt != nil {
		updatedAt = role.UpdatedAt.Format(time.RFC3339)
	}

	return &entity.RoleDetailResponse{
		ID:          role.ID,
		Code:        role.Code,
		Name:        role.Name,
		Description: role.Description,
		Permissions: perms,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}, nil
}

func (r *rbacRepo) CreateRole(ctx context.Context, code, name, description string) (string, error) {
	roleID := uuid.New().String()
	now := time.Now().UTC()
	role := userrepo.Role{
		ID:          roleID,
		Code:        strings.TrimSpace(code),
		Name:        strings.TrimSpace(name),
		Description: strings.TrimSpace(description),
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}
	if err := r.db.WithContext(ctx).Create(&role).Error; err != nil {
		return "", err
	}
	return roleID, nil
}

func (r *rbacRepo) UpdateRole(ctx context.Context, roleID, code, name, description string) error {
	updates := map[string]interface{}{
		"code":        strings.TrimSpace(code),
		"name":        strings.TrimSpace(name),
		"description": strings.TrimSpace(description),
		"updated_at":  time.Now().UTC(),
	}
	return r.db.WithContext(ctx).Model(&userrepo.Role{}).Where("id = ?", roleID).Updates(updates).Error
}

func (r *rbacRepo) DeleteRole(ctx context.Context, roleID, reason string) error {
	now := time.Now().UTC()
	return r.db.WithContext(ctx).Model(&userrepo.Role{}).
		Where("id = ?", roleID).
		Updates(map[string]interface{}{
			"deleted_at":     now,
			"deleted_reason": reason,
		}).Error
}

func (r *rbacRepo) CountUsersWithRole(ctx context.Context, roleID string) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Table("user_roles").Where("role_id = ?", roleID).Count(&count).Error
	return count, err
}

func (r *rbacRepo) SyncRolePermissions(ctx context.Context, roleID string, permissionIDs []string) error {
	if err := r.db.WithContext(ctx).Where("role_id = ?", roleID).Delete(&RolePermission{}).Error; err != nil {
		return err
	}
	now := time.Now().UTC()
	for _, permID := range permissionIDs {
		if permID == "" {
			continue
		}
		rp := RolePermission{RoleID: roleID, PermissionID: permID, CreatedAt: &now}
		if err := r.db.WithContext(ctx).Create(&rp).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r *rbacRepo) getPermissionCodesByRoleID(ctx context.Context, roleID string) ([]string, error) {
	var codes []string
	err := r.db.WithContext(ctx).Table("permissions").
		Select("permissions.code").
		Joins("INNER JOIN role_permissions ON role_permissions.permission_id = permissions.id").
		Where("role_permissions.role_id = ?", roleID).
		Pluck("permissions.code", &codes).Error
	return codes, err
}

func (r *rbacRepo) getPermissionsByRoleID(ctx context.Context, roleID string) ([]entity.PermissionInfo, error) {
	var perms []Permission
	err := r.db.WithContext(ctx).Table("permissions").
		Select("permissions.*").
		Joins("INNER JOIN role_permissions ON role_permissions.permission_id = permissions.id").
		Where("role_permissions.role_id = ?", roleID).
		Find(&perms).Error
	if err != nil {
		return nil, err
	}
	result := make([]entity.PermissionInfo, 0, len(perms))
	for i := range perms {
		result = append(result, perms[i].ToEntity())
	}
	return result, nil
}
