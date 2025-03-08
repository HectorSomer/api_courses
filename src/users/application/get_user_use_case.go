package application

import (
	"event_driven/src/users/domain/entities"
	"event_driven/src/users/domain/repositories"
)

type GetUserUseCase struct {
	db repositories.IUser
}

func NewGetUserUseCase(db repositories.IUser) *GetUserUseCase {
    return &GetUserUseCase{db: db}
}

func (gu *GetUserUseCase) Execute(id int) (*entities.User, error) {
	user, err := gu.db.GetUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}