package dokumenpersyaratanrepo

import (
	"github.com/ProjectWidyaprada/backend/core/repository"
	"gorm.io/gorm"
)

type dokumenPersyaratanRepo struct {
	db *gorm.DB
}

func NewDokumenPersyaratanRepo(db *gorm.DB) repository.DokumenPersyaratanRepo {
	return &dokumenPersyaratanRepo{db: db}
}
