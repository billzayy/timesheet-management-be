package helper

import (
	"strings"
	"time"

	"github.com/billzayy/timesheet-management-be/internal/dto"
	"github.com/billzayy/timesheet-management-be/internal/models"
	"github.com/google/uuid"
)

const (
	DefaultMorningStart   = "08:00"
	DefaultMorningEnd     = "12:00"
	DefaultAfternoonStart = "13:00"
	DefaultAfternoonEnd   = "17:00"
)

func ConvertUserReadToDTO(r models.UserRead) dto.GetUserDTO {
	return dto.GetUserDTO{
		FullName:              r.FullName,
		Email:                 r.Email,
		DOB:                   r.DOB,
		Gender:                r.Gender,
		Phone:                 r.Phone,
		CurrentAddress:        r.CurrentAddress,
		Address:               r.Address,
		AvatarPath:            r.AvatarPath,
		BankAccount:           *r.BankAccount,
		IdentifyNumber:        *r.IdentifyNumber,
		IdentifyIssueDate:     *r.IdentifyIssueDate,
		IdentifyPlace:         *r.IdentifyPlace,
		EmergencyContact:      r.EmergencyContact,
		EmergencyContactPhone: r.EmergencyContactPhone,
		TaxCode:               r.TaxCode,
		MezonID:               r.MezonID,
		Roles:                 r.RoleName,
		LevelID:               r.LevelID,
		BranchID:              r.BranchID,
		PositionID:            r.PositionID,
		UserTypeID:            r.UserTypeID,
		BranchName:            *r.BranchName,
		LevelName:             *r.LevelName,
		PositionName:          *r.PositionName,
		UserTypeName:          *r.UserTypeName,
		MorningStartAt:        r.MorningStartAt,
		MorningEndAt:          r.MorningEndAt,
		MorningWorkingTime:    r.MorningWorkingTime,
		AfternoonStartAt:      r.AfternoonStartAt,
		AfternoonEndAt:        r.AfternoonEndAt,
		AfternoonWorkingTime:  r.AfternoonWorkingTime,
		ID:                    r.UserID,
	}
}

func ConvertShiftTime(input *dto.RequestUserDTO, creatorId uuid.UUID) (*models.WorkingTime, *models.WorkingTime, error) {
	morningStartAt, err := time.Parse(
		"15:04", withDefault(input.MorningStartAt, DefaultMorningStart))

	if err != nil {
		return nil, nil, err
	}

	morningEndAt, err := time.Parse(
		"15:04", withDefault(input.MorningEndAt, DefaultMorningEnd))

	if err != nil {
		return nil, nil, err
	}

	afternoonStartAt, err := time.Parse(
		"15:04", withDefault(input.AfternoonStartAt, DefaultAfternoonStart))

	if err != nil {
		return nil, nil, err
	}

	afternoonEndAt, err := time.Parse(
		"15:04", withDefault(input.AfternoonEndAt, DefaultAfternoonEnd))

	if err != nil {
		return nil, nil, err
	}

	morningShift := dto.WorkingTimeDTO{
		ShiftName:    "morning",
		StartTime:    morningStartAt,
		EndTime:      morningEndAt,
		WorkingHours: input.MorningWorkingTime,
		CreatedBy:    creatorId,
	}

	afternoonShift := dto.WorkingTimeDTO{
		ShiftName:    "afternoon",
		StartTime:    afternoonStartAt,
		EndTime:      afternoonEndAt,
		WorkingHours: input.AfternoonWorkingTime,
		CreatedBy:    creatorId,
	}

	return morningShift.ToWorkingTime(), afternoonShift.ToWorkingTime(), nil
}

func withDefault(value, defaultValue string) string {
	if strings.TrimSpace(value) == "" {
		return defaultValue
	}
	return value
}
