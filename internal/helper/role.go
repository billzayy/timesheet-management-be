package helper

import (
	"github.com/billzayy/timesheet-management-be/internal/dto"
	"github.com/billzayy/timesheet-management-be/internal/models"
)

func ConverUserDTO(input []models.UserRead) []dto.RoleUserDTO {
	result := make([]dto.RoleUserDTO, 0, len(input))

	for _, m := range input {
		result = append(result,
			dto.RoleUserDTO{
				FullName:   m.FullName,
				Email:      m.Email,
				AvatarPath: m.AvatarPath,
				LevelID:    m.LevelID,
				LevelName:  m.LevelName,
				BranchID:   m.BranchID,
				BranchName: m.LevelName,
				// BranchColor:  m.BranchColor,
				PositionID:   m.PositionID,
				PositionName: m.PositionName,
				UserTypeID:   m.UserTypeID,
				UserTypeName: m.UserTypeName,
				UserID:       m.UserID,
			})
	}

	return result
}
