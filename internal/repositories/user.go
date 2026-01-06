package repositories

import (
	"context"
	"errors"

	backend "github.com/billzayy/timesheet-management-be"
	"github.com/billzayy/timesheet-management-be/internal/dto"
	"github.com/billzayy/timesheet-management-be/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindAll(ctx context.Context, limit, offset int) ([]models.UserRead, error)
	FindByEmail(ctx context.Context, email string) (dto.GetUserDTO, error)
	FindById(ctx context.Context, id uuid.UUID) (dto.GetUserDTO, error)
	Delete(ctx context.Context, email string) error
	CheckEmailAndPassword(ctx context.Context, email string) (*models.User, error)
	FindWorkingTime(ctx context.Context, userId uuid.UUID) (dto.WorkingTimeDTO, error)
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
		Table("users u").
		Select(`
			u.id,
			COALESCE(u.sur_name, '') || ' ' || COALESCE(u.last_name, '') AS full_name,
			u.email,
			u.dob,
			u.gender,
			u.phone,
			u.current_address,
			u.address,
			u.avatar_path,
			u.bank_account,
			u.identify_number,
			u.identify_issue_date,
			u.identify_place,
			u.emergency_contact,
			u.emergency_contact_phone,
			u.tax_code,
			u.mezon_id,
			u.level_id,
			u.branch_id,
			u.position_id,
			u.user_type_id,

			b.name  AS branch_name,
			l.name  AS level_name,
			p.name  AS position_name,
			ut.name AS user_type_name,

			MAX(CASE WHEN wt.shift_name = 'morning' THEN wt.start_time END) AS morning_start_at,
			MAX(CASE WHEN wt.shift_name = 'morning' THEN wt.end_time   END) AS morning_end_at,
			ROUND(
				MAX(CASE WHEN wt.shift_name = 'morning'
					THEN EXTRACT(EPOCH FROM (wt.end_time - wt.start_time)) / 3600 END
				), 1
			) AS morning_working_time,

			MAX(CASE WHEN wt.shift_name = 'afternoon' THEN wt.start_time END) AS afternoon_start_at,
			MAX(CASE WHEN wt.shift_name = 'afternoon' THEN wt.end_time   END) AS afternoon_end_at,
			ROUND(
				MAX(CASE WHEN wt.shift_name = 'afternoon'
					THEN EXTRACT(EPOCH FROM (wt.end_time - wt.start_time)) / 3600 END
				), 1
			) AS afternoon_working_time
		`).
		Joins("LEFT JOIN branches b ON b.id = u.branch_id").
		Joins("LEFT JOIN levels l ON l.id = u.level_id").
		Joins("LEFT JOIN positions p ON p.id = u.position_id").
		Joins("LEFT JOIN user_type ut ON ut.id = u.user_type_id").
		Joins("LEFT JOIN working_times wt ON wt.entity_id = u.id AND wt.entity_type = 'user'").
		Group(`u.id,b.name,l.name,p.name,ut.name`).
		Order("u.id").
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

func (r *userRepo) FindById(ctx context.Context, id uuid.UUID) (dto.GetUserDTO, error) {
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
		Where("users.id = ?", id).Find(&result).Error

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
