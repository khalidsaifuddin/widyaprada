package repository

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

type ExamRepo interface {
	List(ctx context.Context, req entity.GetExamListRequest) (*entity.GetExamListResponse, error)
	GetByID(ctx context.Context, id string) (*entity.ExamDetail, error)
	GetByCode(ctx context.Context, code string, excludeID string) (*entity.ExamDetail, error)
	Create(ctx context.Context, e *entity.ExamDetail) error
	Update(ctx context.Context, e *entity.ExamDetail) error
	Delete(ctx context.Context, id string, reason string) error

	GetContentsByExamID(ctx context.Context, examID string) ([]entity.ExamContentItem, error)
	SetContents(ctx context.Context, examID string, items []entity.ExamContentItem) error

	GetParticipantsByExamID(ctx context.Context, examID string) ([]entity.ExamParticipantItem, error)
	SetParticipants(ctx context.Context, examID string, userIDs []string) error
}
