package models

import (
	"time"

	"github.com/google/uuid"
)

type Position struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null;uniqueIndex" json:"name"`
	ShortName string    `gorm:"type:varchar(100);not null" json:"short_name"`
	Code      string    `gorm:"type:varchar(10);not null;uniqueIndex" json:"code"`
	Color     string    `gorm:"type:varchar(20);not null" json:"color"`
	CreatedAt time.Time `gorm:"type:timestamptz;not null;default:now()" json:"created_at"`
	CreatedBy uuid.UUID `gorm:"type:uuid;not null" json:"created_by"`
}

func (Position) TableName() string {
	return "positions"
}
