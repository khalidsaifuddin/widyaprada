package ujikomrepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/google/uuid"
)

func (r *ujikomRepo) CreateApplication(ctx context.Context, userID, jenisUjikom, statusID string, docs []entity.ApplyUjikomDocumentInput) error {
	now := time.Now().UTC()
	id := uuid.New().String()

	app := UjikomApplication{
		ID:          id,
		UserID:      userID,
		JenisUjikom: jenisUjikom,
		StatusKode:  statusID, // we pass status kode directly, e.g. ujikom_menunggu_verifikasi
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}
	if err := r.db.WithContext(ctx).Create(&app).Error; err != nil {
		return err
	}

	for _, d := range docs {
		doc := UjikomApplicationDocument{
			ID:                  uuid.New().String(),
			UjikomApplicationID: id,
			DocumentType:        d.DocumentType,
			FilePath:            d.FilePath,
			PortofolioText:      d.PortofolioText,
			CreatedAt:           &now,
			UpdatedAt:           &now,
		}
		if err := r.db.WithContext(ctx).Create(&doc).Error; err != nil {
			return err
		}
	}
	return nil
}
