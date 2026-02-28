package entity

import "time"

type Example struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type GetExampleListRequest struct {
	Page            int64  `json:"page" form:"page"`
	PageSize        int64  `json:"page_size" form:"page_size"`
	Search          string `json:"search" form:"search"`
	IsReturnAllData bool   `json:"is_return_all_data" form:"is_return_all_data"`
}

type GetExampleListResponse struct {
	Items     []Example `json:"items"`
	TotalPage int64     `json:"total_page"`
	TotalData int64     `json:"total_data"`
	Page      int64     `json:"page"`
	PageSize  int64     `json:"page_size"`
}
