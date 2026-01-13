package dto

import (
	"time"

	"github.com/billzayy/timesheet-management-be/internal/models"
	"github.com/google/uuid"
)

type RequestUserDTO struct {
	SurName               string    `json:"sur_name" binding:"required,max=100"`
	LastName              string    `json:"last_name" binding:"required,max=100"`
	Email                 string    `json:"email" binding:"required,email"`
	Password              string    `json:"password" binding:"required"`
	DOB                   time.Time `json:"dob" binding:"required"`
	Gender                string    `json:"gender"`
	Phone                 string    `json:"phone" binding:"required,len=11"`
	CurrentAddress        *string   `json:"current_address,omitempty"`
	Address               string    `json:"address" binding:"required"`
	AvatarPath            *string   `json:"avatar_path,omitempty"`
	BankAccount           string    `json:"bank_account" binding:"required"`
	IdentifyNumber        string    `json:"identify_number" binding:"required,len=12"`
	IdentifyIssueDate     time.Time `json:"identify_issue_date" binding:"required"`
	IdentifyPlace         string    `json:"identify_place" binding:"required"`
	EmergencyContact      *string   `json:"emergency_contact,omitempty"`
	EmergencyContactPhone *string   `json:"emergency_contact_phone,omitempty"`
	TaxCode               *string   `json:"tax_code,omitempty"`
	MezonID               string    `json:"mezon_id" binding:"required"`

	LevelID    int64 `json:"level_id" binding:"required"`
	BranchID   int64 `json:"branch_id" binding:"required"`
	PositionID int64 `json:"position_id" binding:"required"`
	UserTypeID int64 `json:"user_type_id" binding:"required"`

	MorningWorkingTime   float64 `json:"morning_working_time"`
	MorningStartAt       string  `json:"morning_start_at"`
	MorningEndAt         string  `json:"morning_end_at"`
	AfternoonStartAt     string  `json:"afternoon_start_at"`
	AfternoonEndAt       string  `json:"afternoon_end_at"`
	AfternoonWorkingTime float64 `json:"afternoon_working_time"`

	RoleId int64 `json:"role_id"`
}

type GetUserDTO struct {
	FullName              string    `json:"full_name"`
	Email                 string    `json:"email"`
	DOB                   time.Time `json:"dob"`
	Gender                string    `json:"gender"`
	Phone                 string    `json:"phone"`
	CurrentAddress        *string   `json:"current_address,omitempty"`
	Address               string    `json:"address"`
	AvatarPath            *string   `json:"avatar_path,omitempty"`
	BankAccount           string    `json:"bank_account"`
	IdentifyNumber        string    `json:"identify_number"`
	IdentifyIssueDate     time.Time `json:"identify_issue_date" `
	IdentifyPlace         string    `json:"identify_place" `
	EmergencyContact      *string   `json:"emergency_contact,omitempty"`
	EmergencyContactPhone *string   `json:"emergency_contact_phone,omitempty"`
	TaxCode               *string   `json:"tax_code,omitempty"`
	MezonID               string    `json:"mezon_id" `
	Roles                 []string  `json:"roles,omitempty"`
	LevelID               int64     `json:"level_id" `
	LevelName             string    `json:"level_name"`
	BranchID              int64     `json:"branch_id" `
	BranchName            string    `json:"branch_name"`
	PositionID            int64     `json:"position_id" `
	PositionName          string    `json:"position_name"`
	UserTypeID            int64     `json:"user_type_id" `
	UserTypeName          string    `json:"user_type_name"`
	MorningWorkingTime    float64   `json:"morning_working_time"`
	MorningStartAt        string    `json:"morning_start_at"`
	MorningEndAt          string    `json:"morning_end_at"`
	AfternoonStartAt      string    `json:"afternoon_start_at"`
	AfternoonEndAt        string    `json:"afternoon_end_at"`
	AfternoonWorkingTime  float64   `json:"afternoon_working_time"`
	ID                    uuid.UUID `json:"id"`
}

type RoleUserDTO struct {
	UserID     uuid.UUID
	FullName   string
	Email      string
	AvatarPath *string

	LevelID    int64
	BranchID   int64
	PositionID int64
	UserTypeID int64

	BranchName   *string
	BranchColor  *string
	LevelName    *string
	PositionName *string
	UserTypeName *string
}

func (dto RequestUserDTO) ToUser() models.User {
	return models.User{
		SurName:               dto.SurName,
		LastName:              dto.LastName,
		Email:                 dto.Email,
		Password:              dto.Password,
		DOB:                   dto.DOB,
		Gender:                dto.Gender,
		Phone:                 dto.Phone,
		CurrentAddress:        dto.CurrentAddress,
		Address:               dto.Address,
		AvatarPath:            dto.AvatarPath,
		BankAccount:           dto.BankAccount,
		IdentifyNumber:        dto.IdentifyNumber,
		IdentifyIssueDate:     dto.IdentifyIssueDate,
		IdentifyPlace:         dto.IdentifyPlace,
		EmergencyContact:      dto.EmergencyContact,
		EmergencyContactPhone: dto.EmergencyContactPhone,
		TaxCode:               dto.TaxCode,
		MezonID:               dto.MezonID,
		LevelID:               dto.LevelID,
		BranchID:              dto.BranchID,
		PositionID:            dto.PositionID,
		UserTypeID:            dto.UserTypeID,
	}
}
