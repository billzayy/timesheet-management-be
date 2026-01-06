package services

import (
	"context"

	"github.com/billzayy/timesheet-management-be/internal/dto"
	"github.com/billzayy/timesheet-management-be/internal/repositories"
	"github.com/google/uuid"
)

type OrganizeService interface {
	/* Levels */
	GetAllLevels(ctx context.Context, limit, offset int) ([]dto.LevelDTO, error)
	CreateLevel(ctx context.Context, input dto.LevelDTO, id uuid.UUID) error
	UpdateLevel(ctx context.Context, input dto.LevelDTO) error
	DeleteLevel(ctx context.Context, id int64) error

	/* Position */
	GetAllPositions(ctx context.Context, limit, offset int) ([]dto.PositionDTO, error)
	CreatePosition(ctx context.Context, input dto.PositionDTO, id uuid.UUID) error
	UpdatePosition(ctx context.Context, input dto.PositionDTO) error
	DeletePosition(ctx context.Context, id int64) error

	/* User Type */
	GetAllUserTypes(ctx context.Context, limit, offset int) ([]dto.UserTypeDTO, error)
	CreateUserType(ctx context.Context, input dto.UserTypeDTO, id uuid.UUID) error
	UpdateUserType(ctx context.Context, input dto.UserTypeDTO) error
	DeleteUserType(ctx context.Context, id int64) error
}

type organizeService struct {
	repo repositories.OrganizeRepository
}

func NewOrganizeService(repo repositories.OrganizeRepository) OrganizeService {
	return &organizeService{repo}
}
