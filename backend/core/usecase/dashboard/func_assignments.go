package dashboard

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *dashboardUsecase) GetAssignments(ctx context.Context, userID string, limit, page int64) (*entity.DashboardAssignmentsResponse, error) {
	items, total, err := u.assignmentRepo.ListForDashboard(ctx, userID, limit, page)
	if err != nil {
		return nil, err
	}
	return &entity.DashboardAssignmentsResponse{
		Data: items,
		Meta: struct {
			Total int64 `json:"total"`
		}{Total: total},
	}, nil
}
