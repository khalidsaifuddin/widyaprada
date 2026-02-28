package entity

// GetRoleListRequest untuk GET /api/v1/rbac/roles
type GetRoleListRequest struct {
	Q        string `form:"q"`
	Page     int64  `form:"page"`
	PageSize int64  `form:"page_size"`
}

// RoleListItem untuk response list roles
type RoleListItem struct {
	ID          string   `json:"id"`
	Code        string   `json:"code"`
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Permissions []string `json:"permissions,omitempty"` // permission codes
}

// GetRoleListResponse response list roles
type GetRoleListResponse struct {
	Items     []RoleListItem `json:"items"`
	TotalPage int64          `json:"total_page"`
	TotalData int64          `json:"total_data"`
	Page      int64          `json:"page"`
	PageSize  int64          `json:"page_size"`
}

// RoleDetailResponse untuk GET /api/v1/rbac/roles/:id
type RoleDetailResponse struct {
	ID           string           `json:"id"`
	Code         string           `json:"code"`
	Name         string           `json:"name"`
	Description  string           `json:"description,omitempty"`
	Permissions  []PermissionInfo `json:"permissions"`
	CreatedAt    string           `json:"created_at,omitempty"`
	UpdatedAt    string           `json:"updated_at,omitempty"`
}

// PermissionInfo untuk permission di response
type PermissionInfo struct {
	ID    string `json:"id"`
	Code  string `json:"code"`
	Name  string `json:"name"`
	Group string `json:"group,omitempty"`
}

// CreateRoleRequest untuk POST /api/v1/rbac/roles
type CreateRoleRequest struct {
	Code         string   `json:"code" binding:"required"`
	Name         string   `json:"name" binding:"required"`
	Description  string   `json:"description"`
	PermissionIDs []string `json:"permission_ids"`
}

// UpdateRoleRequest untuk PUT /api/v1/rbac/roles/:id
type UpdateRoleRequest struct {
	Code          string   `json:"code"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	PermissionIDs []string `json:"permission_ids"`
}

// DeleteRoleRequest untuk DELETE /api/v1/rbac/roles/:id
type DeleteRoleRequest struct {
	Reason string `json:"reason" binding:"required"`
}

// GetPermissionListRequest untuk GET /api/v1/rbac/permissions
type GetPermissionListRequest struct {
	Q     string `form:"q"`
	Group string `form:"group"`
	Page  int64  `form:"page"`
	PageSize int64 `form:"page_size"`
}

// PermissionListItem untuk response list permissions
type PermissionListItem struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Group       string `json:"group,omitempty"`
	Description string `json:"description,omitempty"`
}

// GetPermissionListResponse response list permissions
type GetPermissionListResponse struct {
	Items     []PermissionListItem `json:"items"`
	TotalPage int64                `json:"total_page"`
	TotalData int64                `json:"total_data"`
	Page      int64                `json:"page"`
	PageSize  int64                `json:"page_size"`
}

// PermissionDetailResponse untuk GET /api/v1/rbac/permissions/:id
type PermissionDetailResponse struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Group       string `json:"group,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

// CreatePermissionRequest untuk POST /api/v1/rbac/permissions
type CreatePermissionRequest struct {
	Code        string `json:"code" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Group       string `json:"group"`
	Description string `json:"description"`
}

// UpdatePermissionRequest untuk PUT /api/v1/rbac/permissions/:id
type UpdatePermissionRequest struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Group       string `json:"group"`
	Description string `json:"description"`
}

// DeletePermissionRequest untuk DELETE /api/v1/rbac/permissions/:id
type DeletePermissionRequest struct {
	Reason string `json:"reason" binding:"required"`
}
