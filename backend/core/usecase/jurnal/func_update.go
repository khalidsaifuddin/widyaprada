package jurnal

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *jurnalUsecase) Update(ctx context.Context, j *entity.Jurnal) error {
	return u.journalRepo.Update(ctx, j)
}
