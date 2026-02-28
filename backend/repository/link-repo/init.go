package linkrepo

import (
	"github.com/ProjectWidyaprada/backend/core/repository"
	"gorm.io/gorm"
)

type linkRepo struct {
	db *gorm.DB
}

func NewLinkRepo(db *gorm.DB) repository.LinkRepo {
	return &linkRepo{db: db}
}
