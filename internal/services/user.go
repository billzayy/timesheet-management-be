package services

import (
	"context"
	"fmt"
	"strconv"

	"github.com/billzayy/timesheet-management-be/internal/dto"
	"github.com/billzayy/timesheet-management-be/internal/middleware"
	"github.com/billzayy/timesheet-management-be/internal/models"
	"github.com/billzayy/timesheet-management-be/internal/repositories"
	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(ctx context.Context, dto *dto.RequestUserDTO) error
	GetAllUsers(ctx context.Context, limit, offset string) ([]dto.GetUserDTO, error)
	GetByEmail(ctx context.Context, email string) (dto.GetUserDTO, error)
	GetById(ctx context.Context, id uuid.UUID) (dto.GetUserDTO, error)
	DeleteByEmail(ctx context.Context, email string) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) CreateUser(ctx context.Context, dto *dto.RequestUserDTO) error {
	requestUser := dto.ToUser()

	hashedPass, err := middleware.HashPassword(dto.Password)

	if err != nil {
		return fmt.Errorf("failed on hash password")
	}

	requestUser.Password = hashedPass

	return s.repo.Create(ctx, &requestUser)
}

func (s *userService) GetAllUsers(ctx context.Context, limitStr, offsetStr string) ([]dto.GetUserDTO, error) {
	limit, err := strconv.Atoi(limitStr)

	if err != nil {
		return nil, err
	}

	offset, err := strconv.Atoi(offsetStr)

	if err != nil {
		return nil, err
	}

	rows, err := s.repo.FindAll(ctx, limit, offset)

	if err != nil {
		return nil, err
	}

	result := make([]dto.GetUserDTO, 0, len(rows))

	for _, r := range rows {
		result = append(result, convertUserReadToDTO(r))
	}

	return result, nil
}

func (s *userService) GetByEmail(ctx context.Context, email string) (dto.GetUserDTO, error) {
	row, err := s.repo.FindByEmail(ctx, email)

	if err != nil {
		return dto.GetUserDTO{}, err
	}

	result := convertUserReadToDTO(row)

	return result, nil
}

func (s *userService) GetById(ctx context.Context, id uuid.UUID) (dto.GetUserDTO, error) {
	row, err := s.repo.FindById(ctx, id)

	if err != nil {
		return dto.GetUserDTO{}, err
	}

	result := convertUserReadToDTO(row)

	return result, nil
}

func (s *userService) DeleteByEmail(ctx context.Context, email string) error {
	return s.repo.Delete(ctx, email)
}

func convertUserReadToDTO(r models.UserRead) dto.GetUserDTO {
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
	}
}
