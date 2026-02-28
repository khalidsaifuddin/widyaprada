package assignment

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *assignmentUsecase) GetApplyStatus(ctx context.Context, userID string) (*entity.ApplyStatusResponse, error) {
	return u.ujikomRepo.GetLatestApplicationByUser(ctx, userID)
}
