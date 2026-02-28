package wpdatarepo

import (
	"github.com/ProjectWidyaprada/backend/core/repository"
	"gorm.io/gorm"
)

type wpdataRepo struct {
	db *gorm.DB
}

func NewWPDataRepo(db *gorm.DB) repository.WPDataRepo {
	return &wpdataRepo{db: db}
}
