package repository

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

type JournalRepo interface {
	ListPublished(ctx context.Context, req entity.GetJurnalListRequest) (*entity.GetJurnalListResponse, error)
	GetByID(ctx context.Context, id string) (*entity.Jurnal, error)
	ListPublishedForLanding(ctx context.Context, limit int) ([]entity.JurnalPublicItem, error)
	Create(ctx context.Context, j *entity.Jurnal) error

	// ListMyJournals jurnal milik user (untuk dashboard)
	ListMyJournals(ctx context.Context, userID string, limit, page int64) ([]entity.DashboardJournalItem, int64, error)
}
