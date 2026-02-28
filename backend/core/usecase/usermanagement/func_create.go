package usermanagement

import (
	"context"
	"strings"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/auth"
	"github.com/google/uuid"
)

func (u *userUsecase) Create(ctx context.Context, req entity.CreateUserRequest, actor *ActorContext) (*entity.UserDetailResponse, error) {
	if len(req.RoleIDs) == 0 {
		return nil, entity.ErrAtLeastOneRole
	}

	email := strings.TrimSpace(strings.ToLower(req.Email))
	username := strings.TrimSpace(req.Username)

	// Validasi unik email
	exist, _ := u.userRepo.FindByEmail(ctx, email)
	if exist != nil {
		return nil, entity.ErrEmailExists
	}

	// Validasi unik username
	exist, _ = u.userRepo.FindByUsername(ctx, username)
	if exist != nil {
		return nil, entity.ErrUsernameExists
	}

	// Scope: Admin Satker hanya bisa create user dengan satker sendiri
	if actor != nil && !actor.IsSuperAdmin() && actor.SatkerID != nil {
		if req.SatkerID == nil || *req.SatkerID != *actor.SatkerID {
			return nil, entity.ErrUserNotFound
		}
	}

	passwordHash, err := auth.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		ID:           uuid.New().String(),
		Name:         strings.TrimSpace(req.Name),
		Email:        email,
		Username:     username,
		NIP:          strings.TrimSpace(req.NIP),
		PasswordHash: passwordHash,
		SatkerID:     req.SatkerID,
		IsActive:     req.IsActive,
	}

	if err := u.userRepo.CreateUserWithRoles(ctx, user, req.RoleIDs); err != nil {
		return nil, err
	}

	return u.Get(ctx, user.ID, actor)
}
