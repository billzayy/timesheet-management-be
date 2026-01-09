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
	GetRoleById(ctx context.Context, id int64) (dto.RolePermissionDTO, error)
	CreateRole(ctx context.Context, input dto.RoleDTO, id uuid.UUID) error
	DeleteRole(ctx context.Context, id int64) error
}

type roleService struct {
	roleRepo       repositories.RoleRepository
	permissionRepo repositories.PermissionRepository
	userRepo       repositories.UserRepository
}

func NewRoleService(
	roleRep repositories.RoleRepository,
	perRep repositories.PermissionRepository,
	userRep repositories.UserRepository) RoleService {
	return &roleService{
		roleRepo:       roleRep,
		permissionRepo: perRep,
		userRepo:       userRep}
}

func (s *roleService) GetList(ctx context.Context, limit, offset int) ([]dto.RoleDTO, error) {
	data, err := s.roleRepo.FindAll(ctx, limit, offset)

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

func (s *roleService) GetRoleById(ctx context.Context, id int64) (dto.RolePermissionDTO, error) {
	roleData, err := s.roleRepo.FindRoleById(ctx, id)

	if err != nil {
		return dto.RolePermissionDTO{}, err
	}

	permissionDatas, err := s.permissionRepo.FindAllPermissions()

	if err != nil {
		return dto.RolePermissionDTO{}, err
	}

	grantedNames, err := s.permissionRepo.FindGrantedPermission(id)

	if err != nil {
		return dto.RolePermissionDTO{}, err
	}

	userDatas, err := s.userRepo.FindByRoleId(ctx, 1)

	if err != nil {
		return dto.RolePermissionDTO{}, err
	}

	userList := converUserDTO(userDatas)

	return dto.RolePermissionDTO{
		Name:               roleData.Name,
		DisplayName:        roleData.DisplayName,
		Description:        roleData.Description,
		Permissions:        permissionDatas,
		GrantedPermissions: grantedNames,
		Users:              userList,
	}, nil
}

func (s *roleService) CreateRole(ctx context.Context, input dto.RoleDTO, id uuid.UUID) error {
	convert := models.Role{
		Name:        input.Name,
		DisplayName: input.DisplayName,
		Description: input.DisplayName,
		CreatedBy:   id,
	}

	return s.roleRepo.Create(ctx, convert)
}

func (s *roleService) DeleteRole(ctx context.Context, id int64) error {
	return s.roleRepo.Delete(ctx, id)
}

func converUserDTO(input []models.UserRead) []dto.RoleUserDTO {
	result := make([]dto.RoleUserDTO, 0, len(input))

	for _, m := range input {
		result = append(result,
			dto.RoleUserDTO{
				FullName:   m.FullName,
				Email:      m.Email,
				AvatarPath: m.AvatarPath,
				LevelID:    m.LevelID,
				LevelName:  m.LevelName,
				BranchID:   m.BranchID,
				BranchName: m.LevelName,
				// BranchColor:  m.BranchColor,
				PositionID:   m.PositionID,
				PositionName: m.PositionName,
				UserTypeID:   m.UserTypeID,
				UserTypeName: m.UserTypeName,
				UserID:       m.UserID,
			})
	}

	return result
}
