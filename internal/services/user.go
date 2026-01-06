package services

import (
	"context"
	"fmt"
	"strconv"

	"github.com/billzayy/timesheet-management-be/internal/dto"
	"github.com/billzayy/timesheet-management-be/internal/middleware"
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
		return []dto.GetUserDTO{}, err
	}

	offset, err := strconv.Atoi(offsetStr)

	if err != nil {
		return []dto.GetUserDTO{}, err
	}

	return s.repo.FindAll(ctx, limit, offset)
}

func (s *userService) GetByEmail(ctx context.Context, email string) (dto.GetUserDTO, error) {
	return s.repo.FindByEmail(ctx, email)
}

func (s *userService) GetById(ctx context.Context, id uuid.UUID) (dto.GetUserDTO, error) {
	data, err := s.repo.FindById(ctx, id)

	if err != nil {
		return dto.GetUserDTO{}, err
	}

	time, err := s.repo.FindWorkingTime(ctx, id)

	if err != nil {
		return dto.GetUserDTO{}, err
	}

	data.MorningStartAt = time.MorningStartAt
	data.MorningEndAt = time.MorningEndAt
	data.MoringWorkingTime = time.MorningWorkingTime

	data.AfternoonStartAt = time.MorningStartAt
	data.AfternoonEndAt = time.MorningEndAt
	data.AfternoonWorkingTime = time.MorningWorkingTime

	return data, nil
}

func (s *userService) DeleteByEmail(ctx context.Context, email string) error {
	return s.repo.Delete(ctx, email)
}
