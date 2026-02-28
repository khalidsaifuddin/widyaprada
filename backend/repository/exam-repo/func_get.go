package examrepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"gorm.io/gorm"
)

func (r *examRepo) GetByID(ctx context.Context, id string) (*entity.ExamDetail, error) {
	var e Exam
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&e).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return r.toEntity(&e), nil
}

func (r *examRepo) GetByCode(ctx context.Context, code string, excludeID string) (*entity.ExamDetail, error) {
	q := r.db.WithContext(ctx).Where("code = ?", code)
	if excludeID != "" {
		q = q.Where("id != ?", excludeID)
	}
	var e Exam
	err := q.First(&e).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return r.toEntity(&e), nil
}

func (r *examRepo) toEntity(e *Exam) *entity.ExamDetail {
	createdAt, updatedAt := "", ""
	if e.CreatedAt != nil {
		createdAt = e.CreatedAt.UTC().Format(time.RFC3339)
	}
	if e.UpdatedAt != nil {
		updatedAt = e.UpdatedAt.UTC().Format(time.RFC3339)
	}
	jadwalMulai, jadwalSelesai := "", ""
	if e.JadwalMulai != nil {
		jadwalMulai = e.JadwalMulai.UTC().Format(time.RFC3339)
	}
	if e.JadwalSelesai != nil {
		jadwalSelesai = e.JadwalSelesai.UTC().Format(time.RFC3339)
	}
	return &entity.ExamDetail{
		ID:                   e.ID,
		Code:                 e.Code,
		Name:                 e.Name,
		JadwalMulai:          jadwalMulai,
		JadwalSelesai:        jadwalSelesai,
		DurasiMenit:          e.DurasiMenit,
		Status:               e.Status,
		VerificationStatus:   e.VerificationStatus,
		ShuffleQuestions:     e.ShuffleQuestions,
		TampilkanLeaderboard: e.TampilkanLeaderboard,
		DeletedReason:        e.DeletedReason,
		CreatedAt:            createdAt,
		UpdatedAt:            updatedAt,
	}
}
