package assignment

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *assignmentUsecase) GetLeaderboard(ctx context.Context, userID, examID string) (*entity.GetLeaderboardResponse, error) {
	ok, err := u.assignRepo.IsParticipant(ctx, examID, userID)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, entity.ErrAssignmentForbidden
	}

	name, tampilkan, err := u.assignRepo.GetExamForLeaderboard(ctx, examID)
	if err != nil {
		return nil, err
	}
	if name == "" {
		return nil, entity.ErrAssignmentForbidden
	}
	if !tampilkan {
		return nil, entity.ErrLeaderboardPrivate
	}

	items, err := u.assignRepo.GetLeaderboard(ctx, examID)
	if err != nil {
		return nil, err
	}

	return &entity.GetLeaderboardResponse{
		ExamID:   examID,
		ExamName: name,
		Items:    items,
	}, nil
}
