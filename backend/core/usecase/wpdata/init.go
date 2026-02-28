package wpdata

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/core/repository"
)

type ActorContext struct {
	UserID    string
	SatkerID  *string
	RoleCodes []string
}

func (a *ActorContext) IsSuperAdmin() bool {
	for _, c := range a.RoleCodes {
		if c == "SUPER_ADMIN" {
			return true
		}
	}
	return false
}

type WPDataUsecase interface {
	List(ctx context.Context, req entity.GetWPDataListRequest, actor *ActorContext) (*entity.GetWPDataListResponse, error)
	Get(ctx context.Context, id string, actor *ActorContext) (*entity.WPDataDetailResponse, error)
	Create(ctx context.Context, req entity.CreateWPDataRequest, actor *ActorContext) (*entity.WPDataDetailResponse, error)
	Update(ctx context.Context, id string, req entity.UpdateWPDataRequest, actor *ActorContext) (*entity.WPDataDetailResponse, error)
	Delete(ctx context.Context, id string, reason string, actor *ActorContext) error
}

type CalonPesertaUsecase interface {
	List(ctx context.Context, req entity.GetCalonPesertaListRequest, actor *ActorContext) (*entity.GetCalonPesertaListResponse, error)
	Get(ctx context.Context, id string, actor *ActorContext) (*entity.CalonPesertaDetailResponse, error)
	Verify(ctx context.Context, id string, req entity.VerifyCalonPesertaRequest, actor *ActorContext) error
}

type wpdataUsecase struct {
	wpdataRepo repository.WPDataRepo
}

type calonPesertaUsecase struct {
	ujikomRepo repository.UjikomRepo
}

func NewWPDataUsecase(wpdataRepo repository.WPDataRepo) WPDataUsecase {
	return &wpdataUsecase{wpdataRepo: wpdataRepo}
}

func NewCalonPesertaUsecase(ujikomRepo repository.UjikomRepo) CalonPesertaUsecase {
	return &calonPesertaUsecase{ujikomRepo: ujikomRepo}
}
