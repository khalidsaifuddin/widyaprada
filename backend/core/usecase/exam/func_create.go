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

	// Build contents: prefer Contents (ordered), else fallback to QuestionIDs + PackageIDs
	var contents []entity.ExamContentItem
	if len(req.Contents) > 0 {
		for i, c := range req.Contents {
			if c.SourceType != entity.ExamContentSourceQuestion && c.SourceType != entity.ExamContentSourcePackage {
				continue
			}
			contents = append(contents, entity.ExamContentItem{
				SourceType: c.SourceType,
				SourceID:   c.SourceID,
				SortOrder:  i + 1,
			})
		}
	} else {
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
	}
	if len(contents) == 0 {
		return nil, entity.ErrExamMinContent
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
