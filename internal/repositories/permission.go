package repositories

import (
	"encoding/json"

	"github.com/billzayy/timesheet-management-be/internal/models"
	"gorm.io/gorm"
)

type PermissionRepository interface {
	FindAllPermissions() ([]models.PermissionNode, error)
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
