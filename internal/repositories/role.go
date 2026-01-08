package repositories

import (
	"context"

	backend "github.com/billzayy/timesheet-management-be"
	"github.com/billzayy/timesheet-management-be/internal/models"
	"gorm.io/gorm"
)

type RoleRepository interface {
	FindAll(ctx context.Context, limit, offset int) ([]models.Role, error)
	FindRoleById(ctx context.Context, id int64) (models.Role, error)
	Create(ctx context.Context, input models.Role) error
	Delete(ctx context.Context, id int64) error
}

type roleRepo struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepo{db: db}
}

func (r *roleRepo) FindAll(ctx context.Context, limit, offset int) ([]models.Role, error) {
	return gorm.G[models.Role](r.db).
		Limit(limit).Offset(offset).
		Order("name ASC").
		Find(ctx)
}

func (r *roleRepo) FindRoleById(ctx context.Context, id int64) (models.Role, error) {
	return gorm.G[models.Role](r.db).Where("id = ?", id).First(ctx)
}

func (r *roleRepo) Create(ctx context.Context, input models.Role) error {
	return gorm.G[models.Role](r.db).Create(ctx, &models.Role{
		Name:        input.Name,
		DisplayName: input.DisplayName,
		Description: input.Description,
		CreatedBy:   input.CreatedBy,
	})
}

func (r *roleRepo) Delete(ctx context.Context, id int64) error {
	rowAffected, err := gorm.G[models.Role](r.db).Where("id = ?", id).Delete(ctx)

	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return backend.ErrRoleNotFound
	}

	return nil
}
