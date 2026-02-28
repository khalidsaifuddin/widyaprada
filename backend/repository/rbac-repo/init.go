package rbacrepo

import (
	"github.com/ProjectWidyaprada/backend/core/repository"
	"gorm.io/gorm"
)

type rbacRepo struct {
	db *gorm.DB
}

func NewRBACRepo(db *gorm.DB) repository.RBACRepo {
	return &rbacRepo{db: db}
}
