package services

import (
	"context"

	"github.com/billzayy/timesheet-management-be/internal/dto"
	"github.com/google/uuid"
)

func (s *organizeService) GetAllBranches(ctx context.Context, limit, offset int) ([]dto.BranchDTO, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	rawData, err := s.repo.FindAllBranches(ctx, limit, offset)

	if err != nil {
		return nil, err
	}

	result := make([]dto.BranchDTO, 0, len(rawData))

	for _, v := range rawData {
		result = append(result, dto.BranchDTO{
			ID:                   v.ID,
			DisplayName:          v.DisplayName,
			Name:                 v.Name,
			Code:                 v.Code,
			Color:                v.Color,
			MorningWorkingTime:   v.MorningWorkingTime,
			MorningStartAt:       v.MorningStartAt,
			MorningEndAt:         v.MorningEndAt,
			AfternoonWorkingTime: v.AfternoonWorkingTime,
			AfternoonStartAt:     v.MorningStartAt,
			AfternoonEndAt:       v.MorningEndAt,
		})
	}

	return result, nil
}

func (s *organizeService) CreateBranch(ctx context.Context, input dto.BranchDTO, id uuid.UUID) error {
	convert := input.ToCreateBranch(id)

	return s.repo.CreateBranch(ctx, convert)
}

func (s *organizeService) UpdateBranch(ctx context.Context, input dto.BranchDTO) error {
	convert := input.ToUpdateBranch()

	return s.repo.UpdateBranch(ctx, convert)
}

func (s *organizeService) DeleteBranch(ctx context.Context, id int64) error {
	return s.repo.DeleteBranch(ctx, id)
}
