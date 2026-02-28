package packagerepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/google/uuid"
)

func (r *packageRepo) Create(ctx context.Context, pkg *entity.PackageDetail) error {
	if pkg.ID == "" {
		pkg.ID = uuid.New().String()
	}
	now := time.Now().UTC()
	dto := QuestionPackage{
		ID:                 pkg.ID,
		Code:               pkg.Code,
		Name:               pkg.Name,
		Description:        pkg.Description,
		Status:             pkg.Status,
		VerificationStatus: pkg.VerificationStatus,
		CreatedAt:          &now,
		UpdatedAt:          &now,
	}
	if dto.Status == "" {
		dto.Status = entity.QuestionStatusDraft
	}
	if dto.VerificationStatus == "" {
		dto.VerificationStatus = entity.QuestionVerifBelum
	}
	return r.db.WithContext(ctx).Create(&dto).Error
}
