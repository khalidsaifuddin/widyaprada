package packagerepo

import (
	"github.com/ProjectWidyaprada/backend/core/repository"
	"gorm.io/gorm"
)

type packageRepo struct {
	db *gorm.DB
}

func NewPackageRepo(db *gorm.DB) repository.PackageRepo {
	return &packageRepo{db: db}
}
