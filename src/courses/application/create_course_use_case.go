package application

import (
	
	"event_driven/src/courses/domain/entities"
	"event_driven/src/courses/domain/repositories"
)

type CreateCourseUseCase struct {
	db repositories.ICourse
}

func NewCreateCourseUseCase(db repositories.ICourse) *CreateCourseUseCase {
    return &CreateCourseUseCase{db: db}
}

func (uc *CreateCourseUseCase) Execute(course *entities.Course) (*entities.Course, error) {
	courseCreated, err := uc.db.CreateCourse(*course)
	if err != nil {
		return nil, err
	}
	return courseCreated, nil
}