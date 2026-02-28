package cms

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/google/uuid"
)

func (u *sliderUsecase) getSatkerFilter(actor *ActorContext) *string {
	if actor == nil || actor.IsSuperAdmin() {
		return nil
	}
	return actor.SatkerID
}

func (u *sliderUsecase) List(ctx context.Context, req entity.GetSlideListRequest, actor *ActorContext) (*entity.GetSlideListResponse, error) {
	return u.slideRepo.List(ctx, req, u.getSatkerFilter(actor))
}

func (u *sliderUsecase) Get(ctx context.Context, id string, actor *ActorContext) (*entity.Slide, error) {
	return u.slideRepo.GetByID(ctx, id)
}

func (u *sliderUsecase) Create(ctx context.Context, req entity.CreateSlideRequest, actor *ActorContext) (*entity.Slide, error) {
	id := uuid.New().String()
	status := req.Status
	if status == "" {
		status = entity.SlideStatusDraft
	}
	s := &entity.Slide{
		ID:        id,
		ImageURL:  req.ImageURL,
		Title:     req.Title,
		Subtitle:  req.Subtitle,
		LinkURL:   req.LinkURL,
		CTALabel:  req.CTALabel,
		SortOrder: req.SortOrder,
		Status:    status,
		DateStart: req.DateStart,
		DateEnd:   req.DateEnd,
		SatkerID:  req.SatkerID,
	}
	if !actor.IsSuperAdmin() && actor.SatkerID != nil {
		s.SatkerID = actor.SatkerID
	}
	if err := u.slideRepo.Create(ctx, s); err != nil {
		return nil, err
	}
	return u.slideRepo.GetByID(ctx, id)
}

func (u *sliderUsecase) Update(ctx context.Context, id string, req entity.UpdateSlideRequest, actor *ActorContext) (*entity.Slide, error) {
	existing, err := u.slideRepo.GetByID(ctx, id)
	if err != nil || existing == nil {
		return nil, entity.ErrRecordNotFound
	}
	s := &entity.Slide{
		ID:        id,
		ImageURL:  req.ImageURL,
		Title:     req.Title,
		Subtitle:  req.Subtitle,
		LinkURL:   req.LinkURL,
		CTALabel:  req.CTALabel,
		SortOrder: existing.SortOrder,
		Status:    req.Status,
		DateStart: req.DateStart,
		DateEnd:   req.DateEnd,
	}
	if req.SortOrder != nil {
		s.SortOrder = *req.SortOrder
	}
	if s.ImageURL == "" {
		s.ImageURL = existing.ImageURL
	}
	if s.Status == "" {
		s.Status = existing.Status
	}
	if err := u.slideRepo.Update(ctx, s); err != nil {
		return nil, err
	}
	return u.slideRepo.GetByID(ctx, id)
}

func (u *sliderUsecase) Delete(ctx context.Context, id string, actor *ActorContext) error {
	existing, err := u.slideRepo.GetByID(ctx, id)
	if err != nil || existing == nil {
		return entity.ErrRecordNotFound
	}
	return u.slideRepo.Delete(ctx, id)
}
