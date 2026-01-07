package repositories

import (
	"context"

	backend "github.com/billzayy/timesheet-management-be"
	"github.com/billzayy/timesheet-management-be/internal/models"
	"gorm.io/gorm"
)

func (r *organizeRepo) FindAllBranches(ctx context.Context, limit, offset int) ([]models.BranchRead, error) {
	var result []models.BranchRead

	err := r.db.WithContext(ctx).
		Table("branch_working_time_summary_v").
		Order("id").
		Limit(limit).
		Offset(offset).
		Find(&result).Error

	return result, err
}

func (r *organizeRepo) CreateBranch(ctx context.Context, input models.Branch) error {
	return gorm.G[models.Branch](r.db).Create(ctx, &input)
}

func (r *organizeRepo) UpdateBranch(ctx context.Context, input models.Branch) error {
	rowAffected, err := gorm.G[models.Branch](r.db).Where("id = ?", input.ID).Updates(ctx, models.Branch{
		Name:        input.Name,
		DisplayName: input.DisplayName,
		Code:        input.Code,
		Color:       input.Color,
	})

	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return backend.ErrPositionNotFound
	}
	return nil
}

func (r *organizeRepo) DeleteBranch(ctx context.Context, id int64) error {
	rowAffected, err := gorm.G[models.Branch](r.db).Where("id = ?", id).Delete(ctx)

	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return backend.ErrPositionNotFound
	}
	return nil
}
