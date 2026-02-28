package exam

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/google/uuid"
)

func (u *examUsecase) Create(ctx context.Context, req entity.CreateExamRequest) (*entity.ExamDetailResponse, error) {
	existing, err := u.examRepo.GetByCode(ctx, req.Code, "")
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, entity.ErrExamCodeExists
	}

	// Build contents from question_ids and package_ids
	var contents []entity.ExamContentItem
	sortOrder := 1
	for _, qid := range req.QuestionIDs {
		contents = append(contents, entity.ExamContentItem{
			SourceType: entity.ExamContentSourceQuestion,
			SourceID:   qid,
			SortOrder:  sortOrder,
		})
		sortOrder++
	}
	for _, pid := range req.PackageIDs {
		contents = append(contents, entity.ExamContentItem{
			SourceType: entity.ExamContentSourcePackage,
			SourceID:   pid,
			SortOrder:  sortOrder,
		})
		sortOrder++
	}
	if len(contents) == 0 {
		return nil, entity.ErrExamMinContent
	}
	if len(req.ParticipantIDs) == 0 {
		return nil, entity.ErrExamMinParticipant
	}

	e := &entity.ExamDetail{
		ID:                   uuid.New().String(),
		Code:                 req.Code,
		Name:                 req.Name,
		JadwalMulai:          req.JadwalMulai,
		JadwalSelesai:        req.JadwalSelesai,
		DurasiMenit:          req.DurasiMenit,
		Status:               entity.ExamStatusDraft,
		VerificationStatus:   entity.QuestionVerifBelum,
		ShuffleQuestions:     req.ShuffleQuestions,
		TampilkanLeaderboard: req.TampilkanLeaderboard,
	}
	if err := u.examRepo.Create(ctx, e); err != nil {
		return nil, err
	}
	if err := u.examRepo.SetContents(ctx, e.ID, contents); err != nil {
		return nil, err
	}
	if err := u.examRepo.SetParticipants(ctx, e.ID, req.ParticipantIDs); err != nil {
		return nil, err
	}
	return u.Get(ctx, e.ID)
}
