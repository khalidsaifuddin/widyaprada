package jurnal

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *jurnalUsecase) ListPublished(ctx context.Context, req entity.GetJurnalListRequest) (*entity.GetJurnalListResponse, error) {
	return u.journalRepo.ListPublished(ctx, req)
}
