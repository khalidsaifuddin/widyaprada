package jurnal

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *jurnalUsecase) GetByID(ctx context.Context, id string) (*entity.JurnalDetailResponse, error) {
	j, err := u.journalRepo.GetByID(ctx, id)
	if err != nil || j == nil {
		return nil, entity.ErrRecordNotFound
	}
	if j.Status != entity.JurnalStatusPublished {
		return nil, entity.ErrRecordNotFound
	}
	return &entity.JurnalDetailResponse{
		ID:          j.ID,
		Title:       j.Title,
		Author:      j.Author,
		Abstract:    j.Abstract,
		Content:     j.Content,
		PdfURL:      j.PdfURL,
		PublishedAt: j.PublishedAt,
		Status:      j.Status,
		Year:        j.Year,
		Category:    j.Category,
		CreatedAt:   j.CreatedAt,
		UpdatedAt:   j.UpdatedAt,
	}, nil
}
