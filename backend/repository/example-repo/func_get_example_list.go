package examplerepo

import (
	"context"
	"fmt"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/helper"
)

func (r *exampleRepo) GetExampleList(ctx context.Context, request entity.GetExampleListRequest) (entity.GetExampleListResponse, error) {
	var resp entity.GetExampleListResponse
	db := r.db.WithContext(ctx).Model(&Example{})

	if request.Search != "" {
		searchString := fmt.Sprintf("%%%s%%", request.Search)
		db = db.Where("name LIKE ?", searchString)
	}

	if err := db.Count(&resp.TotalData).Error; err != nil {
		return resp, err
	}

	if !request.IsReturnAllData {
		offset, limit := helper.GetOffsetAndLimit(request.Page, request.PageSize)
		db = db.Offset(int(offset)).Limit(int(limit))
	}

	var results []Example
	if err := db.Find(&results).Error; err != nil {
		return resp, err
	}

	for _, row := range results {
		resp.Items = append(resp.Items, row.ToEntity())
	}
	resp.TotalPage = helper.GenerateTotalPage(resp.TotalData, request.PageSize)
	resp.Page = request.Page
	resp.PageSize = request.PageSize
	return resp, nil
}
