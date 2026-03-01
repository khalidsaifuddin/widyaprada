package cbt

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *cbtUsecase) GetExamDetail(ctx context.Context, userID, examID string) (*entity.CBTExamDetailResponse, error) {
	return u.cbtRepo.GetExamByIDForInstructions(ctx, examID, userID)
}
