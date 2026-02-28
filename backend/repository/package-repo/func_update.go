package packagerepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (r *packageRepo) Update(ctx context.Context, pkg *entity.PackageDetail) error {
	now := time.Now().UTC()
	return r.db.WithContext(ctx).Model(&QuestionPackage{}).Where("id = ?", pkg.ID).Updates(map[string]interface{}{
		"code":                pkg.Code,
		"name":                pkg.Name,
		"description":         pkg.Description,
		"status":              pkg.Status,
		"verification_status": pkg.VerificationStatus,
		"updated_at":          now,
	}).Error
}
