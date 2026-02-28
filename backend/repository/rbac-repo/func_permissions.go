package rbacrepo

import (
	"context"
	"strings"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *rbacRepo) ListPermissions(ctx context.Context, req entity.GetPermissionListRequest) (entity.GetPermissionListResponse, error) {
	var resp entity.GetPermissionListResponse
	db := r.db.WithContext(ctx).Model(&Permission{})

	if req.Q != "" {
		q := "%" + strings.TrimSpace(req.Q) + "%"
		db = db.Where("LOWER(code) LIKE LOWER(?) OR LOWER(name) LIKE LOWER(?)", q, q)
	}
	if req.Group != "" {
		db = db.Where("group = ?", req.Group)
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
	db = db.Order("group ASC, code ASC").Offset(int(offset)).Limit(int(limit))

	var perms []Permission
	if err := db.Find(&perms).Error; err != nil {
		return resp, err
	}

	for i := range perms {
		resp.Items = append(resp.Items, entity.PermissionListItem{
			ID:          perms[i].ID,
			Code:        perms[i].Code,
			Name:        perms[i].Name,
			Group:       perms[i].Group,
			Description: perms[i].Description,
		})
	}

	resp.TotalPage = helper.GenerateTotalPage(resp.TotalData, req.PageSize)
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	return resp, nil
}

func (r *rbacRepo) GetPermissionByID(ctx context.Context, permID string) (*entity.PermissionDetailResponse, error) {
	var p Permission
	err := r.db.WithContext(ctx).Where("id = ?", permID).First(&p).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	createdAt, updatedAt := "", ""
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt.Format(time.RFC3339)
	}
	if p.UpdatedAt != nil {
		updatedAt = p.UpdatedAt.Format(time.RFC3339)
	}
	return &entity.PermissionDetailResponse{
		ID:          p.ID,
		Code:        p.Code,
		Name:        p.Name,
		Group:       p.Group,
		Description: p.Description,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}, nil
}

func (r *rbacRepo) CreatePermission(ctx context.Context, code, name, group, description string) (string, error) {
	permID := uuid.New().String()
	now := time.Now().UTC()
	p := Permission{
		ID:        permID,
		Code:      strings.TrimSpace(code),
		Name:      strings.TrimSpace(name),
		Group:     strings.TrimSpace(group),
		Description: strings.TrimSpace(description),
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	if err := r.db.WithContext(ctx).Create(&p).Error; err != nil {
		return "", err
	}
	return permID, nil
}

func (r *rbacRepo) UpdatePermission(ctx context.Context, permID, code, name, group, description string) error {
	updates := map[string]interface{}{
		"code":        strings.TrimSpace(code),
		"name":        strings.TrimSpace(name),
		"group":       strings.TrimSpace(group),
		"description": strings.TrimSpace(description),
		"updated_at":  time.Now().UTC(),
	}
	return r.db.WithContext(ctx).Model(&Permission{}).Where("id = ?", permID).Updates(updates).Error
}

func (r *rbacRepo) DeletePermission(ctx context.Context, permID, reason string) error {
	now := time.Now().UTC()
	return r.db.WithContext(ctx).Model(&Permission{}).
		Where("id = ?", permID).
		Updates(map[string]interface{}{
			"deleted_at":     now,
			"deleted_reason": reason,
		}).Error
}

func (r *rbacRepo) CountRolesWithPermission(ctx context.Context, permID string) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Table("role_permissions").Where("permission_id = ?", permID).Count(&count).Error
	return count, err
}
