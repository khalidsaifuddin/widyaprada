package forgotpassword

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ProjectWidyaprada/backend/config"
	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/core/repository"
	"github.com/ProjectWidyaprada/backend/pkg/auth"
	"github.com/ProjectWidyaprada/backend/pkg/email"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockUserRepo struct {
	findByEmail     func(ctx context.Context, email string) (*entity.User, error)
	updatePassword  func(ctx context.Context, userID, passwordHash string) error
}

func (m *mockUserRepo) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	if m.findByEmail != nil {
		return m.findByEmail(ctx, email)
	}
	return nil, nil
}
func (m *mockUserRepo) UpdatePassword(ctx context.Context, userID, passwordHash string) error {
	if m.updatePassword != nil {
		return m.updatePassword(ctx, userID, passwordHash)
	}
	return nil
}

func (m *mockUserRepo) FindByEmailOrUsername(ctx context.Context, identifier string) (*entity.User, error) {
	return nil, nil
}
func (m *mockUserRepo) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	return nil, nil
}
func (m *mockUserRepo) GetUserRoles(ctx context.Context, userID string) ([]entity.RoleInfo, error) {
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

type mockTokenRepo struct {
	create        func(ctx context.Context, userID, tokenHash string, expiresAt time.Time) error
	findByTokenHash func(ctx context.Context, tokenHash string) (*repository.PasswordResetToken, error)
	delete        func(ctx context.Context, id string) error
}

func (m *mockTokenRepo) Create(ctx context.Context, userID, tokenHash string, expiresAt time.Time) error {
	if m.create != nil {
		return m.create(ctx, userID, tokenHash, expiresAt)
	}
	return nil
}
func (m *mockTokenRepo) FindByTokenHash(ctx context.Context, tokenHash string) (*repository.PasswordResetToken, error) {
	if m.findByTokenHash != nil {
		return m.findByTokenHash(ctx, tokenHash)
	}
	return nil, nil
}
func (m *mockTokenRepo) Delete(ctx context.Context, id string) error {
	if m.delete != nil {
		return m.delete(ctx, id)
	}
	return nil
}

type mockEmailService struct{}

func (m *mockEmailService) SendPasswordEmail(params email.SendPasswordEmailParams) error {
	return nil
}
func (m *mockEmailService) SendPasswordResetLink(params email.SendPasswordResetLinkParams) error {
	return nil
}

func TestForgotPasswordUsecase_RequestReset(t *testing.T) {
	cfg := config.Config{
		FrontendURL:        "http://localhost:3000",
		ResetTokenExpiryHr: 1,
	}
	emailSvc := &mockEmailService{}

	t.Run("empty email returns ErrInvalidEmailFormat", func(t *testing.T) {
		uc := NewForgotPasswordUsecase(&mockUserRepo{}, &mockTokenRepo{}, emailSvc, cfg)
		res, err := uc.RequestReset(context.Background(), "")
		assert.ErrorIs(t, err, entity.ErrInvalidEmailFormat)
		assert.Nil(t, res)
	})

	t.Run("invalid email format returns ErrInvalidEmailFormat", func(t *testing.T) {
		uc := NewForgotPasswordUsecase(&mockUserRepo{}, &mockTokenRepo{}, emailSvc, cfg)
		res, err := uc.RequestReset(context.Background(), "invalid-email")
		assert.ErrorIs(t, err, entity.ErrInvalidEmailFormat)
		assert.Nil(t, res)
	})

	t.Run("email not registered returns success (security)", func(t *testing.T) {
		repo := &mockUserRepo{
			findByEmail: func(ctx context.Context, email string) (*entity.User, error) {
				return nil, nil
			},
		}
		uc := NewForgotPasswordUsecase(repo, &mockTokenRepo{}, emailSvc, cfg)
		res, err := uc.RequestReset(context.Background(), "unknown@test.com")
		require.NoError(t, err)
		require.NotNil(t, res)
		assert.Contains(t, res.Message, "Jika email Anda terdaftar")
	})

	t.Run("success creates token and returns message", func(t *testing.T) {
		var createdTokenHash string
		repo := &mockUserRepo{
			findByEmail: func(ctx context.Context, email string) (*entity.User, error) {
				return &entity.User{ID: "u1", Name: "Test", Email: email}, nil
			},
		}
		tokenRepo := &mockTokenRepo{
			create: func(ctx context.Context, userID, tokenHash string, expiresAt time.Time) error {
				createdTokenHash = tokenHash
				return nil
			},
		}
		uc := NewForgotPasswordUsecase(repo, tokenRepo, emailSvc, cfg)
		res, err := uc.RequestReset(context.Background(), "user@test.com")
		require.NoError(t, err)
		require.NotNil(t, res)
		assert.NotEmpty(t, createdTokenHash)
	})
}

func TestForgotPasswordUsecase_ResetPassword(t *testing.T) {
	plainToken, hashedToken, _ := auth.GenerateResetToken()
	cfg := config.Config{}
	emailSvc := &mockEmailService{}

	t.Run("empty token returns ErrResetTokenInvalid", func(t *testing.T) {
		uc := NewForgotPasswordUsecase(&mockUserRepo{}, &mockTokenRepo{}, emailSvc, cfg)
		res, err := uc.ResetPassword(context.Background(), "", "newpass123", "newpass123")
		assert.ErrorIs(t, err, entity.ErrResetTokenInvalid)
		assert.Nil(t, res)
	})

	t.Run("password too short returns ErrPasswordTooShort", func(t *testing.T) {
		uc := NewForgotPasswordUsecase(&mockUserRepo{}, &mockTokenRepo{}, emailSvc, cfg)
		res, err := uc.ResetPassword(context.Background(), "token", "short", "short")
		assert.ErrorIs(t, err, entity.ErrPasswordTooShort)
		assert.Nil(t, res)
	})

	t.Run("password mismatch returns ErrPasswordMismatch", func(t *testing.T) {
		uc := NewForgotPasswordUsecase(&mockUserRepo{}, &mockTokenRepo{}, emailSvc, cfg)
		res, err := uc.ResetPassword(context.Background(), "token", "newpass123", "different")
		assert.ErrorIs(t, err, entity.ErrPasswordMismatch)
		assert.Nil(t, res)
	})

	t.Run("invalid token returns ErrResetTokenInvalid", func(t *testing.T) {
		tokenRepo := &mockTokenRepo{
			findByTokenHash: func(ctx context.Context, tokenHash string) (*repository.PasswordResetToken, error) {
				return nil, nil
			},
		}
		uc := NewForgotPasswordUsecase(&mockUserRepo{}, tokenRepo, emailSvc, cfg)
		res, err := uc.ResetPassword(context.Background(), "invalid-token", "newpass123", "newpass123")
		assert.ErrorIs(t, err, entity.ErrResetTokenInvalid)
		assert.Nil(t, res)
	})

	t.Run("success updates password", func(t *testing.T) {
		var updatedUserID, updatedHash string
		tokenRepo := &mockTokenRepo{
			findByTokenHash: func(ctx context.Context, tokenHash string) (*repository.PasswordResetToken, error) {
				if tokenHash == hashedToken {
					return &repository.PasswordResetToken{ID: "t1", UserID: "u1", TokenHash: tokenHash}, nil
				}
				return nil, nil
			},
			delete: func(ctx context.Context, id string) error {
				return nil
			},
		}
		userRepo := &mockUserRepo{
			updatePassword: func(ctx context.Context, userID, passwordHash string) error {
				updatedUserID = userID
				updatedHash = passwordHash
				return nil
			},
		}
		uc := NewForgotPasswordUsecase(userRepo, tokenRepo, emailSvc, cfg)
		res, err := uc.ResetPassword(context.Background(), plainToken, "newpass123", "newpass123")
		require.NoError(t, err)
		require.NotNil(t, res)
		assert.Equal(t, "u1", updatedUserID)
		assert.NotEmpty(t, updatedHash)
		assert.True(t, auth.VerifyPassword("newpass123", updatedHash))
	})
}

func TestForgotPasswordUsecase_ResetPassword_FindByTokenHashError(t *testing.T) {
	tokenRepo := &mockTokenRepo{
		findByTokenHash: func(ctx context.Context, tokenHash string) (*repository.PasswordResetToken, error) {
			return nil, errors.New("db error")
		},
	}
	uc := NewForgotPasswordUsecase(&mockUserRepo{}, tokenRepo, &mockEmailService{}, config.Config{})
	res, err := uc.ResetPassword(context.Background(), "any-token", "newpass123", "newpass123")
	assert.Error(t, err)
	assert.Nil(t, res)
}
