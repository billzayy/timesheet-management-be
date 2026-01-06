package services

import (
	"context"

	"github.com/billzayy/timesheet-management-be/internal/dto"
	"github.com/google/uuid"
)

func (s *organizeService) GetAllPositions(ctx context.Context, limit, offset int) ([]dto.PositionDTO, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	rawData, err := s.repo.FindAllPosition(ctx, limit, offset)

	if err != nil {
		return nil, err
	}

	result := make([]dto.PositionDTO, 0, len(rawData))

	for _, v := range rawData {
		result = append(result, dto.PositionDTO{
			ID:        v.ID,
			Name:      v.Name,
			ShortName: v.ShortName,
			Code:      v.Code,
			Color:     v.Color,
		})
	}

	return result, nil
}

func (s *organizeService) CreatePosition(ctx context.Context, input dto.PositionDTO, id uuid.UUID) error {
	convert := input.ToCreatePosition(id)

	return s.repo.CreatePosition(ctx, convert)
}

func (s *organizeService) UpdatePosition(ctx context.Context, input dto.PositionDTO) error {
	return s.repo.UpdatePosition(ctx, input)
}

func (s *organizeService) DeletePosition(ctx context.Context, id int64) error {
	return s.repo.DeletePosition(ctx, id)
}
