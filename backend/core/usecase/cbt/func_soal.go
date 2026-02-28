package cbt

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *cbtUsecase) GetSoal(ctx context.Context, userID, attemptID string) (*entity.CBTListQuestionsResponse, error) {
	examID, attemptUserID, submitted, err := u.cbtRepo.GetAttemptByID(ctx, attemptID)
	if err != nil || examID == "" {
		return nil, entity.ErrCBTAttemptNotFound
	}
	if attemptUserID != userID {
		return nil, entity.ErrCBTAttemptNotOwned
	}
	if submitted {
		return nil, entity.ErrCBTAlreadySubmitted
	}

	questionIDs, err := u.cbtRepo.GetQuestionIDsForAttempt(ctx, attemptID)
	if err != nil {
		return nil, err
	}

	questions := make([]entity.CBTQuestionItem, 0, len(questionIDs))
	for i, qid := range questionIDs {
		q, err := u.cbtRepo.GetQuestionForCBT(ctx, qid)
		if err != nil {
			continue
		}
		q.Num = i + 1
		questions = append(questions, *q)
	}

	return &entity.CBTListQuestionsResponse{
		AttemptID: attemptID,
		Total:     len(questions),
		Questions: questions,
	}, nil
}

func (u *cbtUsecase) GetSoalByNomor(ctx context.Context, userID, attemptID string, num int) (*entity.CBTQuestionItem, error) {
	_, attemptUserID, submitted, err := u.cbtRepo.GetAttemptByID(ctx, attemptID)
	if err != nil || attemptUserID == "" {
		return nil, entity.ErrCBTAttemptNotFound
	}
	if attemptUserID != userID {
		return nil, entity.ErrCBTAttemptNotOwned
	}
	if submitted {
		return nil, entity.ErrCBTAlreadySubmitted
	}

	questionIDs, err := u.cbtRepo.GetQuestionIDsForAttempt(ctx, attemptID)
	if err != nil {
		return nil, err
	}
	if num < 1 || num > len(questionIDs) {
		return nil, entity.ErrCBTAttemptNotFound
	}

	qid := questionIDs[num-1]
	q, err := u.cbtRepo.GetQuestionForCBT(ctx, qid)
	if err != nil {
		return nil, err
	}
	q.Num = num
	return q, nil
}
