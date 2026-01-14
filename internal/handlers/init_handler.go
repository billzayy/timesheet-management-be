package handlers

import "github.com/billzayy/timesheet-management-be/internal/services"

type Handlers struct {
	UserHandler       *UserHandler
	AuthHandler       *AuthHandler
	OrganizeHandler   *OrganizeHandler
	PermissionHandler *PermissionHandler
	RoleHandler       *RoleHandler
}

func NewHandlers(s *services.Services) *Handlers {
	return &Handlers{
		UserHandler:       NewUserHandler(s.UserService),
		AuthHandler:       NewAuthHandler(s.AuthService, s.PermissionService),
		OrganizeHandler:   NewOrganizeHandler(s.OrganizeService),
		PermissionHandler: NewPermissionHandler(s.PermissionService),
		RoleHandler:       NewRoleHandler(s.RoleService),
	}
}
