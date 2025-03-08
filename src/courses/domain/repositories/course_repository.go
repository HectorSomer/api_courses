package repositories

import "event_driven/src/courses/domain/entities"

type ICourse interface {
	CreateCourse(course entities.Course) (*entities.Course, error)
	RegistrateToCourse(registration entities.Registration) (*entities.RegistrationInfo, error)
	GetCourse(id int) (*entities.Course, error)
	GetCourses() (*[]entities.Course, error)
}