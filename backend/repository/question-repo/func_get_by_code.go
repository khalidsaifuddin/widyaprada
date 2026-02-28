package questionrepo

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"gorm.io/gorm"
)

func (r *questionRepo) GetByCode(ctx context.Context, code string, excludeID string) (*entity.QuestionDetail, error) {
	q := r.db.WithContext(ctx).Where("code = ?", code)
	if excludeID != "" {
		q = q.Where("id != ?", excludeID)
	}
	var row Question
	err := q.First(&row).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	e := row.ToEntity()
	return &e, nil
}
