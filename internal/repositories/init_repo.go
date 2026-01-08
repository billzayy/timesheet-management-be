package repositories

import "gorm.io/gorm"

type Repositories struct {
	UserRepository       UserRepository
	OrganizeRepository   OrganizeRepository
	PermissionRepository PermissionRepository
	RoleRepository       RoleRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepository:       NewUserRepository(db),
		OrganizeRepository:   NewOrganizeRepository(db),
		PermissionRepository: NewPermissionRepository(db),
		RoleRepository:       NewRoleRepository(db),
	}
}
