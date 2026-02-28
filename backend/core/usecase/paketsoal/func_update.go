package paketsoal

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *paketSoalUsecase) Update(ctx context.Context, id string, req entity.UpdatePackageRequest) (*entity.PackageDetailResponse, error) {
	pkg, err := u.packageRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if pkg == nil {
		return nil, entity.ErrPackageNotFound
	}

	if req.Code != "" {
		existing, _ := u.packageRepo.GetByCode(ctx, req.Code, id)
		if existing != nil {
			return nil, entity.ErrPackageCodeExists
		}
		pkg.Code = req.Code
	}
	if req.Name != "" {
		pkg.Name = req.Name
	}
	pkg.Description = req.Description
	if req.Status != "" {
		pkg.Status = req.Status
	}

	if len(req.QuestionIDs) > 0 {
		items := make([]entity.PackageQuestionItem, len(req.QuestionIDs))
		for i, qid := range req.QuestionIDs {
			items[i] = entity.PackageQuestionItem{
				PackageID:  id,
				QuestionID: qid,
				SortOrder:  i + 1,
			}
		}
		if err := u.packageRepo.SetItems(ctx, id, items); err != nil {
			return nil, err
		}
	}

	if err := u.packageRepo.Update(ctx, pkg); err != nil {
		return nil, err
	}
	return u.Get(ctx, id)
}
