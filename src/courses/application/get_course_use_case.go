package application

import (
	"event_driven/src/courses/domain/entities"
	"event_driven/src/courses/domain/repositories"
)

type GetCourseUseCase struct {
	db repositories.ICourse
}

func NewGetCourseUseCase(db repositories.ICourse) *GetCourseUseCase {
	return &GetCourseUseCase{db: db}
}

func (gc *GetCourseUseCase) Execute(id int) (*entities.Course, error){
	course, err := gc.db.GetCourse(id)
    if err != nil {
        return nil, err
    }
    return course, nil
}