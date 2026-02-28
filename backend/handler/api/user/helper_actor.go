package user

import (
	usermanagement_usecase "github.com/ProjectWidyaprada/backend/core/usecase/usermanagement"
	"github.com/ProjectWidyaprada/backend/pkg/auth"
	"github.com/gin-gonic/gin"
)

// getActorContext builds ActorContext from JWT claims and DB (for satker_id)
func (h *userHTTPHandler) getActorContext(c *gin.Context) *usermanagement_usecase.ActorContext {
	claims := auth.GetClaimsFromContext(c)
	if claims == nil {
		return nil
	}

	actor := &usermanagement_usecase.ActorContext{
		UserID:    claims.UserID,
		RoleCodes: claims.Roles,
	}

	// Super Admin: no satker filter. Admin Satker: fetch satker from user
	if !actor.IsSuperAdmin() {
		u, err := h.userRepo.GetByID(c.Request.Context(), claims.UserID)
		if err == nil && u != nil && u.SatkerID != nil {
			actor.SatkerID = u.SatkerID
		}
	}
	return actor
}
