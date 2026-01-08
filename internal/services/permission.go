package services

import (
	"github.com/billzayy/timesheet-management-be/internal/models"
	"github.com/billzayy/timesheet-management-be/internal/repositories"
)

type PermissionService interface {
	GeAllPermission() ([]models.PermissionNode, error)
}

type permissionService struct {
	repo repositories.PermissionRepository
}

func NewPermissionService(repo repositories.PermissionRepository) PermissionService {
	return &permissionService{repo}
}

func (s *permissionService) GeAllPermission() ([]models.PermissionNode, error) {
	return s.repo.FindAllPermissions()
}
