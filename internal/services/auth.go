package services

import (
	"context"
	"fmt"
	"time"

	"github.com/billzayy/timesheet-management-be/internal/dto"
	"github.com/billzayy/timesheet-management-be/internal/middleware"
	"github.com/billzayy/timesheet-management-be/internal/repositories"
)

type AuthService interface {
	Login(ctx context.Context, dto *dto.LoginDTO) (*dto.RespLoginDTO, error)
}

type authService struct {
	repo repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) AuthService {
	return &userService{repo}
}

func (s *userService) Login(ctx context.Context, input *dto.LoginDTO) (*dto.RespLoginDTO, error) {
	data, err := s.repo.CheckEmailAndPassword(ctx, input.Email)

	if err != nil {
		return nil, fmt.Errorf("check email and password failed")
	}

	valid, err := middleware.ValidatePassword(input.Password, data.Password)

	if err != nil {
		return nil, fmt.Errorf("failed on validate password")
	}

	if !valid {
		return nil, fmt.Errorf("password is not valid")
	}

	expiredTime := time.Now().Add(time.Hour * 24).Unix()

	accessToken, err := middleware.AccessToken(data.ID.String(), expiredTime)

	if err != nil {
		return nil, err
	}

	refreshToken, err := middleware.RefreshToken(data.ID.String())

	if err != nil {
		return nil, err
	}

	return &dto.RespLoginDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiredTime:  int(expiredTime),
		TokenType:    "Bearer",
	}, nil
}
