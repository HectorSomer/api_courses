package application

import (
	"event_driven/src/users/domain/entities"
	"event_driven/src/users/domain/repositories"
)
type CreateUserUseCase struct {
 db repositories.IUser
}

func NewCreateUserUseCase(db repositories.IUser) *CreateUserUseCase {
 return &CreateUserUseCase{db: db}
}

func (uc *CreateUserUseCase) Execute(user *entities.User) (*entities.User, error) {
 userCreated, err := uc.db.CreateUser(*user)
 if err != nil {
     return nil, err
 }
 return userCreated, nil
}