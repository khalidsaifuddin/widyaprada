package cms

import (
	"context"
	"regexp"
	"strings"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

var slugRegex = regexp.MustCompile(`[^a-z0-9]+`)

func slugify(s string) string {
	return strings.Trim(slugRegex.ReplaceAllString(strings.ToLower(s), "-"), "-")
}

func (u *beritaCMSUsecase) getSatkerFilter(actor *ActorContext) *string {
	if actor == nil || actor.IsSuperAdmin() {
		return nil
	}
	return actor.SatkerID
}

func (u *beritaCMSUsecase) List(ctx context.Context, req entity.GetArticleListRequest, actor *ActorContext) (*entity.GetArticleListResponse, error) {
	return u.articleRepo.List(ctx, req, u.getSatkerFilter(actor))
}

func (u *beritaCMSUsecase) Get(ctx context.Context, id string, actor *ActorContext) (*entity.Article, error) {
	return u.articleRepo.GetByID(ctx, id)
}

func (u *beritaCMSUsecase) Create(ctx context.Context, req entity.CreateArticleRequest, actor *ActorContext) (*entity.Article, error) {
	slug := req.Slug
	if slug == "" {
		slug = slugify(req.Title)
	}
	exists, _ := u.articleRepo.SlugExists(ctx, slug, "")
	if exists {
		slug = slug + "-" + randomSuffix()
	}
	status := req.Status
	if status == "" {
		status = entity.ArticleStatusDraft
	}
	a := &entity.Article{
		Title:        req.Title,
		Slug:         slug,
		Content:      req.Content,
		Excerpt:      req.Excerpt,
		ThumbnailURL: req.Thumbnail,
		PublishedAt:  req.PublishedAt,
		Status:       status,
		AuthorName:   req.AuthorName,
		Category:     req.Category,
		SatkerID:     req.SatkerID,
	}
	if !actor.IsSuperAdmin() && actor.SatkerID != nil {
		a.SatkerID = actor.SatkerID
	}
	id, err := u.articleRepo.Create(ctx, a)
	if err != nil {
		return nil, err
	}
	return u.articleRepo.GetByID(ctx, id)
}

func (u *beritaCMSUsecase) Update(ctx context.Context, id string, req entity.UpdateArticleRequest, actor *ActorContext) (*entity.Article, error) {
	existing, err := u.articleRepo.GetByID(ctx, id)
	if err != nil || existing == nil {
		return nil, entity.ErrRecordNotFound
	}
	slug := req.Slug
	if slug == "" {
		slug = existing.Slug
	} else {
		exists, _ := u.articleRepo.SlugExists(ctx, slug, id)
		if exists {
			slug = slug + "-" + randomSuffix()
		}
	}
	a := &entity.Article{
		ID:           id,
		Title:        req.Title,
		Slug:         slug,
		Content:      req.Content,
		Excerpt:      req.Excerpt,
		ThumbnailURL: req.Thumbnail,
		PublishedAt:  req.PublishedAt,
		Status:       req.Status,
		AuthorName:   req.AuthorName,
		Category:     req.Category,
	}
	if a.Title == "" {
		a.Title = existing.Title
	}
	if a.Status == "" {
		a.Status = existing.Status
	}
	if err := u.articleRepo.Update(ctx, a); err != nil {
		return nil, err
	}
	return u.articleRepo.GetByID(ctx, id)
}

func (u *beritaCMSUsecase) Delete(ctx context.Context, id string, actor *ActorContext) error {
	existing, err := u.articleRepo.GetByID(ctx, id)
	if err != nil || existing == nil {
		return entity.ErrRecordNotFound
	}
	return u.articleRepo.Delete(ctx, id)
}
