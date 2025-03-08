package application

import (
	"event_driven/src/users/application/entities"
	"event_driven/src/users/domain/repositories"
)

type LoginUseCase struct{
	db repositories.IUser
}

func NewLoginUseCase(db repositories.IUser) *LoginUseCase {
	return &LoginUseCase{db: db}
}

func (uc *LoginUseCase) Execute(email, password string) (*entities.UserLogin, error) {
	user, err := uc.db.Login(email, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
