package exam

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *examUsecase) Update(ctx context.Context, id string, req entity.UpdateExamRequest) (*entity.ExamDetailResponse, error) {
	e, err := u.examRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if e == nil {
		return nil, entity.ErrExamNotFound
	}
	if e.Status != entity.ExamStatusDraft {
		return nil, entity.ErrExamNotDraft
	}

	if req.Code != "" {
		existing, _ := u.examRepo.GetByCode(ctx, req.Code, id)
		if existing != nil {
			return nil, entity.ErrExamCodeExists
		}
		e.Code = req.Code
	}
	if req.Name != "" {
		e.Name = req.Name
	}
	if req.JadwalMulai != "" {
		e.JadwalMulai = req.JadwalMulai
	}
	if req.JadwalSelesai != "" {
		e.JadwalSelesai = req.JadwalSelesai
	}
	if req.DurasiMenit > 0 {
		e.DurasiMenit = req.DurasiMenit
	}
	if req.ShuffleQuestions != nil {
		e.ShuffleQuestions = *req.ShuffleQuestions
	}
	if req.TampilkanLeaderboard != nil {
		e.TampilkanLeaderboard = *req.TampilkanLeaderboard
	}

	if len(req.QuestionIDs) > 0 || len(req.PackageIDs) > 0 {
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
		if err := u.examRepo.SetContents(ctx, id, contents); err != nil {
			return nil, err
		}
	}
	if len(req.ParticipantIDs) > 0 {
		if err := u.examRepo.SetParticipants(ctx, id, req.ParticipantIDs); err != nil {
			return nil, err
		}
	}

	if err := u.examRepo.Update(ctx, e); err != nil {
		return nil, err
	}
	return u.Get(ctx, id)
}
