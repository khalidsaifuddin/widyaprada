package usermanagement

import (
	"context"
	"strings"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *userUsecase) Delete(ctx context.Context, userID string, reason string, actor *ActorContext) error {
	reason = strings.TrimSpace(reason)
	if reason == "" {
		return entity.ErrDeleteReasonRequired
	}

	// Super Admin tidak boleh hapus diri sendiri
	if actor != nil && actor.UserID == userID {
		return entity.ErrCannotDeleteSelf
	}

	user, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return err
	}
	if user == nil {
		return entity.ErrUserNotFound
	}

	// Scope: Admin Satker hanya bisa delete user dalam satker sendiri
	if actor != nil && !actor.IsSuperAdmin() && actor.SatkerID != nil {
		if user.SatkerID == nil || *user.SatkerID != *actor.SatkerID {
			return entity.ErrUserNotFound
		}
	}

	return u.userRepo.DeleteUser(ctx, userID, reason)
}
