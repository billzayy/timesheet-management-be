package services

import (
	"context"

	"github.com/billzayy/timesheet-management-be/internal/dto"
	"github.com/billzayy/timesheet-management-be/internal/models"
	"github.com/billzayy/timesheet-management-be/internal/repositories"
	"github.com/google/uuid"
)

type RoleService interface {
	GetList(ctx context.Context, limit, offset int) ([]dto.RoleDTO, error)
	CreateRole(ctx context.Context, input dto.RoleDTO, id uuid.UUID) error
	DeleteRole(ctx context.Context, id int64) error
}

type roleService struct {
	repo repositories.RoleRepository
}

func NewRoleService(repo repositories.RoleRepository) RoleService {
	return &roleService{repo: repo}
}

func (s *roleService) GetList(ctx context.Context, limit, offset int) ([]dto.RoleDTO, error) {
	data, err := s.repo.FindAll(ctx, limit, offset)

	if err != nil {
		return nil, err
	}

	result := make([]dto.RoleDTO, 0, len(data))

	for _, v := range data {
		result = append(result, dto.RoleDTO{
			ID:          v.ID,
			Name:        v.Name,
			DisplayName: v.DisplayName,
			Description: v.Description,
		})
	}

	return result, nil
}

func (s *roleService) CreateRole(ctx context.Context, input dto.RoleDTO, id uuid.UUID) error {
	convert := models.Role{
		Name:        input.Name,
		DisplayName: input.DisplayName,
		Description: input.DisplayName,
		CreatedBy:   id,
	}

	return s.repo.Create(ctx, convert)
}

func (s *roleService) DeleteRole(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
