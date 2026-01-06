package services

import (
	"github.com/billzayy/timesheet-management-be/internal/repositories"
)

type Services struct {
	UserService     UserService
	AuthService     AuthService
	OrganizeService OrganizeService
}

func NewServices(r *repositories.Repositories) *Services {
	return &Services{
		UserService:     NewUserService(r.UserRepository),
		AuthService:     NewAuthService(r.UserRepository),
		OrganizeService: NewOrganizeService(r.OrganizeRepository),
	}
}
