package models

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"type:varchar(200);not null;uniqueIndex" json:"name"`
	DisplayName string    `gorm:"type:varchar(200);not null" json:"display_name"`
	Description string    `gorm:"type:varchar(200);not null" json:"description"`
	CreatedAt   time.Time `gorm:"type:timestamptz;not null;default:now()" json:"created_at"`
	CreatedBy   uuid.UUID `gorm:"type:uuid;not null" json:"created_by"`
}

type Permission struct {
	ID              int64  `gorm:"primaryKey"`
	Name            string `gorm:"size:200;unique;not null"`
	DisplayName     string `gorm:"size:200;not null"`
	IsConfiguration bool   `gorm:"not null;default:false"`

	ParentID *int64
	Parent   *Permission `gorm:"foreignKey:ParentID"`

	CreatedBy *uuid.UUID
	CreatedAt time.Time
}

type PermissionNode struct {
	Name        string           `json:"name"`
	DisplayName string           `json:"display_name"`
	Children    []PermissionNode `json:"children"`
}

type UserRole struct {
	UserId uuid.UUID `gorm:"type:uuid; not null" json:"user_id"`
	RoleId int64     `gorm:"type:bigint;not null" json:"role_id"`

	CreatedAt time.Time `gorm:"type:timestamptz;not null;default:now()" json:"created_at"`
	CreatedBy uuid.UUID `gorm:"type:uuid;not null" json:"created_by"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
