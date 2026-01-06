package repositories

import (
	"context"

	"github.com/billzayy/timesheet-management-be/internal/dto"
	"github.com/billzayy/timesheet-management-be/internal/models"
	"gorm.io/gorm"
)

type OrganizeRepository interface {
	/* Levels */
	FindAllLevel(ctx context.Context, limit, offset int) ([]models.Level, error)
	CreateLevel(ctx context.Context, input models.Level) error
	UpdateLevel(ctx context.Context, input dto.LevelDTO) error
	DeleteLevel(ctx context.Context, id int64) error

	/* Positions */
	FindAllPosition(ctx context.Context, limit, offset int) ([]models.Position, error)
	CreatePosition(ctx context.Context, input models.Position) error
	UpdatePosition(ctx context.Context, input dto.PositionDTO) error
	DeletePosition(ctx context.Context, id int64) error
}

type organizeRepo struct {
	db *gorm.DB
}

func NewOrganizeRepository(db *gorm.DB) OrganizeRepository {
	return &organizeRepo{db}
}
