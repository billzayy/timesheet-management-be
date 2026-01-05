package repositories

import (
	"context"
	"errors"

	backend "github.com/billzayy/timesheet-management-be"
	"github.com/billzayy/timesheet-management-be/internal/dto"
	"github.com/billzayy/timesheet-management-be/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindAll(ctx context.Context, limit, offset int) ([]dto.GetUserDTO, error)
	FindByEmail(ctx context.Context, email string) (dto.GetUserDTO, error)
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

func (r *userRepo) FindAll(ctx context.Context, limit, offset int) ([]dto.GetUserDTO, error) {
	var result []dto.GetUserDTO

	err := r.db.WithContext(ctx).
		Table("users").
		Select(`
			COALESCE(users.sur_name, '') || ' ' || COALESCE(users.last_name, '') AS full_name,
			users.email,
			users.dob,
			users.gender,
			users.phone,
			users.current_address,
			users.address,
			users.avatar_path,
			users.bank_account,
			users.identify_number,
			users.identify_issue_date,
			users.identify_place,
			users.emergency_contact,
			users.emergency_contact_phone,
			users.tax_code,
			users.mezon_id,
			users.level_id,
			users.branch_id,
			users.position_id,
			users.user_type_id,
			branches.name  AS branch_name,
			levels.name    AS level_name,
			positions.name AS position_name,
			user_type.name AS user_type_name
		`).
		Joins("LEFT JOIN branches ON branches.id = users.branch_id").
		Joins("LEFT JOIN levels ON levels.id = users.level_id").
		Joins("LEFT JOIN positions ON positions.id = users.position_id").
		Joins("LEFT JOIN user_type ON user_type.id = users.user_type_id").
		Limit(limit).
		Offset(offset).
		Scan(&result).Error

	return result, err
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (dto.GetUserDTO, error) {
	var result dto.GetUserDTO

	err := r.db.WithContext(ctx).
		Table("users").
		Select(`
			COALESCE(users.sur_name, '') || ' ' || COALESCE(users.last_name, '') AS full_name,
			users.email,
			users.dob,
			users.gender,
			users.phone,
			users.current_address,
			users.address,
			users.avatar_path,
			users.bank_account,
			users.identify_number,
			users.identify_issue_date,
			users.identify_place,
			users.emergency_contact,
			users.emergency_contact_phone,
			users.tax_code,
			users.mezon_id,
			users.level_id,
			users.branch_id,
			users.position_id,
			users.user_type_id,
			branches.name  AS branch_name,
			levels.name    AS level_name,
			positions.name AS position_name,
			user_type.name AS user_type_name
		`).
		Joins("LEFT JOIN branches ON branches.id = users.branch_id").
		Joins("LEFT JOIN levels ON levels.id = users.level_id").
		Joins("LEFT JOIN positions ON positions.id = users.position_id").
		Joins("LEFT JOIN user_type ON user_type.id = users.user_type_id").
		Where("users.email = ?", email).Find(&result).Error

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
