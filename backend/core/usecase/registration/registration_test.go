package registration

import (
	"context"
	"errors"
	"testing"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/email"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockUserRepo struct {
	findByEmail   func(ctx context.Context, email string) (*entity.User, error)
	getRoleByCode func(ctx context.Context, code string) (*entity.RoleInfo, error)
	createUser    func(ctx context.Context, user *entity.User, roleID string) error
}

func (m *mockUserRepo) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	if m.findByEmail != nil {
		return m.findByEmail(ctx, email)
	}
	return nil, nil
}

func (m *mockUserRepo) GetRoleByCode(ctx context.Context, code string) (*entity.RoleInfo, error) {
	if m.getRoleByCode != nil {
		return m.getRoleByCode(ctx, code)
	}
	return nil, nil
}

func (m *mockUserRepo) CreateUser(ctx context.Context, user *entity.User, roleID string) error {
	if m.createUser != nil {
		return m.createUser(ctx, user, roleID)
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

type mockEmailService struct{}

func (m *mockEmailService) SendPasswordEmail(params email.SendPasswordEmailParams) error {
	return nil
}

func (m *mockEmailService) SendPasswordResetLink(params email.SendPasswordResetLinkParams) error {
	return nil
}

func TestRegistrationUsecase_Register(t *testing.T) {
	emailSvc := &mockEmailService{}

	t.Run("empty name returns ErrValidation", func(t *testing.T) {
		uc := NewRegistrationUsecase(&mockUserRepo{}, emailSvc)
		res, err := uc.Register(context.Background(), entity.RegisterRequest{
			Name:  "",
			Email: "user@test.com",
		})
		assert.ErrorIs(t, err, entity.ErrValidation)
		assert.Nil(t, res)
	})

	t.Run("empty email returns ErrInvalidEmailFormat", func(t *testing.T) {
		uc := NewRegistrationUsecase(&mockUserRepo{}, emailSvc)
		res, err := uc.Register(context.Background(), entity.RegisterRequest{
			Name:  "Test",
			Email: "",
		})
		assert.ErrorIs(t, err, entity.ErrInvalidEmailFormat)
		assert.Nil(t, res)
	})

	t.Run("invalid email format returns ErrInvalidEmailFormat", func(t *testing.T) {
		uc := NewRegistrationUsecase(&mockUserRepo{}, emailSvc)
		res, err := uc.Register(context.Background(), entity.RegisterRequest{
			Name:  "Test",
			Email: "not-an-email",
		})
		assert.ErrorIs(t, err, entity.ErrInvalidEmailFormat)
		assert.Nil(t, res)
	})

	t.Run("email already registered returns ErrEmailAlreadyRegistered", func(t *testing.T) {
		repo := &mockUserRepo{
			findByEmail: func(ctx context.Context, email string) (*entity.User, error) {
				return &entity.User{ID: "u1", Email: email}, nil
			},
		}
		uc := NewRegistrationUsecase(repo, emailSvc)
		res, err := uc.Register(context.Background(), entity.RegisterRequest{
			Name:  "Test",
			Email: "existing@test.com",
		})
		assert.ErrorIs(t, err, entity.ErrEmailAlreadyRegistered)
		assert.Nil(t, res)
	})

	t.Run("role not found returns ErrRoleNotFound", func(t *testing.T) {
		repo := &mockUserRepo{
			findByEmail: func(ctx context.Context, email string) (*entity.User, error) {
				return nil, nil
			},
			getRoleByCode: func(ctx context.Context, code string) (*entity.RoleInfo, error) {
				return nil, nil
			},
		}
		uc := NewRegistrationUsecase(repo, emailSvc)
		res, err := uc.Register(context.Background(), entity.RegisterRequest{
			Name:  "Test",
			Email: "new@test.com",
		})
		assert.ErrorIs(t, err, entity.ErrRoleNotFound)
		assert.Nil(t, res)
	})

	t.Run("success creates user and returns message", func(t *testing.T) {
		var createdUser *entity.User
		repo := &mockUserRepo{
			findByEmail: func(ctx context.Context, email string) (*entity.User, error) {
				return nil, nil
			},
			getRoleByCode: func(ctx context.Context, code string) (*entity.RoleInfo, error) {
				return &entity.RoleInfo{ID: "r1", Code: entity.DefaultRoleCodePeserta}, nil
			},
			createUser: func(ctx context.Context, user *entity.User, roleID string) error {
				createdUser = user
				return nil
			},
		}
		uc := NewRegistrationUsecase(repo, emailSvc)
		res, err := uc.Register(context.Background(), entity.RegisterRequest{
			Name:  "New User",
			Email: "newuser@test.com",
			NIP:   "12345",
		})
		require.NoError(t, err)
		require.NotNil(t, res)
		assert.Contains(t, res.Message, "Registrasi berhasil")
		require.NotNil(t, createdUser)
		assert.Equal(t, "New User", createdUser.Name)
		assert.Equal(t, "newuser@test.com", createdUser.Email)
		assert.Equal(t, "newuser@test.com", createdUser.Username)
		assert.Equal(t, "12345", createdUser.NIP)
		assert.True(t, createdUser.IsActive)
		assert.NotEmpty(t, createdUser.PasswordHash)
	})
}

func TestRegistrationUsecase_Register_FindByEmailError(t *testing.T) {
	repo := &mockUserRepo{
		findByEmail: func(ctx context.Context, email string) (*entity.User, error) {
			return nil, errors.New("db error")
		},
	}
	uc := NewRegistrationUsecase(repo, &mockEmailService{})
	res, err := uc.Register(context.Background(), entity.RegisterRequest{
		Name:  "Test",
		Email: "user@test.com",
	})
	assert.Error(t, err)
	assert.Nil(t, res)
}
