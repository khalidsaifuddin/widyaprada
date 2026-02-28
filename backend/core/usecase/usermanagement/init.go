package usermanagement

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/core/repository"
)

// ActorContext konteks pengguna yang melakukan aksi (untuk RBAC scope)
type ActorContext struct {
	UserID    string
	SatkerID  *string
	RoleCodes []string
}

// IsSuperAdmin true jika actor punya role SUPER_ADMIN
func (a *ActorContext) IsSuperAdmin() bool {
	for _, c := range a.RoleCodes {
		if c == "SUPER_ADMIN" {
			return true
		}
	}
	return false
}

// UserUsecase interface untuk manajemen pengguna (SDD_Auth_Manajemen_Pengguna)
type UserUsecase interface {
	List(ctx context.Context, req entity.GetUserListRequest, actor *ActorContext) (entity.GetUserListResponse, error)
	Get(ctx context.Context, userID string, actor *ActorContext) (*entity.UserDetailResponse, error)
	Create(ctx context.Context, req entity.CreateUserRequest, actor *ActorContext) (*entity.UserDetailResponse, error)
	Update(ctx context.Context, userID string, req entity.UpdateUserRequest, actor *ActorContext) (*entity.UserDetailResponse, error)
	Delete(ctx context.Context, userID string, reason string, actor *ActorContext) error
}

type userUsecase struct {
	userRepo repository.UserRepo
}

func NewUserUsecase(userRepo repository.UserRepo) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}
