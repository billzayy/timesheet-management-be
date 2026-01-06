package repositories

import (
	"context"

	"github.com/billzayy/timesheet-management-be/internal/dto"
	"github.com/google/uuid"
)

func (r *userRepo) FindWorkingTime(ctx context.Context, userId uuid.UUID) (dto.WorkingTimeDTO, error) {
	var result dto.WorkingTimeDTO

	err := r.db.WithContext(ctx).
		Table("working_times wt").
		Select(`
  		MAX(CASE WHEN wt.shift_name = 'morning' THEN wt.start_time END) AS morning_start_at,
  		MAX(CASE WHEN wt.shift_name = 'morning' THEN wt.end_time   END) AS morning_end_at,
  		ROUND(
    		MAX(CASE WHEN wt.shift_name = 'morning'
      		THEN EXTRACT(EPOCH FROM (wt.end_time - wt.start_time)) / 3600 END),1) 
				AS morning_working_time,

  	MAX(CASE WHEN wt.shift_name = 'afternoon' THEN wt.start_time END) AS afternoon_start_at,
  	MAX(CASE WHEN wt.shift_name = 'afternoon' THEN wt.end_time   END) AS afternoon_end_at,
  	ROUND(MAX(CASE WHEN wt.shift_name = 'afternoon'
      THEN EXTRACT(EPOCH FROM (wt.end_time - wt.start_time)) / 3600 END),1) 
		AS afternoon_working_time`).
		Joins("LEFT JOIN users u ON u.id = wt.entity_id").
		Where("entity_id = ?", userId).
		Scan(&result).Error

	return result, err
}
