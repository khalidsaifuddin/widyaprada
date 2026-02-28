package questionrepo

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"gorm.io/gorm"
)

func (r *questionRepo) GetCategoryByID(ctx context.Context, id string) (*entity.QuestionCategory, error) {
	var c QuestionCategory
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&c).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	e := c.ToEntity()
	return &e, nil
}

func (r *questionRepo) ListCategories(ctx context.Context) ([]entity.QuestionCategory, error) {
	var rows []QuestionCategory
	err := r.db.WithContext(ctx).Find(&rows).Error
	if err != nil {
		return nil, err
	}
	out := make([]entity.QuestionCategory, len(rows))
	for i := range rows {
		out[i] = rows[i].ToEntity()
	}
	return out, nil
}

// ensureCategoryExists validasi category_id saat create/update
func (r *questionRepo) ensureCategoryExists(ctx context.Context, categoryID string) error {
	if categoryID == "" {
		return nil
	}
	cat, err := r.GetCategoryByID(ctx, categoryID)
	if err != nil {
		return err
	}
	if cat == nil {
		return entity.ErrCategoryNotFound
	}
	return nil
}
