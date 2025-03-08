package controllers

import (
	"context"
	"event_driven/src/courses/application"
	"event_driven/src/courses/domain/entities"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RegistrateToCourseController struct {
	useCase *application.RegistrateToCourseUseCase
}

func NewRegistrateToCourseController(useCase *application.RegistrateToCourseUseCase) *RegistrateToCourseController {
    return &RegistrateToCourseController{useCase: useCase}
}

func (rtcc *RegistrateToCourseController) RegistrateToCourse(c *gin.Context){
	var registration entities.Registration
	if err := c.ShouldBindJSON(&registration); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
        return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
	registrationToCourse, err := rtcc.useCase.Execute(ctx, &registration)
	if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
	response := gin.H{
		"post": gin.H{
			"course_id": registrationToCourse.IDCourse,
			"user_id": registrationToCourse.IDUserStudent,
			"teacher_id": registrationToCourse.IDUserTeacher,
        },
	}
	c.JSON(http.StatusCreated, response)
}