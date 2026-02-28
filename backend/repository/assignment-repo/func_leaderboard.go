package assignmentrepo

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/repository/exam-repo"
	"github.com/ProjectWidyaprada/backend/repository/user-repo"
	"gorm.io/gorm"
)

func (r *assignmentRepo) GetLeaderboard(ctx context.Context, examID string) ([]entity.LeaderboardItem, error) {
	var attempts []examrepo.ExamAttempt
	err := r.db.WithContext(ctx).
		Where("exam_id = ? AND submitted_at IS NOT NULL AND score IS NOT NULL", examID).
		Order("score DESC").
		Find(&attempts).Error
	if err != nil {
		return nil, err
	}

	items := make([]entity.LeaderboardItem, 0, len(attempts))
	for i := range attempts {
		att := &attempts[i]
		score := 0.0
		if att.Score != nil {
			score = *att.Score
		}
		var u userrepo.User
		if err := r.db.WithContext(ctx).Where("id = ?", att.UserID).First(&u).Error; err == nil {
			items = append(items, entity.LeaderboardItem{
				Rank:   i + 1,
				UserID: att.UserID,
				Name:   u.Name,
				Score:  score,
			})
		} else {
			items = append(items, entity.LeaderboardItem{
				Rank:   i + 1,
				UserID: att.UserID,
				Name:   "-",
				Score:  score,
			})
		}
	}
	return items, nil
}

func (r *assignmentRepo) IsParticipant(ctx context.Context, examID, userID string) (bool, error) {
	var n int64
	err := r.db.WithContext(ctx).Table("exam_participants").
		Where("exam_id = ? AND user_id = ?", examID, userID).
		Count(&n).Error
	return n > 0, err
}

func (r *assignmentRepo) GetExamForLeaderboard(ctx context.Context, examID string) (name string, tampilkan bool, err error) {
	var e examrepo.Exam
	if err := r.db.WithContext(ctx).Where("id = ?", examID).First(&e).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", false, nil
		}
		return "", false, err
	}
	return e.Name, e.TampilkanLeaderboard, nil
}
