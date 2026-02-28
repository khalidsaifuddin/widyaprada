package entity

// GetUserListRequest untuk GET /api/v1/users
type GetUserListRequest struct {
	Q         string `form:"q"`
	RoleID    string `form:"role_id"`
	SatkerID  string `form:"satker_id"`
	Status    string `form:"status"` // active, inactive, all
	Page      int64  `form:"page"`
	PageSize  int64  `form:"page_size"`
	SortBy    string `form:"sort_by"`
	SortOrder string `form:"sort_order"` // asc, desc
}

// GetUserListResponse response list users
type GetUserListResponse struct {
	Items      []UserListItem `json:"items"`
	TotalPage  int64          `json:"total_page"`
	TotalData  int64          `json:"total_data"`
	Page       int64          `json:"page"`
	PageSize   int64          `json:"page_size"`
}

// UserListItem item dalam list
type UserListItem struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	Roles     []RoleInfo `json:"roles"`
	SatkerID  *string    `json:"satker_id,omitempty"`
	IsActive  bool       `json:"is_active"`
	CreatedAt string     `json:"created_at,omitempty"`
}

// UserDetailResponse untuk GET /api/v1/users/:id
type UserDetailResponse struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	NIP       string     `json:"nip,omitempty"`
	Roles     []RoleInfo `json:"roles"`
	SatkerID  *string    `json:"satker_id,omitempty"`
	IsActive  bool       `json:"is_active"`
	CreatedAt string     `json:"created_at,omitempty"`
	UpdatedAt string     `json:"updated_at,omitempty"`
}

// CreateUserRequest untuk POST /api/v1/users
type CreateUserRequest struct {
	Name     string   `json:"name" binding:"required"`
	Email    string   `json:"email" binding:"required"`
	Username string   `json:"username" binding:"required"`
	Password string   `json:"password" binding:"required"`
	NIP      string   `json:"nip"`
	RoleIDs  []string `json:"role_ids" binding:"required"`
	SatkerID *string  `json:"satker_id"`
	IsActive bool     `json:"is_active"`
}

// UpdateUserRequest untuk PUT /api/v1/users/:id
type UpdateUserRequest struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Username string   `json:"username"`
	Password string   `json:"password"` // kosong = tidak diubah
	NIP      string   `json:"nip"`
	RoleIDs  []string `json:"role_ids"`
	SatkerID *string  `json:"satker_id"`
	IsActive *bool    `json:"is_active"`
}

// DeleteUserRequest untuk DELETE /api/v1/users/:id
type DeleteUserRequest struct {
	Reason string `json:"reason" binding:"required"`
}
