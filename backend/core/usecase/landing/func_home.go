package landing

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

const (
	landingSliderLimit  = 10
	landingBeritaLimit  = 5
	landingTautanLimit  = 10
	landingJurnalLimit  = 5
)

func (u *landingUsecase) GetHome(ctx context.Context) (*entity.LandingHomeResponse, error) {
	slider, _ := u.slideRepo.ListPublishedForLanding(ctx, landingSliderLimit)
	berita, _, _ := u.articleRepo.ListPublished(ctx, entity.GetArticleListRequest{PageSize: landingBeritaLimit}, landingBeritaLimit)
	tautan, _ := u.linkRepo.ListActiveForLanding(ctx, landingTautanLimit)
	jurnal, _ := u.journalRepo.ListPublishedForLanding(ctx, landingJurnalLimit)

	return &entity.LandingHomeResponse{
		Slider: slider,
		Berita: berita,
		Tautan: tautan,
		Jurnal: jurnal,
	}, nil
}
