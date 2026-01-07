package repositories

import (
	"context"

	backend "github.com/billzayy/timesheet-management-be"
	"github.com/billzayy/timesheet-management-be/internal/models"
	"gorm.io/gorm"
)

func (r *organizeRepo) FindAllPosition(ctx context.Context, limit, offset int) ([]models.Position, error) {
	data, err := gorm.G[models.Position](r.db).Limit(limit).Offset(offset).Find(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *organizeRepo) CreatePosition(ctx context.Context, input models.Position) error {
	return gorm.G[models.Position](r.db).Create(ctx, &input)
}

func (r *organizeRepo) UpdatePosition(ctx context.Context, input models.Position) error {
	rowAffected, err := gorm.G[models.Position](r.db).Where("id = ?", input.ID).Updates(ctx, models.Position{
		Name:      input.Name,
		ShortName: input.ShortName,
		Code:      input.Code,
		Color:     input.Color,
	})

	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return backend.ErrPositionNotFound
	}
	return nil
}

func (r *organizeRepo) DeletePosition(ctx context.Context, id int64) error {
	rowAffected, err := gorm.G[models.Position](r.db).Where("id = ?", id).Delete(ctx)

	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return backend.ErrPositionNotFound
	}
	return nil
}
