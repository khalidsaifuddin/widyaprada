package rbacrepo

import (
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"gorm.io/gorm"
)

// Permission model untuk tabel permissions
type Permission struct {
	ID            string         `gorm:"column:id;primaryKey;type:uuid"`
	Code          string         `gorm:"column:code;size:100;uniqueIndex;not null"`
	Name          string         `gorm:"column:name;size:255;not null"`
	Group         string         `gorm:"column:group;size:100"`
	Description   string         `gorm:"column:description"`
	DeletedReason string         `gorm:"column:deleted_reason"`
	CreatedAt     *time.Time     `gorm:"column:created_at"`
	UpdatedAt     *time.Time     `gorm:"column:updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (Permission) TableName() string {
	return "permissions"
}

// RolePermission model untuk tabel role_permissions
type RolePermission struct {
	RoleID       string     `gorm:"column:role_id;primaryKey;type:uuid"`
	PermissionID string     `gorm:"column:permission_id;primaryKey;type:uuid"`
	CreatedAt    *time.Time `gorm:"column:created_at"`
}

func (RolePermission) TableName() string {
	return "role_permissions"
}

// ToEntity converts Permission DTO to entity.PermissionInfo
func (p *Permission) ToEntity() entity.PermissionInfo {
	return entity.PermissionInfo{
		ID:    p.ID,
		Code:  p.Code,
		Name:  p.Name,
		Group: p.Group,
	}
}

// FromEntity converts entity to Permission DTO (for create/update)
func (Permission) FromEntity(code, name, group, description string) Permission {
	return Permission{
		Code:        code,
		Name:        name,
		Group:       group,
		Description: description,
	}
}
