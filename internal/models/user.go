package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                    uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	SurName               string    `gorm:"type:varchar(100);not null" json:"sur_name"`
	LastName              string    `gorm:"type:varchar(100);not null" json:"last_name"`
	Email                 string    `gorm:"type:varchar(255);not null;uniqueIndex" json:"email"`
	DOB                   time.Time `gorm:"type:date;not null" json:"dob"`
	Gender                string    `gorm:"type:gendertype;default:'do not tell';not null" json:"gender"`
	Phone                 string    `gorm:"type:varchar(11);not null" json:"phone"`
	CurrentAddress        *string   `gorm:"type:varchar(255)" json:"current_address,omitempty"`
	Address               string    `gorm:"type:varchar(255);not null" json:"address"`
	AvatarPath            *string   `gorm:"type:varchar(255)" json:"avatar_path,omitempty"`
	BankAccount           string    `gorm:"type:varchar(14);not null" json:"bank_account"`
	IdentifyNumber        string    `gorm:"type:varchar(12);not null;uniqueIndex" json:"identify_number"`
	IdentifyIssueDate     time.Time `gorm:"type:date;not null" json:"identify_issue_date"`
	IdentifyPlace         string    `gorm:"type:varchar(255);not null" json:"identify_place"`
	EmergencyContact      *string   `gorm:"type:varchar(255)" json:"emergency_contact,omitempty"`
	EmergencyContactPhone *string   `gorm:"type:varchar(11)" json:"emergency_contact_phone,omitempty"`
	TaxCode               *string   `gorm:"type:varchar(10);uniqueIndex" json:"tax_code,omitempty"`
	IsActive              bool      `gorm:"not null;default:false" json:"is_active"`
	MezonID               string    `gorm:"type:varchar(255);not null;uniqueIndex" json:"mezon_id"`

	CreatedAt time.Time  `gorm:"type:timestamptz;not null;default:now()" json:"created_at"`
	CreatedBy *uuid.UUID `gorm:"type:uuid" json:"created_by,omitempty"`

	LevelID    int64  `gorm:"type:bigint;not null" json:"level_id"`
	BranchID   int64  `gorm:"type:bigint;not null" json:"branch_id"`
	PositionID int64  `gorm:"type:bigint;not null" json:"position_id"`
	UserTypeID int64  `gorm:"type:bigint;not null" json:"user_type_id"`
	Password   string `gorm:"type:varchar(255); not null" json:"password"`
}

type UserRead struct {
	UserID                uuid.UUID
	FullName              string
	Email                 string
	DOB                   time.Time
	Gender                string
	Phone                 string
	CurrentAddress        *string
	Address               string
	AvatarPath            *string
	BankAccount           *string
	IdentifyNumber        *string
	IdentifyIssueDate     *time.Time
	IdentifyPlace         *string
	EmergencyContact      *string
	EmergencyContactPhone *string
	TaxCode               *string
	MezonID               string

	LevelID    int64
	BranchID   int64
	PositionID int64
	UserTypeID int64

	BranchName   *string
	LevelName    *string
	PositionName *string
	UserTypeName *string

	MorningWorkingTime   float64 `json:"morning_working_time"`
	MorningStartAt       string  `json:"morning_start_at"`
	MorningEndAt         string  `json:"morning_end_at"`
	AfternoonStartAt     string  `json:"afternoon_start_at"`
	AfternoonEndAt       string  `json:"afternoon_end_at"`
	AfternoonWorkingTime float64 `json:"afternoon_working_time"`
	BranchColor          string  `json:"branch_color"`
}

func (User) TableName() string {
	return "users"
}
