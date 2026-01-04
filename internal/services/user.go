package services

import (
	"context"

	"github.com/billzayy/timesheet-management-be/internal/dto"
	"github.com/billzayy/timesheet-management-be/internal/repositories"
)

type UserService interface {
	CreateUser(ctx context.Context, dto *dto.RequestUserDTO) error
	GetAllUsers(ctx context.Context) ([]dto.GetUserDTO, error)
	GetUserByEmail(ctx context.Context, email string) (dto.GetUserDTO, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) CreateUser(ctx context.Context, dto *dto.RequestUserDTO) error {
	requestUser := dto.ToUser()

	return s.repo.Create(ctx, &requestUser)
}

func (s *userService) GetAllUsers(ctx context.Context) ([]dto.GetUserDTO, error) {
	return s.repo.FindAll(ctx)
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (dto.GetUserDTO, error) {
	return s.repo.FindByEmail(ctx, email)
}
