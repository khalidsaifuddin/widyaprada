package banksoal

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/core/repository"
)

type BankSoalUsecase interface {
	List(ctx context.Context, req entity.GetQuestionListRequest) (*entity.GetQuestionListResponse, error)
	Get(ctx context.Context, id string) (*entity.QuestionDetailResponse, error)
	Create(ctx context.Context, req entity.CreateQuestionRequest) (*entity.QuestionDetailResponse, error)
	Update(ctx context.Context, id string, req entity.UpdateQuestionRequest) (*entity.QuestionDetailResponse, error)
	Delete(ctx context.Context, id string, reason string) error
	Verify(ctx context.Context, id string) error
	Unverify(ctx context.Context, id string) error

	// Categories (untuk dropdown dll)
	ListCategories(ctx context.Context) ([]entity.QuestionCategory, error)
}

type bankSoalUsecase struct {
	questionRepo repository.QuestionRepo
}

func NewBankSoalUsecase(questionRepo repository.QuestionRepo) BankSoalUsecase {
	return &bankSoalUsecase{questionRepo: questionRepo}
}
