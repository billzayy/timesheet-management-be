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
	Create(ctx context.Context, user *models.User) error
	FindAll(ctx context.Context, limit, offset int) ([]models.UserRead, error)
	FindByEmail(ctx context.Context, email string) (models.UserRead, error)
	FindById(ctx context.Context, id uuid.UUID) (models.UserRead, error)
	Delete(ctx context.Context, email string) error
	CheckEmailAndPassword(ctx context.Context, email string) (*models.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db}
}

func (r *userRepo) Create(ctx context.Context, user *models.User) error {
	return gorm.G[models.User](r.db).Create(ctx, user)
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

func (r *userRepo) Delete(ctx context.Context, email string) error {
	rowAffected, err := gorm.G[models.User](r.db).Where("email = ?", email).Delete(ctx)

	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return backend.ErrUserNotFound
	}

	return nil
}

func (r *userRepo) CheckEmailAndPassword(ctx context.Context, email string) (*models.User, error) {
	data, err := gorm.G[models.User](r.db).Where("email = ?", email).First(ctx)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, backend.ErrUserNotFound
	}

	return &data, nil
}
