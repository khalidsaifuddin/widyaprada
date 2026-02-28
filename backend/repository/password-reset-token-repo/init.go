package passwordresettokenrepo

import (
	"github.com/ProjectWidyaprada/backend/core/repository"
	"gorm.io/gorm"
)

type passwordResetTokenRepo struct {
	db *gorm.DB
}

func NewPasswordResetTokenRepo(db *gorm.DB) repository.PasswordResetTokenRepo {
	return &passwordResetTokenRepo{db: db}
}
