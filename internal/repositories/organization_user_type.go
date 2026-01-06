package repositories

import (
	"context"

	backend "github.com/billzayy/timesheet-management-be"
	"github.com/billzayy/timesheet-management-be/internal/dto"
	"github.com/billzayy/timesheet-management-be/internal/models"
	"gorm.io/gorm"
)

func (r *organizeRepo) FindAllUserType(ctx context.Context, limit, offset int) ([]models.UserType, error) {
	data, err := gorm.G[models.UserType](r.db).Limit(limit).Offset(offset).Find(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *organizeRepo) CreateUserType(ctx context.Context, input models.UserType) error {
	return gorm.G[models.UserType](r.db).Create(ctx, &input)
}

func (r *organizeRepo) UpdateUserType(ctx context.Context, input dto.UserTypeDTO) error {
	rowAffected, err := gorm.G[models.UserType](r.db).Where("id = ?", input.ID).Updates(ctx, models.UserType{
		Name: input.Name,
		Code: input.Code,
	})

	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return backend.ErrUserTypeNotFound
	}
	return nil
}

func (r *organizeRepo) DeleteUserType(ctx context.Context, id int64) error {
	rowAffected, err := gorm.G[models.UserType](r.db).Where("id = ?", id).Delete(ctx)

	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return backend.ErrUserTypeNotFound
	}
	return nil
}
