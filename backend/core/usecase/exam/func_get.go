package exam

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *examUsecase) Get(ctx context.Context, id string) (*entity.ExamDetailResponse, error) {
	e, err := u.examRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if e == nil {
		return nil, entity.ErrExamNotFound
	}
	contents, _ := u.examRepo.GetContentsByExamID(ctx, id)
	participants, _ := u.examRepo.GetParticipantsByExamID(ctx, id)

	contentInfos := make([]entity.ExamContentInfo, len(contents))
	for i := range contents {
		contentInfos[i] = entity.ExamContentInfo{
			SourceType: contents[i].SourceType,
			SourceID:   contents[i].SourceID,
			SortOrder:  contents[i].SortOrder,
		}
	}

	partInfos := make([]entity.ExamParticipantInfo, len(participants))
	for i := range participants {
		user, _ := u.userRepo.GetByID(ctx, participants[i].UserID)
		name := ""
		if user != nil {
			name = user.Name
		}
		partInfos[i] = entity.ExamParticipantInfo{
			UserID:   participants[i].UserID,
			UserName: name,
		}
	}

	return &entity.ExamDetailResponse{
		ID:                   e.ID,
		Code:                 e.Code,
		Name:                 e.Name,
		JadwalMulai:          e.JadwalMulai,
		JadwalSelesai:        e.JadwalSelesai,
		DurasiMenit:          e.DurasiMenit,
		Status:               e.Status,
		VerificationStatus:   e.VerificationStatus,
		ShuffleQuestions:     e.ShuffleQuestions,
		TampilkanLeaderboard: e.TampilkanLeaderboard,
		Contents:             contentInfos,
		Participants:         partInfos,
		CreatedAt:            e.CreatedAt,
		UpdatedAt:            e.UpdatedAt,
	}, nil
}
