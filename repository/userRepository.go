package repository

import (
	"Databriz-Meetings-API-Go/db/entities"
	"Databriz-Meetings-API-Go/models/azure"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	CreateUsers(members []azure.Member)
	GetByEmail() *entities.UserEntity
}

type userRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{Db: db}
}

func (u *userRepository) CreateUsers(members []azure.Member) {
	tx := u.Db.Begin()
	for _, user := range members {
		userEnt := entities.UserEntity{
			Email: user.Identity.UniqueName,
			Name:  user.Identity.DisplayName,
		}
		tx.Where(entities.UserEntity{Email: userEnt.Email}).Assign(userEnt).FirstOrCreate(&userEnt)
	}
	tx.Commit()
}

func (u *userRepository) GetByEmail() *entities.UserEntity {
	panic("implement me") // TODO
}
