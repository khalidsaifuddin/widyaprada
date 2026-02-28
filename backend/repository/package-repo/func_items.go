package packagerepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (r *packageRepo) GetItemsByPackageID(ctx context.Context, packageID string) ([]entity.PackageQuestionItem, error) {
	var rows []PackageQuestionItem
	err := r.db.WithContext(ctx).Where("package_id = ?", packageID).Order("sort_order").Find(&rows).Error
	if err != nil {
		return nil, err
	}
	out := make([]entity.PackageQuestionItem, len(rows))
	for i := range rows {
		out[i] = entity.PackageQuestionItem{
			PackageID:  rows[i].PackageID,
			QuestionID: rows[i].QuestionID,
			SortOrder:  rows[i].SortOrder,
		}
	}
	return out, nil
}

func (r *packageRepo) SetItems(ctx context.Context, packageID string, items []entity.PackageQuestionItem) error {
	if err := r.db.WithContext(ctx).Where("package_id = ?", packageID).Delete(&PackageQuestionItem{}).Error; err != nil {
		return err
	}
	if len(items) == 0 {
		return nil
	}
	now := time.Now().UTC()
	rows := make([]PackageQuestionItem, len(items))
	for i := range items {
		rows[i] = PackageQuestionItem{
			PackageID:  packageID,
			QuestionID: items[i].QuestionID,
			SortOrder:  items[i].SortOrder,
			CreatedAt:  &now,
		}
	}
	return r.db.WithContext(ctx).Create(&rows).Error
}
