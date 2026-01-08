package dto

type RoleDTO struct {
	ID          int64  `json:"id"`
	Name        string `json:"name" binding:"required"`
	DisplayName string `json:"display_name" binding:"required"`
	Description string `json:"description"`
}
