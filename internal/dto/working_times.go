package dto

type WorkingTimeDTO struct {
	MorningStartAt     string  `json:"morning_start_at"`
	MorningEndAt       string  `json:"morning_end_at"`
	MorningWorkingTime float64 `json:"morning_working_time"`

	AfternoonStartAt     string  `json:"afternoon_start_at"`
	AfternoonEndAt       string  `json:"afternoon_end_at"`
	AfternoonWorkingTime float64 `json:"afternoon_working_time"`
}
