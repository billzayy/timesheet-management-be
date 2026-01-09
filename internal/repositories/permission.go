package repositories

import (
	"encoding/json"
	"fmt"

	"github.com/billzayy/timesheet-management-be/internal/models"
	"gorm.io/gorm"
)

type PermissionRepository interface {
	FindAllPermissions() ([]models.PermissionNode, error)
	FindPermissionWithRoleId(id int64) ([]models.PermissionNode, error)
	FindGrantedPermission(id int64) ([]string, error)
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

func (r *permissionRepo) FindGrantedPermission(id int64) ([]string, error) {
	var names []string

	err := r.db.
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
