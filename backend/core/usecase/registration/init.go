package registration

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/core/repository"
	"github.com/ProjectWidyaprada/backend/pkg/email"
)

type RegistrationUsecase interface {
	Register(ctx context.Context, req entity.RegisterRequest) (*entity.RegisterResponse, error)
}

type registrationUsecase struct {
	userRepo     repository.UserRepo
	emailService email.EmailService
}

func NewRegistrationUsecase(userRepo repository.UserRepo, emailService email.EmailService) RegistrationUsecase {
	return &registrationUsecase{
		userRepo:     userRepo,
		emailService: emailService,
	}
}
