package repositories

import (
	"event_driven/src/users/domain/entities"
	entitiesUser "event_driven/src/users/application/entities"
)

type IUser interface {
	CreateUser(user entities.User) (*entities.User, error)
	GetUser(id int) (*entities.User, error)
	Login(gmail string, password string) (*entitiesUser.UserLogin, error)
}
