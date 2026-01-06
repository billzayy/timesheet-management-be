package services

import (
	"context"

	"github.com/billzayy/timesheet-management-be/internal/dto"
	"github.com/google/uuid"
)

func (s *organizeService) GetAllUserTypes(ctx context.Context, limit, offset int) ([]dto.UserTypeDTO, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	rawData, err := s.repo.FindAllUserType(ctx, limit, offset)

	if err != nil {
		return nil, err
	}

	result := make([]dto.UserTypeDTO, 0, len(rawData))

	for _, v := range rawData {
		result = append(result, dto.UserTypeDTO{
			ID:   v.ID,
			Name: v.Name,
			Code: v.Code,
		})
	}

	return result, nil
}

func (s *organizeService) CreateUserType(ctx context.Context, input dto.UserTypeDTO, id uuid.UUID) error {
	convert := input.ToCreateUserType(id)

	return s.repo.CreateUserType(ctx, convert)
}

func (s *organizeService) UpdateUserType(ctx context.Context, input dto.UserTypeDTO) error {
	return s.repo.UpdateUserType(ctx, input)
}

func (s *organizeService) DeleteUserType(ctx context.Context, id int64) error {
	return s.repo.DeleteUserType(ctx, id)
}
