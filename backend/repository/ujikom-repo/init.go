package ujikomrepo

import (
	"github.com/ProjectWidyaprada/backend/core/repository"
	"gorm.io/gorm"
)

type ujikomRepo struct {
	db *gorm.DB
}

func NewUjikomRepo(db *gorm.DB) repository.UjikomRepo {
	return &ujikomRepo{db: db}
}
