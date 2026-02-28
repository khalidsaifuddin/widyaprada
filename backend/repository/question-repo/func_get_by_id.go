package questionrepo

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"gorm.io/gorm"
)

func (r *questionRepo) GetByID(ctx context.Context, id string) (*entity.QuestionDetail, error) {
	var q Question
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&q).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	e := q.ToEntity()
	return &e, nil
}
