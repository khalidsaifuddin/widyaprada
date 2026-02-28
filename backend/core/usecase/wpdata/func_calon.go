package wpdata

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *calonPesertaUsecase) getSatkerFilter(actor *ActorContext) *string {
	if actor == nil || actor.IsSuperAdmin() {
		return nil
	}
	return actor.SatkerID
}

func (u *calonPesertaUsecase) List(ctx context.Context, req entity.GetCalonPesertaListRequest, actor *ActorContext) (*entity.GetCalonPesertaListResponse, error) {
	return u.ujikomRepo.ListCalonPeserta(ctx, req, u.getSatkerFilter(actor))
}

func (u *calonPesertaUsecase) Get(ctx context.Context, id string, actor *ActorContext) (*entity.CalonPesertaDetailResponse, error) {
	res, err := u.ujikomRepo.GetCalonPesertaByID(ctx, id)
	if err != nil || res == nil {
		return nil, entity.ErrRecordNotFound
	}
	return res, nil
}

func (u *calonPesertaUsecase) Verify(ctx context.Context, id string, req entity.VerifyCalonPesertaRequest, actor *ActorContext) error {
	_, err := u.ujikomRepo.GetCalonPesertaByID(ctx, id)
	if err != nil || actor == nil {
		return entity.ErrRecordNotFound
	}
	statusKode := entity.ApplyStatusLolos
	if !req.Approved {
		statusKode = entity.ApplyStatusTidakLolos
	}
	return u.ujikomRepo.UpdateApplicationStatus(ctx, id, statusKode, req.Catatan)
}
