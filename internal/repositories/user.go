package repositories

import (
	"context"
	"errors"

	backend "github.com/billzayy/timesheet-management-be"
	"github.com/billzayy/timesheet-management-be/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User,
		morningShift *models.WorkingTime, nightShift *models.WorkingTime) error
	FindAll(ctx context.Context, limit, offset int) ([]models.UserRead, error)
	FindByEmail(ctx context.Context, email string) (models.UserRead, error)
	FindById(ctx context.Context, id uuid.UUID) (models.UserRead, error)
	Delete(ctx context.Context, id uuid.UUID) error
	CheckEmailAndPassword(ctx context.Context, email string) (*models.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db}
}

func (s *userRepo) Create(ctx context.Context, user *models.User,
	morningShift *models.WorkingTime, nightShift *models.WorkingTime) error {

	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := gorm.G[models.User](tx).Create(ctx, user); err != nil {
			return err // rollback
		}

		// Morning Shift for user
		morningShift.EntityType = "user"
		morningShift.EntityId = user.ID

		if err := gorm.G[models.WorkingTime](tx).Create(ctx, morningShift); err != nil {
			return err // rollback
		}

		// Afternoon Shift for user
		nightShift.EntityType = "user"
		nightShift.EntityId = user.ID

		if err := gorm.G[models.WorkingTime](tx).Create(ctx, nightShift); err != nil {
			return err // rollback
		}

		return nil // commit
	})
}

func (r *userRepo) FindAll(ctx context.Context, limit, offset int) ([]models.UserRead, error) {
	var result []models.UserRead

	err := r.db.WithContext(ctx).
		Table("user_daily_summary_v").
		Order("user_id").
		Limit(limit).
		Offset(offset).
		Find(&result).Error

	return result, err
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (models.UserRead, error) {
	var result models.UserRead

	err := r.db.WithContext(ctx).
		Table("user_daily_summary_v").
		Where("email = ?", email).
		First(&result).Error

	if err != nil {
		return result, err
	}

	if result.UserID == uuid.Nil {
		return result, backend.ErrUserNotFound
	}

	return result, nil
}

func (r *userRepo) FindById(ctx context.Context, id uuid.UUID) (models.UserRead, error) {
	var result models.UserRead

	err := r.db.WithContext(ctx).
		Table("user_daily_summary_v").
		Where("user_id = ?", id).
		First(&result).Error

	if result.Email == "" {
		return result, backend.ErrUserNotFound
	}

	return result, err
}

func (r *userRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		rowAffected, err := gorm.G[models.User](tx).Where("id = ?", id).Delete(ctx)

		if err != nil {
			return err // rollback
		}

		if rowAffected == 0 {
			return backend.ErrUserNotFound
		}

		rowAffected, err = gorm.G[models.WorkingTime](tx).Where("entity_id = ?", id).Delete(ctx)

		if err != nil {
			return err // rollback
		}

		if rowAffected == 0 {
			return backend.ErrUserNotFound
		}

		return nil
	})
}

func (r *userRepo) CheckEmailAndPassword(ctx context.Context, email string) (*models.User, error) {
	data, err := gorm.G[models.User](r.db).Where("email = ?", email).First(ctx)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, backend.ErrUserNotFound
	}

	return &data, nil
}
