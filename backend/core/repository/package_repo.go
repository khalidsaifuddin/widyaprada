package repository

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

type PackageRepo interface {
	List(ctx context.Context, req entity.GetPackageListRequest) (*entity.GetPackageListResponse, error)
	GetByID(ctx context.Context, id string) (*entity.PackageDetail, error)
	GetByCode(ctx context.Context, code string, excludeID string) (*entity.PackageDetail, error)
	Create(ctx context.Context, pkg *entity.PackageDetail) error
	Update(ctx context.Context, pkg *entity.PackageDetail) error
	Delete(ctx context.Context, id string, reason string) error

	GetItemsByPackageID(ctx context.Context, packageID string) ([]entity.PackageQuestionItem, error)
	SetItems(ctx context.Context, packageID string, items []entity.PackageQuestionItem) error

	IsPackageUsedInExam(ctx context.Context, packageID string) (bool, error)
}
