package sliderepo

import (
	"github.com/ProjectWidyaprada/backend/core/repository"
	"gorm.io/gorm"
)

type slideRepo struct {
	db *gorm.DB
}

func NewSlideRepo(db *gorm.DB) repository.SlideRepo {
	return &slideRepo{db: db}
}
