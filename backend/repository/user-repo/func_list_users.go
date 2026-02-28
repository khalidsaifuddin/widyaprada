package userrepo

import (
	"context"
	"fmt"
	"strings"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
)

// ListUsers daftar users dengan filter dan paginasi
func (r *userRepo) ListUsers(ctx context.Context, req entity.GetUserListRequest, satkerFilter *string) (entity.GetUserListResponse, error) {
	var resp entity.GetUserListResponse
	db := r.db.WithContext(ctx).Model(&User{}).
		Select("users.id, users.name, users.email, users.username, users.satker_id, users.is_active, users.created_at")

	// Filter by satker (Admin Satker scope)
	if satkerFilter != nil && *satkerFilter != "" {
		db = db.Where("users.satker_id = ?", *satkerFilter)
	}

	// Search
	if req.Q != "" {
		q := "%" + strings.TrimSpace(req.Q) + "%"
		db = db.Where(
			"LOWER(users.name) LIKE LOWER(?) OR LOWER(users.email) LIKE LOWER(?) OR LOWER(users.username) LIKE LOWER(?)",
			q, q, q,
		)
	}

	// Filter role
	if req.RoleID != "" {
		db = db.Joins("INNER JOIN user_roles ON user_roles.user_id = users.id").
			Where("user_roles.role_id = ?", req.RoleID)
	}

	// Filter satker
	if req.SatkerID != "" {
		db = db.Where("users.satker_id = ?", req.SatkerID)
	}

	// Filter status
	switch strings.ToLower(req.Status) {
	case "active":
		db = db.Where("users.is_active = ?", true)
	case "inactive":
		db = db.Where("users.is_active = ?", false)
	}

	if err := db.Count(&resp.TotalData).Error; err != nil {
		return resp, err
	}

	// Sort
	sortBy := "users.created_at"
	if req.SortBy != "" {
		switch strings.ToLower(req.SortBy) {
		case "name":
			sortBy = "users.name"
		case "email":
			sortBy = "users.email"
		case "username":
			sortBy = "users.username"
		case "created_at":
			sortBy = "users.created_at"
		}
	}
	sortOrder := "DESC"
	if strings.ToLower(req.SortOrder) == "asc" {
		sortOrder = "ASC"
	}
	db = db.Order(fmt.Sprintf("%s %s", sortBy, sortOrder))

	// Pagination
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	offset, limit := helper.GetOffsetAndLimit(req.Page, req.PageSize)
	db = db.Offset(int(offset)).Limit(int(limit))

	var users []User
	if err := db.Find(&users).Error; err != nil {
		return resp, err
	}

	for i := range users {
		roles, _ := r.GetUserRoles(ctx, users[i].ID)
		resp.Items = append(resp.Items, users[i].ToUserListItem(roles))
	}

	resp.TotalPage = helper.GenerateTotalPage(resp.TotalData, req.PageSize)
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	return resp, nil
}
