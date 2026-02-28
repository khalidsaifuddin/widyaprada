package forgotpassword

import (
	"context"

	"github.com/ProjectWidyaprada/backend/config"
	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/core/repository"
	"github.com/ProjectWidyaprada/backend/pkg/email"
)

type ForgotPasswordUsecase interface {
	RequestReset(ctx context.Context, emailAddr string) (*entity.ForgotPasswordResponse, error)
	ResetPassword(ctx context.Context, token, password, passwordConfirm string) (*entity.ResetPasswordResponse, error)
}

type forgotPasswordUsecase struct {
	userRepo       repository.UserRepo
	tokenRepo      repository.PasswordResetTokenRepo
	emailService   email.EmailService
	cfg            config.Config
}

func NewForgotPasswordUsecase(
	userRepo repository.UserRepo,
	tokenRepo repository.PasswordResetTokenRepo,
	emailService email.EmailService,
	cfg config.Config,
) ForgotPasswordUsecase {
	return &forgotPasswordUsecase{
		userRepo:     userRepo,
		tokenRepo:    tokenRepo,
		emailService: emailService,
		cfg:          cfg,
	}
}
