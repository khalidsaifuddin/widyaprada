package jurnal

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *jurnalUsecase) Create(ctx context.Context, j *entity.Jurnal) error {
	return u.journalRepo.Create(ctx, j)
}
