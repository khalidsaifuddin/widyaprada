package examplerepo

import (
	"github.com/ProjectWidyaprada/backend/core/repository"
	"gorm.io/gorm"
)

type exampleRepo struct {
	db *gorm.DB
}

func NewExampleRepo(db *gorm.DB) repository.ExampleRepo {
	return &exampleRepo{db: db}
}
