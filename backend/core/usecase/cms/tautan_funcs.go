package cms

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *tautanUsecase) getSatkerFilter(actor *ActorContext) *string {
	if actor == nil || actor.IsSuperAdmin() {
		return nil
	}
	return actor.SatkerID
}

func (u *tautanUsecase) List(ctx context.Context, req entity.GetLinkListRequest, actor *ActorContext) (*entity.GetLinkListResponse, error) {
	return u.linkRepo.List(ctx, req, u.getSatkerFilter(actor))
}

func (u *tautanUsecase) Get(ctx context.Context, id string, actor *ActorContext) (*entity.Link, error) {
	return u.linkRepo.GetByID(ctx, id)
}

func (u *tautanUsecase) Create(ctx context.Context, req entity.CreateLinkRequest, actor *ActorContext) (*entity.Link, error) {
	status := req.Status
	if status == "" {
		status = entity.LinkStatusAktif
	}
	l := &entity.Link{
		Title:        req.Title,
		URL:          req.URL,
		Description:  req.Description,
		SortOrder:    req.SortOrder,
		Status:       status,
		OpenInNewTab: req.OpenInNewTab,
		SatkerID:     req.SatkerID,
	}
	if !actor.IsSuperAdmin() && actor.SatkerID != nil {
		l.SatkerID = actor.SatkerID
	}
	id, err := u.linkRepo.Create(ctx, l)
	if err != nil {
		return nil, err
	}
	return u.linkRepo.GetByID(ctx, id)
}

func (u *tautanUsecase) Update(ctx context.Context, id string, req entity.UpdateLinkRequest, actor *ActorContext) (*entity.Link, error) {
	existing, err := u.linkRepo.GetByID(ctx, id)
	if err != nil || existing == nil {
		return nil, entity.ErrRecordNotFound
	}
	l := &entity.Link{
		ID:           id,
		Title:        req.Title,
		URL:          req.URL,
		Description:  req.Description,
		SortOrder:    existing.SortOrder,
		Status:       req.Status,
		OpenInNewTab: existing.OpenInNewTab,
	}
	if req.SortOrder != nil {
		l.SortOrder = *req.SortOrder
	}
	if req.OpenInNewTab != nil {
		l.OpenInNewTab = *req.OpenInNewTab
	}
	if l.Title == "" {
		l.Title = existing.Title
	}
	if l.URL == "" {
		l.URL = existing.URL
	}
	if l.Status == "" {
		l.Status = existing.Status
	}
	if err := u.linkRepo.Update(ctx, l); err != nil {
		return nil, err
	}
	return u.linkRepo.GetByID(ctx, id)
}

func (u *tautanUsecase) Delete(ctx context.Context, id string, actor *ActorContext) error {
	existing, err := u.linkRepo.GetByID(ctx, id)
	if err != nil || existing == nil {
		return entity.ErrRecordNotFound
	}
	return u.linkRepo.Delete(ctx, id)
}
