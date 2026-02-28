package examplerepo

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"gorm.io/gorm"
)

func (r *exampleRepo) GetExampleDetail(ctx context.Context, id int64) (entity.Example, error) {
	var resp entity.Example
	db := r.db.WithContext(ctx).Model(&Example{}).Where("id = ?", id)
	var row Example
	if err := db.First(&row).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return resp, entity.WrapRecordNotFoundf("example not found with id: %d", id)
		}
		return resp, err
	}
	return row.ToEntity(), nil
}
