package dto

import (
	"github.com/billzayy/timesheet-management-be/internal/models"
)

type RoleDTO struct {
	ID          int64  `json:"id"`
	Name        string `json:"name" binding:"required"`
	DisplayName string `json:"display_name" binding:"required"`
	Description string `json:"description"`
}

type RolePermissionDTO struct {
	Name        string `json:"name" binding:"required"`
	DisplayName string `json:"display_name" binding:"required"`
	Description string `json:"description"`

	Permissions []models.PermissionNode `json:"permissions"`

	GrantedPermissions []string `json:"grantedPermissionNames"`

	Users []RoleUserDTO `json:"users"`
}

type UserConfig struct {
	Auth AuthConfig `json:"auth"`
}

type AuthConfig struct {
	AllPermissions     []string `json:"all_permissions"`
	GrantedPermissions []string `json:"granted_permissions"`
}
