package services

import (
	"github.com/billzayy/timesheet-management-be/internal/repositories"
)

type Services struct {
	UserService       UserService
	AuthService       AuthService
	OrganizeService   OrganizeService
	PermissionService PermissionService
	RoleService       RoleService
}

func NewServices(r *repositories.Repositories) *Services {
	return &Services{
		UserService:       NewUserService(r.UserRepository),
		AuthService:       NewAuthService(r.UserRepository),
		OrganizeService:   NewOrganizeService(r.OrganizeRepository),
		PermissionService: NewPermissionService(r.PermissionRepository),
		RoleService:       NewRoleService(r.RoleRepository, r.PermissionRepository, r.UserRepository),
	}
}
