package repositories

import (
	"context"

	"github.com/billzayy/timesheet-management-be/internal/models"
	"gorm.io/gorm"
)

type OrganizeRepository interface {
	/* Branches */
	FindAllBranches(ctx context.Context, limit, offset int) ([]models.BranchRead, error)
	CreateBranch(ctx context.Context, input models.Branch) error
	UpdateBranch(ctx context.Context, input models.Branch) error
	DeleteBranch(ctx context.Context, id int64) error

	/* Levels */
	FindAllLevel(ctx context.Context, limit, offset int) ([]models.Level, error)
	CreateLevel(ctx context.Context, input models.Level) error
	UpdateLevel(ctx context.Context, input models.Level) error
	DeleteLevel(ctx context.Context, id int64) error

	/* Positions */
	FindAllPosition(ctx context.Context, limit, offset int) ([]models.Position, error)
	CreatePosition(ctx context.Context, input models.Position) error
	UpdatePosition(ctx context.Context, input models.Position) error
	DeletePosition(ctx context.Context, id int64) error

	/* User Type */
	FindAllUserType(ctx context.Context, limit, offset int) ([]models.UserType, error)
	CreateUserType(ctx context.Context, input models.UserType) error
	UpdateUserType(ctx context.Context, input models.UserType) error
	DeleteUserType(ctx context.Context, id int64) error
}

type organizeRepo struct {
	db *gorm.DB
}

func NewOrganizeRepository(db *gorm.DB) OrganizeRepository {
	return &organizeRepo{db}
}
