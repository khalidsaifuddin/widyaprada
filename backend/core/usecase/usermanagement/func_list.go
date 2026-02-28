package usermanagement

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *userUsecase) List(ctx context.Context, req entity.GetUserListRequest, actor *ActorContext) (entity.GetUserListResponse, error) {
	var satkerFilter *string
	if actor != nil && !actor.IsSuperAdmin() && actor.SatkerID != nil {
		satkerFilter = actor.SatkerID
	}
	return u.userRepo.ListUsers(ctx, req, satkerFilter)
}
