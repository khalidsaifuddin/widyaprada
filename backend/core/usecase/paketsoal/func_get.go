package paketsoal

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *paketSoalUsecase) Get(ctx context.Context, id string) (*entity.PackageDetailResponse, error) {
	pkg, err := u.packageRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if pkg == nil {
		return nil, entity.ErrPackageNotFound
	}
	items, err := u.packageRepo.GetItemsByPackageID(ctx, id)
	if err != nil {
		return nil, err
	}
	questions := make([]entity.PackageQuestionInfo, len(items))
	for i := range items {
		q, _ := u.questionRepo.GetByID(ctx, items[i].QuestionID)
		code := ""
		if q != nil {
			code = q.Code
		}
		questions[i] = entity.PackageQuestionInfo{
			QuestionID:   items[i].QuestionID,
			QuestionCode: code,
			SortOrder:    items[i].SortOrder,
		}
	}
	return &entity.PackageDetailResponse{
		ID:                 pkg.ID,
		Code:               pkg.Code,
		Name:               pkg.Name,
		Description:        pkg.Description,
		Status:             pkg.Status,
		VerificationStatus: pkg.VerificationStatus,
		Questions:          questions,
		CreatedAt:          pkg.CreatedAt,
		UpdatedAt:          pkg.UpdatedAt,
	}, nil
}
