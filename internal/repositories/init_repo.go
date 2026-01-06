package repositories

import "gorm.io/gorm"

type Repositories struct {
	UserRepository     UserRepository
	OrganizeRepository OrganizeRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepository:     NewUserRepository(db),
		OrganizeRepository: NewOrganizeRepository(db),
	}
}
