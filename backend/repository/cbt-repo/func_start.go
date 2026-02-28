package cbtrepo

import (
	"context"
	"math/rand"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/repository/exam-repo"
	"github.com/google/uuid"
)

func (r *cbtRepo) CreateAttempt(ctx context.Context, examID, userID string) (*entity.CBTStartResponse, error) {
	var exam examrepo.Exam
	if err := r.db.WithContext(ctx).Where("id = ?", examID).First(&exam).Error; err != nil {
		return nil, err
	}

	// Cek sudah submit?
	var cnt int64
	r.db.WithContext(ctx).Model(&examrepo.ExamAttempt{}).
		Where("exam_id = ? AND user_id = ? AND submitted_at IS NOT NULL", examID, userID).
		Count(&cnt)
	if cnt > 0 {
		return nil, entity.ErrCBTAlreadyStarted
	}

	// Sudah ada attempt yang belum submit?
	var existing examrepo.ExamAttempt
	err := r.db.WithContext(ctx).Where("exam_id = ? AND user_id = ? AND submitted_at IS NULL", examID, userID).First(&existing).Error
	if err == nil {
		startedAt := ""
		if existing.StartedAt != nil {
			startedAt = existing.StartedAt.UTC().Format(time.RFC3339)
		}
		jadwalSelesai := ""
		if exam.JadwalSelesai != nil {
			jadwalSelesai = exam.JadwalSelesai.UTC().Format(time.RFC3339)
		}
		return &entity.CBTStartResponse{
			AttemptID:     existing.ID,
			ExamID:        examID,
			StartedAt:     startedAt,
			DurasiMenit:   exam.DurasiMenit,
			JadwalSelesai: jadwalSelesai,
		}, nil
	}

	attemptID := uuid.New().String()
	now := time.Now().UTC()

	attempt := examrepo.ExamAttempt{
		ID:        attemptID,
		ExamID:    examID,
		UserID:    userID,
		StartedAt: &now,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	if err := r.db.WithContext(ctx).Create(&attempt).Error; err != nil {
		return nil, err
	}

	// Expand contents -> question IDs
	questionIDs, err := r.expandContentsToQuestionIDs(ctx, examID)
	if err != nil {
		return nil, err
	}

	if exam.ShuffleQuestions && len(questionIDs) > 0 {
		rnd := rand.New(rand.NewSource(now.UnixNano()))
		rnd.Shuffle(len(questionIDs), func(i, j int) { questionIDs[i], questionIDs[j] = questionIDs[j], questionIDs[i] })
	}

	for i, qid := range questionIDs {
		aq := examrepo.ExamAttemptQuestion{
			AttemptID:  attemptID,
			QuestionID: qid,
			SortOrder:  i + 1,
			CreatedAt:  &now,
		}
		if err := r.db.WithContext(ctx).Create(&aq).Error; err != nil {
			return nil, err
		}
	}

	jadwalSelesai := ""
	if exam.JadwalSelesai != nil {
		jadwalSelesai = exam.JadwalSelesai.UTC().Format(time.RFC3339)
	}

	return &entity.CBTStartResponse{
		AttemptID:     attemptID,
		ExamID:        examID,
		StartedAt:     now.Format(time.RFC3339),
		DurasiMenit:   exam.DurasiMenit,
		JadwalSelesai: jadwalSelesai,
	}, nil
}

func (r *cbtRepo) expandContentsToQuestionIDs(ctx context.Context, examID string) ([]string, error) {
	var contents []examrepo.ExamContent
	if err := r.db.WithContext(ctx).Where("exam_id = ?", examID).Order("sort_order").Find(&contents).Error; err != nil {
		return nil, err
	}

	var questionIDs []string
	for _, c := range contents {
		if c.SourceType == entity.ExamContentSourceQuestion {
			questionIDs = append(questionIDs, c.SourceID)
		} else if c.SourceType == entity.ExamContentSourcePackage {
			var items []struct {
				QuestionID string
			}
			if err := r.db.WithContext(ctx).Table("package_question_items").
				Select("question_id").
				Where("package_id = ?", c.SourceID).
				Order("sort_order").
				Scan(&items).Error; err != nil {
				return nil, err
			}
			for _, it := range items {
				questionIDs = append(questionIDs, it.QuestionID)
			}
		}
	}
	return questionIDs, nil
}
