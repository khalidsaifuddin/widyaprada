package packagerepo

import (
	"context"
)

// IsPackageUsedInExam returns true if package is used in any exam.
// SDD_Manajemen_Uji_Kompetensi not yet implemented, return false.
func (r *packageRepo) IsPackageUsedInExam(ctx context.Context, packageID string) (bool, error) {
	_ = packageID
	return false, nil
}
