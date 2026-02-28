package examrepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (r *examRepo) GetContentsByExamID(ctx context.Context, examID string) ([]entity.ExamContentItem, error) {
	var rows []ExamContent
	err := r.db.WithContext(ctx).Where("exam_id = ?", examID).Order("sort_order").Find(&rows).Error
	if err != nil {
		return nil, err
	}
	out := make([]entity.ExamContentItem, len(rows))
	for i := range rows {
		out[i] = entity.ExamContentItem{
			ExamID:     rows[i].ExamID,
			SourceType: rows[i].SourceType,
			SourceID:   rows[i].SourceID,
			SortOrder:  rows[i].SortOrder,
		}
	}
	return out, nil
}

func (r *examRepo) SetContents(ctx context.Context, examID string, items []entity.ExamContentItem) error {
	if err := r.db.WithContext(ctx).Where("exam_id = ?", examID).Delete(&ExamContent{}).Error; err != nil {
		return err
	}
	if len(items) == 0 {
		return nil
	}
	now := time.Now().UTC()
	rows := make([]ExamContent, len(items))
	for i := range items {
		rows[i] = ExamContent{
			ExamID:     examID,
			SourceType: items[i].SourceType,
			SourceID:   items[i].SourceID,
			SortOrder:  items[i].SortOrder,
			CreatedAt:  &now,
		}
	}
	return r.db.WithContext(ctx).Create(&rows).Error
}
