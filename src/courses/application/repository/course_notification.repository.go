package repository

import (
	"context"
	"event_driven/src/courses/application/models"
)

type ICourseNotification interface {
	Publish(ctx context.Context, event models.CourseRegistredEvent) (*models.CourseRegistredEvent, error)
}