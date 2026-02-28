package userrepo

import (
	"github.com/ProjectWidyaprada/backend/core/repository"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

// NewUserRepo membuat instance UserRepo
func NewUserRepo(db *gorm.DB) repository.UserRepo {
	return &userRepo{db: db}
}
