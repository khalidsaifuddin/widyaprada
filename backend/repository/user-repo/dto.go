package userrepo

import (
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"gorm.io/gorm"
)

// User model untuk tabel users
type User struct {
	ID             string         `gorm:"column:id;primaryKey;type:uuid"`
	Name           string         `gorm:"column:name;size:255;not null"`
	Email          string         `gorm:"column:email;size:255;uniqueIndex;not null"`
	Username       string         `gorm:"column:username;size:100;uniqueIndex;not null"`
	NIP            string         `gorm:"column:nip;size:50"`
	PasswordHash   string         `gorm:"column:password_hash;size:255;not null"`
	SatkerID       *string        `gorm:"column:satker_id;type:uuid"`
	IsActive       bool           `gorm:"column:is_active;default:true"`
	DeletedReason  string         `gorm:"column:deleted_reason"`
	CreatedAt      *time.Time     `gorm:"column:created_at"`
	UpdatedAt      *time.Time     `gorm:"column:updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (User) TableName() string {
	return "users"
}

// Role model untuk tabel roles
type Role struct {
	ID            string         `gorm:"column:id;primaryKey;type:uuid"`
	Code          string         `gorm:"column:code;size:50;uniqueIndex;not null"`
	Name          string         `gorm:"column:name;size:100;not null"`
	Description   string         `gorm:"column:description"`
	DeletedReason string         `gorm:"column:deleted_reason"`
	CreatedAt     *time.Time     `gorm:"column:created_at"`
	UpdatedAt     *time.Time     `gorm:"column:updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (Role) TableName() string {
	return "roles"
}

// UserRole model untuk tabel user_roles (many-to-many)
type UserRole struct {
	UserID    string     `gorm:"column:user_id;primaryKey;type:uuid"`
	RoleID    string     `gorm:"column:role_id;primaryKey;type:uuid"`
	CreatedAt *time.Time `gorm:"column:created_at"`
}

func (UserRole) TableName() string {
	return "user_roles"
}

// ToEntity converts DTO to entity.User
func (u *User) ToEntity() entity.User {
	createdAt, updatedAt := "", ""
	if u.CreatedAt != nil {
		createdAt = u.CreatedAt.Format(time.RFC3339)
	}
	if u.UpdatedAt != nil {
		updatedAt = u.UpdatedAt.Format(time.RFC3339)
	}
	return entity.User{
		ID:           u.ID,
		Name:         u.Name,
		Email:        u.Email,
		Username:     u.Username,
		NIP:          u.NIP,
		PasswordHash: u.PasswordHash,
		SatkerID:     u.SatkerID,
		IsActive:     u.IsActive,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}

// FromEntity converts entity.User to DTO (caller sets CreatedAt, UpdatedAt, ID as needed)
func (User) FromEntity(e entity.User) User {
	return User{
		ID:           e.ID,
		Name:         e.Name,
		Email:        e.Email,
		Username:     e.Username,
		NIP:          e.NIP,
		PasswordHash: e.PasswordHash,
		SatkerID:     e.SatkerID,
		IsActive:     e.IsActive,
	}
}

// ToUserListItem converts User DTO + roles to entity.UserListItem (for list response)
func (u *User) ToUserListItem(roles []entity.RoleInfo) entity.UserListItem {
	createdAt := ""
	if u.CreatedAt != nil {
		createdAt = u.CreatedAt.Format(time.RFC3339)
	}
	return entity.UserListItem{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Username:  u.Username,
		Roles:     roles,
		SatkerID:  u.SatkerID,
		IsActive:  u.IsActive,
		CreatedAt: createdAt,
	}
}

// ToEntity converts Role DTO to entity.RoleInfo
func (r *Role) ToEntity() entity.RoleInfo {
	return entity.RoleInfo{
		ID:   r.ID,
		Code: r.Code,
		Name: r.Name,
	}
}
