package models

import (
	"time"

	"github.com/google/uuid"
)

type UserType struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null;uniqueIndex" json:"name"`
	Code      string    `gorm:"type:varchar(10);not null;uniqueIndex" json:"code"`
	CreatedAt time.Time `gorm:"type:timestamptz;not null;default:now()" json:"created_at"`
	CreatedBy uuid.UUID `gorm:"type:uuid;not null" json:"created_by"`
}

func (UserType) TableName() string {
	return "user_type"
}
