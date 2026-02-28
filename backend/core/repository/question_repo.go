package repository

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

type QuestionRepo interface {
	// Categories
	GetCategoryByID(ctx context.Context, id string) (*entity.QuestionCategory, error)
	ListCategories(ctx context.Context) ([]entity.QuestionCategory, error)

	// Questions
	List(ctx context.Context, req entity.GetQuestionListRequest) (*entity.GetQuestionListResponse, error)
	GetByID(ctx context.Context, id string) (*entity.QuestionDetail, error)
	GetByCode(ctx context.Context, code string, excludeID string) (*entity.QuestionDetail, error)
	Create(ctx context.Context, q *entity.QuestionDetail) error
	Update(ctx context.Context, q *entity.QuestionDetail) error
	Delete(ctx context.Context, id string, reason string) error

	// Options
	GetOptionsByQuestionID(ctx context.Context, questionID string) ([]entity.QuestionOption, error)
	CreateOptions(ctx context.Context, questionID string, opts []entity.QuestionOption) error
	DeleteOptionsByQuestionID(ctx context.Context, questionID string) error

	// Paket check (untuk validasi delete)
	IsQuestionUsedInPackage(ctx context.Context, questionID string) (bool, error)
}
