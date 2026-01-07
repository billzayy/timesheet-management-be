package dto

import (
	"time"

	"github.com/billzayy/timesheet-management-be/internal/models"
)

type WorkingTimeDTO struct {
	ShiftName    string    `json:"shift_name"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	WorkingHours float64   `json:"working_hours"`
}

func (dto WorkingTimeDTO) ToWorkingTime() *models.WorkingTime {
	return &models.WorkingTime{
		ShiftName:    dto.ShiftName,
		StartTime:    dto.StartTime,
		EndTime:      dto.EndTime,
		WorkingHours: dto.WorkingHours,
	}
}
