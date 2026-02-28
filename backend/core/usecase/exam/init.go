package exam

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/core/repository"
)

type ExamUsecase interface {
	List(ctx context.Context, req entity.GetExamListRequest) (*entity.GetExamListResponse, error)
	Get(ctx context.Context, id string) (*entity.ExamDetailResponse, error)
	Create(ctx context.Context, req entity.CreateExamRequest) (*entity.ExamDetailResponse, error)
	Update(ctx context.Context, id string, req entity.UpdateExamRequest) (*entity.ExamDetailResponse, error)
	Delete(ctx context.Context, id string, reason string) error
	Publish(ctx context.Context, id string) error
	Verify(ctx context.Context, id string) error
	Unverify(ctx context.Context, id string) error
}

type examUsecase struct {
	examRepo    repository.ExamRepo
	userRepo    repository.UserRepo
}

func NewExamUsecase(examRepo repository.ExamRepo, userRepo repository.UserRepo) ExamUsecase {
	return &examUsecase{
		examRepo: examRepo,
		userRepo: userRepo,
	}
}
