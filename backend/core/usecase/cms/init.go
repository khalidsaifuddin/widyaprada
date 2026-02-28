package cms

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

type SliderUsecase interface {
	List(ctx context.Context, req entity.GetSlideListRequest, actor *ActorContext) (*entity.GetSlideListResponse, error)
	Get(ctx context.Context, id string, actor *ActorContext) (*entity.Slide, error)
	Create(ctx context.Context, req entity.CreateSlideRequest, actor *ActorContext) (*entity.Slide, error)
	Update(ctx context.Context, id string, req entity.UpdateSlideRequest, actor *ActorContext) (*entity.Slide, error)
	Delete(ctx context.Context, id string, actor *ActorContext) error
}

type BeritaCMSUsecase interface {
	List(ctx context.Context, req entity.GetArticleListRequest, actor *ActorContext) (*entity.GetArticleListResponse, error)
	Get(ctx context.Context, id string, actor *ActorContext) (*entity.Article, error)
	Create(ctx context.Context, req entity.CreateArticleRequest, actor *ActorContext) (*entity.Article, error)
	Update(ctx context.Context, id string, req entity.UpdateArticleRequest, actor *ActorContext) (*entity.Article, error)
	Delete(ctx context.Context, id string, actor *ActorContext) error
}

type TautanUsecase interface {
	List(ctx context.Context, req entity.GetLinkListRequest, actor *ActorContext) (*entity.GetLinkListResponse, error)
	Get(ctx context.Context, id string, actor *ActorContext) (*entity.Link, error)
	Create(ctx context.Context, req entity.CreateLinkRequest, actor *ActorContext) (*entity.Link, error)
	Update(ctx context.Context, id string, req entity.UpdateLinkRequest, actor *ActorContext) (*entity.Link, error)
	Delete(ctx context.Context, id string, actor *ActorContext) error
}

type sliderUsecase struct {
	slideRepo repository.SlideRepo
}

func NewSliderUsecase(slideRepo repository.SlideRepo) SliderUsecase {
	return &sliderUsecase{slideRepo: slideRepo}
}

type beritaCMSUsecase struct {
	articleRepo repository.ArticleRepo
}

func NewBeritaCMSUsecase(articleRepo repository.ArticleRepo) BeritaCMSUsecase {
	return &beritaCMSUsecase{articleRepo: articleRepo}
}

type tautanUsecase struct {
	linkRepo repository.LinkRepo
}

func NewTautanUsecase(linkRepo repository.LinkRepo) TautanUsecase {
	return &tautanUsecase{linkRepo: linkRepo}
}
