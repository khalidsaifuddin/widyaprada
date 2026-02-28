package usermanagement

import (
	"context"
	"strings"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/auth"
)

func (u *userUsecase) Update(ctx context.Context, userID string, req entity.UpdateUserRequest, actor *ActorContext) (*entity.UserDetailResponse, error) {
	user, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, entity.ErrUserNotFound
	}

	// Scope: Admin Satker hanya bisa update user dalam satker sendiri
	if actor != nil && !actor.IsSuperAdmin() && actor.SatkerID != nil {
		if user.SatkerID == nil || *user.SatkerID != *actor.SatkerID {
			return nil, entity.ErrUserNotFound
		}
	}

	// Update fields
	if req.Name != "" {
		user.Name = strings.TrimSpace(req.Name)
	}
	if req.Email != "" {
		email := strings.TrimSpace(strings.ToLower(req.Email))
		exist, _ := u.userRepo.FindByEmail(ctx, email)
		if exist != nil && exist.ID != userID {
			return nil, entity.ErrEmailExists
		}
		user.Email = email
	}
	if req.Username != "" {
		username := strings.TrimSpace(req.Username)
		exist, _ := u.userRepo.FindByUsername(ctx, username)
		if exist != nil && exist.ID != userID {
			return nil, entity.ErrUsernameExists
		}
		user.Username = username
	}
	if req.NIP != "" {
		user.NIP = strings.TrimSpace(req.NIP)
	}
	if req.SatkerID != nil {
		// Scope: Admin Satker hanya bisa set satker sendiri
		if actor != nil && !actor.IsSuperAdmin() && actor.SatkerID != nil {
			if *req.SatkerID != *actor.SatkerID {
				return nil, entity.ErrUserNotFound
			}
		}
		user.SatkerID = req.SatkerID
	}
	if req.IsActive != nil {
		user.IsActive = *req.IsActive
	}

	// Password (optional di update)
	if req.Password != "" {
		hash, err := auth.HashPassword(req.Password)
		if err != nil {
			return nil, err
		}
		user.PasswordHash = hash
	}

	roleIDs := req.RoleIDs
	if len(roleIDs) == 0 {
		roleIDs = nil // tidak update roles
	} else {
		// Validasi minimal 1 role
		validRoles := make([]string, 0)
		for _, r := range roleIDs {
			if r != "" {
				validRoles = append(validRoles, r)
			}
		}
		if len(validRoles) == 0 {
			return nil, entity.ErrAtLeastOneRole
		}
		roleIDs = validRoles
	}

	if err := u.userRepo.UpdateUser(ctx, user, roleIDs); err != nil {
		return nil, err
	}

	return u.Get(ctx, userID, actor)
}
