package articlerepo

import (
	"github.com/ProjectWidyaprada/backend/core/repository"
	"gorm.io/gorm"
)

type articleRepo struct {
	db *gorm.DB
}

func NewArticleRepo(db *gorm.DB) repository.ArticleRepo {
	return &articleRepo{db: db}
}
