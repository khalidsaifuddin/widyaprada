package auth

import (
	"context"
	"strings"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/auth"
)

// Login mengimplementasikan use case login sesuai SDD
func (u *authUsecase) Login(ctx context.Context, req entity.LoginRequest) (*entity.LoginResponse, error) {
	// 1. Normalisasi input
	identifier := strings.TrimSpace(strings.ToLower(req.Identifier))
	password := strings.TrimSpace(req.Password)
	if identifier == "" || password == "" {
		return nil, entity.ErrInvalidCredentials
	}

	// 2. Ambil user via email atau username
	user, err := u.userRepo.FindByEmailOrUsername(ctx, identifier)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, entity.ErrInvalidCredentials
	}

	// 3. Cek is_active
	if !user.IsActive {
		return nil, entity.ErrAccountInactive
	}

	// 4. Verifikasi password
	if !auth.VerifyPassword(password, user.PasswordHash) {
		return nil, entity.ErrInvalidCredentials
	}

	// 5. Rate limit (Redis) - diimplementasikan nanti

	// 6. Ambil roles
	roles, err := u.userRepo.GetUserRoles(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	// 7. Hitung default_home_path
	defaultHomePath := u.getDefaultHomePath(roles)

	// 8. Buat token JWT
	roleCodes := make([]string, 0, len(roles))
	for _, r := range roles {
		roleCodes = append(roleCodes, r.Code)
	}

	accessToken, expiresIn, err := auth.GenerateToken(
		user.ID, user.Email, roleCodes,
		u.cfg.JWTSecret, u.cfg.JWTExpiryHr,
	)
	if err != nil {
		return nil, err
	}

	// 9. Return LoginResult
	return &entity.LoginResponse{
		AccessToken: accessToken,
		TokenType:   "bearer",
		ExpiresIn:   expiresIn,
		User: entity.UserResponse{
			ID:              user.ID,
			Name:            user.Name,
			Email:           user.Email,
			Username:        user.Username,
			Roles:           roles,
			DefaultHomePath: defaultHomePath,
		},
	}, nil
}

// getDefaultHomePath menentukan halaman pertama berdasarkan role (prioritas tertinggi)
func (u *authUsecase) getDefaultHomePath(roles []entity.RoleInfo) string {
	// Prioritas: SUPER_ADMIN > ADMIN_SATKER > WIDYAPRADA
	for _, r := range roles {
		switch strings.ToUpper(r.Code) {
		case "SUPER_ADMIN":
			return "/dashboard"
		case "ADMIN_SATKER":
			return "/dashboard"
		case "WIDYAPRADA":
			return "/dashboard"
		}
	}
	return "/dashboard"
}
