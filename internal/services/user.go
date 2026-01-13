package services

import (
	"context"
	"fmt"
	"strconv"

	"github.com/billzayy/timesheet-management-be/internal/dto"
	"github.com/billzayy/timesheet-management-be/internal/helper"
	"github.com/billzayy/timesheet-management-be/internal/middleware"
	"github.com/billzayy/timesheet-management-be/internal/repositories"
	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(ctx context.Context, dto *dto.RequestUserDTO, id uuid.UUID) error
	GetAllUsers(ctx context.Context, limit, offset string) ([]dto.GetUserDTO, error)
	GetByEmail(ctx context.Context, email string) (dto.GetUserDTO, error)
	GetById(ctx context.Context, id uuid.UUID) (dto.GetUserDTO, error)
	Delete(ctx context.Context, email string) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) CreateUser(ctx context.Context, input *dto.RequestUserDTO, id uuid.UUID) error {
	requestUser := input.ToUser()

	hashedPass, err := middleware.HashPassword(input.Password)

	if err != nil {
		return fmt.Errorf("failed on hash password")
	}

	requestUser.Password = hashedPass
	requestUser.CreatedBy = &id

	morning, afternoon, err := helper.ConvertShiftTime(input, id)

	if err != nil {
		return err
	}

	return s.repo.Create(ctx, &requestUser, morning, afternoon, input.RoleId)
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
		result = append(result, helper.ConvertUserReadToDTO(r))
	}

	return result, nil
}

func (s *userService) GetByEmail(ctx context.Context, email string) (dto.GetUserDTO, error) {
	row, err := s.repo.FindByEmail(ctx, email)

	if err != nil {
		return dto.GetUserDTO{}, err
	}

	result := helper.ConvertUserReadToDTO(row)

	return result, nil
}

func (s *userService) GetById(ctx context.Context, id uuid.UUID) (dto.GetUserDTO, error) {
	row, err := s.repo.FindById(ctx, id)

	if err != nil {
		return dto.GetUserDTO{}, err
	}

	result := helper.ConvertUserReadToDTO(row)

	return result, nil
}

func (s *userService) Delete(ctx context.Context, inputId string) error {
	id, err := uuid.Parse(inputId)

	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, id)
}
