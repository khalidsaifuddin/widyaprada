package auth

import (
	"context"
	"errors"
	"testing"

	"github.com/ProjectWidyaprada/backend/config"
	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockUserRepo struct {
	findByEmailOrUsername func(ctx context.Context, identifier string) (*entity.User, error)
	getUserRoles          func(ctx context.Context, userID string) ([]entity.RoleInfo, error)
}

func (m *mockUserRepo) FindByEmailOrUsername(ctx context.Context, identifier string) (*entity.User, error) {
	if m.findByEmailOrUsername != nil {
		return m.findByEmailOrUsername(ctx, identifier)
	}
	return nil, nil
}

func (m *mockUserRepo) GetUserRoles(ctx context.Context, userID string) ([]entity.RoleInfo, error) {
	if m.getUserRoles != nil {
		return m.getUserRoles(ctx, userID)
	}
	return nil, nil
}

func (m *mockUserRepo) FindByEmail(ctx context.Context, email string) (*entity.User, error) { return nil, nil }
func (m *mockUserRepo) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	return nil, nil
}
func (m *mockUserRepo) GetRoleByCode(ctx context.Context, code string) (*entity.RoleInfo, error) {
	return nil, nil
}
func (m *mockUserRepo) CreateUser(ctx context.Context, user *entity.User, roleID string) error {
	return nil
}
func (m *mockUserRepo) CreateUserWithRoles(ctx context.Context, user *entity.User, roleIDs []string) error {
	return nil
}
func (m *mockUserRepo) UpdatePassword(ctx context.Context, userID, passwordHash string) error {
	return nil
}
func (m *mockUserRepo) ListUsers(ctx context.Context, req entity.GetUserListRequest, satkerFilter *string) (entity.GetUserListResponse, error) {
	return entity.GetUserListResponse{}, nil
}
func (m *mockUserRepo) GetByID(ctx context.Context, userID string) (*entity.User, error) {
	return nil, nil
}
func (m *mockUserRepo) UpdateUser(ctx context.Context, user *entity.User, roleIDs []string) error {
	return nil
}
func (m *mockUserRepo) DeleteUser(ctx context.Context, userID, reason string) error {
	return nil
}

func TestAuthUsecase_Login(t *testing.T) {
	hash, _ := auth.HashPassword("secret123")
	cfg := config.Config{
		JWTSecret:   "test-secret",
		JWTExpiryHr: 1,
	}
	bl := auth.NewMemoryBlacklist()

	t.Run("empty identifier returns ErrInvalidCredentials", func(t *testing.T) {
		uc := NewAuthUsecase(&mockUserRepo{}, bl, cfg)
		res, err := uc.Login(context.Background(), entity.LoginRequest{
			Identifier: "",
			Password:   "any",
		})
		assert.ErrorIs(t, err, entity.ErrInvalidCredentials)
		assert.Nil(t, res)
	})

	t.Run("empty password returns ErrInvalidCredentials", func(t *testing.T) {
		uc := NewAuthUsecase(&mockUserRepo{}, bl, cfg)
		res, err := uc.Login(context.Background(), entity.LoginRequest{
			Identifier: "user@test.com",
			Password:   "",
		})
		assert.ErrorIs(t, err, entity.ErrInvalidCredentials)
		assert.Nil(t, res)
	})

	t.Run("user not found returns ErrInvalidCredentials", func(t *testing.T) {
		repo := &mockUserRepo{
			findByEmailOrUsername: func(ctx context.Context, identifier string) (*entity.User, error) {
				return nil, nil
			},
		}
		uc := NewAuthUsecase(repo, bl, cfg)
		res, err := uc.Login(context.Background(), entity.LoginRequest{
			Identifier: "unknown@test.com",
			Password:   "secret123",
		})
		assert.ErrorIs(t, err, entity.ErrInvalidCredentials)
		assert.Nil(t, res)
	})

	t.Run("inactive user returns ErrAccountInactive", func(t *testing.T) {
		repo := &mockUserRepo{
			findByEmailOrUsername: func(ctx context.Context, identifier string) (*entity.User, error) {
				return &entity.User{
					ID:           "u1",
					Email:        "user@test.com",
					PasswordHash: hash,
					IsActive:     false,
				}, nil
			},
		}
		uc := NewAuthUsecase(repo, bl, cfg)
		res, err := uc.Login(context.Background(), entity.LoginRequest{
			Identifier: "user@test.com",
			Password:   "secret123",
		})
		assert.ErrorIs(t, err, entity.ErrAccountInactive)
		assert.Nil(t, res)
	})

	t.Run("wrong password returns ErrInvalidCredentials", func(t *testing.T) {
		repo := &mockUserRepo{
			findByEmailOrUsername: func(ctx context.Context, identifier string) (*entity.User, error) {
				return &entity.User{
					ID:           "u1",
					Email:        "user@test.com",
					PasswordHash: hash,
					IsActive:     true,
				}, nil
			},
			getUserRoles: func(ctx context.Context, userID string) ([]entity.RoleInfo, error) {
				return []entity.RoleInfo{{Code: "PESERTA"}}, nil
			},
		}
		uc := NewAuthUsecase(repo, bl, cfg)
		res, err := uc.Login(context.Background(), entity.LoginRequest{
			Identifier: "user@test.com",
			Password:   "wrongpassword",
		})
		assert.ErrorIs(t, err, entity.ErrInvalidCredentials)
		assert.Nil(t, res)
	})

	t.Run("success returns token and default_home_path", func(t *testing.T) {
		roles := []entity.RoleInfo{{ID: "r1", Code: "SUPER_ADMIN", Name: "Super Admin"}}
		repo := &mockUserRepo{
			findByEmailOrUsername: func(ctx context.Context, identifier string) (*entity.User, error) {
				return &entity.User{
					ID:           "u1",
					Name:         "Test User",
					Email:        "user@test.com",
					Username:     "user",
					PasswordHash: hash,
					IsActive:     true,
				}, nil
			},
			getUserRoles: func(ctx context.Context, userID string) ([]entity.RoleInfo, error) {
				return roles, nil
			},
		}
		uc := NewAuthUsecase(repo, bl, cfg)
		res, err := uc.Login(context.Background(), entity.LoginRequest{
			Identifier: "user@test.com",
			Password:   "secret123",
		})
		require.NoError(t, err)
		require.NotNil(t, res)
		assert.NotEmpty(t, res.AccessToken)
		assert.Equal(t, "bearer", res.TokenType)
		assert.Equal(t, "/dashboard", res.User.DefaultHomePath)
		assert.Equal(t, "u1", res.User.ID)
		assert.Equal(t, "Test User", res.User.Name)
	})
}

func TestAuthUsecase_Login_GetUserRolesError(t *testing.T) {
	hash, _ := auth.HashPassword("secret123")
	repo := &mockUserRepo{
		findByEmailOrUsername: func(ctx context.Context, identifier string) (*entity.User, error) {
			return &entity.User{ID: "u1", Email: "u@x.com", PasswordHash: hash, IsActive: true}, nil
		},
		getUserRoles: func(ctx context.Context, userID string) ([]entity.RoleInfo, error) {
			return nil, errors.New("db error")
		},
	}
	uc := NewAuthUsecase(repo, auth.NewMemoryBlacklist(), config.Config{JWTSecret: "s", JWTExpiryHr: 1})
	res, err := uc.Login(context.Background(), entity.LoginRequest{
		Identifier: "u@x.com",
		Password:   "secret123",
	})
	assert.Error(t, err)
	assert.Nil(t, res)
}
