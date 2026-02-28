package paketsoal

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/google/uuid"
)

func (u *paketSoalUsecase) Create(ctx context.Context, req entity.CreatePackageRequest) (*entity.PackageDetailResponse, error) {
	existing, err := u.packageRepo.GetByCode(ctx, req.Code, "")
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, entity.ErrPackageCodeExists
	}
	if len(req.QuestionIDs) == 0 {
		return nil, entity.ErrPackageMinOneQuestion
	}

	pkg := &entity.PackageDetail{
		ID:                 uuid.New().String(),
		Code:               req.Code,
		Name:               req.Name,
		Description:        req.Description,
		Status:             req.Status,
		VerificationStatus: entity.QuestionVerifBelum,
	}
	if pkg.Status == "" {
		pkg.Status = entity.QuestionStatusDraft
	}
	if err := u.packageRepo.Create(ctx, pkg); err != nil {
		return nil, err
	}

	items := make([]entity.PackageQuestionItem, len(req.QuestionIDs))
	for i, qid := range req.QuestionIDs {
		items[i] = entity.PackageQuestionItem{
			PackageID:  pkg.ID,
			QuestionID: qid,
			SortOrder:  i + 1,
		}
	}
	if err := u.packageRepo.SetItems(ctx, pkg.ID, items); err != nil {
		return nil, err
	}
	return u.Get(ctx, pkg.ID)
}
