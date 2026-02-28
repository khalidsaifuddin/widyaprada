package examrepo

import (
	"context"
	"time"
)

func (r *examRepo) Delete(ctx context.Context, id string, reason string) error {
	now := time.Now().UTC()
	return r.db.WithContext(ctx).Model(&Exam{}).Where("id = ?", id).Updates(map[string]interface{}{
		"deleted_at":     now,
		"deleted_reason": reason,
		"updated_at":     now,
	}).Error
}
