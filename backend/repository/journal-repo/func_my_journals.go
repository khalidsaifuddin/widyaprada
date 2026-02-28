package journalrepo

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (r *journalRepo) ListMyJournals(ctx context.Context, userID string, limit, page int64) ([]entity.DashboardJournalItem, int64, error) {
	if limit < 1 {
		limit = 10
	}
	if page < 1 {
		page = 1
	}

	db := r.db.WithContext(ctx).Model(&Journal{}).
		Where("deleted_at IS NULL").
		Where("user_id = ?", userID)

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var items []Journal
	offset := (page - 1) * limit
	if err := db.Order("COALESCE(submitted_at, updated_at, created_at) DESC NULLS LAST").
		Offset(int(offset)).
		Limit(int(limit)).
		Find(&items).Error; err != nil {
		return nil, 0, err
	}

	list := make([]entity.DashboardJournalItem, len(items))
	for i := range items {
		submittedAt := ""
		if items[i].SubmittedAt != nil {
			submittedAt = items[i].SubmittedAt.UTC().Format("2006-01-02T15:04:05Z07:00")
		} else if items[i].UpdatedAt != nil {
			submittedAt = items[i].UpdatedAt.UTC().Format("2006-01-02T15:04:05Z07:00")
		} else if items[i].CreatedAt != nil {
			submittedAt = items[i].CreatedAt.UTC().Format("2006-01-02T15:04:05Z07:00")
		}
		status := items[i].Status
		if status == "" {
			status = "Draft"
		}
		list[i] = entity.DashboardJournalItem{
			ID:          items[i].ID,
			Title:       items[i].Title,
			SubmittedAt: submittedAt,
			Status:      status,
		}
	}
	return list, total, nil
}
