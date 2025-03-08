package application

import (
	"event_driven/src/courses/domain/entities"
	"event_driven/src/courses/domain/repositories"
)

type GetCoursesUseCase struct {
	db repositories.ICourse
}

func NewGetCoursesUseCase(db repositories.ICourse) *GetCoursesUseCase {
    return &GetCoursesUseCase{db: db}
}

func (gc *GetCoursesUseCase) Execute() (*[]entities.Course, error) {
    courses, err := gc.db.GetCourses()
    if err != nil {
        return nil, err
    }
    return courses, err
}