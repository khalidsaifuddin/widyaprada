package packagerepo

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"gorm.io/gorm"
)

func (r *packageRepo) GetByID(ctx context.Context, id string) (*entity.PackageDetail, error) {
	var p QuestionPackage
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&p).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	e := p.ToEntity()
	return &e, nil
}

func (r *packageRepo) GetByCode(ctx context.Context, code string, excludeID string) (*entity.PackageDetail, error) {
	q := r.db.WithContext(ctx).Where("code = ?", code)
	if excludeID != "" {
		q = q.Where("id != ?", excludeID)
	}
	var p QuestionPackage
	err := q.First(&p).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	e := p.ToEntity()
	return &e, nil
}
