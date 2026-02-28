package journalrepo

import (
	"github.com/ProjectWidyaprada/backend/core/repository"
	"gorm.io/gorm"
)

type journalRepo struct {
	db *gorm.DB
}

func NewJournalRepo(db *gorm.DB) repository.JournalRepo {
	return &journalRepo{db: db}
}
