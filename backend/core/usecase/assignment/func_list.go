package assignment

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *assignmentUsecase) ListAssignments(ctx context.Context, userID string, req entity.GetAssignmentListRequest) (*entity.GetAssignmentListResponse, error) {
	return u.assignRepo.ListMyAssignments(ctx, userID, req)
}
