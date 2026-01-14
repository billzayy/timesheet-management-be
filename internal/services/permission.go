package services

import (
	"context"

	"github.com/billzayy/timesheet-management-be/internal/dto"
	"github.com/billzayy/timesheet-management-be/internal/models"
	"github.com/billzayy/timesheet-management-be/internal/repositories"
	"github.com/google/uuid"
)

type PermissionService interface {
	GeAllPermission() ([]models.PermissionNode, error)
	GetAuthConfig(ctx context.Context, id uuid.UUID) (*dto.UserConfig, error)
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

func (s *permissionService) GetAuthConfig(ctx context.Context, id uuid.UUID) (*dto.UserConfig, error) {
	allPermissions, err := s.repo.FindAllNamePermission(ctx)

	if err != nil {
		return nil, err
	}

	grantedPermissions, err := s.repo.FindGrantedPermissionByUserId(ctx, id)

	if err != nil {
		return nil, err
	}

	auth := &dto.AuthConfig{
		AllPermissions:     allPermissions,
		GrantedPermissions: grantedPermissions,
	}

	return &dto.UserConfig{Auth: *auth}, nil
}
