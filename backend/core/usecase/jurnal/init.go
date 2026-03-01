package jurnal

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/core/repository"
)

type JurnalUsecase interface {
	ListPublished(ctx context.Context, req entity.GetJurnalListRequest) (*entity.GetJurnalListResponse, error)
	GetByID(ctx context.Context, id string) (*entity.JurnalDetailResponse, error)
	GetByIDForOwner(ctx context.Context, id, userID string) (*entity.Jurnal, error)
	Create(ctx context.Context, j *entity.Jurnal) error
	Update(ctx context.Context, j *entity.Jurnal) error
}

type jurnalUsecase struct {
	journalRepo repository.JournalRepo
}

func NewJurnalUsecase(journalRepo repository.JournalRepo) JurnalUsecase {
	return &jurnalUsecase{journalRepo: journalRepo}
}
