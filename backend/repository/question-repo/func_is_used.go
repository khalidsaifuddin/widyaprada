package questionrepo

import (
	"context"
)

// IsQuestionUsedInPackage returns true if question is used in any question package.
func (r *questionRepo) IsQuestionUsedInPackage(ctx context.Context, questionID string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Table("package_question_items").Where("question_id = ?", questionID).Limit(1).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
