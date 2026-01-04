package services

import (
	"github.com/billzayy/timesheet-management-be/internal/repositories"
)

type Services struct {
	UserService UserService
}

func NewServices(r *repositories.Repositories) *Services {
	return &Services{
		UserService: NewUserService(r.UserRepository),
	}
}
