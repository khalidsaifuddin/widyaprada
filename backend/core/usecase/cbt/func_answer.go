package cbt

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *cbtUsecase) SimpanJawaban(ctx context.Context, userID, attemptID string, req entity.CBTSaveAnswerRequest) error {
	_, attemptUserID, submitted, err := u.cbtRepo.GetAttemptByID(ctx, attemptID)
	if err != nil || attemptUserID == "" {
		return entity.ErrCBTAttemptNotFound
	}
	if attemptUserID != userID {
		return entity.ErrCBTAttemptNotOwned
	}
	if submitted {
		return entity.ErrCBTAlreadySubmitted
	}

	answerValue := req.OptionID
	if req.AnswerText != "" {
		answerValue = req.AnswerText
	}
	return u.cbtRepo.SaveAnswer(ctx, attemptID, req.QuestionID, answerValue)
}
