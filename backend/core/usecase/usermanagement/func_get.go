package usermanagement

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *userUsecase) Get(ctx context.Context, userID string, actor *ActorContext) (*entity.UserDetailResponse, error) {
	user, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, entity.ErrUserNotFound
	}

	// Scope: Admin Satker hanya bisa akses user dalam satker sendiri
	if actor != nil && !actor.IsSuperAdmin() && actor.SatkerID != nil {
		if user.SatkerID == nil || *user.SatkerID != *actor.SatkerID {
			return nil, entity.ErrUserNotFound
		}
	}

	roles, err := u.userRepo.GetUserRoles(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &entity.UserDetailResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Username:  user.Username,
		NIP:       user.NIP,
		Roles:     roles,
		SatkerID:  user.SatkerID,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
