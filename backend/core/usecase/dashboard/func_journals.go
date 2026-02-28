package dashboard

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *dashboardUsecase) GetMyJournals(ctx context.Context, userID string, limit, page int64) (*entity.DashboardJournalsResponse, error) {
	items, total, err := u.journalRepo.ListMyJournals(ctx, userID, limit, page)
	if err != nil {
		return nil, err
	}
	return &entity.DashboardJournalsResponse{
		Data: items,
		Meta: struct {
			Total int64 `json:"total"`
		}{Total: total},
	}, nil
}
