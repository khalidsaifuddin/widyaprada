package assignment

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/core/repository"
)

// AssignmentUsecase interface
type AssignmentUsecase interface {
	ListDokumenPersyaratan(ctx context.Context, jenisUjikom string) (*entity.ListDokumenPersyaratanResponse, error)
	ApplyUjikom(ctx context.Context, userID string, req entity.ApplyUjikomRequest, docs []entity.ApplyUjikomDocumentInput) error
	GetApplyStatus(ctx context.Context, userID string) (*entity.ApplyStatusResponse, error)
	ListAssignments(ctx context.Context, userID string, req entity.GetAssignmentListRequest) (*entity.GetAssignmentListResponse, error)
	GetAssignmentResult(ctx context.Context, userID, examID string) (*entity.AssignmentResultResponse, error)
	GetLeaderboard(ctx context.Context, userID, examID string) (*entity.GetLeaderboardResponse, error)
}

type assignmentUsecase struct {
	dokumenRepo repository.DokumenPersyaratanRepo
	ujikomRepo  repository.UjikomRepo
	assignRepo  repository.AssignmentRepo
}

// NewAssignmentUsecase creates assignment usecase
func NewAssignmentUsecase(
	dokumenRepo repository.DokumenPersyaratanRepo,
	ujikomRepo repository.UjikomRepo,
	assignRepo repository.AssignmentRepo,
) AssignmentUsecase {
	return &assignmentUsecase{
		dokumenRepo: dokumenRepo,
		ujikomRepo:  ujikomRepo,
		assignRepo:  assignRepo,
	}
}
