package examrepo

import (
	"github.com/ProjectWidyaprada/backend/core/repository"
	"gorm.io/gorm"
)

type examRepo struct {
	db *gorm.DB
}

func NewExamRepo(db *gorm.DB) repository.ExamRepo {
	return &examRepo{db: db}
}
