package services

import (
	"context"

	"github.com/billzayy/timesheet-management-be/internal/dto"
	"github.com/google/uuid"
)

func (s *organizeService) GetAllLevels(ctx context.Context, limit, offset int) ([]dto.LevelDTO, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	rawData, err := s.repo.FindAllLevel(ctx, limit, offset)

	if err != nil {
		return nil, err
	}

	result := make([]dto.LevelDTO, 0, len(rawData))

	for _, v := range rawData {
		result = append(result, dto.LevelDTO{
			ID:          v.ID,
			Name:        v.Name,
			DisplayName: v.DisplayName,
			Code:        v.Code,
			Color:       v.Color,
		})
	}

	return result, nil
}

func (s *organizeService) CreateLevel(ctx context.Context, input dto.LevelDTO, id uuid.UUID) error {
	convert := input.ToCreateLevel(id)

	return s.repo.CreateLevel(ctx, convert)
}

func (s *organizeService) UpdateLevel(ctx context.Context, input dto.LevelDTO) error {
	return s.repo.UpdateLevel(ctx, input)
}

func (s *organizeService) DeleteLevel(ctx context.Context, id int64) error {
	return s.repo.DeleteLevel(ctx, id)
}
