package cbtrepo

import (
	"github.com/ProjectWidyaprada/backend/core/repository"
	"gorm.io/gorm"
)

type cbtRepo struct {
	db *gorm.DB
}

func NewCBTRepo(db *gorm.DB) repository.CBTRepo {
	return &cbtRepo{db: db}
}
