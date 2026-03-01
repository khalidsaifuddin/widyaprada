package jurnal

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *jurnalUsecase) GetByIDForOwner(ctx context.Context, id, userID string) (*entity.Jurnal, error) {
	j, err := u.journalRepo.GetByID(ctx, id)
	if err != nil || j == nil {
		return nil, err
	}
	if j.UserID == nil || *j.UserID != userID {
		return nil, nil
	}
	return j, nil
}
