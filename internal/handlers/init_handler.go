package handlers

import "github.com/billzayy/timesheet-management-be/internal/services"

type Handlers struct {
	UserHandler     *UserHandler
	AuthHandler     *AuthHandler
	OrganizeHandler *OrganizeHandler
}

func NewHandlers(s *services.Services) *Handlers {
	return &Handlers{
		UserHandler:     NewUserHandler(s.UserService),
		AuthHandler:     NewAuthHandler(s.AuthService),
		OrganizeHandler: NewOrganizeHandler(s.OrganizeService),
	}
}
