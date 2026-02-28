package auth

import (
	"context"

	"github.com/ProjectWidyaprada/backend/config"
	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/core/repository"
	"github.com/ProjectWidyaprada/backend/pkg/auth"
)

type AuthUsecase interface {
	Login(ctx context.Context, req entity.LoginRequest) (*entity.LoginResponse, error)
	Logout(ctx context.Context, tokenString string) (*entity.LogoutResponse, error)
}

type authUsecase struct {
	userRepo  repository.UserRepo
	blacklist *auth.MemoryBlacklist
	cfg       config.Config
}

func NewAuthUsecase(userRepo repository.UserRepo, blacklist *auth.MemoryBlacklist, cfg config.Config) AuthUsecase {
	return &authUsecase{
		userRepo:  userRepo,
		blacklist: blacklist,
		cfg:       cfg,
	}
}
