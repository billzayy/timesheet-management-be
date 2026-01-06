package models

import (
	"time"

	"github.com/google/uuid"
)

type WorkingTime struct {
	ID           uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	EntityType   string     `gorm:"type:entitytype;not null" json:"entity_type"`
	EntityId     uuid.UUID  `gorm:"type:uuid;not null" json:"entity_id"`
	ShiftName    string     `gorm:"type:shifttype;not null" json:"shift_name"`
	StartTime    string     `gorm:"type:time;not null" json:"start_time"`
	EndTime      string     `gorm:"type:time;not null" json:"end_time"`
	WorkingHours float64    `gorm:"type:numeric(3,1);->" json:"working_hours"`
	CreatedAt    time.Time  `gorm:"type:timestamptz;not null;default:now()" json:"created_at"`
	CreatedBy    *uuid.UUID `gorm:"type:uuid" json:"created_by,omitempty"`
}
