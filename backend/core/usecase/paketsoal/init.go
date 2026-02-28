package paketsoal

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/core/repository"
)

type PaketSoalUsecase interface {
	List(ctx context.Context, req entity.GetPackageListRequest) (*entity.GetPackageListResponse, error)
	Get(ctx context.Context, id string) (*entity.PackageDetailResponse, error)
	Create(ctx context.Context, req entity.CreatePackageRequest) (*entity.PackageDetailResponse, error)
	Update(ctx context.Context, id string, req entity.UpdatePackageRequest) (*entity.PackageDetailResponse, error)
	Delete(ctx context.Context, id string, reason string) error
	Verify(ctx context.Context, id string) error
	Unverify(ctx context.Context, id string) error
}

type paketSoalUsecase struct {
	packageRepo repository.PackageRepo
	questionRepo  repository.QuestionRepo
}

func NewPaketSoalUsecase(packageRepo repository.PackageRepo, questionRepo repository.QuestionRepo) PaketSoalUsecase {
	return &paketSoalUsecase{
		packageRepo:  packageRepo,
		questionRepo: questionRepo,
	}
}
