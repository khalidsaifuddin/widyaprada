package ujikomrepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"gorm.io/gorm"
)

func (r *ujikomRepo) GetLatestApplicationByUser(ctx context.Context, userID string) (*entity.ApplyStatusResponse, error) {
	var app UjikomApplication
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("created_at DESC").
		First(&app).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	statusHuman := "menunggu_verifikasi"
	switch app.StatusKode {
	case entity.ApplyStatusLolos:
		statusHuman = "lolos"
	case entity.ApplyStatusTidakLolos:
		statusHuman = "tidak_lolos"
	default:
		statusHuman = "menunggu_verifikasi"
	}

	appliedAt := ""
	if app.CreatedAt != nil {
		appliedAt = app.CreatedAt.UTC().Format(time.RFC3339)
	}

	return &entity.ApplyStatusResponse{
		Status:       statusHuman,
		StatusKode:   app.StatusKode,
		CatatanTolak: app.CatatanTolak,
		JenisUjikom:  app.JenisUjikom,
		AppliedAt:    appliedAt,
	}, nil
}

func (r *ujikomRepo) HasPendingOrApproved(ctx context.Context, userID, jenisUjikom string) (bool, error) {
	var n int64
	err := r.db.WithContext(ctx).Model(&UjikomApplication{}).
		Where("user_id = ? AND jenis_ujikom = ? AND status_kode IN ?",
			userID, jenisUjikom, []string{entity.ApplyStatusMenungguVerifikasi, entity.ApplyStatusLolos}).
		Count(&n).Error
	return n > 0, err
}
