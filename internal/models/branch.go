package models

import (
	"time"

	"github.com/google/uuid"
)

type Branch struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"type:varchar(255);not null;uniqueIndex" json:"name"`
	DisplayName string    `gorm:"type:varchar(100);not null" json:"display_name"`
	Code        string    `gorm:"type:varchar(10);not null;uniqueIndex" json:"code"`
	Color       string    `gorm:"type:varchar(20);not null" json:"color"`
	CreatedAt   time.Time `gorm:"type:timestamptz;not null;default:now()" json:"created_at"`
	CreatedBy   uuid.UUID `gorm:"type:uuid;not null" json:"created_by"`
}

type BranchRead struct {
	ID          int64
	Name        string
	DisplayName string
	Code        string
	Color       string

	MorningWorkingTime   float64 `json:"morning_working_time"`
	MorningStartAt       string  `json:"morning_start_at"`
	MorningEndAt         string  `json:"morning_end_at"`
	AfternoonStartAt     string  `json:"afternoon_start_at"`
	AfternoonEndAt       string  `json:"afternoon_end_at"`
	AfternoonWorkingTime float64 `json:"afternoon_working_time"`
}

func (Branch) TableName() string {
	return "branches"
}
