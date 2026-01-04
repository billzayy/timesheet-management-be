package handlers

import "github.com/billzayy/timesheet-management-be/internal/services"

type Handlers struct {
	UserHandler *UserHandler
}

func NewHandlers(s *services.Services) *Handlers {
	return &Handlers{
		UserHandler: NewUserHandler(s.UserService),
	}
}
