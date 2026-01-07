package models

import (
	"time"

	"github.com/google/uuid"
)

type WorkingTime struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	EntityType   string    `gorm:"type:entityType;not null" json:"entity_type"`
	EntityId     uuid.UUID `gorm:"type:uuid;not null" json:"entity_id"`
	ShiftName    string    `gorm:"type:shiftType;not null" json:"shift_name"`
	StartTime    time.Time `gorm:"type:time;not null" json:"start_time"`
	EndTime      time.Time `gorm:"type:time;not null" json:"end_time"`
	WorkingHours float64   `gorm:"->" json:"working_hours"`
	// CreatedAt    time.Time `gorm:"type:timestamptz;not null;default:now()" json:"created_at"`
	// CreatedBy    uuid.UUID `gorm:"type:uuid;not null" json:"created_by,omitempty"`
}
