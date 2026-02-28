package assignmentrepo

import (
	"github.com/ProjectWidyaprada/backend/core/repository"
	"gorm.io/gorm"
)

type assignmentRepo struct {
	db *gorm.DB
}

func NewAssignmentRepo(db *gorm.DB) repository.AssignmentRepo {
	return &assignmentRepo{db: db}
}
