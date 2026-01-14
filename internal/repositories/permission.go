package repositories

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/billzayy/timesheet-management-be/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PermissionRepository interface {
	FindAllPermissions() ([]models.PermissionNode, error)
	FindPermissionWithRoleId(id int64) ([]models.PermissionNode, error)
	FindGrantedPermissionByRoleId(ctx context.Context, id int64) ([]string, error)
	FindAllNamePermission(ctx context.Context) ([]string, error)
	FindGrantedPermissionByUserId(ctx context.Context, id uuid.UUID) ([]string, error)
}

type permissionRepo struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return &permissionRepo{db}
}

func (r *permissionRepo) FindAllPermissions() ([]models.PermissionNode, error) {
	var raw []byte
	var result []models.PermissionNode

	err := r.db.Raw(`
		SELECT jsonb_agg(
			jsonb_build_object(
    	'name', p.name,
    	'display_name', p.display_name,
    	'children', get_children(p.id)
  		)
		)
		FROM permissions p
		WHERE parent_id IS NULL;
	`).Row().Scan(&raw)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *permissionRepo) FindAllNamePermission(ctx context.Context) ([]string, error) {
	var result []string

	err := r.db.WithContext(ctx).
		Table("permissions p").
		Select("p.name").
		Scan(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *permissionRepo) FindPermissionWithRoleId(id int64) ([]models.PermissionNode, error) {
	var raw []byte
	var result []models.PermissionNode

	query := fmt.Sprintf(`
		SELECT jsonb_agg(
			jsonb_build_object(
    	'name', p.name,
    	'display_name', p.display_name,
    	'children', get_children(p.id)
  		)
		)
		FROM permissions p
		INNER JOIN role_permissions rp ON rp.permission_id = p.id
		WHERE parent_id IS NULL AND rp.role_id=?;
		`)

	err := r.db.Raw(query, id).Row().Scan(&raw)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *permissionRepo) FindGrantedPermissionByRoleId(ctx context.Context, id int64) ([]string, error) {
	var names []string

	err := r.db.WithContext(ctx).
		Table("permissions p").
		Select("p.name").
		Joins("INNER JOIN role_permissions rp ON rp.permission_id = p.id").
		Where("rp.role_id = ?", id).
		Scan(&names).Error

	if err != nil {
		return nil, err
	}

	return names, nil
}

func (r *permissionRepo) FindGrantedPermissionByUserId(ctx context.Context, id uuid.UUID) ([]string, error) {
	var names []string

	err := r.db.WithContext(ctx).
		Table("permissions p").
		Select("p.name").
		Joins("LEFT JOIN role_permissions rp ON rp.permission_id = p.id").
		Joins("LEFT JOIN roles r ON r.id = rp.role_id").
		Joins("LEFT JOIN user_roles ur ON ur.role_id = r.id").
		Where("ur.user_id = ?", id).
		Scan(&names).Error

	if err != nil {
		return nil, err
	}

	return names, nil
}
