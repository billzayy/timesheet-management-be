package repositories

import (
	"context"

	backend "github.com/billzayy/timesheet-management-be"
	"github.com/billzayy/timesheet-management-be/internal/models"
	"gorm.io/gorm"
)

func (r *organizeRepo) FindAllLevel(ctx context.Context, limit, offset int) ([]models.Level, error) {
	data, err := gorm.G[models.Level](r.db).Limit(limit).Offset(offset).Find(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *organizeRepo) CreateLevel(ctx context.Context, input models.Level) error {
	return gorm.G[models.Level](r.db).Create(ctx, &input)
}

func (r *organizeRepo) UpdateLevel(ctx context.Context, input models.Level) error {
	rowAffected, err := gorm.G[models.Level](r.db).Where("id = ?", input.ID).Updates(ctx, models.Level{
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

func (r *organizeRepo) DeleteLevel(ctx context.Context, id int64) error {
	rowAffected, err := gorm.G[models.Level](r.db).Where("id = ?", id).Delete(ctx)

	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return backend.ErrPositionNotFound
	}
	return nil
}
