package repository

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

// UserRepo interface untuk operasi user (sesuai SDD)
type UserRepo interface {
	FindByEmailOrUsername(ctx context.Context, identifier string) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
	GetUserRoles(ctx context.Context, userID string) ([]entity.RoleInfo, error)
	GetRoleByCode(ctx context.Context, code string) (*entity.RoleInfo, error)
	CreateUser(ctx context.Context, user *entity.User, roleID string) error
	CreateUserWithRoles(ctx context.Context, user *entity.User, roleIDs []string) error
	UpdatePassword(ctx context.Context, userID, passwordHash string) error

	// User Management (SDD_Auth_Manajemen_Pengguna)
	ListUsers(ctx context.Context, req entity.GetUserListRequest, satkerFilter *string) (entity.GetUserListResponse, error)
	GetByID(ctx context.Context, userID string) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User, roleIDs []string) error
	DeleteUser(ctx context.Context, userID, reason string) error
}
