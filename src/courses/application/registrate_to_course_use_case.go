package application

import (
    "context"
    "event_driven/src/courses/application/models"
    "event_driven/src/courses/application/repository"
    "event_driven/src/courses/domain/entities"
    "event_driven/src/courses/domain/repositories"
    "fmt"
)

type RegistrateToCourseUseCase struct {
    db repositories.ICourse
    rb repository.ICourseNotification
}

func NewRegistrateToCourseUseCase(db repositories.ICourse, rb repository.ICourseNotification) *RegistrateToCourseUseCase {
    return &RegistrateToCourseUseCase{db: db, rb: rb}
}

func (uc *RegistrateToCourseUseCase) Execute(ctx context.Context, registration *entities.Registration) (*entities.RegistrationInfo, error) {
    registrationCreated, err := uc.db.RegistrateToCourse(*registration)
    if err != nil {
        return nil, err
    }
    var info models.CourseRegistredEvent
    info.IDUserTeacher = registrationCreated.IDUserTeacher
    info.Message = "El usuario " + registrationCreated.UserStudent + " se ha registrado a su curso"
    info.PersonEmit = registrationCreated.UserStudent
    registred, err := uc.rb.Publish(ctx, info)
    if err != nil {
        return nil, err
    }
    if registred == nil {
		return nil, fmt.Errorf("¿Qué pasó?")
	}
    fmt.Println("Publicación exitosa:", info)

    return registrationCreated, nil
}
