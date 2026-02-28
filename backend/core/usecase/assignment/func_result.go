package assignment

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *assignmentUsecase) GetAssignmentResult(ctx context.Context, userID, examID string) (*entity.AssignmentResultResponse, error) {
	ok, err := u.assignRepo.IsParticipant(ctx, examID, userID)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, entity.ErrAssignmentForbidden
	}
	return u.assignRepo.GetAttemptByExamUser(ctx, examID, userID)
}
