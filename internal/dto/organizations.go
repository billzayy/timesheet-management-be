package dto

import (
	"github.com/billzayy/timesheet-management-be/internal/models"
	"github.com/google/uuid"
)

type BranchDTO struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Code        string `json:"code"`
	Color       string `json:"color"`
}

type PositionDTO struct {
	ID        int64  `json:"id"`
	Name      string `json:"name" binding:"required"`
	ShortName string `json:"short_name" binding:"required"`
	Code      string `json:"code" binding:"required"`
	Color     string `json:"color" binding:"required"`
}

type LevelDTO struct {
	ID          int64  `json:"id"`
	Name        string `json:"name" binding:"required"`
	DisplayName string `json:"display_name" binding:"required"`
	Code        string `json:"code" binding:"required"`
	Color       string `json:"color" binding:"required"`
}

type UserTypeDTO struct {
	ID   int64  `json:"id"`
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
}

func (dto PositionDTO) ToCreatePosition(id uuid.UUID) models.Position {
	return models.Position{
		Name:      dto.Name,
		ShortName: dto.ShortName,
		Code:      dto.Code,
		Color:     dto.Color,
		CreatedBy: id,
	}
}

func (dto PositionDTO) ToUpdatePosition() models.Position {
	return models.Position{
		ID:        dto.ID,
		Name:      dto.Name,
		ShortName: dto.ShortName,
		Code:      dto.Code,
		Color:     dto.Color,
	}
}

func (dto LevelDTO) ToCreateLevel(id uuid.UUID) models.Level {
	return models.Level{
		Name:        dto.Name,
		DisplayName: dto.DisplayName,
		Code:        dto.Code,
		Color:       dto.Color,
		CreatedBy:   id,
	}
}

func (dto LevelDTO) ToUpdateLevel() models.Level {
	return models.Level{
		ID:          dto.ID,
		Name:        dto.Name,
		DisplayName: dto.DisplayName,
		Code:        dto.Code,
		Color:       dto.Color,
	}
}

func (dto UserTypeDTO) ToCreateUserType(id uuid.UUID) models.UserType {
	return models.UserType{
		Name:      dto.Name,
		Code:      dto.Code,
		CreatedBy: id,
	}
}

func (dto UserTypeDTO) ToUpdateUserType() models.UserType {
	return models.UserType{
		ID:   dto.ID,
		Name: dto.Name,
		Code: dto.Code,
	}
}
